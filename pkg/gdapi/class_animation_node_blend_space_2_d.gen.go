// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeBlendSpace2D struct {
  obj gdc.ObjectPtr
}

func createAnimationNodeBlendSpace2D(obj gdc.ObjectPtr) *AnimationNodeBlendSpace2D {
  return &AnimationNodeBlendSpace2D{
    obj: obj,
  }
}

func (me *AnimationNodeBlendSpace2D) BaseClass() string {
  return "AnimationNodeBlendSpace2D"
}
