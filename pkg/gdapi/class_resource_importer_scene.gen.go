// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporterScene struct {
  ResourceImporter
}

func (me *ResourceImporterScene) BaseClass() string {
  return "ResourceImporterScene"
}



// Enums

func (me *ResourceImporterScene) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterScene) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterScene) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
