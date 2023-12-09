// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiplayerAPI struct {
  obj gdc.ObjectPtr
}

func (me *MultiplayerAPI) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiplayerAPI) BaseClass() string {
  return "MultiplayerAPI"
}



// Enums

type MultiplayerAPIRPCMode int
const (
  MultiplayerAPIRPCModeRpcModeDisabled MultiplayerAPIRPCMode = 0
  MultiplayerAPIRPCModeRpcModeAnyPeer MultiplayerAPIRPCMode = 1
  MultiplayerAPIRPCModeRpcModeAuthority MultiplayerAPIRPCMode = 2
)

func (me *MultiplayerAPI) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiplayerAPI) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MultiplayerAPI) HasMultiplayerPeer()  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPI) GetMultiplayerPeer()  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPI) SetMultiplayerPeer(peer MultiplayerPeer, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPI) GetUniqueId()  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPI) IsServer()  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPI) GetRemoteSenderId()  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPI) Poll()  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPI) Rpc(peer int, object Object, method StringName, arguments Array, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPI) ObjectConfigurationAdd(object Object, configuration Variant, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPI) ObjectConfigurationRemove(object Object, configuration Variant, )  {
  panic("TODO: implement")
}

func  (me *MultiplayerAPI) GetPeers()  {
  panic("TODO: implement")
}

func  MultiplayerAPISetDefaultInterface(interface_name StringName, )  {
  panic("TODO: implement")
}

func  MultiplayerAPIGetDefaultInterface()  {
  panic("TODO: implement")
}

func  MultiplayerAPICreateDefaultInterface()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
