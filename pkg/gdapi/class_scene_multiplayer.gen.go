// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SceneMultiplayer struct {
  obj gdc.ObjectPtr
}

func (me *SceneMultiplayer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SceneMultiplayer) BaseClass() string {
  return "SceneMultiplayer"
}



// Enums

func (me *SceneMultiplayer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SceneMultiplayer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SceneMultiplayer) SetRootPath(path NodePath, )  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) GetRootPath()  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) Clear()  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) DisconnectPeer(id int, )  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) GetAuthenticatingPeers()  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) SendAuth(id int, data PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) CompleteAuth(id int, )  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) SetAuthCallback(callback Callable, )  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) GetAuthCallback()  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) SetAuthTimeout(timeout float32, )  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) GetAuthTimeout()  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) SetRefuseNewConnections(refuse bool, )  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) IsRefusingNewConnections()  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) SetAllowObjectDecoding(enable bool, )  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) IsObjectDecodingAllowed()  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) SetServerRelayEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) IsServerRelayEnabled()  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) SendBytes(bytes PackedByteArray, id int, mode MultiplayerPeerTransferMode, channel int, )  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) GetMaxSyncPacketSize()  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) SetMaxSyncPacketSize(size int, )  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) GetMaxDeltaPacketSize()  {
  panic("TODO: implement")
}

func  (me *SceneMultiplayer) SetMaxDeltaPacketSize(size int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
