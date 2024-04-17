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

type ptrsForAnimationNodeOneShotList struct {
	fnSetFadeinTime             gdc.MethodBindPtr
	fnGetFadeinTime             gdc.MethodBindPtr
	fnSetFadeinCurve            gdc.MethodBindPtr
	fnGetFadeinCurve            gdc.MethodBindPtr
	fnSetFadeoutTime            gdc.MethodBindPtr
	fnGetFadeoutTime            gdc.MethodBindPtr
	fnSetFadeoutCurve           gdc.MethodBindPtr
	fnGetFadeoutCurve           gdc.MethodBindPtr
	fnSetBreakLoopAtEnd         gdc.MethodBindPtr
	fnIsLoopBrokenAtEnd         gdc.MethodBindPtr
	fnSetAutorestart            gdc.MethodBindPtr
	fnHasAutorestart            gdc.MethodBindPtr
	fnSetAutorestartDelay       gdc.MethodBindPtr
	fnGetAutorestartDelay       gdc.MethodBindPtr
	fnSetAutorestartRandomDelay gdc.MethodBindPtr
	fnGetAutorestartRandomDelay gdc.MethodBindPtr
	fnSetMixMode                gdc.MethodBindPtr
	fnGetMixMode                gdc.MethodBindPtr
}

var ptrsForAnimationNodeOneShot ptrsForAnimationNodeOneShotList

func initAnimationNodeOneShotPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationNodeOneShot")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_fadein_time")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnSetFadeinTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_fadein_time")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnGetFadeinTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_fadein_curve")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnSetFadeinCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
	}
	{
		methodName := StringNameFromStr("get_fadein_curve")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnGetFadeinCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
	}
	{
		methodName := StringNameFromStr("set_fadeout_time")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnSetFadeoutTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_fadeout_time")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnGetFadeoutTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_fadeout_curve")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnSetFadeoutCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
	}
	{
		methodName := StringNameFromStr("get_fadeout_curve")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnGetFadeoutCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
	}
	{
		methodName := StringNameFromStr("set_break_loop_at_end")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnSetBreakLoopAtEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_loop_broken_at_end")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnIsLoopBrokenAtEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_autorestart")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnSetAutorestart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("has_autorestart")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnHasAutorestart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_autorestart_delay")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnSetAutorestartDelay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_autorestart_delay")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnGetAutorestartDelay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_autorestart_random_delay")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnSetAutorestartRandomDelay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_autorestart_random_delay")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnGetAutorestartRandomDelay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_mix_mode")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnSetMixMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1018899799))
	}
	{
		methodName := StringNameFromStr("get_mix_mode")
		defer methodName.Destroy()
		ptrsForAnimationNodeOneShot.fnGetMixMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3076550526))
	}

}

type AnimationNodeOneShot struct {
	AnimationNodeSync
}

func (me *AnimationNodeOneShot) BaseClass() string {
	return "AnimationNodeOneShot"
}

func NewAnimationNodeOneShot() *AnimationNodeOneShot {
	str := StringNameFromStr("AnimationNodeOneShot") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimationNodeOneShot{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type AnimationNodeOneShotOneShotRequest int

const (
	AnimationNodeOneShotOneShotRequestOneShotRequestNone    AnimationNodeOneShotOneShotRequest = 0
	AnimationNodeOneShotOneShotRequestOneShotRequestFire    AnimationNodeOneShotOneShotRequest = 1
	AnimationNodeOneShotOneShotRequestOneShotRequestAbort   AnimationNodeOneShotOneShotRequest = 2
	AnimationNodeOneShotOneShotRequestOneShotRequestFadeOut AnimationNodeOneShotOneShotRequest = 3
)

type AnimationNodeOneShotMixMode int

const (
	AnimationNodeOneShotMixModeMixModeBlend AnimationNodeOneShotMixMode = 0
	AnimationNodeOneShotMixModeMixModeAdd   AnimationNodeOneShotMixMode = 1
)

func (me *AnimationNodeOneShot) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimationNodeOneShot) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeOneShot) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AnimationNodeOneShot) SetFadeinTime(time float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnSetFadeinTime), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeOneShot) GetFadeinTime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnGetFadeinTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeOneShot) SetFadeinCurve(curve Curve) {
	cargs := []gdc.ConstTypePtr{curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnSetFadeinCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeOneShot) GetFadeinCurve() Curve {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnGetFadeinCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationNodeOneShot) SetFadeoutTime(time float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnSetFadeoutTime), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeOneShot) GetFadeoutTime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnGetFadeoutTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeOneShot) SetFadeoutCurve(curve Curve) {
	cargs := []gdc.ConstTypePtr{curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnSetFadeoutCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeOneShot) GetFadeoutCurve() Curve {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnGetFadeoutCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationNodeOneShot) SetBreakLoopAtEnd(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnSetBreakLoopAtEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeOneShot) IsLoopBrokenAtEnd() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnIsLoopBrokenAtEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeOneShot) SetAutorestart(active bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnSetAutorestart), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeOneShot) HasAutorestart() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnHasAutorestart), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeOneShot) SetAutorestartDelay(time float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnSetAutorestartDelay), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeOneShot) GetAutorestartDelay() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnGetAutorestartDelay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeOneShot) SetAutorestartRandomDelay(time float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnSetAutorestartRandomDelay), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeOneShot) GetAutorestartRandomDelay() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnGetAutorestartRandomDelay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeOneShot) SetMixMode(mode AnimationNodeOneShotMixMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnSetMixMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeOneShot) GetMixMode() AnimationNodeOneShotMixMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AnimationNodeOneShotMixMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeOneShot.fnGetMixMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
