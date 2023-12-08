// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

type AnimationNodeStateMachineTransitionSwitchMode int
const (
  AnimationNodeStateMachineTransitionSwitchModeImmediate AnimationNodeStateMachineTransitionSwitchMode = 0
  AnimationNodeStateMachineTransitionSwitchModeSync AnimationNodeStateMachineTransitionSwitchMode = 1
  AnimationNodeStateMachineTransitionSwitchModeAtEnd AnimationNodeStateMachineTransitionSwitchMode = 2
)

type AnimationNodeStateMachineTransitionAdvanceMode int
const (
  AnimationNodeStateMachineTransitionAdvanceModeDisabled AnimationNodeStateMachineTransitionAdvanceMode = 0
  AnimationNodeStateMachineTransitionAdvanceModeEnabled AnimationNodeStateMachineTransitionAdvanceMode = 1
  AnimationNodeStateMachineTransitionAdvanceModeAuto AnimationNodeStateMachineTransitionAdvanceMode = 2
)
