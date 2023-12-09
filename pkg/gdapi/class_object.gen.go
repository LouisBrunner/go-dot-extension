// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Object struct {
  obj gdc.ObjectPtr
}

func (me *Object) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Object) BaseClass() string {
  return "Object"
}



// Constants

var (
  ObjectNotificationPostinitialize = "0" // TODO: construct correctly
  ObjectNotificationPredelete = "1" // TODO: construct correctly
)

// Enums

type ObjectConnectFlags int
const (
  ObjectConnectFlagsConnectDeferred ObjectConnectFlags = 1
  ObjectConnectFlagsConnectPersist ObjectConnectFlags = 2
  ObjectConnectFlagsConnectOneShot ObjectConnectFlags = 4
  ObjectConnectFlagsConnectReferenceCounted ObjectConnectFlags = 8
)

func (me *Object) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Object) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Object) GetClass()  {
  panic("TODO: implement")
}

func  (me *Object) IsClass(class String, )  {
  panic("TODO: implement")
}

func  (me *Object) Set(property StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *Object) Get(property StringName, )  {
  panic("TODO: implement")
}

func  (me *Object) SetIndexed(property_path NodePath, value Variant, )  {
  panic("TODO: implement")
}

func  (me *Object) GetIndexed(property_path NodePath, )  {
  panic("TODO: implement")
}

func  (me *Object) GetPropertyList()  {
  panic("TODO: implement")
}

func  (me *Object) GetMethodList()  {
  panic("TODO: implement")
}

func  (me *Object) PropertyCanRevert(property StringName, )  {
  panic("TODO: implement")
}

func  (me *Object) PropertyGetRevert(property StringName, )  {
  panic("TODO: implement")
}

func  (me *Object) Notification(what int, reversed bool, )  {
  panic("TODO: implement")
}

func  (me *Object) ToString()  {
  panic("TODO: implement")
}

func  (me *Object) GetInstanceId()  {
  panic("TODO: implement")
}

func  (me *Object) SetScript(script Variant, )  {
  panic("TODO: implement")
}

func  (me *Object) GetScript()  {
  panic("TODO: implement")
}

func  (me *Object) SetMeta(name StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *Object) RemoveMeta(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Object) GetMeta(name StringName, default_ Variant, )  {
  panic("TODO: implement")
}

func  (me *Object) HasMeta(name StringName, )  {
  panic("TODO: implement")
}

func  (me *Object) GetMetaList()  {
  panic("TODO: implement")
}

func  (me *Object) AddUserSignal(signal String, arguments Array, )  {
  panic("TODO: implement")
}

func  (me *Object) HasUserSignal(signal StringName, )  {
  panic("TODO: implement")
}

func  (me *Object) EmitSignal(signal StringName, )  {
  panic("TODO: implement")
}

func  (me *Object) Call(method StringName, )  {
  panic("TODO: implement")
}

func  (me *Object) CallDeferred(method StringName, )  {
  panic("TODO: implement")
}

func  (me *Object) SetDeferred(property StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *Object) Callv(method StringName, arg_array Array, )  {
  panic("TODO: implement")
}

func  (me *Object) HasMethod(method StringName, )  {
  panic("TODO: implement")
}

func  (me *Object) HasSignal(signal StringName, )  {
  panic("TODO: implement")
}

func  (me *Object) GetSignalList()  {
  panic("TODO: implement")
}

func  (me *Object) GetSignalConnectionList(signal StringName, )  {
  panic("TODO: implement")
}

func  (me *Object) GetIncomingConnections()  {
  panic("TODO: implement")
}

func  (me *Object) Connect(signal StringName, callable Callable, flags int, )  {
  panic("TODO: implement")
}

func  (me *Object) Disconnect(signal StringName, callable Callable, )  {
  panic("TODO: implement")
}

func  (me *Object) IsConnected(signal StringName, callable Callable, )  {
  panic("TODO: implement")
}

func  (me *Object) SetBlockSignals(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Object) IsBlockingSignals()  {
  panic("TODO: implement")
}

func  (me *Object) NotifyPropertyListChanged()  {
  panic("TODO: implement")
}

func  (me *Object) SetMessageTranslation(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Object) CanTranslateMessages()  {
  panic("TODO: implement")
}

func  (me *Object) Tr(message StringName, context StringName, )  {
  panic("TODO: implement")
}

func  (me *Object) TrN(message StringName, plural_message StringName, n int, context StringName, )  {
  panic("TODO: implement")
}

func  (me *Object) IsQueuedForDeletion()  {
  panic("TODO: implement")
}

func  (me *Object) CancelFree()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
