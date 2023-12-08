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

func  (me *ColorPicker) SetPickColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) GetPickColor() { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) SetDeferredMode(mode bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) IsDeferredMode() { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) SetColorMode(color_mode ColorPickerColorModeType, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) GetColorMode() { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) SetEditAlpha(show bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) IsEditingAlpha() { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) SetCanAddSwatches(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) AreSwatchesEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) SetPresetsVisible(visible bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) ArePresetsVisible() { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) SetModesVisible(visible bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) AreModesVisible() { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) SetSamplerVisible(visible bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) IsSamplerVisible() { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) SetSlidersVisible(visible bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) AreSlidersVisible() { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) SetHexVisible(visible bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) IsHexVisible() { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) AddPreset(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) ErasePreset(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) GetPresets() { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) AddRecentPreset(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) EraseRecentPreset(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) GetRecentPresets() { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) SetPickerShape(shape ColorPickerPickerShapeType, ) { // TODO: return value
  // TODO: implement
}

func  (me *ColorPicker) GetPickerShape() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
