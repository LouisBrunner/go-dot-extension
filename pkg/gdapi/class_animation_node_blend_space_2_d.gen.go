// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  AnimationNodeBlendSpace2DBlendModeBlendModeInterpolated AnimationNodeBlendSpace2DBlendMode = 0
  AnimationNodeBlendSpace2DBlendModeBlendModeDiscrete AnimationNodeBlendSpace2DBlendMode = 1
  AnimationNodeBlendSpace2DBlendModeBlendModeDiscreteCarry AnimationNodeBlendSpace2DBlendMode = 2
)

func  (me *AnimationNodeBlendSpace2D) AddBlendPoint(node AnimationRootNode, pos Vector2, at_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) SetBlendPointPosition(point int, pos Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointPosition(point int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) SetBlendPointNode(point int, node AnimationRootNode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointNode(point int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) RemoveBlendPoint(point int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointCount() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) AddTriangle(x int, y int, z int, at_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) GetTrianglePoint(triangle int, point int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) RemoveTriangle(triangle int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) GetTriangleCount() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) SetMinSpace(min_space Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) GetMinSpace() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) SetMaxSpace(max_space Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) GetMaxSpace() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) SetSnap(snap Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) GetSnap() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) SetXLabel(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) GetXLabel() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) SetYLabel(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) GetYLabel() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) SetAutoTriangles(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) GetAutoTriangles() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) SetBlendMode(mode AnimationNodeBlendSpace2DBlendMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) GetBlendMode() { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) SetUseSync(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AnimationNodeBlendSpace2D) IsUsingSync() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
