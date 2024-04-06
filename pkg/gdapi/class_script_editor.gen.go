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

type ptrsForScriptEditorList struct {
	fnGetCurrentEditor            gdc.MethodBindPtr
	fnGetOpenScriptEditors        gdc.MethodBindPtr
	fnRegisterSyntaxHighlighter   gdc.MethodBindPtr
	fnUnregisterSyntaxHighlighter gdc.MethodBindPtr
	fnGotoLine                    gdc.MethodBindPtr
	fnGetCurrentScript            gdc.MethodBindPtr
	fnGetOpenScripts              gdc.MethodBindPtr
	fnOpenScriptCreateDialog      gdc.MethodBindPtr
}

var ptrsForScriptEditor ptrsForScriptEditorList

func initScriptEditorPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ScriptEditor")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_current_editor")
		defer methodName.Destroy()
		ptrsForScriptEditor.fnGetCurrentEditor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1906266726))
	}
	{
		methodName := StringNameFromStr("get_open_script_editors")
		defer methodName.Destroy()
		ptrsForScriptEditor.fnGetOpenScriptEditors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("register_syntax_highlighter")
		defer methodName.Destroy()
		ptrsForScriptEditor.fnRegisterSyntaxHighlighter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1092774468))
	}
	{
		methodName := StringNameFromStr("unregister_syntax_highlighter")
		defer methodName.Destroy()
		ptrsForScriptEditor.fnUnregisterSyntaxHighlighter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1092774468))
	}
	{
		methodName := StringNameFromStr("goto_line")
		defer methodName.Destroy()
		ptrsForScriptEditor.fnGotoLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_current_script")
		defer methodName.Destroy()
		ptrsForScriptEditor.fnGetCurrentScript = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2146468882))
	}
	{
		methodName := StringNameFromStr("get_open_scripts")
		defer methodName.Destroy()
		ptrsForScriptEditor.fnGetOpenScripts = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("open_script_create_dialog")
		defer methodName.Destroy()
		ptrsForScriptEditor.fnOpenScriptCreateDialog = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3186203200))
	}

}

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

func (me *ScriptEditor) GetCurrentEditor() ScriptEditorBase {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewScriptEditorBase()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScriptEditor.fnGetCurrentEditor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ScriptEditor) GetOpenScriptEditors() []ScriptEditorBase {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScriptEditor.fnGetOpenScriptEditors), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[ScriptEditorBase](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *ScriptEditor) RegisterSyntaxHighlighter(syntax_highlighter EditorSyntaxHighlighter) {
	cargs := []gdc.ConstTypePtr{syntax_highlighter.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScriptEditor.fnRegisterSyntaxHighlighter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ScriptEditor) UnregisterSyntaxHighlighter(syntax_highlighter EditorSyntaxHighlighter) {
	cargs := []gdc.ConstTypePtr{syntax_highlighter.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScriptEditor.fnUnregisterSyntaxHighlighter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ScriptEditor) GotoLine(line_number int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScriptEditor.fnGotoLine), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ScriptEditor) GetCurrentScript() Script {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewScript()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScriptEditor.fnGetCurrentScript), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ScriptEditor) GetOpenScripts() []Script {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScriptEditor.fnGetOpenScripts), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Script](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *ScriptEditor) OpenScriptCreateDialog(base_name String, base_path String) {
	cargs := []gdc.ConstTypePtr{base_name.AsCTypePtr(), base_path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScriptEditor.fnOpenScriptCreateDialog), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type ScriptEditorEditorScriptChangedSignalFn func(script Script)

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

type ScriptEditorScriptCloseSignalFn func(script Script)

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
