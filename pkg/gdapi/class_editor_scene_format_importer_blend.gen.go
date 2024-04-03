// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorSceneFormatImporterBlend struct {
  EditorSceneFormatImporter
}

func (me *EditorSceneFormatImporterBlend) BaseClass() string {
  return "EditorSceneFormatImporterBlend"
}

func NewEditorSceneFormatImporterBlend() *EditorSceneFormatImporterBlend {
  str := StringNameFromStr("EditorSceneFormatImporterBlend") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorSceneFormatImporterBlend{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorSceneFormatImporterBlend) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorSceneFormatImporterBlend) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorSceneFormatImporterBlend) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
