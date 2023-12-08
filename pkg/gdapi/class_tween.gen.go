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

func  (me *Tween) TweenProperty(object Object, property NodePath, final_val Variant, duration float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tween) TweenInterval(time float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tween) TweenCallback(callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tween) TweenMethod(method Callable, from Variant, to Variant, duration float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tween) CustomStep(delta float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tween) Stop() { // TODO: return value
  // TODO: implement
}

func  (me *Tween) Pause() { // TODO: return value
  // TODO: implement
}

func  (me *Tween) Play() { // TODO: return value
  // TODO: implement
}

func  (me *Tween) Kill() { // TODO: return value
  // TODO: implement
}

func  (me *Tween) GetTotalElapsedTime() { // TODO: return value
  // TODO: implement
}

func  (me *Tween) IsRunning() { // TODO: return value
  // TODO: implement
}

func  (me *Tween) IsValid() { // TODO: return value
  // TODO: implement
}

func  (me *Tween) BindNode(node Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tween) SetProcessMode(mode TweenTweenProcessMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tween) SetPauseMode(mode TweenTweenPauseMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tween) SetParallel(parallel bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tween) SetLoops(loops int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tween) GetLoopsLeft() { // TODO: return value
  // TODO: implement
}

func  (me *Tween) SetSpeedScale(speed float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tween) SetTrans(trans TweenTransitionType, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tween) SetEase(ease TweenEaseType, ) { // TODO: return value
  // TODO: implement
}

func  (me *Tween) Parallel() { // TODO: return value
  // TODO: implement
}

func  (me *Tween) Chain() { // TODO: return value
  // TODO: implement
}

func  TweenInterpolateValue(initial_value Variant, delta_value Variant, elapsed_time float32, duration float32, trans_type TweenTransitionType, ease_type TweenEaseType, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
