package gdextension

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
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
		properties := me.makeProperties(instance.class)
		*rCount = uint(len(properties))
		return unsafe.SliceData(properties)
	})
	gdc.Callbacks.SetClassCreationInfoFreePropertyListFuncHandler(func(pInstance gdc.ClassInstancePtr, pList *gdc.PropertyInfo) {
		if pList == nil {
			return
		}
		instance, err := restore[*classInstance](unsafe.Pointer(pInstance))
		if err != nil {
			me.Logf(LogLevelError, "could not restore class entry: %s", err.Error())
			return
		}
		for _, entry := range unsafe.Slice(pList, len(instance.class.properties)) {
			ptr := gdapi.StringNameFromPtr(entry.Name)
			ptr.Destroy()
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
}

type classEntry struct {
	name        string
	constructor ClassConstructor
	properties  []string

	parentNamePtr gdapi.StringName
	namePtr       gdapi.StringName

	mutex     sync.Mutex
	instances map[uint64]*classInstance
}

func (me *extension) makeProperties(class *classEntry) []gdc.PropertyInfo {
	properties := gdc.CNewPropertyInfoArray(len(class.properties))
	for i, name := range class.properties {
		name := gdapi.StringNameFromStr(name)
		properties[i] = gdc.PropertyInfo{
			Name: name.AsPtr(),
			Type: gdc.VariantTypeBool, // TODO: wrong
			// TODO: missing loads of fields
		}
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

func (me *extension) prepareClass(ctr ClassConstructor) (*classEntry, error) {
	instance := ctr()
	typ := reflect.TypeOf(instance)
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct but got %s", typ.Kind())
	}

	name := typ.Name()
	properties := make([]string, 0, typ.NumField())
	for _, field := range reflect.VisibleFields(typ) {
		if !field.IsExported() {
			continue
		}
		properties = append(properties, field.Name)
	}

	namePtr := gdapi.StringNameFromStr(name)
	parentPtr := gdapi.StringNameFromStr(instance.BaseClass())

	return &classEntry{
		name:          name,
		constructor:   ctr,
		properties:    properties,
		parentNamePtr: parentPtr,
		namePtr:       namePtr,
		instances:     make(map[uint64]*classInstance),
	}, nil
}

func (me *extension) registerClass(class *classEntry) {
	me.iface.ClassdbRegisterExtensionClass(me.pLibrary, class.namePtr.AsCPtr(), class.parentNamePtr.AsCPtr(), &gdc.ClassCreationInfo{
		ClassUserdata:      store(class),
		CreateInstanceFunc: gdc.Callbacks.GetClassCreationInfoCreateInstanceFuncCallback(),
		FreeInstanceFunc:   gdc.Callbacks.GetClassCreationInfoFreeInstanceFuncCallback(),
		SetFunc:            gdc.Callbacks.GetClassCreationInfoSetFuncCallback(),
		GetFunc:            gdc.Callbacks.GetClassCreationInfoGetFuncCallback(),
		// TODO: broken
		// GetPropertyListFunc:   gdc.Callbacks.GetClassCreationInfoGetPropertyListFuncCallback(),
		FreePropertyListFunc:  gdc.Callbacks.GetClassCreationInfoFreePropertyListFuncCallback(),
		PropertyCanRevertFunc: gdc.Callbacks.GetClassCreationInfoPropertyCanRevertFuncCallback(),
		PropertyGetRevertFunc: gdc.Callbacks.GetClassCreationInfoPropertyGetRevertFuncCallback(),
		ToStringFunc:          gdc.Callbacks.GetClassCreationInfoToStringFuncCallback(),
	})
	// TODO: register methods
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

	// TODO: call constructor?
	class.mutex.Lock()
	defer class.mutex.Unlock()
	class.instances[id] = instance
	return obj
}

func (me *extension) freeClass(instance *classInstance) {
	// TODO: call destructor?
	instance.class.mutex.Lock()
	defer instance.class.mutex.Unlock()
	delete(instance.class.instances, instance.id)
}

func (me *extension) unregisterClass(class *classEntry) {
	defer class.namePtr.Destroy()
	defer class.parentNamePtr.Destroy()
	me.iface.ClassdbUnregisterExtensionClass(me.pLibrary, class.namePtr.AsCPtr())
}
