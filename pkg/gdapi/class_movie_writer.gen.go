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

type ptrsForMovieWriterList struct {
	fnXGetAudioMixRate     gdc.MethodBindPtr
	fnXGetAudioSpeakerMode gdc.MethodBindPtr
	fnXHandlesFile         gdc.MethodBindPtr
	fnXWriteBegin          gdc.MethodBindPtr
	fnXWriteFrame          gdc.MethodBindPtr
	fnXWriteEnd            gdc.MethodBindPtr
	fnAddWriter            gdc.MethodBindPtr
}

var ptrsForMovieWriter ptrsForMovieWriterList

func initMovieWriterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("MovieWriter")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_writer")
		defer methodName.Destroy()
		ptrsForMovieWriter.fnAddWriter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4023702871))
	}
}

type MovieWriter struct {
	Object
}

func (me *MovieWriter) BaseClass() string {
	return "MovieWriter"
}

func NewMovieWriter() *MovieWriter {
	str := StringNameFromStr("MovieWriter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &MovieWriter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *MovieWriter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *MovieWriter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *MovieWriter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func MovieWriterAddWriter(writer MovieWriter) {
	cargs := []gdc.ConstTypePtr{writer.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMovieWriter.fnAddWriter), nil, unsafe.SliceData(cargs), nil)

}

// Signals
