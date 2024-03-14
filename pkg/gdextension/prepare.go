package gdextension

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
	"github.com/LouisBrunner/go-dot-extension/pkg/utils"
	"github.com/iancoleman/strcase"
)

type classProperty struct {
	name     string
	property *gdc.PropertyInfo
	getter   gdapi.StringName
	setter   gdapi.StringName

	objClassNamePtr *gdapi.StringName
}

type classMethod struct {
	name        string
	method      gdc.ClassMethodInfo
	fn          interface{}
	pinner      runtime.Pinner
	saveArgInfo []gdc.PropertyInfo
}

type classSignal struct {
	name     string
	namePtr  gdapi.StringName
	argsPtr  *gdc.PropertyInfo
	argCount uint
}

type classEntry struct {
	name        string
	constructor ClassConstructor
	properties  []classProperty
	methods     []classMethod
	signals     []classSignal
	subscribers []string

	parentNamePtr gdapi.StringName
	namePtr       gdapi.StringName

	mutex     sync.Mutex
	instances map[uint64]*classInstance
}

func (me *classEntry) makeProperties() []gdc.PropertyInfo {
	properties := gdc.CNewPropertyInfoArray(len(me.properties))
	for i, property := range me.properties {
		properties[i] = *property.property
	}
	return properties
}

func typeToVariant(val reflect.Type) (gdc.VariantType, error) {
	switch val.Kind() {
	case reflect.Bool:
		return gdc.VariantTypeBool, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return gdc.VariantTypeInt, nil
	case reflect.Float32, reflect.Float64:
		return gdc.VariantTypeFloat, nil
	case reflect.String:
		return gdc.VariantTypeString, nil
	case reflect.Array, reflect.Slice:
		return gdc.VariantTypeArray, nil
	case reflect.Struct, reflect.Interface:
		bclazz, isBclass := reflect.New(val).Interface().(gdapi.BClass)
		if isBclass {
			return bclazz.Type(), nil
		}
		return gdc.VariantTypeObject, nil
	case reflect.Pointer:
		return typeToVariant(val.Elem())
	default:
		return gdc.VariantTypeNil, fmt.Errorf("unsupported type %s", val.Kind())
	}
}

type tagData struct {
	name   string
	getter string
	setter string
}

func parseTag(tag string) *tagData {
	if tag == "" {
		return nil
	}
	fields := strings.Split(tag, ",")
	data := &tagData{}
	for _, field := range fields {
		parts := strings.Split(field, ":")
		if len(parts) != 2 {
			continue
		}
		switch parts[0] {
		case "name":
			data.name = parts[1]
		case "get":
			data.getter = parts[1]
		case "set":
			data.setter = parts[1]
		}
	}
	return data
}

func (me *extension) prepareClass(ctr ClassConstructor) (*classEntry, error) {
	instance := ctr()
	typ := reflect.TypeOf(instance)
	if typ.Kind() != reflect.Pointer && typ.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a pointer to struct but got %s (%T)", typ.Kind(), instance)
	}

	className := typ.Elem().Name()
	namePtr := gdapi.StringNameFromStr(className)
	parentPtr := gdapi.StringNameFromStr(instance.BaseClass())

	properties := make([]classProperty, 0, typ.Elem().NumField())
	methods := make([]classMethod, 0, typ.NumMethod())
	signals := make([]classSignal, 0)
	subscribers := make([]string, 0)

	ignoredMethods := map[string]struct{}{}
	for _, field := range reflect.VisibleFields(typ.Elem()) {
		ftyp := field.Type
		if field.Anonymous {
			for m := 0; m < ftyp.NumMethod(); m++ {
				method := ftyp.Method(m)
				if method.IsExported() {
					ignoredMethods[method.Name] = struct{}{}
				}
			}
			ftyp = reflect.PointerTo(ftyp)
			for m := 0; m < ftyp.NumMethod(); m++ {
				method := ftyp.Method(m)
				if method.IsExported() {
					ignoredMethods[method.Name] = struct{}{}
				}
			}
			continue
		}

		_, isSubs := reflect.New(ftyp).Interface().(*gdapi.SignalSubscribers)
		if isSubs {
			subscribers = append(subscribers, field.Name)
			continue
		}

		if !field.IsExported() {
			continue
		}

		_, isSignal := reflect.New(ftyp).Interface().(*gdapi.Signal)
		if isSignal {
			sigName := strcase.ToSnake(field.Name)
			signals = append(signals, classSignal{
				name:    field.Name,
				namePtr: gdapi.StringNameFromStr(sigName),
				// TODO: no clue how to get the args ATM
				argsPtr:  nil,
				argCount: 0,
			})
			continue
		}

		getter, setter, prop, err := makeProperty(&namePtr, len(methods), field, ignoredMethods)
		if err != nil {
			return nil, fmt.Errorf("property %q of class %q: %w", field.Name, className, err)
		}

		methods = append(methods, *getter, *setter)
		properties = append(properties, *prop)
	}

	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		if !method.IsExported() {
			continue
		}
		if _, found := ignoredMethods[method.Name]; found {
			continue
		}

		methodPtr, err := makeMethod(&namePtr, len(methods), method)
		if err != nil {
			return nil, fmt.Errorf("method %q of class %q: %w", method.Name, className, err)
		}

		methods = append(methods, *methodPtr)
	}

	return &classEntry{
		name:          className,
		constructor:   ctr,
		properties:    properties,
		methods:       methods,
		signals:       signals,
		subscribers:   subscribers,
		parentNamePtr: parentPtr,
		namePtr:       namePtr,
		instances:     make(map[uint64]*classInstance),
	}, nil
}

