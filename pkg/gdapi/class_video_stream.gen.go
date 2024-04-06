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

type ptrsForVideoStreamList struct {
	fnXInstantiatePlayback gdc.MethodBindPtr
	fnSetFile              gdc.MethodBindPtr
	fnGetFile              gdc.MethodBindPtr
}

var ptrsForVideoStream ptrsForVideoStreamList

func initVideoStreamPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VideoStream")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_file")
		defer methodName.Destroy()
		ptrsForVideoStream.fnSetFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_file")
		defer methodName.Destroy()
		ptrsForVideoStream.fnGetFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
	}
}

type VideoStream struct {
	Resource
}

func (me *VideoStream) BaseClass() string {
	return "VideoStream"
}

func NewVideoStream() *VideoStream {
	str := StringNameFromStr("VideoStream") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VideoStream{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VideoStream) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VideoStream) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VideoStream) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VideoStream) SetFile(file String) {
	cargs := []gdc.ConstTypePtr{file.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStream.fnSetFile), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VideoStream) GetFile() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStream.fnGetFile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
