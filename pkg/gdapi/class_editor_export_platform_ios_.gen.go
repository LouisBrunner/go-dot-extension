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

type EditorExportPlatformIOS struct {
  EditorExportPlatform
}

func (me *EditorExportPlatformIOS) BaseClass() string {
  return "EditorExportPlatformIOS"
}

func NewEditorExportPlatformIOS() *EditorExportPlatformIOS {
  str := StringNameFromStr("EditorExportPlatformIOS") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorExportPlatformIOS{}
  obj.SetBaseObject(objPtr)
  return obj
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
