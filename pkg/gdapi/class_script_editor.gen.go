// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  var ret ScriptEditorBase
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScriptEditor) GetOpenScriptEditors() ScriptEditorBase {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_open_script_editors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret ScriptEditorBase
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScriptEditor) RegisterSyntaxHighlighter(syntax_highlighter EditorSyntaxHighlighter, )  {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_syntax_highlighter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1092774468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(syntax_highlighter.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ScriptEditor) UnregisterSyntaxHighlighter(syntax_highlighter EditorSyntaxHighlighter, )  {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unregister_syntax_highlighter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1092774468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(syntax_highlighter.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ScriptEditor) GotoLine(line_number int, )  {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("goto_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ScriptEditor) GetCurrentScript() Script {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_script")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2146468882) // FIXME: should cache?
  var ret Script
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScriptEditor) GetOpenScripts() Script {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_open_scripts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Script
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ScriptEditor) OpenScriptCreateDialog(base_name String, base_path String, )  {
  classNameV := StringNameFromStr("ScriptEditor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("open_script_create_dialog")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3186203200) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(base_name.AsCTypePtr()), gdc.ConstTypePtr(base_path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals

type ScriptEditorEditorScriptChangedSignalFn func(script Script, )

func (me *ScriptEditor) ConnectEditorScriptChanged(subs SignalSubscribers, fn ScriptEditorEditorScriptChangedSignalFn) {
  sig := StringNameFromStr("editor_script_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *ScriptEditor) DisconnectEditorScriptChanged(subs SignalSubscribers, fn ScriptEditorEditorScriptChangedSignalFn) {
  sig := StringNameFromStr("editor_script_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type ScriptEditorScriptCloseSignalFn func(script Script, )

func (me *ScriptEditor) ConnectScriptClose(subs SignalSubscribers, fn ScriptEditorScriptCloseSignalFn) {
  sig := StringNameFromStr("script_close")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *ScriptEditor) DisconnectScriptClose(subs SignalSubscribers, fn ScriptEditorScriptCloseSignalFn) {
  sig := StringNameFromStr("script_close")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
