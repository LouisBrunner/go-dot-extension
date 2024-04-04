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

type WebSocketMultiplayerPeer struct {
  MultiplayerPeer
}

func (me *WebSocketMultiplayerPeer) BaseClass() string {
  return "WebSocketMultiplayerPeer"
}

func NewWebSocketMultiplayerPeer() *WebSocketMultiplayerPeer {
  str := StringNameFromStr("WebSocketMultiplayerPeer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &WebSocketMultiplayerPeer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *WebSocketMultiplayerPeer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WebSocketMultiplayerPeer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WebSocketMultiplayerPeer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *WebSocketMultiplayerPeer) CreateClient(url String, tls_client_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_client")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1966198364) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{url.AsCTypePtr(), tls_client_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebSocketMultiplayerPeer) CreateServer(port int64, bind_address String, tls_server_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_server")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2400822951) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , bind_address.AsCTypePtr(), tls_server_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&port)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebSocketMultiplayerPeer) GetPeer(peer_id int64, ) WebSocketPeer {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1381378851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewWebSocketPeer()
  pinner.Pin(&peer_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketMultiplayerPeer) GetPeerAddress(id int64, ) String {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peer_address")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketMultiplayerPeer) GetPeerPort(id int64, ) int64 {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peer_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketMultiplayerPeer) GetSupportedProtocols() PackedStringArray {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_supported_protocols")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketMultiplayerPeer) SetSupportedProtocols(protocols PackedStringArray, )  {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_supported_protocols")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{protocols.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketMultiplayerPeer) GetHandshakeHeaders() PackedStringArray {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_handshake_headers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketMultiplayerPeer) SetHandshakeHeaders(protocols PackedStringArray, )  {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_handshake_headers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{protocols.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketMultiplayerPeer) GetInboundBufferSize() int64 {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inbound_buffer_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketMultiplayerPeer) SetInboundBufferSize(buffer_size int64, )  {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_inbound_buffer_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketMultiplayerPeer) GetOutboundBufferSize() int64 {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outbound_buffer_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketMultiplayerPeer) SetOutboundBufferSize(buffer_size int64, )  {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_outbound_buffer_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketMultiplayerPeer) GetHandshakeTimeout() float64 {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_handshake_timeout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketMultiplayerPeer) SetHandshakeTimeout(timeout float64, )  {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_handshake_timeout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketMultiplayerPeer) SetMaxQueuedPackets(max_queued_packets int64, )  {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_queued_packets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_queued_packets) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketMultiplayerPeer) GetMaxQueuedPackets() int64 {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_queued_packets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
