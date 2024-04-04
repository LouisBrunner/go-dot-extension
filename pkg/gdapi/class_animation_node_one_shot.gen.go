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

type AnimationNodeOneShot struct {
  AnimationNodeSync
}

func (me *AnimationNodeOneShot) BaseClass() string {
  return "AnimationNodeOneShot"
}

func NewAnimationNodeOneShot() *AnimationNodeOneShot {
  str := StringNameFromStr("AnimationNodeOneShot") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeOneShot{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *AnimationNodeOneShot) SetFadeinTime(time float64, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fadein_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeOneShot) GetFadeinTime() float64 {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fadein_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeOneShot) SetFadeinCurve(curve Curve, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fadein_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{curve.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeOneShot) GetFadeinCurve() Curve {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fadein_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurve()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeOneShot) SetFadeoutTime(time float64, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fadeout_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeOneShot) GetFadeoutTime() float64 {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fadeout_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeOneShot) SetFadeoutCurve(curve Curve, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fadeout_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{curve.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeOneShot) GetFadeoutCurve() Curve {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fadeout_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurve()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeOneShot) SetAutorestart(active bool, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autorestart")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeOneShot) HasAutorestart() bool {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_autorestart")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeOneShot) SetAutorestartDelay(time float64, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autorestart_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeOneShot) GetAutorestartDelay() float64 {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autorestart_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeOneShot) SetAutorestartRandomDelay(time float64, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autorestart_random_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeOneShot) GetAutorestartRandomDelay() float64 {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autorestart_random_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeOneShot) SetMixMode(mode AnimationNodeOneShotMixMode, )  {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mix_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1018899799) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeOneShot) GetMixMode() AnimationNodeOneShotMixMode {
  classNameV := StringNameFromStr("AnimationNodeOneShot")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mix_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3076550526) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationNodeOneShotMixMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
