// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ColorPickerButton struct {
  obj gdc.ObjectPtr
}

func (me *ColorPickerButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

func (me *ColorPickerButton) GetPropColor() Color {
  panic("TODO: implement")
}

func (me *ColorPickerButton) SetPropColor(value Color) {
  panic("TODO: implement")
}

func (me *ColorPickerButton) GetPropEditAlpha() bool {
  panic("TODO: implement")
}

func (me *ColorPickerButton) SetPropEditAlpha(value bool) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API