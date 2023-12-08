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

type MultiplayerAPIRPCMode int
const (
  MultiplayerAPIRPCModeRpcModeDisabled MultiplayerAPIRPCMode = 0
  MultiplayerAPIRPCModeRpcModeAnyPeer MultiplayerAPIRPCMode = 1
  MultiplayerAPIRPCModeRpcModeAuthority MultiplayerAPIRPCMode = 2
)

func  (me *MultiplayerAPI) HasMultiplayerPeer() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerAPI) GetMultiplayerPeer() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerAPI) SetMultiplayerPeer(peer MultiplayerPeer, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerAPI) GetUniqueId() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerAPI) IsServer() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerAPI) GetRemoteSenderId() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerAPI) Poll() { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerAPI) Rpc(peer int, object Object, method StringName, arguments Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerAPI) ObjectConfigurationAdd(object Object, configuration Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerAPI) ObjectConfigurationRemove(object Object, configuration Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *MultiplayerAPI) GetPeers() { // TODO: return value
  // TODO: implement
}

func  MultiplayerAPISetDefaultInterface(interface_name StringName, ) { // TODO: return value
  // TODO: implement
}

func  MultiplayerAPIGetDefaultInterface() { // TODO: return value
  // TODO: implement
}

func  MultiplayerAPICreateDefaultInterface() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
