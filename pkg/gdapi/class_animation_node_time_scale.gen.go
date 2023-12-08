// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeTimeScale struct {
  obj gdc.ObjectPtr
}

func createAnimationNodeTimeScale(obj gdc.ObjectPtr) *AnimationNodeTimeScale {
  return &AnimationNodeTimeScale{
    obj: obj,
  }
}

func (me *AnimationNodeTimeScale) BaseClass() string {
  return "AnimationNodeTimeScale"
}
