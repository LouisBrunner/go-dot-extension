// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ColorPicker struct {
  VBoxContainer
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

func (me *ColorPicker) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ColorPicker) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ColorPicker) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ColorPicker) SetPickColor(color Color, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pick_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) GetPickColor() Color {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pick_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPicker) SetDeferredMode(mode bool, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_deferred_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) IsDeferredMode() bool {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_deferred_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPicker) SetColorMode(color_mode ColorPickerColorModeType, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1579114136) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&color_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) GetColorMode() ColorPickerColorModeType {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 392907674) // FIXME: should cache?
  var ret ColorPickerColorModeType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPicker) SetEditAlpha(show bool, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_edit_alpha")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&show), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) IsEditingAlpha() bool {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_editing_alpha")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPicker) SetCanAddSwatches(enabled bool, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_can_add_swatches")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) AreSwatchesEnabled() bool {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("are_swatches_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPicker) SetPresetsVisible(visible bool, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_presets_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) ArePresetsVisible() bool {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("are_presets_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPicker) SetModesVisible(visible bool, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_modes_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) AreModesVisible() bool {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("are_modes_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPicker) SetSamplerVisible(visible bool, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sampler_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) IsSamplerVisible() bool {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_sampler_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPicker) SetSlidersVisible(visible bool, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sliders_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) AreSlidersVisible() bool {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("are_sliders_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPicker) SetHexVisible(visible bool, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hex_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) IsHexVisible() bool {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_hex_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPicker) AddPreset(color Color, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_preset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) ErasePreset(color Color, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase_preset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) GetPresets() PackedColorArray {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_presets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1392750486) // FIXME: should cache?
  var ret PackedColorArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPicker) AddRecentPreset(color Color, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_recent_preset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) EraseRecentPreset(color Color, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase_recent_preset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) GetRecentPresets() PackedColorArray {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_recent_presets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1392750486) // FIXME: should cache?
  var ret PackedColorArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ColorPicker) SetPickerShape(shape ColorPickerPickerShapeType, )  {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_picker_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3981373861) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ColorPicker) GetPickerShape() ColorPickerPickerShapeType {
  classNameV := StringNameFromStr("ColorPicker")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_picker_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1143229889) // FIXME: should cache?
  var ret ColorPickerPickerShapeType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ColorPickerColorChangedSignalFn func(color Color, )

func (me *ColorPicker) ConnectColorChanged(subs SignalSubscribers, fn ColorPickerColorChangedSignalFn) {
  sig := StringNameFromStr("color_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ColorPicker) DisconnectColorChanged(subs SignalSubscribers, fn ColorPickerColorChangedSignalFn) {
  sig := StringNameFromStr("color_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ColorPickerPresetAddedSignalFn func(color Color, )

func (me *ColorPicker) ConnectPresetAdded(subs SignalSubscribers, fn ColorPickerPresetAddedSignalFn) {
  sig := StringNameFromStr("preset_added")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ColorPicker) DisconnectPresetAdded(subs SignalSubscribers, fn ColorPickerPresetAddedSignalFn) {
  sig := StringNameFromStr("preset_added")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ColorPickerPresetRemovedSignalFn func(color Color, )

func (me *ColorPicker) ConnectPresetRemoved(subs SignalSubscribers, fn ColorPickerPresetRemovedSignalFn) {
  sig := StringNameFromStr("preset_removed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ColorPicker) DisconnectPresetRemoved(subs SignalSubscribers, fn ColorPickerPresetRemovedSignalFn) {
  sig := StringNameFromStr("preset_removed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
