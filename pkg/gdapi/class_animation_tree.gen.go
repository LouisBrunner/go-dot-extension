// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationTree struct {
  obj gdc.ObjectPtr
}

func createAnimationTree(obj gdc.ObjectPtr) *AnimationTree {
  return &AnimationTree{
    obj: obj,
  }
}

func (me *AnimationTree) BaseClass() string {
  return "AnimationTree"
}
