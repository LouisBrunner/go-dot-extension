// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorExportPlatformLinuxBSD struct {
  obj gdc.ObjectPtr
}

func (me *EditorExportPlatformLinuxBSD) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorExportPlatformLinuxBSD) BaseClass() string {
  return "EditorExportPlatformLinuxBSD"
}



// Enums

func (me *EditorExportPlatformLinuxBSD) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorExportPlatformLinuxBSD) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlatformLinuxBSD) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties