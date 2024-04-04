// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type ColorPickerButton struct {
  Button
}

func (me *ColorPickerButton) BaseClass() string {
  return "ColorPickerButton"
}

func NewColorPickerButton() *ColorPickerButton {
  str := StringNameFromStr("ColorPickerButton") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ColorPickerButton{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ColorPickerButton) GetPickColor() Color {
  classNameV := StringNameFromStr("ColorPickerButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pick_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ColorPickerButton) GetPicker() ColorPicker {
  classNameV := StringNameFromStr("ColorPickerButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_picker")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 331835996) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColorPicker()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ColorPickerButton) GetPopup() PopupPanel {
  classNameV := StringNameFromStr("ColorPickerButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_popup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1322440207) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPopupPanel()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ColorPickerButton) SetEditAlpha(show bool, )  {
  classNameV := StringNameFromStr("ColorPickerButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_edit_alpha")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&show) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ColorPickerButton) IsEditingAlpha() bool {
  classNameV := StringNameFromStr("ColorPickerButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_editing_alpha")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ColorPickerButtonColorChangedSignalFn func(color Color, )

func (me *ColorPickerButton) ConnectColorChanged(subs SignalSubscribers, fn ColorPickerButtonColorChangedSignalFn) {
  sig := StringNameFromStr("color_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ColorPickerButton) DisconnectColorChanged(subs SignalSubscribers, fn ColorPickerButtonColorChangedSignalFn) {
  sig := StringNameFromStr("color_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ColorPickerButtonPopupClosedSignalFn func()

func (me *ColorPickerButton) ConnectPopupClosed(subs SignalSubscribers, fn ColorPickerButtonPopupClosedSignalFn) {
  sig := StringNameFromStr("popup_closed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ColorPickerButton) DisconnectPopupClosed(subs SignalSubscribers, fn ColorPickerButtonPopupClosedSignalFn) {
  sig := StringNameFromStr("popup_closed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ColorPickerButtonPickerCreatedSignalFn func()

func (me *ColorPickerButton) ConnectPickerCreated(subs SignalSubscribers, fn ColorPickerButtonPickerCreatedSignalFn) {
  sig := StringNameFromStr("picker_created")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ColorPickerButton) DisconnectPickerCreated(subs SignalSubscribers, fn ColorPickerButtonPickerCreatedSignalFn) {
  sig := StringNameFromStr("picker_created")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
