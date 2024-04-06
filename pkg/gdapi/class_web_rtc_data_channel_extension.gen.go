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

type ptrsForWebRTCDataChannelExtensionList struct {
	fnXGetPacket               gdc.MethodBindPtr
	fnXPutPacket               gdc.MethodBindPtr
	fnXGetAvailablePacketCount gdc.MethodBindPtr
	fnXGetMaxPacketSize        gdc.MethodBindPtr
	fnXPoll                    gdc.MethodBindPtr
	fnXClose                   gdc.MethodBindPtr
	fnXSetWriteMode            gdc.MethodBindPtr
	fnXGetWriteMode            gdc.MethodBindPtr
	fnXWasStringPacket         gdc.MethodBindPtr
	fnXGetReadyState           gdc.MethodBindPtr
	fnXGetLabel                gdc.MethodBindPtr
	fnXIsOrdered               gdc.MethodBindPtr
	fnXGetId                   gdc.MethodBindPtr
	fnXGetMaxPacketLifeTime    gdc.MethodBindPtr
	fnXGetMaxRetransmits       gdc.MethodBindPtr
	fnXGetProtocol             gdc.MethodBindPtr
	fnXIsNegotiated            gdc.MethodBindPtr
	fnXGetBufferedAmount       gdc.MethodBindPtr
}

var ptrsForWebRTCDataChannelExtension ptrsForWebRTCDataChannelExtensionList

func initWebRTCDataChannelExtensionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("WebRTCDataChannelExtension")
	defer className.Destroy()
}

type WebRTCDataChannelExtension struct {
	WebRTCDataChannel
}

func (me *WebRTCDataChannelExtension) BaseClass() string {
	return "WebRTCDataChannelExtension"
}

func NewWebRTCDataChannelExtension() *WebRTCDataChannelExtension {
	str := StringNameFromStr("WebRTCDataChannelExtension") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &WebRTCDataChannelExtension{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *WebRTCDataChannelExtension) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *WebRTCDataChannelExtension) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *WebRTCDataChannelExtension) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
