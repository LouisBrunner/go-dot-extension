// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MovieWriter struct {
  obj gdc.ObjectPtr
}

func (me *MovieWriter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MovieWriter) BaseClass() string {
  return "MovieWriter"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(writer.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), nil)
}

// Properties