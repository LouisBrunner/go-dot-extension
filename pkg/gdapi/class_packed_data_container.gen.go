// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PackedDataContainer struct {
  obj gdc.ObjectPtr
}

func (me *PackedDataContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PackedDataContainer) BaseClass() string {
  return "PackedDataContainer"
}



// Enums

func (me *PackedDataContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PackedDataContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PackedDataContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PackedDataContainer) Pack(value Variant, ) Error {
  classNameV := StringNameFromStr("PackedDataContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pack")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 966674026) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PackedDataContainer) Size() int {
  classNameV := StringNameFromStr("PackedDataContainer")
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