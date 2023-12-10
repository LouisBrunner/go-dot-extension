// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorExportPlatform struct {
  obj gdc.ObjectPtr
}

func (me *EditorExportPlatform) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorExportPlatform) BaseClass() string {
  return "EditorExportPlatform"
}



// Enums

func (me *EditorExportPlatform) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorExportPlatform) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlatform) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties