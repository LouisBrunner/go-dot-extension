// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type WebRTCMultiplayerPeer struct {
  obj gdc.ObjectPtr
}

func (me *WebRTCMultiplayerPeer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WebRTCMultiplayerPeer) BaseClass() string {
  return "WebRTCMultiplayerPeer"
}



// Enums

func (me *WebRTCMultiplayerPeer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WebRTCMultiplayerPeer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WebRTCMultiplayerPeer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *WebRTCMultiplayerPeer) CreateServer(channels_config Array, ) Error {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_server")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2865356025) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(channels_config.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebRTCMultiplayerPeer) CreateClient(peer_id int, channels_config Array, ) Error {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_client")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1777354631) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id), gdc.ConstTypePtr(channels_config.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebRTCMultiplayerPeer) CreateMesh(peer_id int, channels_config Array, ) Error {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1777354631) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id), gdc.ConstTypePtr(channels_config.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebRTCMultiplayerPeer) AddPeer(peer WebRTCPeerConnection, peer_id int, unreliable_lifetime int, ) Error {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2555866323) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(peer.AsCTypePtr()), gdc.ConstTypePtr(&peer_id), gdc.ConstTypePtr(&unreliable_lifetime), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebRTCMultiplayerPeer) RemovePeer(peer_id int, )  {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebRTCMultiplayerPeer) HasPeer(peer_id int, ) bool {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3067735520) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebRTCMultiplayerPeer) GetPeer(peer_id int, ) Dictionary {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3554694381) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebRTCMultiplayerPeer) GetPeers() Dictionary {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2382534195) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
