// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CompressedCubemapArray struct {
  obj gdc.ObjectPtr
}

func (me *CompressedCubemapArray) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CompressedCubemapArray) BaseClass() string {
  return "CompressedCubemapArray"
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

// Properties