// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Tween struct {
  obj gdc.ObjectPtr
}

func (me *Tween) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Tween) BaseClass() string {
  return "Tween"
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

func  (me *Tween) TweenProperty(object Object, property NodePath, final_val Variant, duration float32, ) PropertyTweener {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tween_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4049770449) // FIXME: should cache?
  var ret PropertyTweener
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(object.AsCTypePtr()), gdc.ConstTypePtr(property.AsCTypePtr()), gdc.ConstTypePtr(final_val.AsCTypePtr()), gdc.ConstTypePtr(&duration), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) TweenInterval(time float32, ) IntervalTweener {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tween_interval")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 413360199) // FIXME: should cache?
  var ret IntervalTweener
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) TweenCallback(callback Callable, ) CallbackTweener {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tween_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1540176488) // FIXME: should cache?
  var ret CallbackTweener
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) TweenMethod(method Callable, from Variant, to Variant, duration float32, ) MethodTweener {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tween_method")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2337877153) // FIXME: should cache?
  var ret MethodTweener
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(method.AsCTypePtr()), gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(to.AsCTypePtr()), gdc.ConstTypePtr(&duration), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) CustomStep(delta float32, ) bool {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("custom_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 330693286) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) Stop()  {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tween) Pause()  {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pause")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tween) Play()  {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tween) Kill()  {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("kill")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Tween) GetTotalElapsedTime() float32 {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_total_elapsed_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) IsRunning() bool {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_running")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) IsValid() bool {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_valid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) BindNode(node Node, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bind_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2946786331) // FIXME: should cache?
  var ret Tween
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) SetProcessMode(mode TweenTweenProcessMode, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 855258840) // FIXME: should cache?
  var ret Tween
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) SetPauseMode(mode TweenTweenPauseMode, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pause_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3363368837) // FIXME: should cache?
  var ret Tween
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) SetParallel(parallel bool, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_parallel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1942052223) // FIXME: should cache?
  var ret Tween
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parallel), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) SetLoops(loops int, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loops")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2670836414) // FIXME: should cache?
  var ret Tween
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loops), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) GetLoopsLeft() int {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_loops_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) SetSpeedScale(speed float32, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3961971106) // FIXME: should cache?
  var ret Tween
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) SetTrans(trans TweenTransitionType, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_trans")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3965963875) // FIXME: should cache?
  var ret Tween
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&trans), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) SetEase(ease TweenEaseType, ) Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ease")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1208117252) // FIXME: should cache?
  var ret Tween
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ease), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) Parallel() Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("parallel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3426978995) // FIXME: should cache?
  var ret Tween
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Tween) Chain() Tween {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("chain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3426978995) // FIXME: should cache?
  var ret Tween
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  TweenInterpolateValue(initial_value Variant, delta_value Variant, elapsed_time float32, duration float32, trans_type TweenTransitionType, ease_type TweenEaseType, ) Variant {
  classNameV := StringNameFromStr("Tween")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("interpolate_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3452526450) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(initial_value.AsCTypePtr()), gdc.ConstTypePtr(delta_value.AsCTypePtr()), gdc.ConstTypePtr(&elapsed_time), gdc.ConstTypePtr(&duration), gdc.ConstTypePtr(&trans_type), gdc.ConstTypePtr(&ease_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API