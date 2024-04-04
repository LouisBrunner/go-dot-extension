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

type CompressedTexture2DArray struct {
  CompressedTextureLayered
}

func (me *CompressedTexture2DArray) BaseClass() string {
  return "CompressedTexture2DArray"
}

func NewCompressedTexture2DArray() *CompressedTexture2DArray {
  str := StringNameFromStr("CompressedTexture2DArray") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CompressedTexture2DArray{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CompressedTexture2DArray) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CompressedTexture2DArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CompressedTexture2DArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
