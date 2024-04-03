// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorSceneFormatImporterFBX struct {
  EditorSceneFormatImporter
}

func (me *EditorSceneFormatImporterFBX) BaseClass() string {
  return "EditorSceneFormatImporterFBX"
}

func NewEditorSceneFormatImporterFBX() *EditorSceneFormatImporterFBX {
  str := StringNameFromStr("EditorSceneFormatImporterFBX") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorSceneFormatImporterFBX{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorSceneFormatImporterFBX) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorSceneFormatImporterFBX) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorSceneFormatImporterFBX) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
