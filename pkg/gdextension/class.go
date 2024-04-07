package gdextension

import (
	"fmt"
	"reflect"

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

func (me *extension) createClass(class *classEntry) (fobj gdc.ObjectPtr) {
	obj := me.iface.ClassdbConstructObject(class.parentNamePtr.AsCPtr())
	defer func() {
		if fobj == nil {
			me.iface.ObjectDestroy(obj)
		}
	}()

	id := me.iface.ObjectGetInstanceId(gdc.ConstObjectPtr(obj))
	instance := &classInstance{
		class:    class,
		instance: class.constructor(),
		id:       id,
	}
	me.iface.ObjectSetInstance(obj, class.namePtr.AsCPtr(), gdc.ClassInstancePtr(store(instance)))
	instance.instance.SetBaseObject(obj)

	for _, property := range class.properties {
		err := property.create(me, instance.instance)
		if err != nil {
			me.Logf(LogLevelError, "error: %s", err.Error())
			return nil
		}
	}
	for _, signal := range class.signals {
		signal.create(me, instance.instance, obj)
	}
	for _, subscriber := range class.subscribers {
		err := subscriber.create(me, instance.instance)
		if err != nil {
			me.Logf(LogLevelError, "error: %s", err.Error())
			return nil
		}
	}

	initializable, is := instance.instance.(Initializable)
	if is {
		initializable.Init()
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

	for _, constant := range class.constants {
		me.Logf(LogLevelDebug, "registering constant %s of class %q", constant.getName(), class.name)
		constant.register(me, class)
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
	for _, constant := range class.constants {
		me.Logf(LogLevelDebug, "unregistering constant %s of class %q", constant.getName(), class.name)
		constant.unregister(me, class)
	}

	me.iface.ClassdbUnregisterExtensionClass(me.pLibrary, class.namePtr.AsCPtr())
	class.namePtr.Destroy()
	class.parentNamePtr.Destroy()
}

func nameForClass(typ reflect.Type, instance any) (string, error) {
	switch typ.Kind() {
	case reflect.Pointer:
		if typ.Elem().Kind() == reflect.Struct {
			return typ.Elem().Name(), nil
		}
	case reflect.Struct:
		return typ.Name(), nil
	}
	return "", fmt.Errorf("expected a struct or a pointer to struct but got %s (%T)", typ.Kind(), instance)
}

func (me *extension) CreateClass(typ reflect.Type) (any, error) {
	name, err := nameForClass(typ, reflect.New(typ).Elem().Interface())
	if err != nil {
		return nil, err
	}

	entry, found := me.registered[name]
	if !found {
		return nil, fmt.Errorf("class %q is not registered", name)
	}

	obj := me.iface.ClassdbConstructObject(entry.namePtr.AsCPtr())
	id := me.iface.ObjectGetInstanceId(gdc.ConstObjectPtr(obj))

	inst, err := entry.lookupInstance(id)
	if err != nil {
		return nil, err
	}
	return inst.instance, nil
}

func CreateClass[T Class](me Extension) (*T, error) {
	typ := reflect.TypeFor[T]()

	instance, err := me.CreateClass(typ)
	if err != nil {
		return nil, err
	}

	actualInstance, cast := instance.(T)
	if !cast {
		return nil, fmt.Errorf("could not cast %T to %T", instance, actualInstance)
	}
	return &actualInstance, nil
}
