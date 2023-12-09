// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeStateMachineTransition struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeStateMachineTransition) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeStateMachineTransition) BaseClass() string {
  return "AnimationNodeStateMachineTransition"
}



// Enums

type AnimationNodeStateMachineTransitionSwitchMode int
const (
  AnimationNodeStateMachineTransitionSwitchModeSwitchModeImmediate AnimationNodeStateMachineTransitionSwitchMode = 0
  AnimationNodeStateMachineTransitionSwitchModeSwitchModeSync AnimationNodeStateMachineTransitionSwitchMode = 1
  AnimationNodeStateMachineTransitionSwitchModeSwitchModeAtEnd AnimationNodeStateMachineTransitionSwitchMode = 2
)

type AnimationNodeStateMachineTransitionAdvanceMode int
const (
  AnimationNodeStateMachineTransitionAdvanceModeAdvanceModeDisabled AnimationNodeStateMachineTransitionAdvanceMode = 0
  AnimationNodeStateMachineTransitionAdvanceModeAdvanceModeEnabled AnimationNodeStateMachineTransitionAdvanceMode = 1
  AnimationNodeStateMachineTransitionAdvanceModeAdvanceModeAuto AnimationNodeStateMachineTransitionAdvanceMode = 2
)

func (me *AnimationNodeStateMachineTransition) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeStateMachineTransition) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeStateMachineTransition) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AnimationNodeStateMachineTransition) SetSwitchMode(mode AnimationNodeStateMachineTransitionSwitchMode, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) GetSwitchMode()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) SetAdvanceMode(mode AnimationNodeStateMachineTransitionAdvanceMode, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) GetAdvanceMode()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) SetAdvanceCondition(name StringName, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) GetAdvanceCondition()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) SetXfadeTime(secs float32, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) GetXfadeTime()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) SetXfadeCurve(curve Curve, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) GetXfadeCurve()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) SetReset(reset bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) IsReset()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) SetPriority(priority int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) GetPriority()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) SetAdvanceExpression(text String, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeStateMachineTransition) GetAdvanceExpression()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
