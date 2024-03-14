// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorSpinSlider struct {
  Range
}

func (me *EditorSpinSlider) BaseClass() string {
  return "EditorSpinSlider"
}



// Enums

func (me *EditorSpinSlider) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorSpinSlider) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorSpinSlider) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorSpinSlider) SetLabel(label String, )  {
  classNameV := StringNameFromStr("EditorSpinSlider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(label.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorSpinSlider) GetLabel() String {
  classNameV := StringNameFromStr("EditorSpinSlider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorSpinSlider) SetSuffix(suffix String, )  {
  classNameV := StringNameFromStr("EditorSpinSlider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_suffix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(suffix.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorSpinSlider) GetSuffix() String {
  classNameV := StringNameFromStr("EditorSpinSlider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_suffix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorSpinSlider) SetReadOnly(read_only bool, )  {
  classNameV := StringNameFromStr("EditorSpinSlider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_read_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&read_only), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorSpinSlider) IsReadOnly() bool {
  classNameV := StringNameFromStr("EditorSpinSlider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_read_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorSpinSlider) SetFlat(flat bool, )  {
  classNameV := StringNameFromStr("EditorSpinSlider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flat), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorSpinSlider) IsFlat() bool {
  classNameV := StringNameFromStr("EditorSpinSlider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_flat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorSpinSlider) SetHideSlider(hide_slider bool, )  {
  classNameV := StringNameFromStr("EditorSpinSlider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hide_slider")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hide_slider), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorSpinSlider) IsHidingSlider() bool {
  classNameV := StringNameFromStr("EditorSpinSlider")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_hiding_slider")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type EditorSpinSliderGrabbedSignalFn func()

func (me *EditorSpinSlider) ConnectGrabbed(subs SignalSubscribers, fn EditorSpinSliderGrabbedSignalFn) {
  sig := StringNameFromStr("grabbed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *EditorSpinSlider) DisconnectGrabbed(subs SignalSubscribers, fn EditorSpinSliderGrabbedSignalFn) {
  sig := StringNameFromStr("grabbed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type EditorSpinSliderUngrabbedSignalFn func()

func (me *EditorSpinSlider) ConnectUngrabbed(subs SignalSubscribers, fn EditorSpinSliderUngrabbedSignalFn) {
  sig := StringNameFromStr("ungrabbed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *EditorSpinSlider) DisconnectUngrabbed(subs SignalSubscribers, fn EditorSpinSliderUngrabbedSignalFn) {
  sig := StringNameFromStr("ungrabbed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type EditorSpinSliderValueFocusEnteredSignalFn func()

func (me *EditorSpinSlider) ConnectValueFocusEntered(subs SignalSubscribers, fn EditorSpinSliderValueFocusEnteredSignalFn) {
  sig := StringNameFromStr("value_focus_entered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *EditorSpinSlider) DisconnectValueFocusEntered(subs SignalSubscribers, fn EditorSpinSliderValueFocusEnteredSignalFn) {
  sig := StringNameFromStr("value_focus_entered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type EditorSpinSliderValueFocusExitedSignalFn func()

func (me *EditorSpinSlider) ConnectValueFocusExited(subs SignalSubscribers, fn EditorSpinSliderValueFocusExitedSignalFn) {
  sig := StringNameFromStr("value_focus_exited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *EditorSpinSlider) DisconnectValueFocusExited(subs SignalSubscribers, fn EditorSpinSliderValueFocusExitedSignalFn) {
  sig := StringNameFromStr("value_focus_exited")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
