// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type EditorExportPlatformLinuxBSD struct {
  EditorExportPlatformPC
}

func (me *EditorExportPlatformLinuxBSD) BaseClass() string {
  return "EditorExportPlatformLinuxBSD"
}

func NewEditorExportPlatformLinuxBSD() *EditorExportPlatformLinuxBSD {
  str := StringNameFromStr("EditorExportPlatformLinuxBSD") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorExportPlatformLinuxBSD{}
  obj.SetBaseObject(objPtr)
  return obj
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

// Signals