func makeProperty(className *gdapi.StringName, idx int, field reflect.StructField, ignoredMethods map[string]struct{}) (*classMethod, *classMethod, *classProperty, error) {
	fieldName := strcase.ToSnake(field.Name)
	info := parseTag(field.Tag.Get("gde"))
	if info != nil && info.name != "" {
		fieldName = info.name
	}

	getterFn := func(me interface{}) interface{} {
		return reflect.ValueOf(me).Elem().FieldByName(field.Name).Interface()
	}
	getterName := fmt.Sprintf("_get_property_%s", fieldName)
	if info != nil && info.getter != "" {
		getterFn = func(me interface{}) interface{} {
			// TODO: need a lot of checking
			return reflect.ValueOf(me).MethodByName(info.getter).Call(nil)[0].Interface()
		}
		ignoredMethods[info.getter] = struct{}{}
	}
	getter := gdapi.StringNameFromStr(getterName)
	setterFn := func(me interface{}, value interface{}) {
		reflect.ValueOf(me).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
	}
	setterName := fmt.Sprintf("_set_property_%s", fieldName)
	if info != nil && info.setter != "" {
		setterFn = func(me interface{}, value interface{}) {
			reflect.ValueOf(me).MethodByName(info.setter).Call([]reflect.Value{reflect.ValueOf(value)})
		}
		ignoredMethods[info.setter] = struct{}{}
	}
	setter := gdapi.StringNameFromStr(setterName)

	varType, err := typeToVariant(field.Type)
	if err != nil {
		return nil, nil, nil, err
	}

	pinner := runtime.Pinner{}

	propCpy := &gdc.PropertyInfo{
		ClassName:  className.AsPtr(),
		Name:       gdapi.PStringNameFromStr(fieldName).AsPtr(),
		Type:       varType,
		Hint:       uint(gdapi.PropertyHintNone),
		HintString: gdapi.StringFromStr("").AsPtr(),
		Usage:      uint(gdapi.PropertyUsageDefault),
	}
	pinner.Pin(propCpy)
	argMetadata := []gdc.ClassMethodArgumentMetadata{gdc.MethodArgumentMetadataNone}
	argMetadataP := unsafe.SliceData(argMetadata)
	pinner.Pin(argMetadataP)

	var objClassName *gdapi.StringName
	if varType == gdc.VariantTypeObject {
		inst, ok := reflect.New(field.Type).Interface().(gdapi.Object)
		if ok {
			objClassName = gdapi.PStringNameFromStr(inst.BaseClass())
		}
	}

	return &classMethod{
			name:   getterName,
			fn:     getterFn,
			pinner: pinner,
			method: gdc.ClassMethodInfo{
				Name:                getter.AsPtr(),
				MethodUserdata:      *(*unsafe.Pointer)(unsafe.Pointer(utils.ToPointer(idx))),
				PtrcallFunc:         gdc.Callbacks.GetClassMethodInfoPtrcallFuncCallback(),
				CallFunc:            gdc.Callbacks.GetClassMethodInfoCallFuncCallback(),
				MethodFlags:         uint(gdapi.MethodFlagsDefault),
				HasReturnValue:      gdc.Bool(1),
				ReturnValueMetadata: gdc.MethodArgumentMetadataNone,
				ReturnValueInfo:     propCpy,
			},
		}, &classMethod{
			name: setterName,
			fn:   setterFn,
			method: gdc.ClassMethodInfo{
				Name:              setter.AsPtr(),
				MethodUserdata:    *(*unsafe.Pointer)(unsafe.Pointer(utils.ToPointer(idx + 1))),
				PtrcallFunc:       gdc.Callbacks.GetClassMethodInfoPtrcallFuncCallback(),
				CallFunc:          gdc.Callbacks.GetClassMethodInfoCallFuncCallback(),
				MethodFlags:       uint(gdapi.MethodFlagsDefault),
				ArgumentCount:     1,
				ArgumentsInfo:     propCpy,
				ArgumentsMetadata: argMetadataP,
			},
		}, &classProperty{
			name:            field.Name,
			property:        propCpy,
			getter:          getter,
			setter:          setter,
			objClassNamePtr: objClassName,
		}, nil
}

