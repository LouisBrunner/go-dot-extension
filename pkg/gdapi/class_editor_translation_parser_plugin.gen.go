// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorTranslationParserPlugin struct {
  RefCounted
}

func (me *EditorTranslationParserPlugin) BaseClass() string {
  return "EditorTranslationParserPlugin"
}

func NewEditorTranslationParserPlugin() *EditorTranslationParserPlugin {
  str := StringNameFromStr("EditorTranslationParserPlugin") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorTranslationParserPlugin{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorTranslationParserPlugin) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorTranslationParserPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorTranslationParserPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
