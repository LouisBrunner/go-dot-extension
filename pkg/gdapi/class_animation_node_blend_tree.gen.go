// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeBlendTree struct {
  obj gdc.ObjectPtr
}

func createAnimationNodeBlendTree(obj gdc.ObjectPtr) *AnimationNodeBlendTree {
  return &AnimationNodeBlendTree{
    obj: obj,
  }
}

func (me *AnimationNodeBlendTree) BaseClass() string {
  return "AnimationNodeBlendTree"
}
