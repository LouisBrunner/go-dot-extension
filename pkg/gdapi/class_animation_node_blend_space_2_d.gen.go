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



// Enums

type AnimationNodeBlendSpace2DBlendMode int
const (
  AnimationNodeBlendSpace2DBlendModeBlendModeInterpolated AnimationNodeBlendSpace2DBlendMode = 0
  AnimationNodeBlendSpace2DBlendModeBlendModeDiscrete AnimationNodeBlendSpace2DBlendMode = 1
  AnimationNodeBlendSpace2DBlendModeBlendModeDiscreteCarry AnimationNodeBlendSpace2DBlendMode = 2
)

func (me *AnimationNodeBlendSpace2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeBlendSpace2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AnimationNodeBlendSpace2D) AddBlendPoint(node AnimationRootNode, pos Vector2, at_index int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) SetBlendPointPosition(point int, pos Vector2, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointPosition(point int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) SetBlendPointNode(point int, node AnimationRootNode, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointNode(point int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) RemoveBlendPoint(point int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointCount()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) AddTriangle(x int, y int, z int, at_index int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) GetTrianglePoint(triangle int, point int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) RemoveTriangle(triangle int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) GetTriangleCount()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) SetMinSpace(min_space Vector2, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) GetMinSpace()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) SetMaxSpace(max_space Vector2, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) GetMaxSpace()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) SetSnap(snap Vector2, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) GetSnap()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) SetXLabel(text String, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) GetXLabel()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) SetYLabel(text String, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) GetYLabel()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) SetAutoTriangles(enable bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) GetAutoTriangles()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) SetBlendMode(mode AnimationNodeBlendSpace2DBlendMode, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) GetBlendMode()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) SetUseSync(enable bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace2D) IsUsingSync()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
