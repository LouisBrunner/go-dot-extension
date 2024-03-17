// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorExportPlatformWindows struct {
  EditorExportPlatformPC
}

func (me *EditorExportPlatformWindows) BaseClass() string {
  return "EditorExportPlatformWindows"
}



// Enums

func (me *EditorExportPlatformWindows) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorExportPlatformWindows) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlatformWindows) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
