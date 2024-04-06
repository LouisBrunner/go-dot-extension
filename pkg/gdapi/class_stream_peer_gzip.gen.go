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

type ptrsForStreamPeerGZIPList struct {
	fnStartCompression   gdc.MethodBindPtr
	fnStartDecompression gdc.MethodBindPtr
	fnFinish             gdc.MethodBindPtr
	fnClear              gdc.MethodBindPtr
}

var ptrsForStreamPeerGZIP ptrsForStreamPeerGZIPList

func initStreamPeerGZIPPtrs(iface gdc.Interface) {

	className := StringNameFromStr("StreamPeerGZIP")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("start_compression")
		defer methodName.Destroy()
		ptrsForStreamPeerGZIP.fnStartCompression = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 781582770))
	}
	{
		methodName := StringNameFromStr("start_decompression")
		defer methodName.Destroy()
		ptrsForStreamPeerGZIP.fnStartDecompression = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 781582770))
	}
	{
		methodName := StringNameFromStr("finish")
		defer methodName.Destroy()
		ptrsForStreamPeerGZIP.fnFinish = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForStreamPeerGZIP.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
}

type StreamPeerGZIP struct {
	StreamPeer
}

func (me *StreamPeerGZIP) BaseClass() string {
	return "StreamPeerGZIP"
}

func NewStreamPeerGZIP() *StreamPeerGZIP {
	str := StringNameFromStr("StreamPeerGZIP") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &StreamPeerGZIP{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *StreamPeerGZIP) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *StreamPeerGZIP) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *StreamPeerGZIP) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *StreamPeerGZIP) StartCompression(use_deflate bool, buffer_size int64) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_deflate), gdc.ConstTypePtr(&buffer_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&use_deflate)
	pinner.Pin(&buffer_size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerGZIP.fnStartCompression), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *StreamPeerGZIP) StartDecompression(use_deflate bool, buffer_size int64) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_deflate), gdc.ConstTypePtr(&buffer_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&use_deflate)
	pinner.Pin(&buffer_size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerGZIP.fnStartDecompression), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *StreamPeerGZIP) Finish() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerGZIP.fnFinish), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *StreamPeerGZIP) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerGZIP.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
