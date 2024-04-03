// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("connect_to_url")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1966198364) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(url.AsCTypePtr()), gdc.ConstTypePtr(tls_client_options.AsCTypePtr()), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebSocketPeer) AcceptStream(stream StreamPeer, ) Error {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("accept_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 255125695) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(stream.AsCTypePtr()), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebSocketPeer) Send(message PackedByteArray, write_mode WebSocketPeerWriteMode, ) Error {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("send")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2780360567) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(message.AsCTypePtr()), gdc.ConstTypePtr(&write_mode), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebSocketPeer) SendText(message String, ) Error {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("send_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(message.AsCTypePtr()), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebSocketPeer) WasStringPacket() bool {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("was_string_packet")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketPeer) Poll()  {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("poll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) Close(code int64, reason String, )  {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("close")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1047156615) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&code), gdc.ConstTypePtr(reason.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) GetConnectedHost() String {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connected_host")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketPeer) GetConnectedPort() int64 {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connected_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketPeer) GetSelectedProtocol() String {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected_protocol")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketPeer) GetRequestedUrl() String {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_requested_url")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketPeer) SetNoDelay(enabled bool, )  {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_no_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) GetCurrentOutboundBufferedAmount() int64 {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_outbound_buffered_amount")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketPeer) GetReadyState() WebSocketPeerState {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ready_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 346482985) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret WebSocketPeerState

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebSocketPeer) GetCloseCode() int64 {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_close_code")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketPeer) GetCloseReason() String {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_close_reason")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketPeer) GetSupportedProtocols() PackedStringArray {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_supported_protocols")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketPeer) SetSupportedProtocols(protocols PackedStringArray, )  {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_supported_protocols")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(protocols.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) GetHandshakeHeaders() PackedStringArray {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_handshake_headers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebSocketPeer) SetHandshakeHeaders(protocols PackedStringArray, )  {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_handshake_headers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4015028928) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(protocols.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) GetInboundBufferSize() int64 {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inbound_buffer_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketPeer) SetInboundBufferSize(buffer_size int64, )  {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_inbound_buffer_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_size), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) GetOutboundBufferSize() int64 {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outbound_buffer_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebSocketPeer) SetOutboundBufferSize(buffer_size int64, )  {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_outbound_buffer_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_size), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) SetMaxQueuedPackets(buffer_size int64, )  {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_queued_packets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_size), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebSocketPeer) GetMaxQueuedPackets() int64 {
  classNameV := StringNameFromStr("WebSocketPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_queued_packets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
