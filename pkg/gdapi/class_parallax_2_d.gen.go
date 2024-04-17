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

type ptrsForParallax2DList struct {
	fnSetScrollScale        gdc.MethodBindPtr
	fnGetScrollScale        gdc.MethodBindPtr
	fnSetRepeatSize         gdc.MethodBindPtr
	fnGetRepeatSize         gdc.MethodBindPtr
	fnSetRepeatTimes        gdc.MethodBindPtr
	fnGetRepeatTimes        gdc.MethodBindPtr
	fnSetAutoscroll         gdc.MethodBindPtr
	fnGetAutoscroll         gdc.MethodBindPtr
	fnSetScrollOffset       gdc.MethodBindPtr
	fnGetScrollOffset       gdc.MethodBindPtr
	fnSetScreenOffset       gdc.MethodBindPtr
	fnGetScreenOffset       gdc.MethodBindPtr
	fnSetLimitBegin         gdc.MethodBindPtr
	fnGetLimitBegin         gdc.MethodBindPtr
	fnSetLimitEnd           gdc.MethodBindPtr
	fnGetLimitEnd           gdc.MethodBindPtr
	fnSetFollowViewport     gdc.MethodBindPtr
	fnGetFollowViewport     gdc.MethodBindPtr
	fnSetIgnoreCameraScroll gdc.MethodBindPtr
	fnIsIgnoreCameraScroll  gdc.MethodBindPtr
}

var ptrsForParallax2D ptrsForParallax2DList

func initParallax2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Parallax2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_scroll_scale")
		defer methodName.Destroy()
		ptrsForParallax2D.fnSetScrollScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_scroll_scale")
		defer methodName.Destroy()
		ptrsForParallax2D.fnGetScrollScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_repeat_size")
		defer methodName.Destroy()
		ptrsForParallax2D.fnSetRepeatSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_repeat_size")
		defer methodName.Destroy()
		ptrsForParallax2D.fnGetRepeatSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_repeat_times")
		defer methodName.Destroy()
		ptrsForParallax2D.fnSetRepeatTimes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_repeat_times")
		defer methodName.Destroy()
		ptrsForParallax2D.fnGetRepeatTimes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_autoscroll")
		defer methodName.Destroy()
		ptrsForParallax2D.fnSetAutoscroll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_autoscroll")
		defer methodName.Destroy()
		ptrsForParallax2D.fnGetAutoscroll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_scroll_offset")
		defer methodName.Destroy()
		ptrsForParallax2D.fnSetScrollOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_scroll_offset")
		defer methodName.Destroy()
		ptrsForParallax2D.fnGetScrollOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_screen_offset")
		defer methodName.Destroy()
		ptrsForParallax2D.fnSetScreenOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_screen_offset")
		defer methodName.Destroy()
		ptrsForParallax2D.fnGetScreenOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_limit_begin")
		defer methodName.Destroy()
		ptrsForParallax2D.fnSetLimitBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_limit_begin")
		defer methodName.Destroy()
		ptrsForParallax2D.fnGetLimitBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_limit_end")
		defer methodName.Destroy()
		ptrsForParallax2D.fnSetLimitEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_limit_end")
		defer methodName.Destroy()
		ptrsForParallax2D.fnGetLimitEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_follow_viewport")
		defer methodName.Destroy()
		ptrsForParallax2D.fnSetFollowViewport = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_follow_viewport")
		defer methodName.Destroy()
		ptrsForParallax2D.fnGetFollowViewport = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_ignore_camera_scroll")
		defer methodName.Destroy()
		ptrsForParallax2D.fnSetIgnoreCameraScroll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_ignore_camera_scroll")
		defer methodName.Destroy()
		ptrsForParallax2D.fnIsIgnoreCameraScroll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}

}

type Parallax2D struct {
	Node2D
}

func (me *Parallax2D) BaseClass() string {
	return "Parallax2D"
}

func NewParallax2D() *Parallax2D {
	str := StringNameFromStr("Parallax2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Parallax2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Parallax2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Parallax2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Parallax2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Parallax2D) SetScrollScale(scale Vector2) {
	cargs := []gdc.ConstTypePtr{scale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnSetScrollScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Parallax2D) GetScrollScale() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnGetScrollScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Parallax2D) SetRepeatSize(repeat_size Vector2) {
	cargs := []gdc.ConstTypePtr{repeat_size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnSetRepeatSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Parallax2D) GetRepeatSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnGetRepeatSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Parallax2D) SetRepeatTimes(repeat_times int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&repeat_times)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnSetRepeatTimes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Parallax2D) GetRepeatTimes() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnGetRepeatTimes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Parallax2D) SetAutoscroll(autoscroll Vector2) {
	cargs := []gdc.ConstTypePtr{autoscroll.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnSetAutoscroll), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Parallax2D) GetAutoscroll() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnGetAutoscroll), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Parallax2D) SetScrollOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnSetScrollOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Parallax2D) GetScrollOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnGetScrollOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Parallax2D) SetScreenOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnSetScreenOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Parallax2D) GetScreenOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnGetScreenOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Parallax2D) SetLimitBegin(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnSetLimitBegin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Parallax2D) GetLimitBegin() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnGetLimitBegin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Parallax2D) SetLimitEnd(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnSetLimitEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Parallax2D) GetLimitEnd() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnGetLimitEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Parallax2D) SetFollowViewport(follow bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&follow)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnSetFollowViewport), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Parallax2D) GetFollowViewport() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnGetFollowViewport), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Parallax2D) SetIgnoreCameraScroll(ignore bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ignore)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnSetIgnoreCameraScroll), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Parallax2D) IsIgnoreCameraScroll() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForParallax2D.fnIsIgnoreCameraScroll), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
