// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CompressedCubemap struct {
  CompressedTextureLayered
}

func (me *CompressedCubemap) BaseClass() string {
  return "CompressedCubemap"
}

func NewCompressedCubemap() *CompressedCubemap {
  str := StringNameFromStr("CompressedCubemap") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CompressedCubemap{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CompressedCubemap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CompressedCubemap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CompressedCubemap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
