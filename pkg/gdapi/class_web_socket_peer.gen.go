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

type ptrsForWebSocketPeerList struct {
  fnConnectToUrl gdc.MethodBindPtr
  fnAcceptStream gdc.MethodBindPtr
  fnSend gdc.MethodBindPtr
  fnSendText gdc.MethodBindPtr
  fnWasStringPacket gdc.MethodBindPtr
  fnPoll gdc.MethodBindPtr
  fnClose gdc.MethodBindPtr
  fnGetConnectedHost gdc.MethodBindPtr
  fnGetConnectedPort gdc.MethodBindPtr
  fnGetSelectedProtocol gdc.MethodBindPtr
  fnGetRequestedUrl gdc.MethodBindPtr
  fnSetNoDelay gdc.MethodBindPtr
  fnGetCurrentOutboundBufferedAmount gdc.MethodBindPtr
  fnGetReadyState gdc.MethodBindPtr
  fnGetCloseCode gdc.MethodBindPtr
  fnGetCloseReason gdc.MethodBindPtr
  fnGetSupportedProtocols gdc.MethodBindPtr
  fnSetSupportedProtocols gdc.MethodBindPtr
  fnGetHandshakeHeaders gdc.MethodBindPtr
  fnSetHandshakeHeaders gdc.MethodBindPtr
  fnGetInboundBufferSize gdc.MethodBindPtr
  fnSetInboundBufferSize gdc.MethodBindPtr
  fnGetOutboundBufferSize gdc.MethodBindPtr
  fnSetOutboundBufferSize gdc.MethodBindPtr
  fnSetMaxQueuedPackets gdc.MethodBindPtr
  fnGetMaxQueuedPackets gdc.MethodBindPtr
}

var ptrsForWebSocketPeer ptrsForWebSocketPeerList

func initWebSocketPeerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("WebSocketPeer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("connect_to_url")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnConnectToUrl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1966198364))
  }
  {
    methodName := StringNameFromStr("accept_stream")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnAcceptStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 255125695))
  }
  {
    methodName := StringNameFromStr("send")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnSend = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2780360567))
  }
  {
    methodName := StringNameFromStr("send_text")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnSendText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
  }
  {
    methodName := StringNameFromStr("was_string_packet")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnWasStringPacket = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("poll")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnPoll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("close")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnClose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1047156615))
  }
  {
    methodName := StringNameFromStr("get_connected_host")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnGetConnectedHost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_connected_port")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnGetConnectedPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_selected_protocol")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnGetSelectedProtocol = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_requested_url")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnGetRequestedUrl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_no_delay")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnSetNoDelay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_current_outbound_buffered_amount")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnGetCurrentOutboundBufferedAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_ready_state")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnGetReadyState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 346482985))
  }
  {
    methodName := StringNameFromStr("get_close_code")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnGetCloseCode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_close_reason")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnGetCloseReason = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_supported_protocols")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnGetSupportedProtocols = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
  {
    methodName := StringNameFromStr("set_supported_protocols")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnSetSupportedProtocols = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015028928))
  }
  {
    methodName := StringNameFromStr("get_handshake_headers")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnGetHandshakeHeaders = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
  }
  {
    methodName := StringNameFromStr("set_handshake_headers")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnSetHandshakeHeaders = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015028928))
  }
  {
    methodName := StringNameFromStr("get_inbound_buffer_size")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnGetInboundBufferSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_inbound_buffer_size")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnSetInboundBufferSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_outbound_buffer_size")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnGetOutboundBufferSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_outbound_buffer_size")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnSetOutboundBufferSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("set_max_queued_packets")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnSetMaxQueuedPackets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_max_queued_packets")
    defer methodName.Destroy()
    ptrsForWebSocketPeer.fnGetMaxQueuedPackets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

type WebSocketPeer struct {
  PacketPeer
}

func (me *WebSocketPeer) BaseClass() string {
  return "WebSocketPeer"
}

func NewWebSocketPeer() *WebSocketPeer {
  str := StringNameFromStr("WebSocketPeer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &WebSocketPeer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type WebSocketPeerWriteMode int
const (
  WebSocketPeerWriteModeWriteModeText WebSocketPeerWriteMode = 0
  WebSocketPeerWriteModeWriteModeBinary WebSocketPeerWriteMode = 1
)

type WebSocketPeerState int
const (
  WebSocketPeerStateStateConnecting WebSocketPeerState = 0
  WebSocketPeerStateStateOpen WebSocketPeerState = 1
  WebSocketPeerStateStateClosing WebSocketPeerState = 2
  WebSocketPeerStateStateClosed WebSocketPeerState = 3
)

func (me *WebSocketPeer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WebSocketPeer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WebSocketPeer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *WebSocketPeer) ConnectToUrl(url String, tls_client_options TLSOptions, ) Error {
  cargs := []gdc.ConstTypePtr{url.AsCTypePtr(), tls_client_options.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnConnectToUrl), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebSocketPeer) AcceptStream(stream StreamPeer, ) Error {
  cargs := []gdc.ConstTypePtr{stream.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnAcceptStream), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebSocketPeer) Send(message PackedByteArray, write_mode WebSocketPeerWriteMode, ) Error {
  cargs := []gdc.ConstTypePtr{message.AsCTypePtr(), gdc.ConstTypePtr(&write_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&write_mode)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnSend), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebSocketPeer) SendText(message String, ) Error {
  cargs := []gdc.ConstTypePtr{message.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnSendText), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebSocketPeer) WasStringPacket() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnWasStringPacket), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketPeer) Poll()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnPoll), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) Close(code int64, reason String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&code) , reason.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnClose), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) GetConnectedHost() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnGetConnectedHost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketPeer) GetConnectedPort() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnGetConnectedPort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketPeer) GetSelectedProtocol() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnGetSelectedProtocol), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketPeer) GetRequestedUrl() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnGetRequestedUrl), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketPeer) SetNoDelay(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnSetNoDelay), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) GetCurrentOutboundBufferedAmount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnGetCurrentOutboundBufferedAmount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketPeer) GetReadyState() WebSocketPeerState {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret WebSocketPeerState

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnGetReadyState), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebSocketPeer) GetCloseCode() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnGetCloseCode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketPeer) GetCloseReason() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnGetCloseReason), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketPeer) GetSupportedProtocols() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnGetSupportedProtocols), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketPeer) SetSupportedProtocols(protocols PackedStringArray, )  {
  cargs := []gdc.ConstTypePtr{protocols.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnSetSupportedProtocols), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) GetHandshakeHeaders() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnGetHandshakeHeaders), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketPeer) SetHandshakeHeaders(protocols PackedStringArray, )  {
  cargs := []gdc.ConstTypePtr{protocols.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnSetHandshakeHeaders), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) GetInboundBufferSize() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnGetInboundBufferSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketPeer) SetInboundBufferSize(buffer_size int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnSetInboundBufferSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) GetOutboundBufferSize() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnGetOutboundBufferSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketPeer) SetOutboundBufferSize(buffer_size int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnSetOutboundBufferSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) SetMaxQueuedPackets(buffer_size int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnSetMaxQueuedPackets), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) GetMaxQueuedPackets() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebSocketPeer.fnGetMaxQueuedPackets), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
