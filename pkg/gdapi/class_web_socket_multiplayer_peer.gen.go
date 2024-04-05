// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForWebSocketMultiplayerPeerList struct {
  fnCreateClient gdc.MethodBindPtr
  fnCreateServer gdc.MethodBindPtr
  fnGetPeer gdc.MethodBindPtr
  fnGetPeerAddress gdc.MethodBindPtr
  fnGetPeerPort gdc.MethodBindPtr
  fnGetSupportedProtocols gdc.MethodBindPtr
  fnSetSupportedProtocols gdc.MethodBindPtr
  fnGetHandshakeHeaders gdc.MethodBindPtr
  fnSetHandshakeHeaders gdc.MethodBindPtr
  fnGetInboundBufferSize gdc.MethodBindPtr
  fnSetInboundBufferSize gdc.MethodBindPtr
  fnGetOutboundBufferSize gdc.MethodBindPtr
  fnSetOutboundBufferSize gdc.MethodBindPtr
  fnGetHandshakeTimeout gdc.MethodBindPtr
  fnSetHandshakeTimeout gdc.MethodBindPtr
  fnSetMaxQueuedPackets gdc.MethodBindPtr
  fnGetMaxQueuedPackets gdc.MethodBindPtr
}

var ptrsForWebSocketMultiplayerPeer ptrsForWebSocketMultiplayerPeerList

func initWebSocketMultiplayerPeerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("WebSocketMultiplayerPeer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("create_client")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnCreateClient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1966198364))
  }
  {
    methodName := StringNameFromStr("create_server")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnCreateServer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2400822951))
  }
  {
    methodName := StringNameFromStr("get_peer")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnGetPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1381378851))
  }
  {
    methodName := StringNameFromStr("get_peer_address")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnGetPeerAddress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
  }
  {
    methodName := StringNameFromStr("get_peer_port")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnGetPeerPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
  }
  {
    methodName := StringNameFromStr("get_supported_protocols")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnGetSupportedProtocols = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
  {
    methodName := StringNameFromStr("set_supported_protocols")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnSetSupportedProtocols = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015028928))
  }
  {
    methodName := StringNameFromStr("get_handshake_headers")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnGetHandshakeHeaders = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
  {
    methodName := StringNameFromStr("set_handshake_headers")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnSetHandshakeHeaders = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015028928))
  }
  {
    methodName := StringNameFromStr("get_inbound_buffer_size")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnGetInboundBufferSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_inbound_buffer_size")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnSetInboundBufferSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_outbound_buffer_size")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnGetOutboundBufferSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_outbound_buffer_size")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnSetOutboundBufferSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_handshake_timeout")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnGetHandshakeTimeout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_handshake_timeout")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnSetHandshakeTimeout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("set_max_queued_packets")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnSetMaxQueuedPackets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_max_queued_packets")
    defer methodName.Destroy()
    ptrsForWebSocketMultiplayerPeer.fnGetMaxQueuedPackets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

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
  cargs := []gdc.ConstTypePtr{url.AsCTypePtr(), tls_client_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnCreateClient), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebSocketMultiplayerPeer) CreateServer(port int64, bind_address String, tls_server_options TLSOptions, ) Error {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , bind_address.AsCTypePtr(), tls_server_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&port)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnCreateServer), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebSocketMultiplayerPeer) GetPeer(peer_id int64, ) WebSocketPeer {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&peer_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewWebSocketPeer()
  pinner.Pin(&peer_id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnGetPeer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketMultiplayerPeer) GetPeerAddress(id int64, ) String {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnGetPeerAddress), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketMultiplayerPeer) GetPeerPort(id int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnGetPeerPort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketMultiplayerPeer) GetSupportedProtocols() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnGetSupportedProtocols), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketMultiplayerPeer) SetSupportedProtocols(protocols PackedStringArray, )  {
  cargs := []gdc.ConstTypePtr{protocols.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnSetSupportedProtocols), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketMultiplayerPeer) GetHandshakeHeaders() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnGetHandshakeHeaders), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketMultiplayerPeer) SetHandshakeHeaders(protocols PackedStringArray, )  {
  cargs := []gdc.ConstTypePtr{protocols.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnSetHandshakeHeaders), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketMultiplayerPeer) GetInboundBufferSize() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnGetInboundBufferSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketMultiplayerPeer) SetInboundBufferSize(buffer_size int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnSetInboundBufferSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketMultiplayerPeer) GetOutboundBufferSize() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnGetOutboundBufferSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketMultiplayerPeer) SetOutboundBufferSize(buffer_size int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnSetOutboundBufferSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketMultiplayerPeer) GetHandshakeTimeout() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnGetHandshakeTimeout), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketMultiplayerPeer) SetHandshakeTimeout(timeout float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeout) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnSetHandshakeTimeout), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketMultiplayerPeer) SetMaxQueuedPackets(max_queued_packets int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_queued_packets) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnSetMaxQueuedPackets), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketMultiplayerPeer) GetMaxQueuedPackets() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketMultiplayerPeer.fnGetMaxQueuedPackets), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
