// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeAnimation struct {
  obj gdc.ObjectPtr
}

func createAnimationNodeAnimation(obj gdc.ObjectPtr) *AnimationNodeAnimation {
  return &AnimationNodeAnimation{
    obj: obj,
  }
}

func (me *AnimationNodeAnimation) BaseClass() string {
  return "AnimationNodeAnimation"
}
