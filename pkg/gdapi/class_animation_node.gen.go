// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNode struct {
  obj gdc.ObjectPtr
}

func createAnimationNode(obj gdc.ObjectPtr) *AnimationNode {
  return &AnimationNode{
    obj: obj,
  }
}

func (me *AnimationNode) BaseClass() string {
  return "AnimationNode"
}
