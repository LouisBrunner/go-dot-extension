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

type ptrsForStyleBoxTextureList struct {
	fnSetTexture          gdc.MethodBindPtr
	fnGetTexture          gdc.MethodBindPtr
	fnSetTextureMargin    gdc.MethodBindPtr
	fnSetTextureMarginAll gdc.MethodBindPtr
	fnGetTextureMargin    gdc.MethodBindPtr
	fnSetExpandMargin     gdc.MethodBindPtr
	fnSetExpandMarginAll  gdc.MethodBindPtr
	fnGetExpandMargin     gdc.MethodBindPtr
	fnSetRegionRect       gdc.MethodBindPtr
	fnGetRegionRect       gdc.MethodBindPtr
	fnSetDrawCenter       gdc.MethodBindPtr
	fnIsDrawCenterEnabled gdc.MethodBindPtr
	fnSetModulate         gdc.MethodBindPtr
	fnGetModulate         gdc.MethodBindPtr
	fnSetHAxisStretchMode gdc.MethodBindPtr
	fnGetHAxisStretchMode gdc.MethodBindPtr
	fnSetVAxisStretchMode gdc.MethodBindPtr
	fnGetVAxisStretchMode gdc.MethodBindPtr
}

var ptrsForStyleBoxTexture ptrsForStyleBoxTextureList

func initStyleBoxTexturePtrs(iface gdc.Interface) {

	className := StringNameFromStr("StyleBoxTexture")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_texture")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_texture_margin")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnSetTextureMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4290182280))
	}
	{
		methodName := StringNameFromStr("set_texture_margin_all")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnSetTextureMarginAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_texture_margin")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnGetTextureMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2869120046))
	}
	{
		methodName := StringNameFromStr("set_expand_margin")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnSetExpandMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4290182280))
	}
	{
		methodName := StringNameFromStr("set_expand_margin_all")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnSetExpandMarginAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_expand_margin")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnGetExpandMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2869120046))
	}
	{
		methodName := StringNameFromStr("set_region_rect")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnSetRegionRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2046264180))
	}
	{
		methodName := StringNameFromStr("get_region_rect")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnGetRegionRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
	}
	{
		methodName := StringNameFromStr("set_draw_center")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnSetDrawCenter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_draw_center_enabled")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnIsDrawCenterEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_modulate")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnSetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_modulate")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnGetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_h_axis_stretch_mode")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnSetHAxisStretchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2965538783))
	}
	{
		methodName := StringNameFromStr("get_h_axis_stretch_mode")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnGetHAxisStretchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3807744063))
	}
	{
		methodName := StringNameFromStr("set_v_axis_stretch_mode")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnSetVAxisStretchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2965538783))
	}
	{
		methodName := StringNameFromStr("get_v_axis_stretch_mode")
		defer methodName.Destroy()
		ptrsForStyleBoxTexture.fnGetVAxisStretchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3807744063))
	}
}

type StyleBoxTexture struct {
	StyleBox
}

func (me *StyleBoxTexture) BaseClass() string {
	return "StyleBoxTexture"
}

func NewStyleBoxTexture() *StyleBoxTexture {
	str := StringNameFromStr("StyleBoxTexture") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &StyleBoxTexture{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type StyleBoxTextureAxisStretchMode int

const (
	StyleBoxTextureAxisStretchModeAxisStretchModeStretch StyleBoxTextureAxisStretchMode = 0
	StyleBoxTextureAxisStretchModeAxisStretchModeTile    StyleBoxTextureAxisStretchMode = 1
	StyleBoxTextureAxisStretchModeAxisStretchModeTileFit StyleBoxTextureAxisStretchMode = 2
)

func (me *StyleBoxTexture) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *StyleBoxTexture) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *StyleBoxTexture) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *StyleBoxTexture) SetTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StyleBoxTexture) GetTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *StyleBoxTexture) SetTextureMargin(margin Side, size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnSetTextureMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StyleBoxTexture) SetTextureMarginAll(size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnSetTextureMarginAll), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StyleBoxTexture) GetTextureMargin(margin Side) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&margin)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnGetTextureMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *StyleBoxTexture) SetExpandMargin(margin Side, size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnSetExpandMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StyleBoxTexture) SetExpandMarginAll(size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnSetExpandMarginAll), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StyleBoxTexture) GetExpandMargin(margin Side) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&margin)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnGetExpandMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *StyleBoxTexture) SetRegionRect(region Rect2) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnSetRegionRect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StyleBoxTexture) GetRegionRect() Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnGetRegionRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *StyleBoxTexture) SetDrawCenter(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnSetDrawCenter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StyleBoxTexture) IsDrawCenterEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnIsDrawCenterEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *StyleBoxTexture) SetModulate(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnSetModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StyleBoxTexture) GetModulate() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnGetModulate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *StyleBoxTexture) SetHAxisStretchMode(mode StyleBoxTextureAxisStretchMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnSetHAxisStretchMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StyleBoxTexture) GetHAxisStretchMode() StyleBoxTextureAxisStretchMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret StyleBoxTextureAxisStretchMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnGetHAxisStretchMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *StyleBoxTexture) SetVAxisStretchMode(mode StyleBoxTextureAxisStretchMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnSetVAxisStretchMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StyleBoxTexture) GetVAxisStretchMode() StyleBoxTextureAxisStretchMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret StyleBoxTextureAxisStretchMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxTexture.fnGetVAxisStretchMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
