// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  AnimationNodeBlendSpace1DBlendModeBlendModeInterpolated AnimationNodeBlendSpace1DBlendMode = 0
  AnimationNodeBlendSpace1DBlendModeBlendModeDiscrete AnimationNodeBlendSpace1DBlendMode = 1
  AnimationNodeBlendSpace1DBlendModeBlendModeDiscreteCarry AnimationNodeBlendSpace1DBlendMode = 2
)

func  (me *AnimationNodeBlendSpace1D) AddBlendPoint(node AnimationRootNode, pos float32, at_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) SetBlendPointPosition(point int, pos float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) GetBlendPointPosition(point int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) SetBlendPointNode(point int, node AnimationRootNode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) GetBlendPointNode(point int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) RemoveBlendPoint(point int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) GetBlendPointCount() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) SetMinSpace(min_space float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) GetMinSpace() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) SetMaxSpace(max_space float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) GetMaxSpace() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) SetSnap(snap float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) GetSnap() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) SetValueLabel(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) GetValueLabel() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) SetBlendMode(mode AnimationNodeBlendSpace1DBlendMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) GetBlendMode() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) SetUseSync(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace1D) IsUsingSync() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
