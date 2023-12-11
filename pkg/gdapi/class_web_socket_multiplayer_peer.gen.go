// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type WebSocketMultiplayerPeer struct {
  obj gdc.ObjectPtr
}

func (me *WebSocketMultiplayerPeer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WebSocketMultiplayerPeer) BaseClass() string {
  return "WebSocketMultiplayerPeer"
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
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3097527179) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(url.AsCTypePtr()), gdc.ConstTypePtr(tls_client_options.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebSocketMultiplayerPeer) CreateServer(port int, bind_address String, tls_server_options TLSOptions, ) Error {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_server")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 337374795) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), gdc.ConstTypePtr(bind_address.AsCTypePtr()), gdc.ConstTypePtr(tls_server_options.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebSocketMultiplayerPeer) GetPeer(peer_id int, ) WebSocketPeer {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1381378851) // FIXME: should cache?
  var ret WebSocketPeer
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebSocketMultiplayerPeer) GetPeerAddress(id int, ) String {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peer_address")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebSocketMultiplayerPeer) GetPeerPort(id int, ) int {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_peer_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebSocketMultiplayerPeer) GetSupportedProtocols() PackedStringArray {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_supported_protocols")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebSocketMultiplayerPeer) SetSupportedProtocols(protocols PackedStringArray, )  {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_supported_protocols")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(protocols.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebSocketMultiplayerPeer) GetHandshakeHeaders() PackedStringArray {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_handshake_headers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebSocketMultiplayerPeer) SetHandshakeHeaders(protocols PackedStringArray, )  {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_handshake_headers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(protocols.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebSocketMultiplayerPeer) GetInboundBufferSize() int {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inbound_buffer_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebSocketMultiplayerPeer) SetInboundBufferSize(buffer_size int, )  {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_inbound_buffer_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebSocketMultiplayerPeer) GetOutboundBufferSize() int {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outbound_buffer_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebSocketMultiplayerPeer) SetOutboundBufferSize(buffer_size int, )  {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_outbound_buffer_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebSocketMultiplayerPeer) GetHandshakeTimeout() float32 {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_handshake_timeout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WebSocketMultiplayerPeer) SetHandshakeTimeout(timeout float32, )  {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_handshake_timeout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebSocketMultiplayerPeer) SetMaxQueuedPackets(max_queued_packets int, )  {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_queued_packets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_queued_packets), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WebSocketMultiplayerPeer) GetMaxQueuedPackets() int {
  classNameV := StringNameFromStr("WebSocketMultiplayerPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_queued_packets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
