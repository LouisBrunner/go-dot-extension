// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PackedDataContainerRef struct {
  RefCounted
}

func (me *PackedDataContainerRef) BaseClass() string {
  return "PackedDataContainerRef"
}

func NewPackedDataContainerRef() *PackedDataContainerRef {
  str := StringNameFromStr("PackedDataContainerRef") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PackedDataContainerRef{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PackedDataContainerRef) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PackedDataContainerRef) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PackedDataContainerRef) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PackedDataContainerRef) Size() int64 {
  classNameV := StringNameFromStr("PackedDataContainerRef")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
