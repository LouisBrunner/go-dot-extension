// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorExportPlatformAndroid struct {
  obj gdc.ObjectPtr
}

func (me *EditorExportPlatformAndroid) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorExportPlatformAndroid) BaseClass() string {
  return "EditorExportPlatformAndroid"
}



// Enums

func (me *EditorExportPlatformAndroid) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorExportPlatformAndroid) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlatformAndroid) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties