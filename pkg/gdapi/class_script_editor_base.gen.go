// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ScriptEditorBase struct {
  VBoxContainer
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

// Signals

type ScriptEditorBaseNameChangedSignalFn func()

func (me *ScriptEditorBase) ConnectNameChanged(subs SignalSubscribers, fn ScriptEditorBaseNameChangedSignalFn) {
  sig := StringNameFromStr("name_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScriptEditorBase) DisconnectNameChanged(subs SignalSubscribers, fn ScriptEditorBaseNameChangedSignalFn) {
  sig := StringNameFromStr("name_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ScriptEditorBaseEditedScriptChangedSignalFn func()

func (me *ScriptEditorBase) ConnectEditedScriptChanged(subs SignalSubscribers, fn ScriptEditorBaseEditedScriptChangedSignalFn) {
  sig := StringNameFromStr("edited_script_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScriptEditorBase) DisconnectEditedScriptChanged(subs SignalSubscribers, fn ScriptEditorBaseEditedScriptChangedSignalFn) {
  sig := StringNameFromStr("edited_script_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ScriptEditorBaseRequestHelpSignalFn func(topic String, )

func (me *ScriptEditorBase) ConnectRequestHelp(subs SignalSubscribers, fn ScriptEditorBaseRequestHelpSignalFn) {
  sig := StringNameFromStr("request_help")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScriptEditorBase) DisconnectRequestHelp(subs SignalSubscribers, fn ScriptEditorBaseRequestHelpSignalFn) {
  sig := StringNameFromStr("request_help")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ScriptEditorBaseRequestOpenScriptAtLineSignalFn func(script Object, line int, )

func (me *ScriptEditorBase) ConnectRequestOpenScriptAtLine(subs SignalSubscribers, fn ScriptEditorBaseRequestOpenScriptAtLineSignalFn) {
  sig := StringNameFromStr("request_open_script_at_line")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScriptEditorBase) DisconnectRequestOpenScriptAtLine(subs SignalSubscribers, fn ScriptEditorBaseRequestOpenScriptAtLineSignalFn) {
  sig := StringNameFromStr("request_open_script_at_line")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ScriptEditorBaseRequestSaveHistorySignalFn func()

func (me *ScriptEditorBase) ConnectRequestSaveHistory(subs SignalSubscribers, fn ScriptEditorBaseRequestSaveHistorySignalFn) {
  sig := StringNameFromStr("request_save_history")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScriptEditorBase) DisconnectRequestSaveHistory(subs SignalSubscribers, fn ScriptEditorBaseRequestSaveHistorySignalFn) {
  sig := StringNameFromStr("request_save_history")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ScriptEditorBaseGoToHelpSignalFn func(what String, )

func (me *ScriptEditorBase) ConnectGoToHelp(subs SignalSubscribers, fn ScriptEditorBaseGoToHelpSignalFn) {
  sig := StringNameFromStr("go_to_help")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScriptEditorBase) DisconnectGoToHelp(subs SignalSubscribers, fn ScriptEditorBaseGoToHelpSignalFn) {
  sig := StringNameFromStr("go_to_help")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ScriptEditorBaseSearchInFilesRequestedSignalFn func(text String, )

func (me *ScriptEditorBase) ConnectSearchInFilesRequested(subs SignalSubscribers, fn ScriptEditorBaseSearchInFilesRequestedSignalFn) {
  sig := StringNameFromStr("search_in_files_requested")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScriptEditorBase) DisconnectSearchInFilesRequested(subs SignalSubscribers, fn ScriptEditorBaseSearchInFilesRequestedSignalFn) {
  sig := StringNameFromStr("search_in_files_requested")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ScriptEditorBaseReplaceInFilesRequestedSignalFn func(text String, )

func (me *ScriptEditorBase) ConnectReplaceInFilesRequested(subs SignalSubscribers, fn ScriptEditorBaseReplaceInFilesRequestedSignalFn) {
  sig := StringNameFromStr("replace_in_files_requested")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScriptEditorBase) DisconnectReplaceInFilesRequested(subs SignalSubscribers, fn ScriptEditorBaseReplaceInFilesRequestedSignalFn) {
  sig := StringNameFromStr("replace_in_files_requested")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ScriptEditorBaseGoToMethodSignalFn func(script Object, method String, )

func (me *ScriptEditorBase) ConnectGoToMethod(subs SignalSubscribers, fn ScriptEditorBaseGoToMethodSignalFn) {
  sig := StringNameFromStr("go_to_method")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScriptEditorBase) DisconnectGoToMethod(subs SignalSubscribers, fn ScriptEditorBaseGoToMethodSignalFn) {
  sig := StringNameFromStr("go_to_method")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
