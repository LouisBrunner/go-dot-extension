package gdextension

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

func (me *extension) addClassCallbacks() {
	gdc.Callbacks.SetClassCreationInfo2CreateInstanceFuncHandler(me.createInstance)
	gdc.Callbacks.SetClassCreationInfo2FreeInstanceFuncHandler(me.freeInstance)
	gdc.Callbacks.SetClassCreationInfo2SetFuncHandler(me.setProperty)
	gdc.Callbacks.SetClassCreationInfo2GetFuncHandler(me.getProperty)
	gdc.Callbacks.SetClassCreationInfo2GetPropertyListFuncHandler(me.getPropertyList)
	gdc.Callbacks.SetClassCreationInfo2FreePropertyListFuncHandler(me.freePropertyList)
	gdc.Callbacks.SetClassCreationInfo2PropertyCanRevertFuncHandler(me.propertyCanRevert)
	gdc.Callbacks.SetClassCreationInfo2PropertyGetRevertFuncHandler(me.propertyRevert)
	gdc.Callbacks.SetClassCreationInfo2ToStringFuncHandler(me.classToString)

	gdc.Callbacks.SetClassMethodInfoCallFuncHandler(me.methodCall)
	gdc.Callbacks.SetClassMethodInfoPtrcallFuncHandler(me.methodPtrcall)
}

type argsGetter func(argTypes []reflect.Type) ([]reflect.Value, error)

func (me *extension) callHelper(methodUserdata unsafe.Pointer, pInstance gdc.ClassInstancePtr, agetter argsGetter) (*classMethod, []reflect.Value, error) {
	instance, err := restore[*classInstance](unsafe.Pointer(pInstance))
	if err != nil {
		return nil, nil, fmt.Errorf("could not restore class entry: %s", err.Error())
	}
	methodID := (uint64)(uintptr(methodUserdata))
	if methodID > uint64(len(instance.class.methods)) {
		return nil, nil, fmt.Errorf("could not find method with ID %d", methodID)
	}
	method := instance.class.methods[methodID]
	me.Logf(LogLevelTrace, "calling: %v", method.name)
	args, err := agetter(method.argTypes)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get arguments for %s: %s", method.name, err.Error())
	}
	me.Logf(LogLevelTrace, "args: %v", args)
	res := reflectCallMethod(instance.instance, &method, args)
	return &method, res, nil
}

func (me *extension) createInstance(pUserdata unsafe.Pointer) gdc.ObjectPtr {
	class, err := restore[*classEntry](pUserdata)
	if err != nil {
		me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
		return nil
	}
	return me.createClass(class)
}

func (me *extension) freeInstance(pUserdata unsafe.Pointer, pInstance gdc.ClassInstancePtr) {
	instance, err := restore[*classInstance](unsafe.Pointer(pInstance))
	if err != nil {
		me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
		return
	}
	me.freeClass(instance)
}

func (me *extension) setProperty(pInstance gdc.ClassInstancePtr, pName gdc.ConstStringNamePtr, pValue gdc.ConstVariantPtr) gdc.Bool {
	instance, err := restore[*classInstance](unsafe.Pointer(pInstance))
	if err != nil {
		me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
		return gdc.Bool(0)
	}

	name := gdapi.StringNameFromCPtr(pName).AsVariant().String()
	me.Logf(LogLevelTrace, "setting: %s::%s to %s", instance.class.name, name, gdapi.NewVariantWithC(pValue))

	return gdc.Bool(0)
}

func (me *extension) getProperty(pInstance gdc.ClassInstancePtr, pName gdc.ConstStringNamePtr, rRet gdc.VariantPtr) gdc.Bool {
	instance, err := restore[*classInstance](unsafe.Pointer(pInstance))
	if err != nil {
		me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
		return gdc.Bool(0)
	}

	sname := gdapi.StringNameFromCPtr(pName)
	name := sname.AsVariant().String()
	me.Logf(LogLevelTrace, "getting: %s::%s", instance.class.name, name)

	return gdc.Bool(0)
}

func (me *extension) getPropertyList(pInstance gdc.ClassInstancePtr, rCount *uint) *gdc.PropertyInfo {
	*rCount = 0
	return nil
}

func (me *extension) freePropertyList(pInstance gdc.ClassInstancePtr, pList *gdc.PropertyInfo) {
}

func (me *extension) propertyCanRevert(pInstance gdc.ClassInstancePtr, pName gdc.ConstStringNamePtr) gdc.Bool {
	instance, err := restore[*classInstance](unsafe.Pointer(pInstance))
	if err != nil {
		me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
		return gdc.Bool(0)
	}

	name := gdapi.StringNameFromCPtr(pName).AsVariant().String()
	me.Logf(LogLevelTrace, "can revert: %s::%s", instance.class.name, name)

	return gdc.Bool(1)
}

