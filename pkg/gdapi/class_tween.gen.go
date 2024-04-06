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

type ptrsForTweenList struct {
	fnTweenProperty       gdc.MethodBindPtr
	fnTweenInterval       gdc.MethodBindPtr
	fnTweenCallback       gdc.MethodBindPtr
	fnTweenMethod         gdc.MethodBindPtr
	fnCustomStep          gdc.MethodBindPtr
	fnStop                gdc.MethodBindPtr
	fnPause               gdc.MethodBindPtr
	fnPlay                gdc.MethodBindPtr
	fnKill                gdc.MethodBindPtr
	fnGetTotalElapsedTime gdc.MethodBindPtr
	fnIsRunning           gdc.MethodBindPtr
	fnIsValid             gdc.MethodBindPtr
	fnBindNode            gdc.MethodBindPtr
	fnSetProcessMode      gdc.MethodBindPtr
	fnSetPauseMode        gdc.MethodBindPtr
	fnSetParallel         gdc.MethodBindPtr
	fnSetLoops            gdc.MethodBindPtr
	fnGetLoopsLeft        gdc.MethodBindPtr
	fnSetSpeedScale       gdc.MethodBindPtr
	fnSetTrans            gdc.MethodBindPtr
	fnSetEase             gdc.MethodBindPtr
	fnParallel            gdc.MethodBindPtr
	fnChain               gdc.MethodBindPtr
	fnInterpolateValue    gdc.MethodBindPtr
}

var ptrsForTween ptrsForTweenList

func initTweenPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Tween")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("tween_property")
		defer methodName.Destroy()
		ptrsForTween.fnTweenProperty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4049770449))
	}
	{
		methodName := StringNameFromStr("tween_interval")
		defer methodName.Destroy()
		ptrsForTween.fnTweenInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 413360199))
	}
	{
		methodName := StringNameFromStr("tween_callback")
		defer methodName.Destroy()
		ptrsForTween.fnTweenCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1540176488))
	}
	{
		methodName := StringNameFromStr("tween_method")
		defer methodName.Destroy()
		ptrsForTween.fnTweenMethod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2337877153))
	}
	{
		methodName := StringNameFromStr("custom_step")
		defer methodName.Destroy()
		ptrsForTween.fnCustomStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 330693286))
	}
	{
		methodName := StringNameFromStr("stop")
		defer methodName.Destroy()
		ptrsForTween.fnStop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("pause")
		defer methodName.Destroy()
		ptrsForTween.fnPause = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("play")
		defer methodName.Destroy()
		ptrsForTween.fnPlay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("kill")
		defer methodName.Destroy()
		ptrsForTween.fnKill = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_total_elapsed_time")
		defer methodName.Destroy()
		ptrsForTween.fnGetTotalElapsedTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("is_running")
		defer methodName.Destroy()
		ptrsForTween.fnIsRunning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("is_valid")
		defer methodName.Destroy()
		ptrsForTween.fnIsValid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("bind_node")
		defer methodName.Destroy()
		ptrsForTween.fnBindNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2946786331))
	}
	{
		methodName := StringNameFromStr("set_process_mode")
		defer methodName.Destroy()
		ptrsForTween.fnSetProcessMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 855258840))
	}
	{
		methodName := StringNameFromStr("set_pause_mode")
		defer methodName.Destroy()
		ptrsForTween.fnSetPauseMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3363368837))
	}
	{
		methodName := StringNameFromStr("set_parallel")
		defer methodName.Destroy()
		ptrsForTween.fnSetParallel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1942052223))
	}
	{
		methodName := StringNameFromStr("set_loops")
		defer methodName.Destroy()
		ptrsForTween.fnSetLoops = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2670836414))
	}
	{
		methodName := StringNameFromStr("get_loops_left")
		defer methodName.Destroy()
		ptrsForTween.fnGetLoopsLeft = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_speed_scale")
		defer methodName.Destroy()
		ptrsForTween.fnSetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3961971106))
	}
	{
		methodName := StringNameFromStr("set_trans")
		defer methodName.Destroy()
		ptrsForTween.fnSetTrans = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3965963875))
	}
	{
		methodName := StringNameFromStr("set_ease")
		defer methodName.Destroy()
		ptrsForTween.fnSetEase = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1208117252))
	}
	{
		methodName := StringNameFromStr("parallel")
		defer methodName.Destroy()
		ptrsForTween.fnParallel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3426978995))
	}
	{
		methodName := StringNameFromStr("chain")
		defer methodName.Destroy()
		ptrsForTween.fnChain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3426978995))
	}
	{
		methodName := StringNameFromStr("interpolate_value")
		defer methodName.Destroy()
		ptrsForTween.fnInterpolateValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3452526450))
	}

}

type Tween struct {
	RefCounted
}

func (me *Tween) BaseClass() string {
	return "Tween"
}

