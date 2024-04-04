// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

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
  TweenTweenProcessModeTweenProcessIdle TweenTweenProcessMode = 1
)

type TweenTweenPauseMode int
const (
  TweenTweenPauseModeTweenPauseBound TweenTweenPauseMode = 0
  TweenTweenPauseModeTweenPauseStop TweenTweenPauseMode = 1
  TweenTweenPauseModeTweenPauseProcess TweenTweenPauseMode = 2
)

type TweenTransitionType int
const (
  TweenTransitionTypeTransLinear TweenTransitionType = 0
  TweenTransitionTypeTransSine TweenTransitionType = 1
  TweenTransitionTypeTransQuint TweenTransitionType = 2
  TweenTransitionTypeTransQuart TweenTransitionType = 3
  TweenTransitionTypeTransQuad TweenTransitionType = 4
  TweenTransitionTypeTransExpo TweenTransitionType = 5
  TweenTransitionTypeTransElastic TweenTransitionType = 6
  TweenTransitionTypeTransCubic TweenTransitionType = 7
  TweenTransitionTypeTransCirc TweenTransitionType = 8
  TweenTransitionTypeTransBounce TweenTransitionType = 9
  TweenTransitionTypeTransBack TweenTransitionType = 10
  TweenTransitionTypeTransSpring TweenTransitionType = 11
)

type TweenEaseType int
const (
  TweenEaseTypeEaseIn TweenEaseType = 0
  TweenEaseTypeEaseOut TweenEaseType = 1
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

func  (me *Tween) TweenProperty(object Object, property NodePath, final_val Variant, duration float64, ) PropertyTweener {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tween_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4049770449) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), property.AsCTypePtr(), final_val.AsCTypePtr(), gdc.ConstTypePtr(&duration) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPropertyTweener()
  pinner.Pin(&duration)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Tween) TweenInterval(time float64, ) IntervalTweener {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tween_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 413360199) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewIntervalTweener()
  pinner.Pin(&time)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Tween) TweenCallback(callback Callable, ) CallbackTweener {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tween_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1540176488) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCallbackTweener()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Tween) TweenMethod(method Callable, from Variant, to Variant, duration float64, ) MethodTweener {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tween_method")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2337877153) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{method.AsCTypePtr(), from.AsCTypePtr(), to.AsCTypePtr(), gdc.ConstTypePtr(&duration) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMethodTweener()
  pinner.Pin(&duration)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Tween) CustomStep(delta float64, ) bool {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("custom_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 330693286) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&delta)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Tween) Stop()  {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Tween) Pause()  {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pause")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Tween) Play()  {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Tween) Kill()  {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("kill")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Tween) GetTotalElapsedTime() float64 {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_total_elapsed_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Tween) IsRunning() bool {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_running")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Tween) IsValid() bool {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_valid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Tween) BindNode(node Node, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bind_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2946786331) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTween()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Tween) SetProcessMode(mode TweenTweenProcessMode, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 855258840) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTween()
  pinner.Pin(&mode)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Tween) SetPauseMode(mode TweenTweenPauseMode, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pause_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3363368837) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTween()
  pinner.Pin(&mode)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Tween) SetParallel(parallel bool, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_parallel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1942052223) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parallel) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTween()
  pinner.Pin(&parallel)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Tween) SetLoops(loops int64, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loops")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2670836414) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loops) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTween()
  pinner.Pin(&loops)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Tween) GetLoopsLeft() int64 {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_loops_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Tween) SetSpeedScale(speed float64, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3961971106) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTween()
  pinner.Pin(&speed)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Tween) SetTrans(trans TweenTransitionType, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_trans")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3965963875) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&trans) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTween()
  pinner.Pin(&trans)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Tween) SetEase(ease TweenEaseType, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ease")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1208117252) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ease) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTween()
  pinner.Pin(&ease)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Tween) Parallel() Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("parallel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3426978995) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTween()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Tween) Chain() Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("chain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3426978995) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTween()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  TweenInterpolateValue(initial_value Variant, delta_value Variant, elapsed_time float64, duration float64, trans_type TweenTransitionType, ease_type TweenEaseType, ) Variant {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("interpolate_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3452526450) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{initial_value.AsCTypePtr(), delta_value.AsCTypePtr(), gdc.ConstTypePtr(&elapsed_time) , gdc.ConstTypePtr(&duration) , gdc.ConstTypePtr(&trans_type) , gdc.ConstTypePtr(&ease_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&elapsed_time)
  pinner.Pin(&duration)
  pinner.Pin(&trans_type)
  pinner.Pin(&ease_type)

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals

type TweenStepFinishedSignalFn func(idx int, )

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

type TweenLoopFinishedSignalFn func(loop_count int, )

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
