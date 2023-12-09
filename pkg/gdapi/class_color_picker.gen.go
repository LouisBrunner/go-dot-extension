// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ColorPicker struct {
  obj gdc.ObjectPtr
}

func (me *ColorPicker) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ColorPicker) BaseClass() string {
  return "ColorPicker"
}



// Enums

type ColorPickerColorModeType int
const (
  ColorPickerColorModeTypeModeRgb ColorPickerColorModeType = 0
  ColorPickerColorModeTypeModeHsv ColorPickerColorModeType = 1
  ColorPickerColorModeTypeModeRaw ColorPickerColorModeType = 2
  ColorPickerColorModeTypeModeOkhsl ColorPickerColorModeType = 3
)

type ColorPickerPickerShapeType int
const (
  ColorPickerPickerShapeTypeShapeHsvRectangle ColorPickerPickerShapeType = 0
  ColorPickerPickerShapeTypeShapeHsvWheel ColorPickerPickerShapeType = 1
  ColorPickerPickerShapeTypeShapeVhsCircle ColorPickerPickerShapeType = 2
  ColorPickerPickerShapeTypeShapeOkhslCircle ColorPickerPickerShapeType = 3
  ColorPickerPickerShapeTypeShapeNone ColorPickerPickerShapeType = 4
)

func (me *ColorPicker) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ColorPicker) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ColorPicker) SetPickColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) GetPickColor()  {
  panic("TODO: implement")
}

func  (me *ColorPicker) SetDeferredMode(mode bool, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) IsDeferredMode()  {
  panic("TODO: implement")
}

func  (me *ColorPicker) SetColorMode(color_mode ColorPickerColorModeType, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) GetColorMode()  {
  panic("TODO: implement")
}

func  (me *ColorPicker) SetEditAlpha(show bool, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) IsEditingAlpha()  {
  panic("TODO: implement")
}

func  (me *ColorPicker) SetCanAddSwatches(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) AreSwatchesEnabled()  {
  panic("TODO: implement")
}

func  (me *ColorPicker) SetPresetsVisible(visible bool, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) ArePresetsVisible()  {
  panic("TODO: implement")
}

func  (me *ColorPicker) SetModesVisible(visible bool, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) AreModesVisible()  {
  panic("TODO: implement")
}

func  (me *ColorPicker) SetSamplerVisible(visible bool, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) IsSamplerVisible()  {
  panic("TODO: implement")
}

func  (me *ColorPicker) SetSlidersVisible(visible bool, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) AreSlidersVisible()  {
  panic("TODO: implement")
}

func  (me *ColorPicker) SetHexVisible(visible bool, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) IsHexVisible()  {
  panic("TODO: implement")
}

func  (me *ColorPicker) AddPreset(color Color, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) ErasePreset(color Color, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) GetPresets()  {
  panic("TODO: implement")
}

func  (me *ColorPicker) AddRecentPreset(color Color, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) EraseRecentPreset(color Color, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) GetRecentPresets()  {
  panic("TODO: implement")
}

func  (me *ColorPicker) SetPickerShape(shape ColorPickerPickerShapeType, )  {
  panic("TODO: implement")
}

func  (me *ColorPicker) GetPickerShape()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
