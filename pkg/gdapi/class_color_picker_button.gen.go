// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ColorPickerButton struct {
  Button
}

func (me *ColorPickerButton) BaseClass() string {
  return "ColorPickerButton"
}



// Enums

func (me *ColorPickerButton) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ColorPickerButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ColorPickerButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ColorPickerButton) SetPickColor(color Color, )  {
  classNameV := StringNameFromStr("ColorPickerButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pick_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPickerButton) GetPickColor() Color {
  classNameV := StringNameFromStr("ColorPickerButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pick_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPickerButton) GetPicker() ColorPicker {
  classNameV := StringNameFromStr("ColorPickerButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_picker")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 331835996) // FIXME: should cache?
  var ret ColorPicker
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPickerButton) GetPopup() PopupPanel {
  classNameV := StringNameFromStr("ColorPickerButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_popup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1322440207) // FIXME: should cache?
  var ret PopupPanel
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPickerButton) SetEditAlpha(show bool, )  {
  classNameV := StringNameFromStr("ColorPickerButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_edit_alpha")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&show), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPickerButton) IsEditingAlpha() bool {
  classNameV := StringNameFromStr("ColorPickerButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_editing_alpha")
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

type ColorPickerButtonColorChangedSignalFn func(color Color, )

func (me *ColorPickerButton) ConnectColorChanged(subs SignalSubscribers, fn ColorPickerButtonColorChangedSignalFn) {
  sig := StringNameFromStr("color_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *ColorPickerButton) DisconnectColorChanged(subs SignalSubscribers, fn ColorPickerButtonColorChangedSignalFn) {
  sig := StringNameFromStr("color_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type ColorPickerButtonPopupClosedSignalFn func()

func (me *ColorPickerButton) ConnectPopupClosed(subs SignalSubscribers, fn ColorPickerButtonPopupClosedSignalFn) {
  sig := StringNameFromStr("popup_closed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *ColorPickerButton) DisconnectPopupClosed(subs SignalSubscribers, fn ColorPickerButtonPopupClosedSignalFn) {
  sig := StringNameFromStr("popup_closed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type ColorPickerButtonPickerCreatedSignalFn func()

func (me *ColorPickerButton) ConnectPickerCreated(subs SignalSubscribers, fn ColorPickerButtonPickerCreatedSignalFn) {
  sig := StringNameFromStr("picker_created")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *ColorPickerButton) DisconnectPickerCreated(subs SignalSubscribers, fn ColorPickerButtonPickerCreatedSignalFn) {
  sig := StringNameFromStr("picker_created")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
