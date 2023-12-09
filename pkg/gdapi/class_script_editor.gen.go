// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScriptEditor struct {
  obj gdc.ObjectPtr
}

func (me *ScriptEditor) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ScriptEditor) BaseClass() string {
  return "ScriptEditor"
}



// Enums

func (me *ScriptEditor) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ScriptEditor) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ScriptEditor) GetCurrentEditor()  {
  panic("TODO: implement")
}

func  (me *ScriptEditor) GetOpenScriptEditors()  {
  panic("TODO: implement")
}

func  (me *ScriptEditor) RegisterSyntaxHighlighter(syntax_highlighter EditorSyntaxHighlighter, )  {
  panic("TODO: implement")
}

func  (me *ScriptEditor) UnregisterSyntaxHighlighter(syntax_highlighter EditorSyntaxHighlighter, )  {
  panic("TODO: implement")
}

func  (me *ScriptEditor) GotoLine(line_number int, )  {
  panic("TODO: implement")
}

func  (me *ScriptEditor) GetCurrentScript()  {
  panic("TODO: implement")
}

func  (me *ScriptEditor) GetOpenScripts()  {
  panic("TODO: implement")
}

func  (me *ScriptEditor) OpenScriptCreateDialog(base_name String, base_path String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
