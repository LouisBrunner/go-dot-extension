// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporterWAV struct {
  ResourceImporter
}

func (me *ResourceImporterWAV) BaseClass() string {
  return "ResourceImporterWAV"
}



// Enums

func (me *ResourceImporterWAV) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterWAV) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterWAV) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
