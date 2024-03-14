// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorExportPlatformMacOS struct {
  EditorExportPlatform
}

func (me *EditorExportPlatformMacOS) BaseClass() string {
  return "EditorExportPlatformMacOS"
}



// Enums

func (me *EditorExportPlatformMacOS) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorExportPlatformMacOS) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlatformMacOS) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
