// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorExportPlatformAndroid struct {
  EditorExportPlatform
}

func (me *EditorExportPlatformAndroid) BaseClass() string {
  return "EditorExportPlatformAndroid"
}

func NewEditorExportPlatformAndroid() *EditorExportPlatformAndroid {
  str := StringNameFromStr("EditorExportPlatformAndroid") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorExportPlatformAndroid{}
  obj.SetBaseObject(objPtr)
  return obj
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

// Signals
