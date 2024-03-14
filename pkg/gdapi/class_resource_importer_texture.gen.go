// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporterTexture struct {
  ResourceImporter
}

func (me *ResourceImporterTexture) BaseClass() string {
  return "ResourceImporterTexture"
}



// Enums

func (me *ResourceImporterTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
