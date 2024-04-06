// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"log"
	"runtime"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForScriptEditorBaseList struct {
	fnGetBaseEditor        gdc.MethodBindPtr
	fnAddSyntaxHighlighter gdc.MethodBindPtr
}

var ptrsForScriptEditorBase ptrsForScriptEditorBaseList

func initScriptEditorBasePtrs(iface gdc.Interface) {

	className := StringNameFromStr("ScriptEditorBase")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_base_editor")
		defer methodName.Destroy()
		ptrsForScriptEditorBase.fnGetBaseEditor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2783021301))
	}
	{
		methodName := StringNameFromStr("add_syntax_highlighter")
		defer methodName.Destroy()
		ptrsForScriptEditorBase.fnAddSyntaxHighlighter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1092774468))
	}
}

type ScriptEditorBase struct {
	VBoxContainer
}

func (me *ScriptEditorBase) BaseClass() string {
	return "ScriptEditorBase"
}

func NewScriptEditorBase() *ScriptEditorBase {
	str := StringNameFromStr("ScriptEditorBase") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ScriptEditorBase{}
	obj.SetBaseObject(objPtr)
	return obj
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

func (me *ScriptEditorBase) GetBaseEditor() Control {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewControl()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScriptEditorBase.fnGetBaseEditor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ScriptEditorBase) AddSyntaxHighlighter(highlighter EditorSyntaxHighlighter) {
	cargs := []gdc.ConstTypePtr{highlighter.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScriptEditorBase.fnAddSyntaxHighlighter), me.obj, unsafe.SliceData(cargs), nil)

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

type ScriptEditorBaseRequestHelpSignalFn func(topic String)

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

type ScriptEditorBaseRequestOpenScriptAtLineSignalFn func(script Object, line int)

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

type ScriptEditorBaseGoToHelpSignalFn func(what String)

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

type ScriptEditorBaseSearchInFilesRequestedSignalFn func(text String)

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

type ScriptEditorBaseReplaceInFilesRequestedSignalFn func(text String)

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

type ScriptEditorBaseGoToMethodSignalFn func(script Object, method String)

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
