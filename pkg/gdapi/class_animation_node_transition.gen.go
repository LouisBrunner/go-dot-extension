// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeTransition struct {
  obj gdc.ObjectPtr
}

func createAnimationNodeTransition(obj gdc.ObjectPtr) *AnimationNodeTransition {
  return &AnimationNodeTransition{
    obj: obj,
  }
}

func (me *AnimationNodeTransition) BaseClass() string {
  return "AnimationNodeTransition"
}
