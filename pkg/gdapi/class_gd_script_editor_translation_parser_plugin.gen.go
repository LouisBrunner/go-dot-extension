// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GDScriptEditorTranslationParserPlugin struct {
  obj gdc.ObjectPtr
}

func (me *GDScriptEditorTranslationParserPlugin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GDScriptEditorTranslationParserPlugin) BaseClass() string {
  return "GDScriptEditorTranslationParserPlugin"
}



// Enums

func (me *GDScriptEditorTranslationParserPlugin) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GDScriptEditorTranslationParserPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GDScriptEditorTranslationParserPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
