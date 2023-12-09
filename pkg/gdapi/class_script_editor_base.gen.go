// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScriptEditorBase struct {
  obj gdc.ObjectPtr
}

func (me *ScriptEditorBase) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ScriptEditorBase) BaseClass() string {
  return "ScriptEditorBase"
}



// Enums

func (me *ScriptEditorBase) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ScriptEditorBase) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ScriptEditorBase) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ScriptEditorBase) GetBaseEditor()  {
  panic("TODO: implement")
}

func  (me *ScriptEditorBase) AddSyntaxHighlighter(highlighter EditorSyntaxHighlighter, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
