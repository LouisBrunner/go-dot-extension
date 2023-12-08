// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeOneShot struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeOneShot) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeOneShot) BaseClass() string {
  return "AnimationNodeOneShot"
}

type AnimationNodeOneShotOneShotRequest int
const (
  AnimationNodeOneShotOneShotRequestOneShotRequestNone AnimationNodeOneShotOneShotRequest = 0
  AnimationNodeOneShotOneShotRequestOneShotRequestFire AnimationNodeOneShotOneShotRequest = 1
  AnimationNodeOneShotOneShotRequestOneShotRequestAbort AnimationNodeOneShotOneShotRequest = 2
  AnimationNodeOneShotOneShotRequestOneShotRequestFadeOut AnimationNodeOneShotOneShotRequest = 3
)

type AnimationNodeOneShotMixMode int
const (
  AnimationNodeOneShotMixModeMixModeBlend AnimationNodeOneShotMixMode = 0
  AnimationNodeOneShotMixModeMixModeAdd AnimationNodeOneShotMixMode = 1
)

func  (me *AnimationNodeOneShot) SetFadeinTime(time float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) GetFadeinTime() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) SetFadeinCurve(curve Curve, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) GetFadeinCurve() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) SetFadeoutTime(time float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) GetFadeoutTime() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) SetFadeoutCurve(curve Curve, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) GetFadeoutCurve() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) SetAutorestart(active bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) HasAutorestart() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) SetAutorestartDelay(time float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) GetAutorestartDelay() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) SetAutorestartRandomDelay(time float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) GetAutorestartRandomDelay() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) SetMixMode(mode AnimationNodeOneShotMixMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeOneShot) GetMixMode() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
