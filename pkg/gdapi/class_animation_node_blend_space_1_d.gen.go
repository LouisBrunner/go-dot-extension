// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeBlendSpace1D struct {
  obj gdc.ObjectPtr
}

func createAnimationNodeBlendSpace1D(obj gdc.ObjectPtr) *AnimationNodeBlendSpace1D {
  return &AnimationNodeBlendSpace1D{
    obj: obj,
  }
}

func (me *AnimationNodeBlendSpace1D) BaseClass() string {
  return "AnimationNodeBlendSpace1D"
}
