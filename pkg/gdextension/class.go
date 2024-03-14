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

func (me *extension) addClassCallbacks() {
	gdc.Callbacks.SetClassCreationInfo2ReferenceFuncHandler(func(pInstance gdc.ClassInstancePtr) {
		me.Logf(LogLevelDebug, "SetClassCreationInfoReferenceFuncHandler")
		// TODO: implement
	})
	gdc.Callbacks.SetClassCreationInfo2UnreferenceFuncHandler(func(pInstance gdc.ClassInstancePtr) {
		me.Logf(LogLevelDebug, "SetClassCreationInfoUnreferenceFuncHandler")
		// TODO: implement
	})
	gdc.Callbacks.SetClassCreationInfo2RecreateInstanceFuncHandler(func(pClassUserdata unsafe.Pointer, pObject gdc.ObjectPtr) gdc.ClassInstancePtr {
		me.Logf(LogLevelDebug, "SetClassCreationInfoRecreateInstanceFuncHandler")
		// TODO: implement
		return nil
	})
	gdc.Callbacks.SetClassCreationInfo2GetVirtualFuncHandler(func(pClassUserdata unsafe.Pointer, pName gdc.ConstStringNamePtr) gdc.ClassCallVirtual {
		me.Logf(LogLevelDebug, "SetClassCreationInfoGetVirtualFuncHandler: %s", gdapi.NewVariantWithC(gdc.ConstVariantPtr(pName)))
		// TODO: implement
		return nil
	})
	gdc.Callbacks.SetClassCreationInfo2GetVirtualCallDataFuncHandler(func(pClassUserdata unsafe.Pointer, pName gdc.ConstStringNamePtr) unsafe.Pointer {
		me.Logf(LogLevelDebug, "SetClassCreationInfoGetVirtualCallDataFuncHandler: %s", gdapi.NewVariantWithC(gdc.ConstVariantPtr(pName)))
		// TODO: implement
		return nil
	})
	gdc.Callbacks.SetClassCreationInfo2CallVirtualWithDataFuncHandler(func(pInstance gdc.ClassInstancePtr, pName gdc.ConstStringNamePtr, pVirtualCallUserdata unsafe.Pointer, pArgs *gdc.ConstTypePtr, rRet gdc.TypePtr) {
		me.Logf(LogLevelDebug, "SetClassCreationInfoCallVirtualWithDataFuncHandler: %s", gdapi.NewVariantWithC(gdc.ConstVariantPtr(pName)))
		// TODO: implement
	})
	gdc.Callbacks.SetClassCreationInfo2GetRidFuncHandler(func(pInstance gdc.ClassInstancePtr) uint64 {
		instance, err := restore[*classInstance](unsafe.Pointer(pInstance))
		if err != nil {
			me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
			return 0
		}
		return instance.id
	})
	gdc.Callbacks.SetClassCreationInfo2CreateInstanceFuncHandler(func(pUserdata unsafe.Pointer) gdc.ObjectPtr {
		class, err := restore[*classEntry](pUserdata)
		if err != nil {
			me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
			return nil
		}
		return me.createClass(class)
	})
	gdc.Callbacks.SetClassCreationInfo2FreeInstanceFuncHandler(func(pUserdata unsafe.Pointer, pInstance gdc.ClassInstancePtr) {
		instance, err := restore[*classInstance](unsafe.Pointer(pInstance))
		if err != nil {
			me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
			return
		}
		me.freeClass(instance)
	})
	gdc.Callbacks.SetClassCreationInfo2SetFuncHandler(func(pInstance gdc.ClassInstancePtr, pName gdc.ConstStringNamePtr, pValue gdc.ConstVariantPtr) gdc.Bool {
		me.Logf(LogLevelDebug, "SetClassCreationInfoSetFuncHandler: %s=%s", gdapi.NewVariantWithC(gdc.ConstVariantPtr(pName)), gdapi.NewVariantWithC(pValue))
		// TODO: implement
		return gdc.Bool(0)
	})
	gdc.Callbacks.SetClassCreationInfo2GetFuncHandler(func(pInstance gdc.ClassInstancePtr, pName gdc.ConstStringNamePtr, rRet gdc.VariantPtr) gdc.Bool {
		me.Logf(LogLevelDebug, "SetClassCreationInfoGetFuncHandler: %s", gdapi.NewVariantWithC(gdc.ConstVariantPtr(pName)))
		// TODO: implement
		return gdc.Bool(0)
	})
	gdc.Callbacks.SetClassCreationInfo2GetPropertyListFuncHandler(func(pInstance gdc.ClassInstancePtr, rCount *uint) *gdc.PropertyInfo {
		instance, err := restore[*classInstance](unsafe.Pointer(pInstance))
		if err != nil {
			me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
			return nil
		}
		properties := instance.class.makeProperties()
		*rCount = uint(len(properties))
		return unsafe.SliceData(properties)
	})
	gdc.Callbacks.SetClassCreationInfo2FreePropertyListFuncHandler(func(pInstance gdc.ClassInstancePtr, pList *gdc.PropertyInfo) {
		if pList == nil {
			return
		}
		gdc.CFreePropertyInfo(pList)
	})
	gdc.Callbacks.SetClassCreationInfo2PropertyCanRevertFuncHandler(func(pInstance gdc.ClassInstancePtr, pName gdc.ConstStringNamePtr) gdc.Bool {
		me.Logf(LogLevelDebug, "SetClassCreationInfoPropertyCanRevertFuncHandler: %s", gdapi.NewVariantWithC(gdc.ConstVariantPtr(pName)))
		// TODO: implement
		return gdc.Bool(0)
	})
	gdc.Callbacks.SetClassCreationInfo2PropertyGetRevertFuncHandler(func(pInstance gdc.ClassInstancePtr, pName gdc.ConstStringNamePtr, rRet gdc.VariantPtr) gdc.Bool {
		me.Logf(LogLevelDebug, "SetClassCreationInfoPropertyGetRevertFuncHandler: %s", gdapi.NewVariantWithC(gdc.ConstVariantPtr(pName)))
		// TODO: implement
		return gdc.Bool(0)
	})
	gdc.Callbacks.SetClassCreationInfo2ToStringFuncHandler(func(pInstance gdc.ClassInstancePtr, rIsValid *uint8, pOut gdc.StringPtr) {
		instance, err := restore[*classInstance](unsafe.Pointer(pInstance))
		if err != nil {
			me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
			*rIsValid = 0
			return
		}
		me.iface.StringOperatorPlusEqCstr(pOut, instance.String())
		*rIsValid = 1
	})
	gdc.Callbacks.SetClassMethodInfoCallFuncHandler(func(methodUserdata unsafe.Pointer, pInstance gdc.ClassInstancePtr, pArgs *gdc.ConstVariantPtr, pArgumentCount gdc.Int, rReturn gdc.VariantPtr, rError *gdc.CallError) {
		me.Logf(LogLevelDebug, "SetClassMethodInfoCallFuncHandler: %v", methodUserdata)
	})
	gdc.Callbacks.SetClassMethodInfoPtrcallFuncHandler(func(methodUserdata unsafe.Pointer, pInstance gdc.ClassInstancePtr, pArgs *gdc.ConstTypePtr, rRet gdc.TypePtr) {
		me.Logf(LogLevelDebug, "SetClassMethodInfoPtrcallFuncHandler: %v", methodUserdata)
	})
}

