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

type ptrsForColorPickerList struct {
	fnSetPickColor       gdc.MethodBindPtr
	fnGetPickColor       gdc.MethodBindPtr
	fnSetDeferredMode    gdc.MethodBindPtr
	fnIsDeferredMode     gdc.MethodBindPtr
	fnSetColorMode       gdc.MethodBindPtr
	fnGetColorMode       gdc.MethodBindPtr
	fnSetEditAlpha       gdc.MethodBindPtr
	fnIsEditingAlpha     gdc.MethodBindPtr
	fnSetCanAddSwatches  gdc.MethodBindPtr
	fnAreSwatchesEnabled gdc.MethodBindPtr
	fnSetPresetsVisible  gdc.MethodBindPtr
	fnArePresetsVisible  gdc.MethodBindPtr
	fnSetModesVisible    gdc.MethodBindPtr
	fnAreModesVisible    gdc.MethodBindPtr
	fnSetSamplerVisible  gdc.MethodBindPtr
	fnIsSamplerVisible   gdc.MethodBindPtr
	fnSetSlidersVisible  gdc.MethodBindPtr
	fnAreSlidersVisible  gdc.MethodBindPtr
	fnSetHexVisible      gdc.MethodBindPtr
	fnIsHexVisible       gdc.MethodBindPtr
	fnAddPreset          gdc.MethodBindPtr
	fnErasePreset        gdc.MethodBindPtr
	fnGetPresets         gdc.MethodBindPtr
	fnAddRecentPreset    gdc.MethodBindPtr
	fnEraseRecentPreset  gdc.MethodBindPtr
	fnGetRecentPresets   gdc.MethodBindPtr
	fnSetPickerShape     gdc.MethodBindPtr
	fnGetPickerShape     gdc.MethodBindPtr
}

var ptrsForColorPicker ptrsForColorPickerList

func initColorPickerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ColorPicker")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_pick_color")
		defer methodName.Destroy()
		ptrsForColorPicker.fnSetPickColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_pick_color")
		defer methodName.Destroy()
		ptrsForColorPicker.fnGetPickColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_deferred_mode")
		defer methodName.Destroy()
		ptrsForColorPicker.fnSetDeferredMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_deferred_mode")
		defer methodName.Destroy()
		ptrsForColorPicker.fnIsDeferredMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_color_mode")
		defer methodName.Destroy()
		ptrsForColorPicker.fnSetColorMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1579114136))
	}
	{
		methodName := StringNameFromStr("get_color_mode")
		defer methodName.Destroy()
		ptrsForColorPicker.fnGetColorMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 392907674))
	}
	{
		methodName := StringNameFromStr("set_edit_alpha")
		defer methodName.Destroy()
		ptrsForColorPicker.fnSetEditAlpha = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_editing_alpha")
		defer methodName.Destroy()
		ptrsForColorPicker.fnIsEditingAlpha = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_can_add_swatches")
		defer methodName.Destroy()
		ptrsForColorPicker.fnSetCanAddSwatches = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("are_swatches_enabled")
		defer methodName.Destroy()
		ptrsForColorPicker.fnAreSwatchesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_presets_visible")
		defer methodName.Destroy()
		ptrsForColorPicker.fnSetPresetsVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("are_presets_visible")
		defer methodName.Destroy()
		ptrsForColorPicker.fnArePresetsVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_modes_visible")
		defer methodName.Destroy()
		ptrsForColorPicker.fnSetModesVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("are_modes_visible")
		defer methodName.Destroy()
		ptrsForColorPicker.fnAreModesVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_sampler_visible")
		defer methodName.Destroy()
		ptrsForColorPicker.fnSetSamplerVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_sampler_visible")
		defer methodName.Destroy()
		ptrsForColorPicker.fnIsSamplerVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_sliders_visible")
		defer methodName.Destroy()
		ptrsForColorPicker.fnSetSlidersVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("are_sliders_visible")
		defer methodName.Destroy()
		ptrsForColorPicker.fnAreSlidersVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_hex_visible")
		defer methodName.Destroy()
		ptrsForColorPicker.fnSetHexVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_hex_visible")
		defer methodName.Destroy()
		ptrsForColorPicker.fnIsHexVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("add_preset")
		defer methodName.Destroy()
		ptrsForColorPicker.fnAddPreset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("erase_preset")
		defer methodName.Destroy()
		ptrsForColorPicker.fnErasePreset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_presets")
		defer methodName.Destroy()
		ptrsForColorPicker.fnGetPresets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1392750486))
	}
	{
		methodName := StringNameFromStr("add_recent_preset")
		defer methodName.Destroy()
		ptrsForColorPicker.fnAddRecentPreset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("erase_recent_preset")
		defer methodName.Destroy()
		ptrsForColorPicker.fnEraseRecentPreset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_recent_presets")
		defer methodName.Destroy()
		ptrsForColorPicker.fnGetRecentPresets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1392750486))
	}
	{
		methodName := StringNameFromStr("set_picker_shape")
		defer methodName.Destroy()
		ptrsForColorPicker.fnSetPickerShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3981373861))
	}
	{
		methodName := StringNameFromStr("get_picker_shape")
		defer methodName.Destroy()
		ptrsForColorPicker.fnGetPickerShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1143229889))
	}
}

type ColorPicker struct {
	VBoxContainer
}

func (me *ColorPicker) BaseClass() string {
	return "ColorPicker"
}

func NewColorPicker() *ColorPicker {
	str := StringNameFromStr("ColorPicker") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ColorPicker{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type ColorPickerColorModeType int

const (
	ColorPickerColorModeTypeModeRgb   ColorPickerColorModeType = 0
	ColorPickerColorModeTypeModeHsv   ColorPickerColorModeType = 1
	ColorPickerColorModeTypeModeRaw   ColorPickerColorModeType = 2
	ColorPickerColorModeTypeModeOkhsl ColorPickerColorModeType = 3
)

type ColorPickerPickerShapeType int

const (
	ColorPickerPickerShapeTypeShapeHsvRectangle ColorPickerPickerShapeType = 0
	ColorPickerPickerShapeTypeShapeHsvWheel     ColorPickerPickerShapeType = 1
	ColorPickerPickerShapeTypeShapeVhsCircle    ColorPickerPickerShapeType = 2
	ColorPickerPickerShapeTypeShapeOkhslCircle  ColorPickerPickerShapeType = 3
	ColorPickerPickerShapeTypeShapeNone         ColorPickerPickerShapeType = 4
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

func (me *ColorPicker) SetPickColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnSetPickColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) GetPickColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnGetPickColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ColorPicker) SetDeferredMode(mode bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnSetDeferredMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) IsDeferredMode() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnIsDeferredMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ColorPicker) SetColorMode(color_mode ColorPickerColorModeType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&color_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnSetColorMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) GetColorMode() ColorPickerColorModeType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ColorPickerColorModeType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnGetColorMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ColorPicker) SetEditAlpha(show bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&show)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnSetEditAlpha), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) IsEditingAlpha() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnIsEditingAlpha), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ColorPicker) SetCanAddSwatches(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnSetCanAddSwatches), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) AreSwatchesEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnAreSwatchesEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ColorPicker) SetPresetsVisible(visible bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnSetPresetsVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) ArePresetsVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnArePresetsVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ColorPicker) SetModesVisible(visible bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnSetModesVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) AreModesVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnAreModesVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ColorPicker) SetSamplerVisible(visible bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnSetSamplerVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) IsSamplerVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnIsSamplerVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ColorPicker) SetSlidersVisible(visible bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnSetSlidersVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) AreSlidersVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnAreSlidersVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ColorPicker) SetHexVisible(visible bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnSetHexVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) IsHexVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnIsHexVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ColorPicker) AddPreset(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnAddPreset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) ErasePreset(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnErasePreset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) GetPresets() PackedColorArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedColorArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnGetPresets), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ColorPicker) AddRecentPreset(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnAddRecentPreset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) EraseRecentPreset(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnEraseRecentPreset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) GetRecentPresets() PackedColorArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedColorArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnGetRecentPresets), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ColorPicker) SetPickerShape(shape ColorPickerPickerShapeType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnSetPickerShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ColorPicker) GetPickerShape() ColorPickerPickerShapeType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ColorPickerPickerShapeType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorPicker.fnGetPickerShape), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ColorPickerColorChangedSignalFn func(color Color)

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

type ColorPickerPresetAddedSignalFn func(color Color)

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

type ColorPickerPresetRemovedSignalFn func(color Color)

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
