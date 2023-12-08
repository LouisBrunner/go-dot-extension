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

// TODO: needed?
// const (
// // )

var (
  ObjectNotificationPostinitialize = "0" // TODO: construct correctly
  ObjectNotificationPredelete = "1" // TODO: construct correctly
)

type ObjectConnectFlags int
const (
  ObjectConnectFlagsConnectDeferred ObjectConnectFlags = 1
  ObjectConnectFlagsConnectPersist ObjectConnectFlags = 2
  ObjectConnectFlagsConnectOneShot ObjectConnectFlags = 4
  ObjectConnectFlagsConnectReferenceCounted ObjectConnectFlags = 8
)

func  (me *Object) GetClass() { // TODO: return value
  // TODO: implement
}

func  (me *Object) IsClass(class String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) Set(property StringName, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) Get(property StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) SetIndexed(property_path NodePath, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) GetIndexed(property_path NodePath, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) GetPropertyList() { // TODO: return value
  // TODO: implement
}

func  (me *Object) GetMethodList() { // TODO: return value
  // TODO: implement
}

func  (me *Object) PropertyCanRevert(property StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) PropertyGetRevert(property StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) Notification(what int, reversed bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) ToString() { // TODO: return value
  // TODO: implement
}

func  (me *Object) GetInstanceId() { // TODO: return value
  // TODO: implement
}

func  (me *Object) SetScript(script Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) GetScript() { // TODO: return value
  // TODO: implement
}

func  (me *Object) SetMeta(name StringName, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) RemoveMeta(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) GetMeta(name StringName, default_ Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) HasMeta(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) GetMetaList() { // TODO: return value
  // TODO: implement
}

func  (me *Object) AddUserSignal(signal String, arguments Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) HasUserSignal(signal StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) EmitSignal(signal StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) Call(method StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) CallDeferred(method StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) SetDeferred(property StringName, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) Callv(method StringName, arg_array Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) HasMethod(method StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) HasSignal(signal StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) GetSignalList() { // TODO: return value
  // TODO: implement
}

func  (me *Object) GetSignalConnectionList(signal StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) GetIncomingConnections() { // TODO: return value
  // TODO: implement
}

func  (me *Object) Connect(signal StringName, callable Callable, flags int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) Disconnect(signal StringName, callable Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) IsConnected(signal StringName, callable Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) SetBlockSignals(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) IsBlockingSignals() { // TODO: return value
  // TODO: implement
}

func  (me *Object) NotifyPropertyListChanged() { // TODO: return value
  // TODO: implement
}

func  (me *Object) SetMessageTranslation(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) CanTranslateMessages() { // TODO: return value
  // TODO: implement
}

func  (me *Object) Tr(message StringName, context StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) TrN(message StringName, plural_message StringName, n int, context StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Object) IsQueuedForDeletion() { // TODO: return value
  // TODO: implement
}

func  (me *Object) CancelFree() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
