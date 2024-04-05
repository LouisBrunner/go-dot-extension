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

type ptrsForEditorSpinSliderList struct {
  fnSetLabel gdc.MethodBindPtr
  fnGetLabel gdc.MethodBindPtr
  fnSetSuffix gdc.MethodBindPtr
  fnGetSuffix gdc.MethodBindPtr
  fnSetReadOnly gdc.MethodBindPtr
  fnIsReadOnly gdc.MethodBindPtr
  fnSetFlat gdc.MethodBindPtr
  fnIsFlat gdc.MethodBindPtr
  fnSetHideSlider gdc.MethodBindPtr
  fnIsHidingSlider gdc.MethodBindPtr
}

var ptrsForEditorSpinSlider ptrsForEditorSpinSliderList

func initEditorSpinSliderPtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorSpinSlider")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_label")
    defer methodName.Destroy()
    ptrsForEditorSpinSlider.fnSetLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_label")
    defer methodName.Destroy()
    ptrsForEditorSpinSlider.fnGetLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_suffix")
    defer methodName.Destroy()
    ptrsForEditorSpinSlider.fnSetSuffix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_suffix")
    defer methodName.Destroy()
    ptrsForEditorSpinSlider.fnGetSuffix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_read_only")
    defer methodName.Destroy()
    ptrsForEditorSpinSlider.fnSetReadOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_read_only")
    defer methodName.Destroy()
    ptrsForEditorSpinSlider.fnIsReadOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_flat")
    defer methodName.Destroy()
    ptrsForEditorSpinSlider.fnSetFlat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_flat")
    defer methodName.Destroy()
    ptrsForEditorSpinSlider.fnIsFlat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_hide_slider")
    defer methodName.Destroy()
    ptrsForEditorSpinSlider.fnSetHideSlider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_hiding_slider")
    defer methodName.Destroy()
    ptrsForEditorSpinSlider.fnIsHidingSlider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type EditorSpinSlider struct {
  Range
}

func (me *EditorSpinSlider) BaseClass() string {
  return "EditorSpinSlider"
}

func NewEditorSpinSlider() *EditorSpinSlider {
  str := StringNameFromStr("EditorSpinSlider") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorSpinSlider{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{label.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSpinSlider.fnSetLabel), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorSpinSlider) GetLabel() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSpinSlider.fnGetLabel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorSpinSlider) SetSuffix(suffix String, )  {
  cargs := []gdc.ConstTypePtr{suffix.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSpinSlider.fnSetSuffix), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorSpinSlider) GetSuffix() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSpinSlider.fnGetSuffix), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorSpinSlider) SetReadOnly(read_only bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&read_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSpinSlider.fnSetReadOnly), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorSpinSlider) IsReadOnly() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSpinSlider.fnIsReadOnly), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorSpinSlider) SetFlat(flat bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flat) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSpinSlider.fnSetFlat), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorSpinSlider) IsFlat() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSpinSlider.fnIsFlat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorSpinSlider) SetHideSlider(hide_slider bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hide_slider) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSpinSlider.fnSetHideSlider), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorSpinSlider) IsHidingSlider() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorSpinSlider.fnIsHidingSlider), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type EditorSpinSliderGrabbedSignalFn func()

func (me *EditorSpinSlider) ConnectGrabbed(subs SignalSubscribers, fn EditorSpinSliderGrabbedSignalFn) {
  sig := StringNameFromStr("grabbed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorSpinSlider) DisconnectGrabbed(subs SignalSubscribers, fn EditorSpinSliderGrabbedSignalFn) {
  sig := StringNameFromStr("grabbed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorSpinSliderUngrabbedSignalFn func()

func (me *EditorSpinSlider) ConnectUngrabbed(subs SignalSubscribers, fn EditorSpinSliderUngrabbedSignalFn) {
  sig := StringNameFromStr("ungrabbed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorSpinSlider) DisconnectUngrabbed(subs SignalSubscribers, fn EditorSpinSliderUngrabbedSignalFn) {
  sig := StringNameFromStr("ungrabbed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorSpinSliderValueFocusEnteredSignalFn func()

func (me *EditorSpinSlider) ConnectValueFocusEntered(subs SignalSubscribers, fn EditorSpinSliderValueFocusEnteredSignalFn) {
  sig := StringNameFromStr("value_focus_entered")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorSpinSlider) DisconnectValueFocusEntered(subs SignalSubscribers, fn EditorSpinSliderValueFocusEnteredSignalFn) {
  sig := StringNameFromStr("value_focus_entered")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorSpinSliderValueFocusExitedSignalFn func()

func (me *EditorSpinSlider) ConnectValueFocusExited(subs SignalSubscribers, fn EditorSpinSliderValueFocusExitedSignalFn) {
  sig := StringNameFromStr("value_focus_exited")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorSpinSlider) DisconnectValueFocusExited(subs SignalSubscribers, fn EditorSpinSliderValueFocusExitedSignalFn) {
  sig := StringNameFromStr("value_focus_exited")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
