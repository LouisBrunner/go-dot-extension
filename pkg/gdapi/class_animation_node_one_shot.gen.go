// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fadein_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeOneShot) GetFadeinTime() float32 {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fadein_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeOneShot) SetFadeinCurve(curve Curve, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fadein_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeOneShot) GetFadeinCurve() Curve {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fadein_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  var ret Curve
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeOneShot) SetFadeoutTime(time float32, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fadeout_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeOneShot) GetFadeoutTime() float32 {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fadeout_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeOneShot) SetFadeoutCurve(curve Curve, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fadeout_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeOneShot) GetFadeoutCurve() Curve {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fadeout_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  var ret Curve
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeOneShot) SetAutorestart(active bool, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autorestart")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeOneShot) HasAutorestart() bool {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_autorestart")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeOneShot) SetAutorestartDelay(time float32, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autorestart_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeOneShot) GetAutorestartDelay() float32 {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autorestart_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeOneShot) SetAutorestartRandomDelay(time float32, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autorestart_random_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeOneShot) GetAutorestartRandomDelay() float32 {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autorestart_random_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeOneShot) SetMixMode(mode AnimationNodeOneShotMixMode, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mix_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1018899799) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeOneShot) GetMixMode() AnimationNodeOneShotMixMode {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mix_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3076550526) // FIXME: should cache?
  var ret AnimationNodeOneShotMixMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *AnimationNodeOneShot) GetPropMixMode() int {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) SetPropMixMode(value int) {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) GetPropFadeinTime() float32 {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) SetPropFadeinTime(value float32) {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) GetPropFadeinCurve() Curve {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) SetPropFadeinCurve(value Curve) {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) GetPropFadeoutTime() float32 {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) SetPropFadeoutTime(value float32) {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) GetPropFadeoutCurve() Curve {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) SetPropFadeoutCurve(value Curve) {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) GetPropAutorestart() bool {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) SetPropAutorestart(value bool) {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) GetPropAutorestartDelay() float32 {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) SetPropAutorestartDelay(value float32) {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) GetPropAutorestartRandomDelay() float32 {
  panic("TODO: implement")
}

func (me *AnimationNodeOneShot) SetPropAutorestartRandomDelay(value float32) {
  panic("TODO: implement")
}