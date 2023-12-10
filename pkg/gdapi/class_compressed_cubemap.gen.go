// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CompressedCubemap struct {
  obj gdc.ObjectPtr
}

func (me *CompressedCubemap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CompressedCubemap) BaseClass() string {
  return "CompressedCubemap"
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

// Properties