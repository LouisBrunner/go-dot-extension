// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeStateMachine struct {
  obj gdc.ObjectPtr
}

func createAnimationNodeStateMachine(obj gdc.ObjectPtr) *AnimationNodeStateMachine {
  return &AnimationNodeStateMachine{
    obj: obj,
  }
}

func (me *AnimationNodeStateMachine) BaseClass() string {
  return "AnimationNodeStateMachine"
}
