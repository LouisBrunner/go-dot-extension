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

type ptrsForNinePatchRectList struct {
	fnSetTexture          gdc.MethodBindPtr
	fnGetTexture          gdc.MethodBindPtr
	fnSetPatchMargin      gdc.MethodBindPtr
	fnGetPatchMargin      gdc.MethodBindPtr
	fnSetRegionRect       gdc.MethodBindPtr
	fnGetRegionRect       gdc.MethodBindPtr
	fnSetDrawCenter       gdc.MethodBindPtr
	fnIsDrawCenterEnabled gdc.MethodBindPtr
	fnSetHAxisStretchMode gdc.MethodBindPtr
	fnGetHAxisStretchMode gdc.MethodBindPtr
	fnSetVAxisStretchMode gdc.MethodBindPtr
	fnGetVAxisStretchMode gdc.MethodBindPtr
}

var ptrsForNinePatchRect ptrsForNinePatchRectList

func initNinePatchRectPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NinePatchRect")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_texture")
		defer methodName.Destroy()
		ptrsForNinePatchRect.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForNinePatchRect.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_patch_margin")
		defer methodName.Destroy()
		ptrsForNinePatchRect.fnSetPatchMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 437707142))
	}
	{
		methodName := StringNameFromStr("get_patch_margin")
		defer methodName.Destroy()
		ptrsForNinePatchRect.fnGetPatchMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1983885014))
	}
	{
		methodName := StringNameFromStr("set_region_rect")
		defer methodName.Destroy()
		ptrsForNinePatchRect.fnSetRegionRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2046264180))
	}
	{
		methodName := StringNameFromStr("get_region_rect")
		defer methodName.Destroy()
		ptrsForNinePatchRect.fnGetRegionRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
	}
	{
		methodName := StringNameFromStr("set_draw_center")
		defer methodName.Destroy()
		ptrsForNinePatchRect.fnSetDrawCenter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_draw_center_enabled")
		defer methodName.Destroy()
		ptrsForNinePatchRect.fnIsDrawCenterEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_h_axis_stretch_mode")
		defer methodName.Destroy()
		ptrsForNinePatchRect.fnSetHAxisStretchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3219608417))
	}
	{
		methodName := StringNameFromStr("get_h_axis_stretch_mode")
		defer methodName.Destroy()
		ptrsForNinePatchRect.fnGetHAxisStretchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3317113799))
	}
	{
		methodName := StringNameFromStr("set_v_axis_stretch_mode")
		defer methodName.Destroy()
		ptrsForNinePatchRect.fnSetVAxisStretchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3219608417))
	}
	{
		methodName := StringNameFromStr("get_v_axis_stretch_mode")
		defer methodName.Destroy()
		ptrsForNinePatchRect.fnGetVAxisStretchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3317113799))
	}

}

type NinePatchRect struct {
	Control
}

func (me *NinePatchRect) BaseClass() string {
	return "NinePatchRect"
}

func NewNinePatchRect() *NinePatchRect {
	str := StringNameFromStr("NinePatchRect") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &NinePatchRect{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type NinePatchRectAxisStretchMode int

const (
	NinePatchRectAxisStretchModeAxisStretchModeStretch NinePatchRectAxisStretchMode = 0
	NinePatchRectAxisStretchModeAxisStretchModeTile    NinePatchRectAxisStretchMode = 1
	NinePatchRectAxisStretchModeAxisStretchModeTileFit NinePatchRectAxisStretchMode = 2
)

func (me *NinePatchRect) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *NinePatchRect) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *NinePatchRect) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *NinePatchRect) SetTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNinePatchRect.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NinePatchRect) GetTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNinePatchRect.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NinePatchRect) SetPatchMargin(margin Side, value int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNinePatchRect.fnSetPatchMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NinePatchRect) GetPatchMargin(margin Side) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&margin)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNinePatchRect.fnGetPatchMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NinePatchRect) SetRegionRect(rect Rect2) {
	cargs := []gdc.ConstTypePtr{rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNinePatchRect.fnSetRegionRect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NinePatchRect) GetRegionRect() Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNinePatchRect.fnGetRegionRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NinePatchRect) SetDrawCenter(draw_center bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_center)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNinePatchRect.fnSetDrawCenter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NinePatchRect) IsDrawCenterEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNinePatchRect.fnIsDrawCenterEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NinePatchRect) SetHAxisStretchMode(mode NinePatchRectAxisStretchMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNinePatchRect.fnSetHAxisStretchMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NinePatchRect) GetHAxisStretchMode() NinePatchRectAxisStretchMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret NinePatchRectAxisStretchMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNinePatchRect.fnGetHAxisStretchMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *NinePatchRect) SetVAxisStretchMode(mode NinePatchRectAxisStretchMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNinePatchRect.fnSetVAxisStretchMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NinePatchRect) GetVAxisStretchMode() NinePatchRectAxisStretchMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret NinePatchRectAxisStretchMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNinePatchRect.fnGetVAxisStretchMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type NinePatchRectTextureChangedSignalFn func()

func (me *NinePatchRect) ConnectTextureChanged(subs SignalSubscribers, fn NinePatchRectTextureChangedSignalFn) {
	sig := StringNameFromStr("texture_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *NinePatchRect) DisconnectTextureChanged(subs SignalSubscribers, fn NinePatchRectTextureChangedSignalFn) {
	sig := StringNameFromStr("texture_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