type classProperty struct {
	name     string
	property gdc.PropertyInfo
	getter   gdapi.StringName
	setter   gdapi.StringName
}

type classMethod struct {
	name            string
	method          gdc.ClassMethodInfo
	fn              interface{}
	pinner          runtime.Pinner
	saveArgInfo     []gdc.PropertyInfo
	saveArgMetadata []gdc.ClassMethodArgumentMetadata
}

type classEntry struct {
	name        string
	constructor ClassConstructor
	properties  []classProperty
	methods     []classMethod

	parentNamePtr gdapi.StringName
	namePtr       gdapi.StringName

	mutex     sync.Mutex
	instances map[uint64]*classInstance
}

func (me *classEntry) makeProperties() []gdc.PropertyInfo {
	properties := gdc.CNewPropertyInfoArray(len(me.properties))
	for i, property := range me.properties {
		properties[i] = property.property
	}
	return properties
}

type classInstance struct {
	class    *classEntry
	instance Class
	id       uint64
}

func (me *classInstance) String() string {
	str := ""
	stringer, cast := me.instance.(fmt.Stringer)
	if !cast {
		str = fmt.Sprintf("%#v", me.instance)
	} else {
		str = stringer.String()
	}
	return str
}

func scalarToName(val reflect.Value) (gdc.StringNamePtr, error) {
	var str gdapi.StringName
	switch val.Kind() {
	case reflect.Bool:
		str = gdapi.StringNameFromStr("bool")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		str = gdapi.StringNameFromStr("int")
	case reflect.Float32, reflect.Float64:
		str = gdapi.StringNameFromStr("float")
	case reflect.String:
		str = gdapi.StringNameFromStr("string")
	case reflect.Array, reflect.Slice:
		str = gdapi.StringNameFromStr("array")
	default:
		return nil, fmt.Errorf("unsupported type %s", val.Kind())
	}
	return str.AsPtr(), nil
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

func infoFromType(name string, val reflect.Value) (*gdc.PropertyInfo, error) {
	namePtr := gdapi.StringNameFromStr(name)
	bcl, isBclass := val.Addr().Interface().(gdapi.BClass)

	varTyp := gdc.VariantTypeObject
	var className gdc.StringNamePtr
	if isBclass {
		varTyp = bcl.Type()
	} else {
		obj, isObj := val.Interface().(Class)
		if isObj {
			classNameStr := gdapi.StringNameFromStr(obj.BaseClass())
			className = (&classNameStr).AsPtr()
		} else {
			var err error
			className, err = scalarToName(val)
			if err != nil {
				return nil, fmt.Errorf("property %q cannot be exported: %w", name, err)
			}
		}
	}
	return &gdc.PropertyInfo{
		Name:      namePtr.AsPtr(),
		Type:      varTyp,
		ClassName: className,
		Usage:     uint(gdapi.PropertyUsageDefault),
		// FIXME: missing hints
	}, nil
}

func (me *extension) prepareClass(ctr ClassConstructor) (*classEntry, error) {
	instance := ctr()
	val := reflect.ValueOf(instance)
	typ := reflect.TypeOf(instance)
	if typ.Kind() != reflect.Pointer && typ.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a pointer to struct but got %s (%T)", typ.Kind(), instance)
	}

	className := typ.Elem().Name()
	namePtr := gdapi.StringNameFromStr(className)
	parentPtr := gdapi.StringNameFromStr(instance.BaseClass())

	properties := make([]classProperty, 0, typ.Elem().NumField())
	methods := make([]classMethod, 0, typ.NumMethod())

	methodsFromParent := map[string]struct{}{}
	for _, field := range reflect.VisibleFields(typ.Elem()) {
		ftyp := field.Type
		if field.Anonymous {
			for m := 0; m < ftyp.NumMethod(); m++ {
				method := ftyp.Method(m)
				if method.IsExported() {
					methodsFromParent[method.Name] = struct{}{}
				}
			}
			ftyp = reflect.PointerTo(ftyp)
			for m := 0; m < ftyp.NumMethod(); m++ {
				method := ftyp.Method(m)
				if method.IsExported() {
					methodsFromParent[method.Name] = struct{}{}
				}
			}
		}

		continue

		if !field.IsExported() || field.Anonymous {
			continue
		}

		fieldName := strcase.ToSnake(field.Name)
		getterName := fmt.Sprintf("_get_property_%s", fieldName)
		getter := gdapi.StringNameFromStr(getterName)
		setterName := fmt.Sprintf("_set_property_%s", fieldName)
		setter := gdapi.StringNameFromStr(setterName)

		prop, err := infoFromType(fieldName, val.Elem().FieldByName(field.Name))
		if err != nil {
			return nil, fmt.Errorf("class %q: %w", className, err)
		}
		propCpy := gdc.CNewPropertyInfo()
		*propCpy = *prop
		methods = append(methods, classMethod{
			name: getterName,
			method: gdc.ClassMethodInfo{
				Name:            getter.AsPtr(),
				MethodUserdata:  *(*unsafe.Pointer)(unsafe.Pointer(utils.ToPointer[int](len(methods)))),
				PtrcallFunc:     gdc.Callbacks.GetClassMethodInfoPtrcallFuncCallback(),
				CallFunc:        gdc.Callbacks.GetClassMethodInfoCallFuncCallback(),
				HasReturnValue:  gdc.Bool(1),
				ReturnValueInfo: propCpy,
				MethodFlags:     uint(gdapi.MethodFlagsDefault | gdapi.MethodFlagConst),
				// FIXME: metadata missing
			},
			fn: func(me interface{}) interface{} {
				return reflect.ValueOf(me).Elem().FieldByName(field.Name).Interface()
			},
		})
		methods = append(methods, classMethod{
			name: setterName,
			method: gdc.ClassMethodInfo{
				Name:           setter.AsPtr(),
				MethodUserdata: *(*unsafe.Pointer)(unsafe.Pointer(utils.ToPointer[int](len(methods)))),
				PtrcallFunc:    gdc.Callbacks.GetClassMethodInfoPtrcallFuncCallback(),
				CallFunc:       gdc.Callbacks.GetClassMethodInfoCallFuncCallback(),
				MethodFlags:    uint(gdapi.MethodFlagsDefault),
				ArgumentCount:  1,
				ArgumentsInfo:  propCpy,
				// FIXME: metadata missing
			},
			fn: func(me interface{}, value interface{}) {
				reflect.ValueOf(me).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
			},
		})
		properties = append(properties, classProperty{
			name:     field.Name,
			property: *prop,
			getter:   getter,
			setter:   setter,
		})
	}

	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		if !method.IsExported() {
			continue
		}
		if _, found := methodsFromParent[method.Name]; found {
			continue
		}

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
					ClassName:  namePtr.AsPtr(),
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
					return nil, fmt.Errorf("argument %d of method %q of class %q: %w", i, method.Name, className, err)
				}
				argInfo[i] = gdc.PropertyInfo{
					ClassName:  namePtr.AsPtr(),
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
			return nil, fmt.Errorf("method %q of class %q: too many return values", method.Name, className)
		} else if method.Type.NumOut() > 0 {
			hasReturn = gdc.Bool(1)
			out := method.Type.Out(0)
			varType, err := typeToVariant(out)
			if err != nil {
				return nil, fmt.Errorf("return value of method %q of class %q: %w", method.Name, className, err)
			}
			retInfo = &gdc.PropertyInfo{
				ClassName:  namePtr.AsPtr(),
				Name:       gdapi.PStringNameFromStr(out.Name()).AsPtr(),
				Type:       varType,
				Hint:       uint(gdapi.PropertyHintNone),
				HintString: gdapi.StringFromStr("").AsPtr(),
				Usage:      uint(gdapi.PropertyUsageDefault),
			}
		}

		methods = append(methods, classMethod{
			name:            method.Name,
			fn:              method.Func.Interface(),
			pinner:          pinner,
			saveArgInfo:     argInfo,
			saveArgMetadata: argMetadata,
			method: gdc.ClassMethodInfo{
				Name:                name.AsPtr(),
				MethodUserdata:      *(*unsafe.Pointer)(unsafe.Pointer(utils.ToPointer(len(methods)))),
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
		})
	}

	return &classEntry{
		name:          className,
		constructor:   ctr,
		properties:    properties,
		methods:       methods,
		parentNamePtr: parentPtr,
		namePtr:       namePtr,
		instances:     make(map[uint64]*classInstance),
	}, nil
}

