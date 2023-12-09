// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeTransition struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeTransition) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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
  panic("TODO: implement")
}

func  (me *AnimationNodeTransition) SetInputAsAutoAdvance(input int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeTransition) IsInputSetAsAutoAdvance(input int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeTransition) SetInputReset(input int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeTransition) IsInputReset(input int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeTransition) SetXfadeTime(time float32, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeTransition) GetXfadeTime()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeTransition) SetXfadeCurve(curve Curve, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeTransition) GetXfadeCurve()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeTransition) SetAllowTransitionToSelf(enable bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeTransition) IsAllowTransitionToSelf()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
