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

type WebRTCDataChannel struct {
  PacketPeer
}

func (me *WebRTCDataChannel) BaseClass() string {
  return "WebRTCDataChannel"
}

func NewWebRTCDataChannel() *WebRTCDataChannel {
  str := StringNameFromStr("WebRTCDataChannel") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &WebRTCDataChannel{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type WebRTCDataChannelWriteMode int
const (
  WebRTCDataChannelWriteModeWriteModeText WebRTCDataChannelWriteMode = 0
  WebRTCDataChannelWriteModeWriteModeBinary WebRTCDataChannelWriteMode = 1
)

type WebRTCDataChannelChannelState int
const (
  WebRTCDataChannelChannelStateStateConnecting WebRTCDataChannelChannelState = 0
  WebRTCDataChannelChannelStateStateOpen WebRTCDataChannelChannelState = 1
  WebRTCDataChannelChannelStateStateClosing WebRTCDataChannelChannelState = 2
  WebRTCDataChannelChannelStateStateClosed WebRTCDataChannelChannelState = 3
)

func (me *WebRTCDataChannel) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WebRTCDataChannel) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WebRTCDataChannel) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *WebRTCDataChannel) Poll() Error {
  classNameV := StringNameFromStr("WebRTCDataChannel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("poll")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebRTCDataChannel) Close()  {
  classNameV := StringNameFromStr("WebRTCDataChannel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("close")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebRTCDataChannel) WasStringPacket() bool {
  classNameV := StringNameFromStr("WebRTCDataChannel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("was_string_packet")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebRTCDataChannel) SetWriteMode(write_mode WebRTCDataChannelWriteMode, )  {
  classNameV := StringNameFromStr("WebRTCDataChannel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_write_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1999768052) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&write_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WebRTCDataChannel) GetWriteMode() WebRTCDataChannelWriteMode {
  classNameV := StringNameFromStr("WebRTCDataChannel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_write_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2848495172) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret WebRTCDataChannelWriteMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebRTCDataChannel) GetReadyState() WebRTCDataChannelChannelState {
  classNameV := StringNameFromStr("WebRTCDataChannel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ready_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3501143017) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret WebRTCDataChannelChannelState

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WebRTCDataChannel) GetLabel() String {
  classNameV := StringNameFromStr("WebRTCDataChannel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebRTCDataChannel) IsOrdered() bool {
  classNameV := StringNameFromStr("WebRTCDataChannel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ordered")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebRTCDataChannel) GetId() int64 {
  classNameV := StringNameFromStr("WebRTCDataChannel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebRTCDataChannel) GetMaxPacketLifeTime() int64 {
  classNameV := StringNameFromStr("WebRTCDataChannel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_packet_life_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebRTCDataChannel) GetMaxRetransmits() int64 {
  classNameV := StringNameFromStr("WebRTCDataChannel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_retransmits")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebRTCDataChannel) GetProtocol() String {
  classNameV := StringNameFromStr("WebRTCDataChannel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_protocol")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WebRTCDataChannel) IsNegotiated() bool {
  classNameV := StringNameFromStr("WebRTCDataChannel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_negotiated")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WebRTCDataChannel) GetBufferedAmount() int64 {
  classNameV := StringNameFromStr("WebRTCDataChannel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffered_amount")
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
