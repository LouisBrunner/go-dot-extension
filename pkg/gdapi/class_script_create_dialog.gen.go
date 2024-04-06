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

type ptrsForScriptCreateDialogList struct {
	fnConfig gdc.MethodBindPtr
}

var ptrsForScriptCreateDialog ptrsForScriptCreateDialogList

func initScriptCreateDialogPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ScriptCreateDialog")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("config")
		defer methodName.Destroy()
		ptrsForScriptCreateDialog.fnConfig = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 869314288))
	}
}

type ScriptCreateDialog struct {
	ConfirmationDialog
}

func (me *ScriptCreateDialog) BaseClass() string {
	return "ScriptCreateDialog"
}

func NewScriptCreateDialog() *ScriptCreateDialog {
	str := StringNameFromStr("ScriptCreateDialog") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ScriptCreateDialog{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ScriptCreateDialog) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ScriptCreateDialog) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ScriptCreateDialog) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ScriptCreateDialog) Config(inherits String, path String, built_in_enabled bool, load_enabled bool) {
	cargs := []gdc.ConstTypePtr{inherits.AsCTypePtr(), path.AsCTypePtr(), gdc.ConstTypePtr(&built_in_enabled), gdc.ConstTypePtr(&load_enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScriptCreateDialog.fnConfig), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type ScriptCreateDialogScriptCreatedSignalFn func(script Script)

func (me *ScriptCreateDialog) ConnectScriptCreated(subs SignalSubscribers, fn ScriptCreateDialogScriptCreatedSignalFn) {
	sig := StringNameFromStr("script_created")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScriptCreateDialog) DisconnectScriptCreated(subs SignalSubscribers, fn ScriptCreateDialogScriptCreatedSignalFn) {
	sig := StringNameFromStr("script_created")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
