// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Popup struct {
  obj gdc.ObjectPtr
}

func (me *Popup) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Popup) BaseClass() string {
  return "Popup"
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
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Popup) DisconnectPopupHide(subs SignalSubscribers, fn PopupPopupHideSignalFn) {
  sig := StringNameFromStr("popup_hide")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
