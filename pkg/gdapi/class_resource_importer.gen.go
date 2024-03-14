// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporter struct {
  RefCounted
}

func (me *ResourceImporter) BaseClass() string {
  return "ResourceImporter"
}



// Enums

type ResourceImporterImportOrder int
const (
  ResourceImporterImportOrderImportOrderDefault ResourceImporterImportOrder = 0
  ResourceImporterImportOrderImportOrderScene ResourceImporterImportOrder = 100
)

func (me *ResourceImporter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
