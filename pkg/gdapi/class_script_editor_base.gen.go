// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *ScriptEditorBase) GetBaseEditor() Control {
  classNameV := StringNameFromStr("ScriptEditorBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_base_editor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2783021301) // FIXME: should cache?
  var ret Control
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScriptEditorBase) AddSyntaxHighlighter(highlighter EditorSyntaxHighlighter, )  {
  classNameV := StringNameFromStr("ScriptEditorBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_syntax_highlighter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1092774468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(highlighter.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API