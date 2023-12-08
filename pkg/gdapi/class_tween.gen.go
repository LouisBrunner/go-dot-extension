// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  TweenTweenProcessPhysics TweenTweenProcessMode = 0
  TweenTweenProcessIdle TweenTweenProcessMode = 1
)

type TweenTweenPauseMode int
const (
  TweenTweenPauseBound TweenTweenPauseMode = 0
  TweenTweenPauseStop TweenTweenPauseMode = 1
  TweenTweenPauseProcess TweenTweenPauseMode = 2
)

type TweenTransitionType int
const (
  TweenTransLinear TweenTransitionType = 0
  TweenTransSine TweenTransitionType = 1
  TweenTransQuint TweenTransitionType = 2
  TweenTransQuart TweenTransitionType = 3
  TweenTransQuad TweenTransitionType = 4
  TweenTransExpo TweenTransitionType = 5
  TweenTransElastic TweenTransitionType = 6
  TweenTransCubic TweenTransitionType = 7
  TweenTransCirc TweenTransitionType = 8
  TweenTransBounce TweenTransitionType = 9
  TweenTransBack TweenTransitionType = 10
  TweenTransSpring TweenTransitionType = 11
)

type TweenEaseType int
const (
  TweenEaseIn TweenEaseType = 0
  TweenEaseOut TweenEaseType = 1
  TweenEaseInOut TweenEaseType = 2
  TweenEaseOutIn TweenEaseType = 3
)
