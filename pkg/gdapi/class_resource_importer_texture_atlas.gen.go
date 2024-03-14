// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporterTextureAtlas struct {
  ResourceImporter
}

func (me *ResourceImporterTextureAtlas) BaseClass() string {
  return "ResourceImporterTextureAtlas"
}



// Enums

func (me *ResourceImporterTextureAtlas) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterTextureAtlas) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterTextureAtlas) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
