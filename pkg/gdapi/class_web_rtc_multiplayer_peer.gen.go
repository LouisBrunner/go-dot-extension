// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type WebRTCMultiplayerPeer struct {
  MultiplayerPeer
}

func (me *WebRTCMultiplayerPeer) BaseClass() string {
  return "WebRTCMultiplayerPeer"
}

func NewWebRTCMultiplayerPeer() *WebRTCMultiplayerPeer {
  str := StringNameFromStr("WebRTCMultiplayerPeer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &WebRTCMultiplayerPeer{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{channels_config.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebRTCMultiplayerPeer) CreateClient(peer_id int64, channels_config Array, ) Error {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_client")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2641732907) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id) , channels_config.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&peer_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebRTCMultiplayerPeer) CreateMesh(peer_id int64, channels_config Array, ) Error {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2641732907) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id) , channels_config.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&peer_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebRTCMultiplayerPeer) AddPeer(peer WebRTCPeerConnection, peer_id int64, unreliable_lifetime int64, ) Error {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4078953270) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{peer.AsCTypePtr(), gdc.ConstTypePtr(&peer_id) , gdc.ConstTypePtr(&unreliable_lifetime) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&peer_id)
  pinner.Pin(&unreliable_lifetime)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebRTCMultiplayerPeer) RemovePeer(peer_id int64, )  {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebRTCMultiplayerPeer) HasPeer(peer_id int64, ) bool {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3067735520) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&peer_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebRTCMultiplayerPeer) GetPeer(peer_id int64, ) Dictionary {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3554694381) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&peer_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebRTCMultiplayerPeer) GetPeers() Dictionary {
  classNameV := StringNameFromStr("WebRTCMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2382534195) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
