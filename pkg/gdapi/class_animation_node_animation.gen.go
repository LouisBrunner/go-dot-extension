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

type ptrsForAnimationNodeAnimationList struct {
	fnSetAnimation          gdc.MethodBindPtr
	fnGetAnimation          gdc.MethodBindPtr
	fnSetPlayMode           gdc.MethodBindPtr
	fnGetPlayMode           gdc.MethodBindPtr
	fnSetUseCustomTimeline  gdc.MethodBindPtr
	fnIsUsingCustomTimeline gdc.MethodBindPtr
	fnSetTimelineLength     gdc.MethodBindPtr
	fnGetTimelineLength     gdc.MethodBindPtr
	fnSetStretchTimeScale   gdc.MethodBindPtr
	fnIsStretchingTimeScale gdc.MethodBindPtr
	fnSetStartOffset        gdc.MethodBindPtr
	fnGetStartOffset        gdc.MethodBindPtr
	fnSetLoopMode           gdc.MethodBindPtr
	fnGetLoopMode           gdc.MethodBindPtr
}

var ptrsForAnimationNodeAnimation ptrsForAnimationNodeAnimationList

func initAnimationNodeAnimationPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationNodeAnimation")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_animation")
		defer methodName.Destroy()
		ptrsForAnimationNodeAnimation.fnSetAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_animation")
		defer methodName.Destroy()
		ptrsForAnimationNodeAnimation.fnGetAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("set_play_mode")
		defer methodName.Destroy()
		ptrsForAnimationNodeAnimation.fnSetPlayMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3347718873))
	}
	{
		methodName := StringNameFromStr("get_play_mode")
		defer methodName.Destroy()
		ptrsForAnimationNodeAnimation.fnGetPlayMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2061244637))
	}
	{
		methodName := StringNameFromStr("set_use_custom_timeline")
		defer methodName.Destroy()
		ptrsForAnimationNodeAnimation.fnSetUseCustomTimeline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_custom_timeline")
		defer methodName.Destroy()
		ptrsForAnimationNodeAnimation.fnIsUsingCustomTimeline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_timeline_length")
		defer methodName.Destroy()
		ptrsForAnimationNodeAnimation.fnSetTimelineLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_timeline_length")
		defer methodName.Destroy()
		ptrsForAnimationNodeAnimation.fnGetTimelineLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_stretch_time_scale")
		defer methodName.Destroy()
		ptrsForAnimationNodeAnimation.fnSetStretchTimeScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_stretching_time_scale")
		defer methodName.Destroy()
		ptrsForAnimationNodeAnimation.fnIsStretchingTimeScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_start_offset")
		defer methodName.Destroy()
		ptrsForAnimationNodeAnimation.fnSetStartOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_start_offset")
		defer methodName.Destroy()
		ptrsForAnimationNodeAnimation.fnGetStartOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_loop_mode")
		defer methodName.Destroy()
		ptrsForAnimationNodeAnimation.fnSetLoopMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3155355575))
	}
	{
		methodName := StringNameFromStr("get_loop_mode")
		defer methodName.Destroy()
		ptrsForAnimationNodeAnimation.fnGetLoopMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1988889481))
	}

}

type AnimationNodeAnimation struct {
	AnimationRootNode
}

func (me *AnimationNodeAnimation) BaseClass() string {
	return "AnimationNodeAnimation"
}

func NewAnimationNodeAnimation() *AnimationNodeAnimation {
	str := StringNameFromStr("AnimationNodeAnimation") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimationNodeAnimation{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type AnimationNodeAnimationPlayMode int

const (
	AnimationNodeAnimationPlayModePlayModeForward  AnimationNodeAnimationPlayMode = 0
	AnimationNodeAnimationPlayModePlayModeBackward AnimationNodeAnimationPlayMode = 1
)

func (me *AnimationNodeAnimation) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimationNodeAnimation) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeAnimation) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AnimationNodeAnimation) SetAnimation(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnSetAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeAnimation) GetAnimation() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnGetAnimation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationNodeAnimation) SetPlayMode(mode AnimationNodeAnimationPlayMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnSetPlayMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeAnimation) GetPlayMode() AnimationNodeAnimationPlayMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AnimationNodeAnimationPlayMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnGetPlayMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AnimationNodeAnimation) SetUseCustomTimeline(use_custom_timeline bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_custom_timeline)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnSetUseCustomTimeline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeAnimation) IsUsingCustomTimeline() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnIsUsingCustomTimeline), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeAnimation) SetTimelineLength(timeline_length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&timeline_length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnSetTimelineLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeAnimation) GetTimelineLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnGetTimelineLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeAnimation) SetStretchTimeScale(stretch_time_scale bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stretch_time_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnSetStretchTimeScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeAnimation) IsStretchingTimeScale() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnIsStretchingTimeScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeAnimation) SetStartOffset(start_offset float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&start_offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnSetStartOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeAnimation) GetStartOffset() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnGetStartOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeAnimation) SetLoopMode(loop_mode AnimationLoopMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnSetLoopMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeAnimation) GetLoopMode() AnimationLoopMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AnimationLoopMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeAnimation.fnGetLoopMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
