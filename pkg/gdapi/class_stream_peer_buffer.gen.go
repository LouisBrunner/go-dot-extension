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

type ptrsForStreamPeerBufferList struct {
	fnSeek         gdc.MethodBindPtr
	fnGetSize      gdc.MethodBindPtr
	fnGetPosition  gdc.MethodBindPtr
	fnResize       gdc.MethodBindPtr
	fnSetDataArray gdc.MethodBindPtr
	fnGetDataArray gdc.MethodBindPtr
	fnClear        gdc.MethodBindPtr
	fnDuplicate    gdc.MethodBindPtr
}

var ptrsForStreamPeerBuffer ptrsForStreamPeerBufferList

func initStreamPeerBufferPtrs(iface gdc.Interface) {

	className := StringNameFromStr("StreamPeerBuffer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("seek")
		defer methodName.Destroy()
		ptrsForStreamPeerBuffer.fnSeek = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForStreamPeerBuffer.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_position")
		defer methodName.Destroy()
		ptrsForStreamPeerBuffer.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("resize")
		defer methodName.Destroy()
		ptrsForStreamPeerBuffer.fnResize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_data_array")
		defer methodName.Destroy()
		ptrsForStreamPeerBuffer.fnSetDataArray = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2971499966))
	}
	{
		methodName := StringNameFromStr("get_data_array")
		defer methodName.Destroy()
		ptrsForStreamPeerBuffer.fnGetDataArray = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2362200018))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForStreamPeerBuffer.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("duplicate")
		defer methodName.Destroy()
		ptrsForStreamPeerBuffer.fnDuplicate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2474064677))
	}
}

type StreamPeerBuffer struct {
	StreamPeer
}

func (me *StreamPeerBuffer) BaseClass() string {
	return "StreamPeerBuffer"
}

func NewStreamPeerBuffer() *StreamPeerBuffer {
	str := StringNameFromStr("StreamPeerBuffer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &StreamPeerBuffer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *StreamPeerBuffer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *StreamPeerBuffer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *StreamPeerBuffer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *StreamPeerBuffer) Seek(position int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerBuffer.fnSeek), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StreamPeerBuffer) GetSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerBuffer.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *StreamPeerBuffer) GetPosition() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerBuffer.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *StreamPeerBuffer) Resize(size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerBuffer.fnResize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StreamPeerBuffer) SetDataArray(data PackedByteArray) {
	cargs := []gdc.ConstTypePtr{data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerBuffer.fnSetDataArray), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StreamPeerBuffer) GetDataArray() PackedByteArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerBuffer.fnGetDataArray), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *StreamPeerBuffer) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerBuffer.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StreamPeerBuffer) Duplicate() StreamPeerBuffer {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStreamPeerBuffer()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeerBuffer.fnDuplicate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
