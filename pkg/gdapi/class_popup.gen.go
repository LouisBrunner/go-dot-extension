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

type ptrsForPopupList struct {
}

var ptrsForPopup ptrsForPopupList

func initPopupPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Popup")
	defer className.Destroy()

}

type Popup struct {
	Window
}

func (me *Popup) BaseClass() string {
	return "Popup"
}

func NewPopup() *Popup {
	str := StringNameFromStr("Popup") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Popup{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Popup) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Popup) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Popup) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals

type PopupPopupHideSignalFn func()

func (me *Popup) ConnectPopupHide(subs SignalSubscribers, fn PopupPopupHideSignalFn) {
	sig := StringNameFromStr("popup_hide")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Popup) DisconnectPopupHide(subs SignalSubscribers, fn PopupPopupHideSignalFn) {
	sig := StringNameFromStr("popup_hide")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
