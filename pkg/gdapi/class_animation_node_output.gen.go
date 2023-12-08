// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeOutput struct {
  obj gdc.ObjectPtr
}

func createAnimationNodeOutput(obj gdc.ObjectPtr) *AnimationNodeOutput {
  return &AnimationNodeOutput{
    obj: obj,
  }
}

func (me *AnimationNodeOutput) BaseClass() string {
  return "AnimationNodeOutput"
}
