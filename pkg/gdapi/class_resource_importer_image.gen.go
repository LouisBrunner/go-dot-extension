// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporterImage struct {
  ResourceImporter
}

func (me *ResourceImporterImage) BaseClass() string {
  return "ResourceImporterImage"
}



// Enums

func (me *ResourceImporterImage) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterImage) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterImage) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