func (me *extension) registerClass(class *classEntry) {
	me.Logf(LogLevelDebug, "registering class %q", class.name)
	me.iface.ClassdbRegisterExtensionClass2(me.pLibrary, class.namePtr.AsCPtr(), class.parentNamePtr.AsCPtr(), &gdc.ClassCreationInfo2{
		ClassUserdata:           store(class),
		IsExposed:               gdc.Bool(1),
		CreateInstanceFunc:      gdc.Callbacks.GetClassCreationInfo2CreateInstanceFuncCallback(),
		FreeInstanceFunc:        gdc.Callbacks.GetClassCreationInfo2FreeInstanceFuncCallback(),
		SetFunc:                 gdc.Callbacks.GetClassCreationInfo2SetFuncCallback(),
		GetFunc:                 gdc.Callbacks.GetClassCreationInfo2GetFuncCallback(),
		GetPropertyListFunc:     gdc.Callbacks.GetClassCreationInfo2GetPropertyListFuncCallback(),
		FreePropertyListFunc:    gdc.Callbacks.GetClassCreationInfo2FreePropertyListFuncCallback(),
		PropertyCanRevertFunc:   gdc.Callbacks.GetClassCreationInfo2PropertyCanRevertFuncCallback(),
		PropertyGetRevertFunc:   gdc.Callbacks.GetClassCreationInfo2PropertyGetRevertFuncCallback(),
		ToStringFunc:            gdc.Callbacks.GetClassCreationInfo2ToStringFuncCallback(),
		ReferenceFunc:           gdc.Callbacks.GetClassCreationInfo2ReferenceFuncCallback(),
		UnreferenceFunc:         gdc.Callbacks.GetClassCreationInfo2UnreferenceFuncCallback(),
		RecreateInstanceFunc:    gdc.Callbacks.GetClassCreationInfo2RecreateInstanceFuncCallback(),
		GetVirtualFunc:          gdc.Callbacks.GetClassCreationInfo2GetVirtualFuncCallback(),
		GetVirtualCallDataFunc:  gdc.Callbacks.GetClassCreationInfo2GetVirtualCallDataFuncCallback(),
		CallVirtualWithDataFunc: gdc.Callbacks.GetClassCreationInfo2CallVirtualWithDataFuncCallback(),
		GetRidFunc:              gdc.Callbacks.GetClassCreationInfo2GetRidFuncCallback(),
	})

	for _, method := range class.methods {
		me.Logf(LogLevelDebug, "registering method %q of class %q", method.name, class.name)
		mi := method.method
		rawMI, clean := mi.ToRaw()
		defer clean()
		me.iface.ClassdbRegisterExtensionClassMethod(me.pLibrary, class.namePtr.AsCPtr(), (*gdc.ClassMethodInfo)(unsafe.Pointer(&rawMI)))
	}

	for _, property := range class.properties {
		me.Logf(LogLevelDebug, "registering property %q of class %q", property.name, class.name)
		me.iface.ClassdbRegisterExtensionClassProperty(me.pLibrary, class.namePtr.AsCPtr(), &property.property, property.setter.AsCPtr(), property.getter.AsCPtr())
	}

	// yes!
	// TODO: register signals

	// not sure we need yet
	// TODO: register sub properties through tags

	// oof, hard
	// TODO: register constants through tags
	// TODO: register enums through tags
}

