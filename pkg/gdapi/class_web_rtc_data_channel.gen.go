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

type ptrsForWebRTCDataChannelList struct {
	fnPoll                 gdc.MethodBindPtr
	fnClose                gdc.MethodBindPtr
	fnWasStringPacket      gdc.MethodBindPtr
	fnSetWriteMode         gdc.MethodBindPtr
	fnGetWriteMode         gdc.MethodBindPtr
	fnGetReadyState        gdc.MethodBindPtr
	fnGetLabel             gdc.MethodBindPtr
	fnIsOrdered            gdc.MethodBindPtr
	fnGetId                gdc.MethodBindPtr
	fnGetMaxPacketLifeTime gdc.MethodBindPtr
	fnGetMaxRetransmits    gdc.MethodBindPtr
	fnGetProtocol          gdc.MethodBindPtr
	fnIsNegotiated         gdc.MethodBindPtr
	fnGetBufferedAmount    gdc.MethodBindPtr
}

var ptrsForWebRTCDataChannel ptrsForWebRTCDataChannelList

func initWebRTCDataChannelPtrs(iface gdc.Interface) {

	className := StringNameFromStr("WebRTCDataChannel")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("poll")
		defer methodName.Destroy()
		ptrsForWebRTCDataChannel.fnPoll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
	}
	{
		methodName := StringNameFromStr("close")
		defer methodName.Destroy()
		ptrsForWebRTCDataChannel.fnClose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("was_string_packet")
		defer methodName.Destroy()
		ptrsForWebRTCDataChannel.fnWasStringPacket = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_write_mode")
		defer methodName.Destroy()
		ptrsForWebRTCDataChannel.fnSetWriteMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1999768052))
	}
	{
		methodName := StringNameFromStr("get_write_mode")
		defer methodName.Destroy()
		ptrsForWebRTCDataChannel.fnGetWriteMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2848495172))
	}
	{
		methodName := StringNameFromStr("get_ready_state")
		defer methodName.Destroy()
		ptrsForWebRTCDataChannel.fnGetReadyState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3501143017))
	}
	{
		methodName := StringNameFromStr("get_label")
		defer methodName.Destroy()
		ptrsForWebRTCDataChannel.fnGetLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("is_ordered")
		defer methodName.Destroy()
		ptrsForWebRTCDataChannel.fnIsOrdered = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_id")
		defer methodName.Destroy()
		ptrsForWebRTCDataChannel.fnGetId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_max_packet_life_time")
		defer methodName.Destroy()
		ptrsForWebRTCDataChannel.fnGetMaxPacketLifeTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_max_retransmits")
		defer methodName.Destroy()
		ptrsForWebRTCDataChannel.fnGetMaxRetransmits = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_protocol")
		defer methodName.Destroy()
		ptrsForWebRTCDataChannel.fnGetProtocol = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("is_negotiated")
		defer methodName.Destroy()
		ptrsForWebRTCDataChannel.fnIsNegotiated = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_buffered_amount")
		defer methodName.Destroy()
		ptrsForWebRTCDataChannel.fnGetBufferedAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

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
	WebRTCDataChannelWriteModeWriteModeText   WebRTCDataChannelWriteMode = 0
	WebRTCDataChannelWriteModeWriteModeBinary WebRTCDataChannelWriteMode = 1
)

type WebRTCDataChannelChannelState int

const (
	WebRTCDataChannelChannelStateStateConnecting WebRTCDataChannelChannelState = 0
	WebRTCDataChannelChannelStateStateOpen       WebRTCDataChannelChannelState = 1
	WebRTCDataChannelChannelStateStateClosing    WebRTCDataChannelChannelState = 2
	WebRTCDataChannelChannelStateStateClosed     WebRTCDataChannelChannelState = 3
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

func (me *WebRTCDataChannel) Poll() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCDataChannel.fnPoll), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCDataChannel) Close() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCDataChannel.fnClose), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *WebRTCDataChannel) WasStringPacket() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCDataChannel.fnWasStringPacket), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *WebRTCDataChannel) SetWriteMode(write_mode WebRTCDataChannelWriteMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&write_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCDataChannel.fnSetWriteMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *WebRTCDataChannel) GetWriteMode() WebRTCDataChannelWriteMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret WebRTCDataChannelWriteMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCDataChannel.fnGetWriteMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCDataChannel) GetReadyState() WebRTCDataChannelChannelState {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret WebRTCDataChannelChannelState

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCDataChannel.fnGetReadyState), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WebRTCDataChannel) GetLabel() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCDataChannel.fnGetLabel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *WebRTCDataChannel) IsOrdered() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCDataChannel.fnIsOrdered), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *WebRTCDataChannel) GetId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCDataChannel.fnGetId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *WebRTCDataChannel) GetMaxPacketLifeTime() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCDataChannel.fnGetMaxPacketLifeTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *WebRTCDataChannel) GetMaxRetransmits() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCDataChannel.fnGetMaxRetransmits), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *WebRTCDataChannel) GetProtocol() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCDataChannel.fnGetProtocol), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *WebRTCDataChannel) IsNegotiated() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCDataChannel.fnIsNegotiated), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *WebRTCDataChannel) GetBufferedAmount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWebRTCDataChannel.fnGetBufferedAmount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
