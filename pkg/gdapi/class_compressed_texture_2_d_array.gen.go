// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CompressedTexture2DArray struct {
  obj gdc.ObjectPtr
}

func (me *CompressedTexture2DArray) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CompressedTexture2DArray) BaseClass() string {
  return "CompressedTexture2DArray"
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

// Properties