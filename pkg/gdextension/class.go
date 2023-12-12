package gdextension

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
	"github.com/LouisBrunner/go-dot-extension/pkg/utils"
	"github.com/iancoleman/strcase"
)

func (me *extension) addClassCallbacks() {
	gdc.Callbacks.SetClassCreationInfoCreateInstanceFuncHandler(func(pUserdata unsafe.Pointer) gdc.ObjectPtr {
		class, err := restore[*classEntry](pUserdata)
		if err != nil {
			me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
			return nil
		}
		return me.createClass(class)
	})
	gdc.Callbacks.SetClassCreationInfoFreeInstanceFuncHandler(func(pUserdata unsafe.Pointer, pInstance gdc.ClassInstancePtr) {
		instance, err := restore[*classInstance](unsafe.Pointer(pInstance))
		if err != nil {
			me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
			return
		}
		me.freeClass(instance)
	})
	gdc.Callbacks.SetClassCreationInfoSetFuncHandler(func(pInstance gdc.ClassInstancePtr, pName gdc.ConstStringNamePtr, pValue gdc.ConstVariantPtr) gdc.Bool {
		me.Logf(LogLevelDebug, "SetClassCreationInfoSetFuncHandler: %s=%s", gdapi.NewVariantWithC(gdc.ConstVariantPtr(pName)), gdapi.NewVariantWithC(pValue))
		// TODO: implement
		return gdc.Bool(0)
	})
	gdc.Callbacks.SetClassCreationInfoGetFuncHandler(func(pInstance gdc.ClassInstancePtr, pName gdc.ConstStringNamePtr, rRet gdc.VariantPtr) gdc.Bool {
		me.Logf(LogLevelDebug, "SetClassCreationInfoGetFuncHandler: %s", gdapi.NewVariantWithC(gdc.ConstVariantPtr(pName)))
		// TODO: implement
		return gdc.Bool(0)
	})
	gdc.Callbacks.SetClassCreationInfoGetPropertyListFuncHandler(func(pInstance gdc.ClassInstancePtr, rCount *uint) *gdc.PropertyInfo {
		instance, err := restore[*classInstance](unsafe.Pointer(pInstance))
		if err != nil {
			me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
			return nil
		}
		properties := instance.class.makeProperties()
		*rCount = uint(len(properties))
		return unsafe.SliceData(properties)
	})
	gdc.Callbacks.SetClassCreationInfoFreePropertyListFuncHandler(func(pInstance gdc.ClassInstancePtr, pList *gdc.PropertyInfo) {
		if pList == nil {
			return
		}
		gdc.CFreePropertyInfo(pList)
	})
	gdc.Callbacks.SetClassCreationInfoPropertyCanRevertFuncHandler(func(pInstance gdc.ClassInstancePtr, pName gdc.ConstStringNamePtr) gdc.Bool {
		me.Logf(LogLevelDebug, "SetClassCreationInfoPropertyCanRevertFuncHandler: %s", gdapi.NewVariantWithC(gdc.ConstVariantPtr(pName)))
		// TODO: implement
		return gdc.Bool(0)
	})
	gdc.Callbacks.SetClassCreationInfoPropertyGetRevertFuncHandler(func(pInstance gdc.ClassInstancePtr, pName gdc.ConstStringNamePtr, rRet gdc.VariantPtr) gdc.Bool {
		me.Logf(LogLevelDebug, "SetClassCreationInfoPropertyGetRevertFuncHandler: %s", gdapi.NewVariantWithC(gdc.ConstVariantPtr(pName)))
		// TODO: implement
		return gdc.Bool(0)
	})
	gdc.Callbacks.SetClassCreationInfoToStringFuncHandler(func(pInstance gdc.ClassInstancePtr, rIsValid *uint8, pOut gdc.StringPtr) {
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
	name   string
	method gdc.ClassMethodInfo
	fn     interface{}
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
			return nil, fmt.Errorf("property %q cannot be exported as it is neither a Godot class or variant", name)
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

	name := typ.Elem().Name()
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
			ftyp = reflect.PtrTo(ftyp)
			for m := 0; m < ftyp.NumMethod(); m++ {
				method := ftyp.Method(m)
				if method.IsExported() {
					methodsFromParent[method.Name] = struct{}{}
				}
			}
		}

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
			return nil, fmt.Errorf("class %q: %w", name, err)
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
		if strings.HasPrefix(methodName, "x_") {
			methodName = strings.TrimPrefix(methodName, "x")
		}
		name := gdapi.StringNameFromStr(methodName)
		methods = append(methods, classMethod{
			name: method.Name,
			method: gdc.ClassMethodInfo{
				Name:           name.AsPtr(),
				MethodUserdata: *(*unsafe.Pointer)(unsafe.Pointer(utils.ToPointer[int](len(methods)))),
				// TODO: missing
			},
		})
	}

	namePtr := gdapi.StringNameFromStr(name)
	parentPtr := gdapi.StringNameFromStr(instance.BaseClass())

	return &classEntry{
		name:          name,
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
	me.iface.ClassdbRegisterExtensionClass(me.pLibrary, class.namePtr.AsCPtr(), class.parentNamePtr.AsCPtr(), &gdc.ClassCreationInfo{
		ClassUserdata:         store(class),
		CreateInstanceFunc:    gdc.Callbacks.GetClassCreationInfoCreateInstanceFuncCallback(),
		FreeInstanceFunc:      gdc.Callbacks.GetClassCreationInfoFreeInstanceFuncCallback(),
		SetFunc:               gdc.Callbacks.GetClassCreationInfoSetFuncCallback(),
		GetFunc:               gdc.Callbacks.GetClassCreationInfoGetFuncCallback(),
		GetPropertyListFunc:   gdc.Callbacks.GetClassCreationInfoGetPropertyListFuncCallback(),
		FreePropertyListFunc:  gdc.Callbacks.GetClassCreationInfoFreePropertyListFuncCallback(),
		PropertyCanRevertFunc: gdc.Callbacks.GetClassCreationInfoPropertyCanRevertFuncCallback(),
		PropertyGetRevertFunc: gdc.Callbacks.GetClassCreationInfoPropertyGetRevertFuncCallback(),
		ToStringFunc:          gdc.Callbacks.GetClassCreationInfoToStringFuncCallback(),
	})

	for _, method := range class.methods {
		me.Logf(LogLevelDebug, "registering method %q of class %q", method.name, class.name)
		cmethod := gdc.CNewClassMethodInfo()
		defer gdc.CFreeClassMethodInfo(cmethod)
		*cmethod = method.method
		me.iface.ClassdbRegisterExtensionClassMethod(me.pLibrary, class.namePtr.AsCPtr(), cmethod)
	}

	for _, property := range class.properties {
		me.Logf(LogLevelDebug, "registering property %q of class %q", property.name, class.name)
		me.iface.ClassdbRegisterExtensionClassProperty(me.pLibrary, class.namePtr.AsCPtr(), &property.property, property.setter.AsCPtr(), property.getter.AsCPtr())
	}

	// TODO: register signals
	// TODO: register constants through tags
	// TODO: register enums through tags
	// TODO: register sub properties through tags
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