func makeMethod(className *gdapi.StringName, idx int, method reflect.Method) (*classMethod, error) {
	methodName := strcase.ToSnake(method.Name)
	flags := uint(gdapi.MethodFlagsDefault)
	if strings.HasPrefix(methodName, "x_") {
		methodName = strings.TrimPrefix(methodName, "x")
		flags = uint(gdapi.MethodFlagVirtual)
	}
	name := gdapi.StringNameFromStr(methodName)

	argCount := method.Type.NumIn() - 1
	var argInfo []gdc.PropertyInfo
	if method.Type.IsVariadic() {
		flags |= uint(gdapi.MethodFlagVararg)
		argCount = 1
		argInfo = []gdc.PropertyInfo{
			{
				ClassName:  className.AsPtr(),
				Name:       gdapi.PStringNameFromStr("varargs").AsPtr(),
				Type:       gdc.VariantTypeNil,
				Hint:       uint(gdapi.PropertyHintNone),
				HintString: gdapi.StringFromStr("").AsPtr(),
				Usage:      uint(gdapi.PropertyUsageDefault),
			},
		}
	} else {
		argInfo = make([]gdc.PropertyInfo, argCount)
		for i := 0; i < argCount; i += 1 {
			arg := method.Type.In(i + 1)
			aname := fmt.Sprintf("arg%d", i)
			varType, err := typeToVariant(arg)
			if err != nil {
				return nil, fmt.Errorf("argument %d: %w", i, err)
			}
			argInfo[i] = gdc.PropertyInfo{
				ClassName:  className.AsPtr(),
				Name:       gdapi.PStringNameFromStr(aname).AsPtr(),
				Type:       varType,
				Hint:       uint(gdapi.PropertyHintNone),
				HintString: gdapi.StringFromStr("").AsPtr(),
				Usage:      uint(gdapi.PropertyUsageDefault),
			}
		}
	}
	argMetadata := make([]gdc.ClassMethodArgumentMetadata, argCount)
	for i := 0; i < argCount; i += 1 {
		argMetadata[i] = gdc.MethodArgumentMetadataNone
	}

	pinner := runtime.Pinner{}
	var argMedataP *gdc.ClassMethodArgumentMetadata
	var argInfoP *gdc.PropertyInfo
	if argCount > 0 {
		argMedataP = unsafe.SliceData(argMetadata)
		argInfoP = unsafe.SliceData(argInfo)
		pinner.Pin(argMedataP)
		pinner.Pin(argInfoP)
	}

	hasReturn := gdc.Bool(0)
	var retInfo *gdc.PropertyInfo
	if method.Type.NumOut() > 1 {
		return nil, fmt.Errorf("too many return values")
	} else if method.Type.NumOut() > 0 {
		hasReturn = gdc.Bool(1)
		out := method.Type.Out(0)
		varType, err := typeToVariant(out)
		if err != nil {
			return nil, fmt.Errorf("return value: %w", err)
		}
		retInfo = &gdc.PropertyInfo{
			ClassName:  className.AsPtr(),
			Name:       gdapi.PStringNameFromStr(out.Name()).AsPtr(),
			Type:       varType,
			Hint:       uint(gdapi.PropertyHintNone),
			HintString: gdapi.StringFromStr("").AsPtr(),
			Usage:      uint(gdapi.PropertyUsageDefault),
		}
		pinner.Pin(retInfo)
	}

	return &classMethod{
		name:        method.Name,
		fn:          method.Func.Interface(),
		pinner:      pinner,
		saveArgInfo: argInfo,
		method: gdc.ClassMethodInfo{
			Name:                name.AsPtr(),
			MethodUserdata:      *(*unsafe.Pointer)(unsafe.Pointer(utils.ToPointer(idx))),
			MethodFlags:         flags,
			CallFunc:            gdc.Callbacks.GetClassMethodInfoCallFuncCallback(),
			PtrcallFunc:         gdc.Callbacks.GetClassMethodInfoPtrcallFuncCallback(),
			HasReturnValue:      hasReturn,
			ReturnValueInfo:     retInfo,
			ReturnValueMetadata: gdc.MethodArgumentMetadataNone,
			ArgumentCount:       uint(argCount),
			ArgumentsMetadata:   argMedataP,
			ArgumentsInfo:       argInfoP,
		},
	}, nil
}
