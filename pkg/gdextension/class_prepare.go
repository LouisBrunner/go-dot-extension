package gdextension

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
	"github.com/LouisBrunner/go-dot-extension/pkg/utils"
	"github.com/iancoleman/strcase"
)

type tagData struct {
	name   string
	getter string
	setter string
}

const tagName = "godot"

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

var reservedMethods = []string{
	"Destroy",
}

func (me *extension) prepareClass(ctr ClassConstructor) (*classEntry, error) {
	instance := ctr()
	typ := reflect.TypeOf(instance)
	if typ.Kind() != reflect.Pointer && typ.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a pointer to struct but got %s (%T)", typ.Kind(), instance)
	}

	className := typ.Elem().Name()
	namePtr := *gdapi.StringNameFromStr(className)
	parentPtr := *gdapi.StringNameFromStr(instance.BaseClass())

	properties := make(map[string]classProperty, typ.Elem().NumField())
	methods := make([]classMethod, 0, typ.NumMethod())
	signals := make([]classSignal, 0)
	subscribers := make([]classSubscriber, 0)

	ignoredMethods := map[string]struct{}{}
	for _, method := range reservedMethods {
		ignoredMethods[method] = struct{}{}
	}

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

		if reflectIsA[gdapi.SignalSubscribers](ftyp) {
			subscribers = append(subscribers, classSubscriber(field.Name))
			continue
		}

		if !field.IsExported() {
			continue
		}

		fieldName := strcase.ToSnake(field.Name)
		info := parseTag(field.Tag.Get(tagName))
		if info != nil && info.name != "" {
			fieldName = info.name
		}

		if reflectIsA[gdapi.Signal](ftyp) {
			signals = append(signals, classSignal{
				name:    field.Name,
				namePtr: *gdapi.StringNameFromStr(fieldName),
				// TODO: no clue how to get the args ATM
				argsPtr:  nil,
				argCount: 0,
			})
			continue
		}

		getter, setter, prop, err := makeProperty(len(methods), fieldName, info, typ, field, ignoredMethods)
		if err != nil {
			return nil, fmt.Errorf("property %q of class %q: %w", field.Name, className, err)
		}

		methods = append(methods, *getter)
		if setter != nil {
			methods = append(methods, *setter)
		}

		_, found := properties[prop.gdName]
		if found {
			return nil, fmt.Errorf("property %q of class %q: duplicate name", field.Name, className)
		}
		properties[prop.gdName] = *prop
	}

	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		if !method.IsExported() {
			continue
		}
		if _, found := ignoredMethods[method.Name]; found {
			continue
		}

		methodPtr, err := makeMethod(len(methods), method)
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

func makeProperty(idx int, fieldName string, info *tagData, clazz reflect.Type, field reflect.StructField, ignoredMethods map[string]struct{}) (*classMethod, *classMethod, *classProperty, error) {
	getterFn := any(func(me interface{}) interface{} {
		return reflectGetField(me, field.Name).Interface()
	})
	getterName := fmt.Sprintf("_get_property_%s", fieldName)
	if info != nil && info.getter != "" {
		method, found := clazz.MethodByName(info.getter)
		if !found {
			return nil, nil, nil, fmt.Errorf("getter %q not found", info.getter)
		}
		if method.Type.NumIn() != 1 || method.Type.NumOut() != 1 {
			return nil, nil, nil, fmt.Errorf("getter %q has wrong signature", info.getter)
		}
		getterFn = method.Func.Interface()
		ignoredMethods[info.getter] = struct{}{}
	}
	getter := gdapi.StringNameFromStr(getterName)

	setterFn := any(func(me interface{}, value interface{}) {
		reflectSetField(me, field.Name, value)
	})
	setterName := fmt.Sprintf("_set_property_%s", fieldName)
	if info != nil && info.setter != "" {
		if info.setter == "nil" {
			setterFn = nil
			setterName = ""
		} else {
			method, found := clazz.MethodByName(info.setter)
			if !found {
				return nil, nil, nil, fmt.Errorf("setter %q not found", info.setter)
			}
			if method.Type.NumIn() != 2 || method.Type.NumOut() != 0 {
				return nil, nil, nil, fmt.Errorf("setter %q has wrong signature", info.setter)
			}
			setterFn = method.Func.Interface()
			ignoredMethods[info.setter] = struct{}{}
		}
	}
	setter := gdapi.StringNameFromStr(setterName)

	varType, objClassName, err := typeToVariantNClass(field.Type)
	if err != nil {
		return nil, nil, nil, err
	}

	pinner := runtime.Pinner{}

	propCpy := &gdc.PropertyInfo{
		ClassName:  objClassName,
		Name:       gdapi.StringNameFromStr(fieldName).AsPtr(),
		Type:       varType,
		Hint:       uint(gdapi.PropertyHintNone),
		HintString: gdapi.StringFromStr("").AsPtr(),
		Usage:      uint(gdapi.PropertyUsageDefault),
	}
	pinner.Pin(propCpy)
	argMetadata := []gdc.ClassMethodArgumentMetadata{gdc.MethodArgumentMetadataNone}
	argMetadataP := unsafe.SliceData(argMetadata)
	pinner.Pin(argMetadataP)

	var objClassNamePtr gdc.ConstStringNamePtr
	if varType == gdc.VariantTypeObject {
		objClassNamePtr = gdc.ConstStringNamePtr(objClassName)
	}

	var setterMethod *classMethod
	if setterFn != nil {
		setterMethod = &classMethod{
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
		}, setterMethod, &classProperty{
			name:            field.Name,
			gdName:          fieldName,
			typ:             field.Type,
			property:        propCpy,
			getter:          *getter,
			setter:          *setter,
			objClassNamePtr: objClassNamePtr,
		}, nil
}

func makeMethod(idx int, method reflect.Method) (*classMethod, error) {
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
				ClassName:  gdapi.NewStringName().AsPtr(),
				Name:       gdapi.StringNameFromStr("varargs").AsPtr(),
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
			varType, className, err := typeToVariantNClass(arg)
			if err != nil {
				return nil, fmt.Errorf("argument %d: %w", i, err)
			}
			argInfo[i] = gdc.PropertyInfo{
				ClassName:  className,
				Name:       gdapi.StringNameFromStr(aname).AsPtr(),
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
		varType, className, err := typeToVariantNClass(out)
		if err != nil {
			return nil, fmt.Errorf("return value: %w", err)
		}
		retInfo = &gdc.PropertyInfo{
			ClassName:  className,
			Name:       gdapi.StringNameFromStr(out.Name()).AsPtr(),
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
