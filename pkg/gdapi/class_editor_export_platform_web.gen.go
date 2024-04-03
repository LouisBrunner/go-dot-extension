// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorExportPlatformWeb struct {
  EditorExportPlatform
}

func (me *EditorExportPlatformWeb) BaseClass() string {
  return "EditorExportPlatformWeb"
}

func NewEditorExportPlatformWeb() *EditorExportPlatformWeb {
  str := StringNameFromStr("EditorExportPlatformWeb") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorExportPlatformWeb{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorExportPlatformWeb) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorExportPlatformWeb) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorExportPlatformWeb) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
