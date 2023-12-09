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



// Enums

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

func  (me *AnimationNodeOneShot) SetFadeinTime(time float32, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) GetFadeinTime()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) SetFadeinCurve(curve Curve, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) GetFadeinCurve()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) SetFadeoutTime(time float32, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) GetFadeoutTime()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) SetFadeoutCurve(curve Curve, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) GetFadeoutCurve()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) SetAutorestart(active bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) HasAutorestart()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) SetAutorestartDelay(time float32, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) GetAutorestartDelay()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) SetAutorestartRandomDelay(time float32, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) GetAutorestartRandomDelay()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) SetMixMode(mode AnimationNodeOneShotMixMode, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeOneShot) GetMixMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
