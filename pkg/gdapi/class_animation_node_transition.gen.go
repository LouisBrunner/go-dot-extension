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

type ptrsForAnimationNodeTransitionList struct {
	fnSetInputCount            gdc.MethodBindPtr
	fnSetInputAsAutoAdvance    gdc.MethodBindPtr
	fnIsInputSetAsAutoAdvance  gdc.MethodBindPtr
	fnSetInputBreakLoopAtEnd   gdc.MethodBindPtr
	fnIsInputLoopBrokenAtEnd   gdc.MethodBindPtr
	fnSetInputReset            gdc.MethodBindPtr
	fnIsInputReset             gdc.MethodBindPtr
	fnSetXfadeTime             gdc.MethodBindPtr
	fnGetXfadeTime             gdc.MethodBindPtr
	fnSetXfadeCurve            gdc.MethodBindPtr
	fnGetXfadeCurve            gdc.MethodBindPtr
	fnSetAllowTransitionToSelf gdc.MethodBindPtr
	fnIsAllowTransitionToSelf  gdc.MethodBindPtr
}

var ptrsForAnimationNodeTransition ptrsForAnimationNodeTransitionList

func initAnimationNodeTransitionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationNodeTransition")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_input_count")
		defer methodName.Destroy()
		ptrsForAnimationNodeTransition.fnSetInputCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_input_as_auto_advance")
		defer methodName.Destroy()
		ptrsForAnimationNodeTransition.fnSetInputAsAutoAdvance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_input_set_as_auto_advance")
		defer methodName.Destroy()
		ptrsForAnimationNodeTransition.fnIsInputSetAsAutoAdvance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_input_break_loop_at_end")
		defer methodName.Destroy()
		ptrsForAnimationNodeTransition.fnSetInputBreakLoopAtEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_input_loop_broken_at_end")
		defer methodName.Destroy()
		ptrsForAnimationNodeTransition.fnIsInputLoopBrokenAtEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_input_reset")
		defer methodName.Destroy()
		ptrsForAnimationNodeTransition.fnSetInputReset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_input_reset")
		defer methodName.Destroy()
		ptrsForAnimationNodeTransition.fnIsInputReset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_xfade_time")
		defer methodName.Destroy()
		ptrsForAnimationNodeTransition.fnSetXfadeTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_xfade_time")
		defer methodName.Destroy()
		ptrsForAnimationNodeTransition.fnGetXfadeTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_xfade_curve")
		defer methodName.Destroy()
		ptrsForAnimationNodeTransition.fnSetXfadeCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
	}
	{
		methodName := StringNameFromStr("get_xfade_curve")
		defer methodName.Destroy()
		ptrsForAnimationNodeTransition.fnGetXfadeCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
	}
	{
		methodName := StringNameFromStr("set_allow_transition_to_self")
		defer methodName.Destroy()
		ptrsForAnimationNodeTransition.fnSetAllowTransitionToSelf = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_allow_transition_to_self")
		defer methodName.Destroy()
		ptrsForAnimationNodeTransition.fnIsAllowTransitionToSelf = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type AnimationNodeTransition struct {
	AnimationNodeSync
}

func (me *AnimationNodeTransition) BaseClass() string {
	return "AnimationNodeTransition"
}

func NewAnimationNodeTransition() *AnimationNodeTransition {
	str := StringNameFromStr("AnimationNodeTransition") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimationNodeTransition{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AnimationNodeTransition) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimationNodeTransition) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeTransition) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AnimationNodeTransition) SetInputCount(input_count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeTransition.fnSetInputCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeTransition) SetInputAsAutoAdvance(input int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeTransition.fnSetInputAsAutoAdvance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeTransition) IsInputSetAsAutoAdvance(input int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&input)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeTransition.fnIsInputSetAsAutoAdvance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeTransition) SetInputBreakLoopAtEnd(input int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeTransition.fnSetInputBreakLoopAtEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeTransition) IsInputLoopBrokenAtEnd(input int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&input)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeTransition.fnIsInputLoopBrokenAtEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeTransition) SetInputReset(input int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeTransition.fnSetInputReset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeTransition) IsInputReset(input int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&input)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeTransition.fnIsInputReset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeTransition) SetXfadeTime(time float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeTransition.fnSetXfadeTime), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeTransition) GetXfadeTime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeTransition.fnGetXfadeTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeTransition) SetXfadeCurve(curve Curve) {
	cargs := []gdc.ConstTypePtr{curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeTransition.fnSetXfadeCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeTransition) GetXfadeCurve() Curve {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeTransition.fnGetXfadeCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationNodeTransition) SetAllowTransitionToSelf(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeTransition.fnSetAllowTransitionToSelf), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeTransition) IsAllowTransitionToSelf() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeTransition.fnIsAllowTransitionToSelf), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
