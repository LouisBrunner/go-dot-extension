package gdextension

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type classProperty struct {
	name     string
	gdName   string
	typ      reflect.Type
	property *gdc.PropertyInfo
	getter   gdapi.StringName
	setter   gdapi.StringName

	isBClass        bool
	objClassNamePtr gdc.ConstStringNamePtr
}

func (me *classProperty) register(ext *extension, class *classEntry) {
	ext.iface.ClassdbRegisterExtensionClassProperty(ext.pLibrary, class.namePtr.AsCPtr(), me.property, me.setter.AsCPtr(), me.getter.AsCPtr())
}

func (me *classProperty) create(ext *extension, instance Class) error {
	switch {
	case me.objClassNamePtr != nil:
		prop := ext.iface.ClassdbConstructObject(me.objClassNamePtr)
		obj, err := ext.objectFromPtr(me.typ, prop)
		if err != nil {
			return err
		}
		reflectSetProperty(instance, me, obj)
	case me.isBClass:
		bclass, err := gdapi.BClassForType(me.property.Type)
		if err != nil {
			return err
		}
		reflectSetProperty(instance, me, reflect.ValueOf(bclass).Elem().Interface())
	}
	return nil
}

func (me *classProperty) free(ext *extension, instance Class) {
	if me.objClassNamePtr == nil {
		return
	}
	val := reflectGetProperty(instance, me).Interface()
	obj, is := val.(gdapi.Object)
	if !is {
		return
	}
	ext.iface.ObjectDestroy(obj.AsPtr())
}

func (me *classProperty) unregister(ext *extension, class *classEntry) {
}

type classMethod struct {
	name        string
	argTypes    []reflect.Type
	returnType  *reflect.Type
	method      gdc.ClassMethodInfo
	fn          interface{}
	pinner      runtime.Pinner
	saveArgInfo []gdc.PropertyInfo
}

func (me *classMethod) register(ext *extension, class *classEntry) {
	mi, clean := me.method.ToRaw()
	defer clean()
	ext.iface.ClassdbRegisterExtensionClassMethod(ext.pLibrary, class.namePtr.AsCPtr(), (*gdc.ClassMethodInfo)(unsafe.Pointer(&mi)))
}

func (me *classMethod) unregister(_ *extension, _ *classEntry) {
	namePtr := gdapi.StringNameFromStringPtr(me.method.Name)
	namePtr.Destroy()
	me.pinner.Unpin()
	for _, arg := range me.saveArgInfo {
		hintPtr := gdapi.StringFromStringPtr(arg.HintString)
		hintPtr.Destroy()
		namePtr := gdapi.StringNameFromStringPtr(arg.Name)
		namePtr.Destroy()
		namePtr = gdapi.StringNameFromStringPtr(arg.ClassName)
		namePtr.Destroy()
	}
	if me.method.ReturnValueInfo != nil {
		hintPtr := gdapi.StringFromStringPtr(me.method.ReturnValueInfo.HintString)
		hintPtr.Destroy()
		namePtr := gdapi.StringNameFromStringPtr(me.method.ReturnValueInfo.Name)
		namePtr.Destroy()
		namePtr = gdapi.StringNameFromStringPtr(me.method.ReturnValueInfo.ClassName)
		namePtr.Destroy()
	}
}

type classSignal struct {
	name     string
	namePtr  gdapi.StringName
	argsPtr  *gdc.PropertyInfo
	argCount uint
}

func (me *classSignal) register(ext *extension, class *classEntry) {
	ext.iface.ClassdbRegisterExtensionClassSignal(ext.pLibrary, class.namePtr.AsCPtr(), me.namePtr.AsCPtr(), me.argsPtr, gdc.Int(me.argCount))
}

func (me *classSignal) create(_ *extension, instance Class, obj gdc.ObjectPtr) {
	reflectSetField(instance, me.name, Signal{
		obj:  *gdapi.ObjectFromPtr(obj),
		name: me.namePtr,
	})
}

func (me *classSignal) free(_ *extension, instance Class) {
	val := reflectGetField(instance, me.name).Interface()
	signal, is := val.(gdapi.Signal)
	if !is {
		return
	}
	signal.Destroy()
}

func (me *classSignal) unregister(_ *extension, _ *classEntry) {
	me.namePtr.Destroy()
}

type classSubscriber string

func (me classSubscriber) create(ext *extension, instance Class) error {
	subInstance, err := CreateClass[*gdapi.SignalSubscribers](ext)
	if err != nil {
		return err
	}
	reflectSetFieldUnsafe(instance, string(me), **subInstance)
	return nil
}

func (me classSubscriber) free(ext *extension, instance Class) {
	val := reflectGetFieldUnsafe(instance, string(me)).Interface()
	subs, is := val.(gdapi.SignalSubscribers)
	if !is {
		ext.iface.ObjectDestroy(subs.AsPtr())
	}
}

type classEntry struct {
	name        string
	constructor ClassConstructor
	properties  map[string]classProperty
	methods     []classMethod
	signals     []classSignal
	subscribers []classSubscriber

	parentNamePtr gdapi.StringName
	namePtr       gdapi.StringName

	mutex     sync.Mutex
	instances map[uint64]*classInstance
}
