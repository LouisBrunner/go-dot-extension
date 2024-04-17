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

type ptrsForAnimationNodeStateMachineTransitionList struct {
	fnSetSwitchMode        gdc.MethodBindPtr
	fnGetSwitchMode        gdc.MethodBindPtr
	fnSetAdvanceMode       gdc.MethodBindPtr
	fnGetAdvanceMode       gdc.MethodBindPtr
	fnSetAdvanceCondition  gdc.MethodBindPtr
	fnGetAdvanceCondition  gdc.MethodBindPtr
	fnSetXfadeTime         gdc.MethodBindPtr
	fnGetXfadeTime         gdc.MethodBindPtr
	fnSetXfadeCurve        gdc.MethodBindPtr
	fnGetXfadeCurve        gdc.MethodBindPtr
	fnSetBreakLoopAtEnd    gdc.MethodBindPtr
	fnIsLoopBrokenAtEnd    gdc.MethodBindPtr
	fnSetReset             gdc.MethodBindPtr
	fnIsReset              gdc.MethodBindPtr
	fnSetPriority          gdc.MethodBindPtr
	fnGetPriority          gdc.MethodBindPtr
	fnSetAdvanceExpression gdc.MethodBindPtr
	fnGetAdvanceExpression gdc.MethodBindPtr
}

var ptrsForAnimationNodeStateMachineTransition ptrsForAnimationNodeStateMachineTransitionList

func initAnimationNodeStateMachineTransitionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationNodeStateMachineTransition")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_switch_mode")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnSetSwitchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2074906633))
	}
	{
		methodName := StringNameFromStr("get_switch_mode")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnGetSwitchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2138562085))
	}
	{
		methodName := StringNameFromStr("set_advance_mode")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnSetAdvanceMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1210869868))
	}
	{
		methodName := StringNameFromStr("get_advance_mode")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnGetAdvanceMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 61101689))
	}
	{
		methodName := StringNameFromStr("set_advance_condition")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnSetAdvanceCondition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_advance_condition")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnGetAdvanceCondition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("set_xfade_time")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnSetXfadeTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_xfade_time")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnGetXfadeTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_xfade_curve")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnSetXfadeCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 270443179))
	}
	{
		methodName := StringNameFromStr("get_xfade_curve")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnGetXfadeCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2460114913))
	}
	{
		methodName := StringNameFromStr("set_break_loop_at_end")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnSetBreakLoopAtEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_loop_broken_at_end")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnIsLoopBrokenAtEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_reset")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnSetReset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_reset")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnIsReset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_priority")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnSetPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_priority")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnGetPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_advance_expression")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnSetAdvanceExpression = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_advance_expression")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachineTransition.fnGetAdvanceExpression = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}

}

type AnimationNodeStateMachineTransition struct {
	Resource
}

func (me *AnimationNodeStateMachineTransition) BaseClass() string {
	return "AnimationNodeStateMachineTransition"
}

func NewAnimationNodeStateMachineTransition() *AnimationNodeStateMachineTransition {
	str := StringNameFromStr("AnimationNodeStateMachineTransition") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimationNodeStateMachineTransition{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type AnimationNodeStateMachineTransitionSwitchMode int

const (
	AnimationNodeStateMachineTransitionSwitchModeSwitchModeImmediate AnimationNodeStateMachineTransitionSwitchMode = 0
	AnimationNodeStateMachineTransitionSwitchModeSwitchModeSync      AnimationNodeStateMachineTransitionSwitchMode = 1
	AnimationNodeStateMachineTransitionSwitchModeSwitchModeAtEnd     AnimationNodeStateMachineTransitionSwitchMode = 2
)

type AnimationNodeStateMachineTransitionAdvanceMode int

const (
	AnimationNodeStateMachineTransitionAdvanceModeAdvanceModeDisabled AnimationNodeStateMachineTransitionAdvanceMode = 0
	AnimationNodeStateMachineTransitionAdvanceModeAdvanceModeEnabled  AnimationNodeStateMachineTransitionAdvanceMode = 1
	AnimationNodeStateMachineTransitionAdvanceModeAdvanceModeAuto     AnimationNodeStateMachineTransitionAdvanceMode = 2
)

func (me *AnimationNodeStateMachineTransition) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimationNodeStateMachineTransition) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeStateMachineTransition) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AnimationNodeStateMachineTransition) SetSwitchMode(mode AnimationNodeStateMachineTransitionSwitchMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnSetSwitchMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeStateMachineTransition) GetSwitchMode() AnimationNodeStateMachineTransitionSwitchMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AnimationNodeStateMachineTransitionSwitchMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnGetSwitchMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AnimationNodeStateMachineTransition) SetAdvanceMode(mode AnimationNodeStateMachineTransitionAdvanceMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnSetAdvanceMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeStateMachineTransition) GetAdvanceMode() AnimationNodeStateMachineTransitionAdvanceMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AnimationNodeStateMachineTransitionAdvanceMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnGetAdvanceMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AnimationNodeStateMachineTransition) SetAdvanceCondition(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnSetAdvanceCondition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeStateMachineTransition) GetAdvanceCondition() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnGetAdvanceCondition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationNodeStateMachineTransition) SetXfadeTime(secs float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnSetXfadeTime), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeStateMachineTransition) GetXfadeTime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnGetXfadeTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeStateMachineTransition) SetXfadeCurve(curve Curve) {
	cargs := []gdc.ConstTypePtr{curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnSetXfadeCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeStateMachineTransition) GetXfadeCurve() Curve {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnGetXfadeCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationNodeStateMachineTransition) SetBreakLoopAtEnd(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnSetBreakLoopAtEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeStateMachineTransition) IsLoopBrokenAtEnd() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnIsLoopBrokenAtEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeStateMachineTransition) SetReset(reset bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&reset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnSetReset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeStateMachineTransition) IsReset() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnIsReset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeStateMachineTransition) SetPriority(priority int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnSetPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeStateMachineTransition) GetPriority() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnGetPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeStateMachineTransition) SetAdvanceExpression(text String) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnSetAdvanceExpression), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeStateMachineTransition) GetAdvanceExpression() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachineTransition.fnGetAdvanceExpression), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AnimationNodeStateMachineTransitionAdvanceConditionChangedSignalFn func()

func (me *AnimationNodeStateMachineTransition) ConnectAdvanceConditionChanged(subs SignalSubscribers, fn AnimationNodeStateMachineTransitionAdvanceConditionChangedSignalFn) {
	sig := StringNameFromStr("advance_condition_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationNodeStateMachineTransition) DisconnectAdvanceConditionChanged(subs SignalSubscribers, fn AnimationNodeStateMachineTransitionAdvanceConditionChangedSignalFn) {
	sig := StringNameFromStr("advance_condition_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
