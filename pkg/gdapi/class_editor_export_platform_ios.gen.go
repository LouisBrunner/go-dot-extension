// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorExportPlatformIOS struct {
  EditorExportPlatform
}

func (me *EditorExportPlatformIOS) BaseClass() string {
  return "EditorExportPlatformIOS"
}



// Enums

func (me *EditorExportPlatformIOS) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorExportPlatformIOS) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlatformIOS) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
