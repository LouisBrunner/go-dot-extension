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

type ptrsForAcceptDialogList struct {
	fnGetOkButton       gdc.MethodBindPtr
	fnGetLabel          gdc.MethodBindPtr
	fnSetHideOnOk       gdc.MethodBindPtr
	fnGetHideOnOk       gdc.MethodBindPtr
	fnSetCloseOnEscape  gdc.MethodBindPtr
	fnGetCloseOnEscape  gdc.MethodBindPtr
	fnAddButton         gdc.MethodBindPtr
	fnAddCancelButton   gdc.MethodBindPtr
	fnRemoveButton      gdc.MethodBindPtr
	fnRegisterTextEnter gdc.MethodBindPtr
	fnSetText           gdc.MethodBindPtr
	fnGetText           gdc.MethodBindPtr
	fnSetAutowrap       gdc.MethodBindPtr
	fnHasAutowrap       gdc.MethodBindPtr
	fnSetOkButtonText   gdc.MethodBindPtr
	fnGetOkButtonText   gdc.MethodBindPtr
}

var ptrsForAcceptDialog ptrsForAcceptDialogList

func initAcceptDialogPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AcceptDialog")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_ok_button")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnGetOkButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1856205918))
	}
	{
		methodName := StringNameFromStr("get_label")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnGetLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 566733104))
	}
	{
		methodName := StringNameFromStr("set_hide_on_ok")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnSetHideOnOk = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_hide_on_ok")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnGetHideOnOk = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_close_on_escape")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnSetCloseOnEscape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_close_on_escape")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnGetCloseOnEscape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("add_button")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnAddButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3328440682))
	}
	{
		methodName := StringNameFromStr("add_cancel_button")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnAddCancelButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 242045556))
	}
	{
		methodName := StringNameFromStr("remove_button")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnRemoveButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1496901182))
	}
	{
		methodName := StringNameFromStr("register_text_enter")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnRegisterTextEnter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1496901182))
	}
	{
		methodName := StringNameFromStr("set_text")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnSetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_text")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnGetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_autowrap")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnSetAutowrap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("has_autowrap")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnHasAutowrap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_ok_button_text")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnSetOkButtonText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_ok_button_text")
		defer methodName.Destroy()
		ptrsForAcceptDialog.fnGetOkButtonText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
}

type AcceptDialog struct {
	Window
}

func (me *AcceptDialog) BaseClass() string {
	return "AcceptDialog"
}

func NewAcceptDialog() *AcceptDialog {
	str := StringNameFromStr("AcceptDialog") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AcceptDialog{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AcceptDialog) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AcceptDialog) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AcceptDialog) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AcceptDialog) GetOkButton() Button {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewButton()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnGetOkButton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AcceptDialog) GetLabel() Label {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewLabel()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnGetLabel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AcceptDialog) SetHideOnOk(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnSetHideOnOk), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AcceptDialog) GetHideOnOk() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnGetHideOnOk), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AcceptDialog) SetCloseOnEscape(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnSetCloseOnEscape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AcceptDialog) GetCloseOnEscape() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnGetCloseOnEscape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AcceptDialog) AddButton(text String, right bool, action String) Button {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), gdc.ConstTypePtr(&right), action.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewButton()
	pinner.Pin(&right)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnAddButton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AcceptDialog) AddCancelButton(name String) Button {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewButton()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnAddCancelButton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AcceptDialog) RemoveButton(button Control) {
	cargs := []gdc.ConstTypePtr{button.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnRemoveButton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AcceptDialog) RegisterTextEnter(line_edit Control) {
	cargs := []gdc.ConstTypePtr{line_edit.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnRegisterTextEnter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AcceptDialog) SetText(text String) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnSetText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AcceptDialog) GetText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnGetText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AcceptDialog) SetAutowrap(autowrap bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&autowrap)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnSetAutowrap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AcceptDialog) HasAutowrap() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnHasAutowrap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AcceptDialog) SetOkButtonText(text String) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnSetOkButtonText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AcceptDialog) GetOkButtonText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAcceptDialog.fnGetOkButtonText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AcceptDialogConfirmedSignalFn func()

func (me *AcceptDialog) ConnectConfirmed(subs SignalSubscribers, fn AcceptDialogConfirmedSignalFn) {
	sig := StringNameFromStr("confirmed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AcceptDialog) DisconnectConfirmed(subs SignalSubscribers, fn AcceptDialogConfirmedSignalFn) {
	sig := StringNameFromStr("confirmed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AcceptDialogCanceledSignalFn func()

func (me *AcceptDialog) ConnectCanceled(subs SignalSubscribers, fn AcceptDialogCanceledSignalFn) {
	sig := StringNameFromStr("canceled")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AcceptDialog) DisconnectCanceled(subs SignalSubscribers, fn AcceptDialogCanceledSignalFn) {
	sig := StringNameFromStr("canceled")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AcceptDialogCustomActionSignalFn func(action StringName)

func (me *AcceptDialog) ConnectCustomAction(subs SignalSubscribers, fn AcceptDialogCustomActionSignalFn) {
	sig := StringNameFromStr("custom_action")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AcceptDialog) DisconnectCustomAction(subs SignalSubscribers, fn AcceptDialogCustomActionSignalFn) {
	sig := StringNameFromStr("custom_action")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
