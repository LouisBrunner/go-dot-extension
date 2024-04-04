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

func  MovieWriterAddWriter(writer MovieWriter, )  {
  classNameV := StringNameFromStr("MovieWriter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_writer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4023702871) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{writer.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), nil)

}

// Signals
