package gdextension

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

	for _, property := range class.properties {
		if property.objClassNamePtr == nil {
			continue
		}
		prop := me.iface.ClassdbConstructObject(property.objClassNamePtr.AsCPtr())
		// FIXME: better error checking
		strct := reflect.ValueOf(instance.instance).Elem()
		strct.FieldByName(property.name).Set(
			reflect.ValueOf(gdapi.ObjectFromPtr(prop)), // TODO: wrong type most likely
		)
	}
	for _, signal := range class.signals {
		// FIXME: better error checking
		strct := reflect.ValueOf(instance.instance).Elem()
		strct.FieldByName(signal.name).Set(
			reflect.ValueOf(gdapi.NewSignalFromObjectStringName(gdapi.ObjectFromPtr(obj), signal.namePtr)),
		)
	}
	subClass := gdapi.StringNameFromStr("SignalSubscribers")
	defer subClass.Destroy()
	for _, subscriber := range class.subscribers {
		// FIXME: better error checking
		subs := me.iface.ClassdbConstructObject(subClass.AsCPtr())
		strct := reflect.ValueOf(instance.instance).Elem()
		fld := strct.FieldByName(subscriber)
		rf := reflect.NewAt(fld.Type(), unsafe.Pointer(fld.UnsafeAddr())).Elem()
		rf.Set(
			reflect.ValueOf(gdapi.NewSignalSubscribers(subs)),
		)
	}

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
		me.iface.ClassdbRegisterExtensionClassProperty(me.pLibrary, class.namePtr.AsCPtr(), property.property, property.setter.AsCPtr(), property.getter.AsCPtr())
	}

	for _, signal := range class.signals {
		me.Logf(LogLevelDebug, "registering signal %q of class %q", signal.name, class.name)
		me.iface.ClassdbRegisterExtensionClassSignal(me.pLibrary, class.namePtr.AsCPtr(), signal.namePtr.AsCPtr(), signal.argsPtr, gdc.Int(signal.argCount))
	}

	// not sure we need yet
	// TODO: register sub properties through tags
	// oof, hard
	// TODO: register constants through tags
	// TODO: register enums through tags
}

func (me *extension) unregisterClass(class *classEntry) {
	me.Logf(LogLevelDebug, "unregistering class %q", class.name)
	class.namePtr.Destroy()
	class.parentNamePtr.Destroy()
	me.iface.ClassdbUnregisterExtensionClass(me.pLibrary, class.namePtr.AsCPtr())
	for _, method := range class.methods {
		me.Logf(LogLevelDebug, "unregistering method %q of class %q", method.name, class.name)
		namePtr := gdapi.StringNameFromPtr(method.method.Name)
		namePtr.Destroy()
		method.pinner.Unpin()
		for _, arg := range method.saveArgInfo {
			hintPtr := gdapi.StringFromPtr(arg.HintString)
			hintPtr.Destroy()
			namePtr := gdapi.StringNameFromPtr(arg.Name)
			namePtr.Destroy()
		}
		if method.method.ReturnValueInfo != nil {
			hintPtr := gdapi.StringFromPtr(method.method.ReturnValueInfo.HintString)
			hintPtr.Destroy()
			namePtr := gdapi.StringNameFromPtr(method.method.ReturnValueInfo.Name)
			namePtr.Destroy()
		}
	}
	for _, property := range class.properties {
		me.Logf(LogLevelDebug, "unregistering property %q of class %q", property.name, class.name)
		if property.objClassNamePtr == nil {
			continue
		}
		property.objClassNamePtr.Destroy()
	}
	for _, signal := range class.signals {
		me.Logf(LogLevelDebug, "unregistering signal %q of class %q", signal.name, class.name)
		signal.namePtr.Destroy()
	}
}

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