func (me *extension) createClass(class *classEntry) gdc.ObjectPtr {
	obj := me.iface.ClassdbConstructObject(class.parentNamePtr.AsCPtr())
	id := me.iface.ObjectGetInstanceId(gdc.ConstObjectPtr(obj))
	instance := &classInstance{
		class:    class,
		instance: class.constructor(),
		id:       id,
	}
	me.iface.ObjectSetInstance(obj, class.namePtr.AsCPtr(), gdc.ClassInstancePtr(store(instance)))
	instance.instance.SetBaseObject(obj)

	class.mutex.Lock()
	defer class.mutex.Unlock()
	class.instances[id] = instance
	return obj
}

func (me *extension) freeClass(instance *classInstance) {
	destroyable, is := instance.instance.(Destroyable)
	if is {
		destroyable.Destroy()
	}
	instance.class.mutex.Lock()
	defer instance.class.mutex.Unlock()
	delete(instance.class.instances, instance.id)
}

func (me *extension) unregisterClass(class *classEntry) {
	me.Logf(LogLevelDebug, "unregistering class %q", class.name)
	class.namePtr.Destroy()
	class.parentNamePtr.Destroy()
	me.iface.ClassdbUnregisterExtensionClass(me.pLibrary, class.namePtr.AsCPtr())
	for i, method := range class.methods {
		me.Logf(LogLevelDebug, "unregistering method %q of class %q", method.name, class.name)
		namePtr := gdapi.StringNameFromPtr(method.method.Name)
		namePtr.Destroy()
		if method.fn != nil && i%2 == 0 { // is getter
			gdc.CFreePropertyInfo(method.method.ReturnValueInfo)
		}
	}
}
