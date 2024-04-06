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

type ptrsForParallaxBackgroundList struct {
	fnSetScrollOffset     gdc.MethodBindPtr
	fnGetScrollOffset     gdc.MethodBindPtr
	fnSetScrollBaseOffset gdc.MethodBindPtr
	fnGetScrollBaseOffset gdc.MethodBindPtr
	fnSetScrollBaseScale  gdc.MethodBindPtr
	fnGetScrollBaseScale  gdc.MethodBindPtr
	fnSetLimitBegin       gdc.MethodBindPtr
	fnGetLimitBegin       gdc.MethodBindPtr
	fnSetLimitEnd         gdc.MethodBindPtr
	fnGetLimitEnd         gdc.MethodBindPtr
	fnSetIgnoreCameraZoom gdc.MethodBindPtr
	fnIsIgnoreCameraZoom  gdc.MethodBindPtr
}

var ptrsForParallaxBackground ptrsForParallaxBackgroundList

func initParallaxBackgroundPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ParallaxBackground")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_scroll_offset")
		defer methodName.Destroy()
		ptrsForParallaxBackground.fnSetScrollOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_scroll_offset")
		defer methodName.Destroy()
		ptrsForParallaxBackground.fnGetScrollOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_scroll_base_offset")
		defer methodName.Destroy()
		ptrsForParallaxBackground.fnSetScrollBaseOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_scroll_base_offset")
		defer methodName.Destroy()
		ptrsForParallaxBackground.fnGetScrollBaseOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_scroll_base_scale")
		defer methodName.Destroy()
		ptrsForParallaxBackground.fnSetScrollBaseScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_scroll_base_scale")
		defer methodName.Destroy()
		ptrsForParallaxBackground.fnGetScrollBaseScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_limit_begin")
		defer methodName.Destroy()
		ptrsForParallaxBackground.fnSetLimitBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_limit_begin")
		defer methodName.Destroy()
		ptrsForParallaxBackground.fnGetLimitBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_limit_end")
		defer methodName.Destroy()
		ptrsForParallaxBackground.fnSetLimitEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_limit_end")
		defer methodName.Destroy()
		ptrsForParallaxBackground.fnGetLimitEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_ignore_camera_zoom")
		defer methodName.Destroy()
		ptrsForParallaxBackground.fnSetIgnoreCameraZoom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_ignore_camera_zoom")
		defer methodName.Destroy()
		ptrsForParallaxBackground.fnIsIgnoreCameraZoom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}

}

type ParallaxBackground struct {
	CanvasLayer
}

func (me *ParallaxBackground) BaseClass() string {
	return "ParallaxBackground"
}

func NewParallaxBackground() *ParallaxBackground {
	str := StringNameFromStr("ParallaxBackground") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ParallaxBackground{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ParallaxBackground) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ParallaxBackground) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ParallaxBackground) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ParallaxBackground) SetScrollOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxBackground.fnSetScrollOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParallaxBackground) GetScrollOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxBackground.fnGetScrollOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParallaxBackground) SetScrollBaseOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxBackground.fnSetScrollBaseOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParallaxBackground) GetScrollBaseOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxBackground.fnGetScrollBaseOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParallaxBackground) SetScrollBaseScale(scale Vector2) {
	cargs := []gdc.ConstTypePtr{scale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxBackground.fnSetScrollBaseScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParallaxBackground) GetScrollBaseScale() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxBackground.fnGetScrollBaseScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParallaxBackground) SetLimitBegin(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxBackground.fnSetLimitBegin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParallaxBackground) GetLimitBegin() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxBackground.fnGetLimitBegin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParallaxBackground) SetLimitEnd(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxBackground.fnSetLimitEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParallaxBackground) GetLimitEnd() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxBackground.fnGetLimitEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ParallaxBackground) SetIgnoreCameraZoom(ignore bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ignore)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxBackground.fnSetIgnoreCameraZoom), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ParallaxBackground) IsIgnoreCameraZoom() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallaxBackground.fnIsIgnoreCameraZoom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
