// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporterLayeredTexture struct {
  ResourceImporter
}

func (me *ResourceImporterLayeredTexture) BaseClass() string {
  return "ResourceImporterLayeredTexture"
}



// Enums

func (me *ResourceImporterLayeredTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterLayeredTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterLayeredTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
