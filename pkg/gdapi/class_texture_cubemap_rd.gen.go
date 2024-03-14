// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextureCubemapRD struct {
  TextureLayeredRD
}

func (me *TextureCubemapRD) BaseClass() string {
  return "TextureCubemapRD"
}



// Enums

func (me *TextureCubemapRD) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextureCubemapRD) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextureCubemapRD) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
