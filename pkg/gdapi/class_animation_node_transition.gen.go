// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationNodeTransition struct {
  AnimationNodeSync
}

func (me *AnimationNodeTransition) BaseClass() string {
  return "AnimationNodeTransition"
}



// Enums

func (me *AnimationNodeTransition) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeTransition) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeTransition) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationNodeTransition) SetInputCount(input_count int, )  {
  classNameV := StringNameFromStr("AnimationNodeTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input_count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeTransition) SetInputAsAutoAdvance(input int, enable bool, )  {
  classNameV := StringNameFromStr("AnimationNodeTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_as_auto_advance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeTransition) IsInputSetAsAutoAdvance(input int, ) bool {
  classNameV := StringNameFromStr("AnimationNodeTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_input_set_as_auto_advance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeTransition) SetInputReset(input int, enable bool, )  {
  classNameV := StringNameFromStr("AnimationNodeTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_reset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeTransition) IsInputReset(input int, ) bool {
  classNameV := StringNameFromStr("AnimationNodeTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_input_reset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&input), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeTransition) SetXfadeTime(time float32, )  {
  classNameV := StringNameFromStr("AnimationNodeTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_xfade_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeTransition) GetXfadeTime() float32 {
  classNameV := StringNameFromStr("AnimationNodeTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_xfade_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeTransition) SetXfadeCurve(curve Curve, )  {
  classNameV := StringNameFromStr("AnimationNodeTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_xfade_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeTransition) GetXfadeCurve() Curve {
  classNameV := StringNameFromStr("AnimationNodeTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_xfade_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  var ret Curve
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeTransition) SetAllowTransitionToSelf(enable bool, )  {
  classNameV := StringNameFromStr("AnimationNodeTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_transition_to_self")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeTransition) IsAllowTransitionToSelf() bool {
  classNameV := StringNameFromStr("AnimationNodeTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_allow_transition_to_self")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
