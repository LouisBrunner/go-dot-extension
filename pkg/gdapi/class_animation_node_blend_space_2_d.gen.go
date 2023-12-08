// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeBlendSpace2D struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeBlendSpace2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeBlendSpace2D) BaseClass() string {
  return "AnimationNodeBlendSpace2D"
}

type AnimationNodeBlendSpace2DBlendMode int
const (
  AnimationNodeBlendSpace2DBlendModeInterpolated AnimationNodeBlendSpace2DBlendMode = 0
  AnimationNodeBlendSpace2DBlendModeDiscrete AnimationNodeBlendSpace2DBlendMode = 1
  AnimationNodeBlendSpace2DBlendModeDiscreteCarry AnimationNodeBlendSpace2DBlendMode = 2
)
