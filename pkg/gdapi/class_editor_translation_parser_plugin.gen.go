// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorTranslationParserPlugin struct {
  obj gdc.ObjectPtr
}

func (me *EditorTranslationParserPlugin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorTranslationParserPlugin) BaseClass() string {
  return "EditorTranslationParserPlugin"
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
