// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeStateMachinePlayback struct {
  obj gdc.ObjectPtr
}

func createAnimationNodeStateMachinePlayback(obj gdc.ObjectPtr) *AnimationNodeStateMachinePlayback {
  return &AnimationNodeStateMachinePlayback{
    obj: obj,
  }
}

func (me *AnimationNodeStateMachinePlayback) BaseClass() string {
  return "AnimationNodeStateMachinePlayback"
}
