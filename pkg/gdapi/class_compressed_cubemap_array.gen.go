// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CompressedCubemapArray struct {
  CompressedTextureLayered
}

func (me *CompressedCubemapArray) BaseClass() string {
  return "CompressedCubemapArray"
}

func NewCompressedCubemapArray() *CompressedCubemapArray {
  str := StringNameFromStr("CompressedCubemapArray") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CompressedCubemapArray{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CompressedCubemapArray) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CompressedCubemapArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CompressedCubemapArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