func (me *extension) propertyRevert(pInstance gdc.ClassInstancePtr, pName gdc.ConstStringNamePtr, rRet gdc.VariantPtr) gdc.Bool {
	instance, err := restore[*classInstance](unsafe.Pointer(pInstance))
	if err != nil {
		me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
		return gdc.Bool(0)
	}

	name := gdapi.StringNameFromCPtr(pName).AsVariant().String()
	me.Logf(LogLevelTrace, "reverting: %s::%s", instance.class.name, name)

	prop, ok := instance.class.properties[name]
	if !ok {
		me.Logf(LogLevelDebug, "could not find property %s::%s", instance.class.name, name)
		return gdc.Bool(0)
	}

	va, err := variantFromReflect(reflect.ValueOf(instance.class.constructor()).Elem().FieldByName(prop.name))
	if err != nil {
		me.Logf(LogLevelError, "could not convert return value for %s::%s: %s", instance.class.name, name, err.Error())
		return gdc.Bool(0)
	}

	defer va.Destroy()
	me.iface.VariantNewCopy(gdc.UninitializedVariantPtr(rRet), va.AsCPtr())
	return gdc.Bool(1)
}

func (me *extension) classToString(pInstance gdc.ClassInstancePtr, rIsValid *uint8, pOut gdc.StringPtr) {
	instance, err := restore[*classInstance](unsafe.Pointer(pInstance))
	if err != nil {
		me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
		*rIsValid = 0
		return
	}
	me.iface.StringOperatorPlusEqCstr(pOut, instance.String())
	*rIsValid = 1
}

func (me *extension) methodCall(methodUserdata unsafe.Pointer, pInstance gdc.ClassInstancePtr, pArgs *gdc.ConstVariantPtr, pArgumentCount gdc.Int, rReturn gdc.VariantPtr, rError *gdc.CallError) {
	method, res, err := me.callHelper(methodUserdata, pInstance, func(argTypes []reflect.Type) ([]reflect.Value, error) {
		if int(pArgumentCount) != len(argTypes) {
			return nil, fmt.Errorf("expected %d arguments, got %d", len(argTypes), pArgumentCount)
		}
		received := unsafe.Slice(pArgs, int(pArgumentCount))
		args := make([]reflect.Value, pArgumentCount)
		for i, pArg := range received {
			val, err := nativeFromVariant(pArg, argTypes[i])
			if err != nil {
				return nil, err
			}
			args[i] = reflect.ValueOf(val)
		}
		return args, nil
	})
	if err != nil {
		me.Logf(LogLevelError, "could not call method: %s", err.Error())
		return
	}
	if method.method.ReturnValueInfo == nil {
		return
	}

	ret := res[0]
	if retRef, is := ret.Interface().(reflect.Value); is {
		ret = retRef
	}
	me.Logf(LogLevelTrace, "return: %v", ret)
	va, err := variantFromReflect(ret)
	if err != nil {
		me.Logf(LogLevelError, "could not convert return value for %s: %s", method.name, err.Error())
		return
	}
	defer va.Destroy()
	me.iface.VariantNewCopy(gdc.UninitializedVariantPtr(rReturn), va.AsCPtr())
}

func (me *extension) methodPtrcall(methodUserdata unsafe.Pointer, pInstance gdc.ClassInstancePtr, pArgs *gdc.ConstTypePtr, rRet gdc.TypePtr) {
	method, res, err := me.callHelper(methodUserdata, pInstance, func(argTypes []reflect.Type) ([]reflect.Value, error) {
		received := unsafe.Slice(pArgs, len(argTypes))
		args := make([]reflect.Value, len(argTypes))
		for i, pArg := range received {
			val, err := me.reflectFromType(pArg, argTypes[i])
			if err != nil {
				return nil, err
			}
			args[i] = val
		}
		return args, nil
	})
	if err != nil {
		me.Logf(LogLevelError, "could not ptrcall method: %s", err.Error())
		return
	}
	if method.method.ReturnValueInfo == nil {
		return
	}

	ret := res[0]
	if retRef, is := ret.Interface().(reflect.Value); is {
		ret = retRef
	}
	me.Logf(LogLevelTrace, "return: %v", ret)
	typ, err := typeFromReflect(ret)
	if err != nil {
		me.Logf(LogLevelError, "could not convert return value for ptrcall of %s: %s", method.name, err.Error())
		return
	}
	*(*gdc.TypePtr)(rRet) = typ
}
