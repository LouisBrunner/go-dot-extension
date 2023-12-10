// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PackedDataContainerRef struct {
  obj gdc.ObjectPtr
}

func (me *PackedDataContainerRef) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PackedDataContainerRef) BaseClass() string {
  return "PackedDataContainerRef"
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

func  (me *PackedDataContainerRef) Size() int {
  classNameV := StringNameFromStr("PackedDataContainerRef")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties