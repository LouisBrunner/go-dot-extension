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

type ptrsForPacketPeerStreamList struct {
	fnSetStreamPeer          gdc.MethodBindPtr
	fnGetStreamPeer          gdc.MethodBindPtr
	fnSetInputBufferMaxSize  gdc.MethodBindPtr
	fnSetOutputBufferMaxSize gdc.MethodBindPtr
	fnGetInputBufferMaxSize  gdc.MethodBindPtr
	fnGetOutputBufferMaxSize gdc.MethodBindPtr
}

var ptrsForPacketPeerStream ptrsForPacketPeerStreamList

func initPacketPeerStreamPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PacketPeerStream")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_stream_peer")
		defer methodName.Destroy()
		ptrsForPacketPeerStream.fnSetStreamPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3281897016))
	}
	{
		methodName := StringNameFromStr("get_stream_peer")
		defer methodName.Destroy()
		ptrsForPacketPeerStream.fnGetStreamPeer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2741655269))
	}
	{
		methodName := StringNameFromStr("set_input_buffer_max_size")
		defer methodName.Destroy()
		ptrsForPacketPeerStream.fnSetInputBufferMaxSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_output_buffer_max_size")
		defer methodName.Destroy()
		ptrsForPacketPeerStream.fnSetOutputBufferMaxSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_input_buffer_max_size")
		defer methodName.Destroy()
		ptrsForPacketPeerStream.fnGetInputBufferMaxSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_output_buffer_max_size")
		defer methodName.Destroy()
		ptrsForPacketPeerStream.fnGetOutputBufferMaxSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
}

type PacketPeerStream struct {
	PacketPeer
}

func (me *PacketPeerStream) BaseClass() string {
	return "PacketPeerStream"
}

func NewPacketPeerStream() *PacketPeerStream {
	str := StringNameFromStr("PacketPeerStream") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PacketPeerStream{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PacketPeerStream) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PacketPeerStream) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PacketPeerStream) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PacketPeerStream) SetStreamPeer(peer StreamPeer) {
	cargs := []gdc.ConstTypePtr{peer.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerStream.fnSetStreamPeer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PacketPeerStream) GetStreamPeer() StreamPeer {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStreamPeer()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerStream.fnGetStreamPeer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PacketPeerStream) SetInputBufferMaxSize(max_size_bytes int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_size_bytes)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerStream.fnSetInputBufferMaxSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PacketPeerStream) SetOutputBufferMaxSize(max_size_bytes int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_size_bytes)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerStream.fnSetOutputBufferMaxSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PacketPeerStream) GetInputBufferMaxSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerStream.fnGetInputBufferMaxSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PacketPeerStream) GetOutputBufferMaxSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPacketPeerStream.fnGetOutputBufferMaxSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