func NewTween() *Tween {
	str := StringNameFromStr("Tween") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Tween{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TweenTweenProcessMode int

const (
	TweenTweenProcessModeTweenProcessPhysics TweenTweenProcessMode = 0
	TweenTweenProcessModeTweenProcessIdle    TweenTweenProcessMode = 1
)

type TweenTweenPauseMode int

const (
	TweenTweenPauseModeTweenPauseBound   TweenTweenPauseMode = 0
	TweenTweenPauseModeTweenPauseStop    TweenTweenPauseMode = 1
	TweenTweenPauseModeTweenPauseProcess TweenTweenPauseMode = 2
)

type TweenTransitionType int

const (
	TweenTransitionTypeTransLinear  TweenTransitionType = 0
	TweenTransitionTypeTransSine    TweenTransitionType = 1
	TweenTransitionTypeTransQuint   TweenTransitionType = 2
	TweenTransitionTypeTransQuart   TweenTransitionType = 3
	TweenTransitionTypeTransQuad    TweenTransitionType = 4
	TweenTransitionTypeTransExpo    TweenTransitionType = 5
	TweenTransitionTypeTransElastic TweenTransitionType = 6
	TweenTransitionTypeTransCubic   TweenTransitionType = 7
	TweenTransitionTypeTransCirc    TweenTransitionType = 8
	TweenTransitionTypeTransBounce  TweenTransitionType = 9
	TweenTransitionTypeTransBack    TweenTransitionType = 10
	TweenTransitionTypeTransSpring  TweenTransitionType = 11
)

type TweenEaseType int

const (
	TweenEaseTypeEaseIn    TweenEaseType = 0
	TweenEaseTypeEaseOut   TweenEaseType = 1
	TweenEaseTypeEaseInOut TweenEaseType = 2
	TweenEaseTypeEaseOutIn TweenEaseType = 3
)

func (me *Tween) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Tween) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Tween) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Tween) TweenProperty(object Object, property NodePath, final_val Variant, duration float64) PropertyTweener {
	cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), property.AsCTypePtr(), final_val.AsCTypePtr(), gdc.ConstTypePtr(&duration)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPropertyTweener()
	pinner.Pin(&duration)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnTweenProperty), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tween) TweenInterval(time float64) IntervalTweener {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewIntervalTweener()
	pinner.Pin(&time)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnTweenInterval), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tween) TweenCallback(callback Callable) CallbackTweener {
	cargs := []gdc.ConstTypePtr{callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCallbackTweener()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnTweenCallback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tween) TweenMethod(method Callable, from Variant, to Variant, duration float64) MethodTweener {
	cargs := []gdc.ConstTypePtr{method.AsCTypePtr(), from.AsCTypePtr(), to.AsCTypePtr(), gdc.ConstTypePtr(&duration)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMethodTweener()
	pinner.Pin(&duration)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnTweenMethod), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tween) CustomStep(delta float64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&delta)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnCustomStep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tween) Stop() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnStop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tween) Pause() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnPause), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tween) Play() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnPlay), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tween) Kill() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnKill), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Tween) GetTotalElapsedTime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnGetTotalElapsedTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tween) IsRunning() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnIsRunning), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tween) IsValid() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnIsValid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tween) BindNode(node Node) Tween {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTween()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnBindNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tween) SetProcessMode(mode TweenTweenProcessMode) Tween {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTween()
	pinner.Pin(&mode)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnSetProcessMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tween) SetPauseMode(mode TweenTweenPauseMode) Tween {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTween()
	pinner.Pin(&mode)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnSetPauseMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tween) SetParallel(parallel bool) Tween {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parallel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTween()
	pinner.Pin(&parallel)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnSetParallel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tween) SetLoops(loops int64) Tween {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loops)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTween()
	pinner.Pin(&loops)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnSetLoops), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tween) GetLoopsLeft() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnGetLoopsLeft), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Tween) SetSpeedScale(speed float64) Tween {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTween()
	pinner.Pin(&speed)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnSetSpeedScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tween) SetTrans(trans TweenTransitionType) Tween {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&trans)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTween()
	pinner.Pin(&trans)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnSetTrans), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tween) SetEase(ease TweenEaseType) Tween {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ease)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTween()
	pinner.Pin(&ease)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnSetEase), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tween) Parallel() Tween {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTween()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnParallel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Tween) Chain() Tween {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTween()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnChain), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func TweenInterpolateValue(initial_value Variant, delta_value Variant, elapsed_time float64, duration float64, trans_type TweenTransitionType, ease_type TweenEaseType) Variant {
	cargs := []gdc.ConstTypePtr{initial_value.AsCTypePtr(), delta_value.AsCTypePtr(), gdc.ConstTypePtr(&elapsed_time), gdc.ConstTypePtr(&duration), gdc.ConstTypePtr(&trans_type), gdc.ConstTypePtr(&ease_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&elapsed_time)
	pinner.Pin(&duration)
	pinner.Pin(&trans_type)
	pinner.Pin(&ease_type)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTween.fnInterpolateValue), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals

type TweenStepFinishedSignalFn func(idx int)

func (me *Tween) ConnectStepFinished(subs SignalSubscribers, fn TweenStepFinishedSignalFn) {
	sig := StringNameFromStr("step_finished")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tween) DisconnectStepFinished(subs SignalSubscribers, fn TweenStepFinishedSignalFn) {
	sig := StringNameFromStr("step_finished")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TweenLoopFinishedSignalFn func(loop_count int)

func (me *Tween) ConnectLoopFinished(subs SignalSubscribers, fn TweenLoopFinishedSignalFn) {
	sig := StringNameFromStr("loop_finished")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tween) DisconnectLoopFinished(subs SignalSubscribers, fn TweenLoopFinishedSignalFn) {
	sig := StringNameFromStr("loop_finished")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TweenFinishedSignalFn func()

func (me *Tween) ConnectFinished(subs SignalSubscribers, fn TweenFinishedSignalFn) {
	sig := StringNameFromStr("finished")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tween) DisconnectFinished(subs SignalSubscribers, fn TweenFinishedSignalFn) {
	sig := StringNameFromStr("finished")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
