// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorExportPlatformPC struct {
  EditorExportPlatform
}

func (me *EditorExportPlatformPC) BaseClass() string {
  return "EditorExportPlatformPC"
}

func NewEditorExportPlatformPC() *EditorExportPlatformPC {
  str := StringNameFromStr("EditorExportPlatformPC") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorExportPlatformPC{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorExportPlatformPC) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorExportPlatformPC) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlatformPC) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
