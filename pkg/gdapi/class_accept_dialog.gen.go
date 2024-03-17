// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AcceptDialog struct {
  Window
}

func (me *AcceptDialog) BaseClass() string {
  return "AcceptDialog"
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

func  (me *AcceptDialog) GetOkButton() Button {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ok_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1856205918) // FIXME: should cache?
  var ret Button
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AcceptDialog) GetLabel() Label {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 566733104) // FIXME: should cache?
  var ret Label
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AcceptDialog) SetHideOnOk(enabled bool, )  {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hide_on_ok")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AcceptDialog) GetHideOnOk() bool {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hide_on_ok")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AcceptDialog) SetCloseOnEscape(enabled bool, )  {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_close_on_escape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AcceptDialog) GetCloseOnEscape() bool {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_close_on_escape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AcceptDialog) AddButton(text String, right bool, action String, ) Button {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3328440682) // FIXME: should cache?
  var ret Button
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(&right), gdc.ConstTypePtr(action.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AcceptDialog) AddCancelButton(name String, ) Button {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_cancel_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 242045556) // FIXME: should cache?
  var ret Button
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AcceptDialog) RemoveButton(button Control, )  {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1496901182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(button.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AcceptDialog) RegisterTextEnter(line_edit Control, )  {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_text_enter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1496901182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(line_edit.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AcceptDialog) SetText(text String, )  {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AcceptDialog) GetText() String {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AcceptDialog) SetAutowrap(autowrap bool, )  {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autowrap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&autowrap), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AcceptDialog) HasAutowrap() bool {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_autowrap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AcceptDialog) SetOkButtonText(text String, )  {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ok_button_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AcceptDialog) GetOkButtonText() String {
  classNameV := StringNameFromStr("AcceptDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ok_button_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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

type AcceptDialogCustomActionSignalFn func(action StringName, )

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
