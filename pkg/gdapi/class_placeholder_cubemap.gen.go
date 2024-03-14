// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PlaceholderCubemap struct {
  PlaceholderTextureLayered
}

func (me *PlaceholderCubemap) BaseClass() string {
  return "PlaceholderCubemap"
}



// Enums

func (me *PlaceholderCubemap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaceholderCubemap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaceholderCubemap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
