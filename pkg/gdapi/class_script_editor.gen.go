// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type ScriptEditor struct {
  PanelContainer
}

func (me *ScriptEditor) BaseClass() string {
  return "ScriptEditor"
}

func NewScriptEditor() *ScriptEditor {
  str := StringNameFromStr("ScriptEditor") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ScriptEditor{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ScriptEditor) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ScriptEditor) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ScriptEditor) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ScriptEditor) GetCurrentEditor() ScriptEditorBase {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_editor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1906266726) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewScriptEditorBase()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ScriptEditor) GetOpenScriptEditors() []ScriptEditorBase {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_open_script_editors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[ScriptEditorBase](ret)
}

func  (me *ScriptEditor) RegisterSyntaxHighlighter(syntax_highlighter EditorSyntaxHighlighter, )  {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_syntax_highlighter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1092774468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{syntax_highlighter.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ScriptEditor) UnregisterSyntaxHighlighter(syntax_highlighter EditorSyntaxHighlighter, )  {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unregister_syntax_highlighter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1092774468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{syntax_highlighter.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ScriptEditor) GotoLine(line_number int64, )  {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("goto_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ScriptEditor) GetCurrentScript() Script {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_script")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2146468882) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewScript()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ScriptEditor) GetOpenScripts() []Script {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_open_scripts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Script](ret)
}

func  (me *ScriptEditor) OpenScriptCreateDialog(base_name String, base_path String, )  {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open_script_create_dialog")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3186203200) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{base_name.AsCTypePtr(), base_path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type ScriptEditorEditorScriptChangedSignalFn func(script Script, )

func (me *ScriptEditor) ConnectEditorScriptChanged(subs SignalSubscribers, fn ScriptEditorEditorScriptChangedSignalFn) {
  sig := StringNameFromStr("editor_script_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScriptEditor) DisconnectEditorScriptChanged(subs SignalSubscribers, fn ScriptEditorEditorScriptChangedSignalFn) {
  sig := StringNameFromStr("editor_script_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ScriptEditorScriptCloseSignalFn func(script Script, )

func (me *ScriptEditor) ConnectScriptClose(subs SignalSubscribers, fn ScriptEditorScriptCloseSignalFn) {
  sig := StringNameFromStr("script_close")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScriptEditor) DisconnectScriptClose(subs SignalSubscribers, fn ScriptEditorScriptCloseSignalFn) {
  sig := StringNameFromStr("script_close")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
