package gdextension

import (
	"fmt"
	"reflect"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type classEntry struct {
	name           string
	parentNamePtr  gdc.StringNamePtr
	parentNameFree func()
	namePtr        gdc.StringNamePtr
	nameFree       func()
	class          Class
	constructor    ClassConstructor
	instances      map[uint64]Class
}

func (me *extension) prepareClass(ctr ClassConstructor) (*classEntry, error) {
	instance := ctr()
	typ := reflect.TypeOf(instance)
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected %T to be a struct but got %s", instance, typ.Kind())
	}

	name := typ.Name()

	namePtr, clean := me.makeStringName(name)
	parentPtr, cleanParent := me.makeStringName(instance.BaseClass())

	return &classEntry{
		name:           name,
		parentNamePtr:  parentPtr,
		parentNameFree: cleanParent,
		namePtr:        namePtr,
		nameFree:       clean,
		class:          instance,
		constructor:    ctr,
		instances:      make(map[uint64]Class),
	}, nil
}

func (me *extension) registerClass(class *classEntry) {
	me.iface.ClassdbRegisterExtensionClass(me.pLibrary, gdc.ConstStringNamePtr(class.namePtr), gdc.ConstStringNamePtr(class.parentNamePtr), &gdc.ClassCreationInfo{
		ClassUserdata:      store(class),
		CreateInstanceFunc: gdc.Callbacks.GetClassCreationInfoCreateInstanceFuncCallback(),
		FreeInstanceFunc:   gdc.Callbacks.GetClassCreationInfoFreeInstanceFuncCallback(),
	})
}

func (me *extension) createClass(class *classEntry) gdc.ObjectPtr {
	obj := me.iface.ClassdbConstructObject(gdc.ConstStringNamePtr(class.parentNamePtr))
	id := me.iface.ObjectGetInstanceId(gdc.ConstObjectPtr(obj))
	me.iface.ObjectSetInstance(obj, gdc.ConstStringNamePtr(class.namePtr), gdc.ClassInstancePtr(&id))
	// TODO: mutex
	class.instances[id] = class.constructor()
	// TODO: call constructor?
	return obj
}

func (me *extension) freeClass(class *classEntry, pInstance gdc.ClassInstancePtr) {
	// TODO: call destructor?
	id := *(*uint64)(pInstance)
	// TODO: mutex
	delete(class.instances, id)
}

func (me *extension) unregisterClass(class *classEntry) {
	defer func() {
		class.parentNameFree()
		class.nameFree()
	}()
	me.iface.ClassdbUnregisterExtensionClass(me.pLibrary, gdc.ConstStringNamePtr(class.namePtr))
}
