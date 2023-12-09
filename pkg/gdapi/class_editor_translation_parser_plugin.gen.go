// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *EditorTranslationParserPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorTranslationParserPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorTranslationParserPlugin) XParseFile(path String, msgids String, msgids_context_plural Array, )  {
  panic("TODO: implement")
}

func  (me *EditorTranslationParserPlugin) XGetRecognizedExtensions()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
