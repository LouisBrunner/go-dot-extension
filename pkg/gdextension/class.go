package gdextension

import (
	"fmt"

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
		property.create(me, instance.instance)
	}
	for _, signal := range class.signals {
		signal.create(me, instance.instance, obj)
	}
	for _, subscriber := range class.subscribers {
		subscriber.create(me, instance.instance)
	}

	class.mutex.Lock()
	defer class.mutex.Unlock()
	class.instances[id] = instance
	return obj
}

func (me *extension) freeClass(instance *classInstance) {
	class := instance.class

	for _, property := range class.properties {
		property.free(me, instance.instance)
	}
	for _, signal := range class.signals {
		signal.free(me, instance.instance)
	}
	for _, subscriber := range class.subscribers {
		subscriber.free(me, instance.instance)
	}

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
		ClassUserdata:         store(class),
		IsExposed:             gdc.Bool(1),
		CreateInstanceFunc:    gdc.Callbacks.GetClassCreationInfo2CreateInstanceFuncCallback(),
		FreeInstanceFunc:      gdc.Callbacks.GetClassCreationInfo2FreeInstanceFuncCallback(),
		SetFunc:               gdc.Callbacks.GetClassCreationInfo2SetFuncCallback(),
		GetFunc:               gdc.Callbacks.GetClassCreationInfo2GetFuncCallback(),
		GetPropertyListFunc:   gdc.Callbacks.GetClassCreationInfo2GetPropertyListFuncCallback(),
		FreePropertyListFunc:  gdc.Callbacks.GetClassCreationInfo2FreePropertyListFuncCallback(),
		PropertyCanRevertFunc: gdc.Callbacks.GetClassCreationInfo2PropertyCanRevertFuncCallback(),
		PropertyGetRevertFunc: gdc.Callbacks.GetClassCreationInfo2PropertyGetRevertFuncCallback(),
		ToStringFunc:          gdc.Callbacks.GetClassCreationInfo2ToStringFuncCallback(),
	})

	for _, method := range class.methods {
		me.Logf(LogLevelDebug, "registering method %q of class %q", method.name, class.name)
		method.register(me, class)
	}

	for _, property := range class.properties {
		me.Logf(LogLevelDebug, "registering property %q of class %q", property.name, class.name)
		property.register(me, class)
	}

	for _, signal := range class.signals {
		me.Logf(LogLevelDebug, "registering signal %q of class %q", signal.name, class.name)
		signal.register(me, class)
	}
}

func (me *extension) unregisterClass(class *classEntry) {
	me.Logf(LogLevelDebug, "unregistering class %q", class.name)

	for _, method := range class.methods {
		me.Logf(LogLevelDebug, "unregistering method %q of class %q", method.name, class.name)
		method.unregister(me, class)
	}
	for _, property := range class.properties {
		me.Logf(LogLevelDebug, "unregistering property %q of class %q", property.name, class.name)
		property.unregister(me, class)
	}
	for _, signal := range class.signals {
		me.Logf(LogLevelDebug, "unregistering signal %q of class %q", signal.name, class.name)
		signal.unregister(me, class)
	}

	me.iface.ClassdbUnregisterExtensionClass(me.pLibrary, class.namePtr.AsCPtr())
	class.namePtr.Destroy()
	class.parentNamePtr.Destroy()
}
