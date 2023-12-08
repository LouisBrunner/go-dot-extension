// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNode struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNode) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNode) BaseClass() string {
  return "AnimationNode"
}

type AnimationNodeFilterAction int
const (
  AnimationNodeFilterIgnore AnimationNodeFilterAction = 0
  AnimationNodeFilterPass AnimationNodeFilterAction = 1
  AnimationNodeFilterStop AnimationNodeFilterAction = 2
  AnimationNodeFilterBlend AnimationNodeFilterAction = 3
)
