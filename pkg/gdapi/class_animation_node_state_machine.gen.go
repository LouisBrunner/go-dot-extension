// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeStateMachine struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeStateMachine) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeStateMachine) BaseClass() string {
  return "AnimationNodeStateMachine"
}

type AnimationNodeStateMachineStateMachineType int
const (
  AnimationNodeStateMachineStateMachineTypeRoot AnimationNodeStateMachineStateMachineType = 0
  AnimationNodeStateMachineStateMachineTypeNested AnimationNodeStateMachineStateMachineType = 1
  AnimationNodeStateMachineStateMachineTypeGrouped AnimationNodeStateMachineStateMachineType = 2
)
