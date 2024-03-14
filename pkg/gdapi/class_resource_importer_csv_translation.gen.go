// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporterCSVTranslation struct {
  ResourceImporter
}

func (me *ResourceImporterCSVTranslation) BaseClass() string {
  return "ResourceImporterCSVTranslation"
}



// Enums

func (me *ResourceImporterCSVTranslation) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterCSVTranslation) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterCSVTranslation) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
