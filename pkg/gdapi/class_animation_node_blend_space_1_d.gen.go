// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeBlendSpace1D struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeBlendSpace1D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeBlendSpace1D) BaseClass() string {
  return "AnimationNodeBlendSpace1D"
}

type AnimationNodeBlendSpace1DBlendMode int
const (
  AnimationNodeBlendSpace1DBlendModeInterpolated AnimationNodeBlendSpace1DBlendMode = 0
  AnimationNodeBlendSpace1DBlendModeDiscrete AnimationNodeBlendSpace1DBlendMode = 1
  AnimationNodeBlendSpace1DBlendModeDiscreteCarry AnimationNodeBlendSpace1DBlendMode = 2
)
