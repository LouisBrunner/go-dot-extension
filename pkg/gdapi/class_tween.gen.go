// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func  (me *Tween) TweenProperty(object Object, property NodePath, final_val Variant, duration float32, )  {
  panic("TODO: implement")
}

func  (me *Tween) TweenInterval(time float32, )  {
  panic("TODO: implement")
}

func  (me *Tween) TweenCallback(callback Callable, )  {
  panic("TODO: implement")
}

func  (me *Tween) TweenMethod(method Callable, from Variant, to Variant, duration float32, )  {
  panic("TODO: implement")
}

func  (me *Tween) CustomStep(delta float32, )  {
  panic("TODO: implement")
}

func  (me *Tween) Stop()  {
  panic("TODO: implement")
}

func  (me *Tween) Pause()  {
  panic("TODO: implement")
}

func  (me *Tween) Play()  {
  panic("TODO: implement")
}

func  (me *Tween) Kill()  {
  panic("TODO: implement")
}

func  (me *Tween) GetTotalElapsedTime()  {
  panic("TODO: implement")
}

func  (me *Tween) IsRunning()  {
  panic("TODO: implement")
}

func  (me *Tween) IsValid()  {
  panic("TODO: implement")
}

func  (me *Tween) BindNode(node Node, )  {
  panic("TODO: implement")
}

func  (me *Tween) SetProcessMode(mode TweenTweenProcessMode, )  {
  panic("TODO: implement")
}

func  (me *Tween) SetPauseMode(mode TweenTweenPauseMode, )  {
  panic("TODO: implement")
}

func  (me *Tween) SetParallel(parallel bool, )  {
  panic("TODO: implement")
}

func  (me *Tween) SetLoops(loops int, )  {
  panic("TODO: implement")
}

func  (me *Tween) GetLoopsLeft()  {
  panic("TODO: implement")
}

func  (me *Tween) SetSpeedScale(speed float32, )  {
  panic("TODO: implement")
}

func  (me *Tween) SetTrans(trans TweenTransitionType, )  {
  panic("TODO: implement")
}

func  (me *Tween) SetEase(ease TweenEaseType, )  {
  panic("TODO: implement")
}

func  (me *Tween) Parallel()  {
  panic("TODO: implement")
}

func  (me *Tween) Chain()  {
  panic("TODO: implement")
}

func  TweenInterpolateValue(initial_value Variant, delta_value Variant, elapsed_time float32, duration float32, trans_type TweenTransitionType, ease_type TweenEaseType, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
