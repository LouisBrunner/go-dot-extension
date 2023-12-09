// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiplayerAPIExtension struct {
  obj gdc.ObjectPtr
}

func (me *MultiplayerAPIExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiplayerAPIExtension) BaseClass() string {
  return "MultiplayerAPIExtension"
}



// Enums

func (me *MultiplayerAPIExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MultiplayerAPIExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiplayerAPIExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MultiplayerAPIExtension) XPoll()  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPIExtension) XSetMultiplayerPeer(multiplayer_peer MultiplayerPeer, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPIExtension) XGetMultiplayerPeer()  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPIExtension) XGetUniqueId()  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPIExtension) XGetPeerIds()  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPIExtension) XRpc(peer int, object Object, method StringName, args Array, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPIExtension) XGetRemoteSenderId()  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPIExtension) XObjectConfigurationAdd(object Object, configuration Variant, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPIExtension) XObjectConfigurationRemove(object Object, configuration Variant, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
