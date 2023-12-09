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



// Enums

type AnimationNodeBlendSpace1DBlendMode int
const (
  AnimationNodeBlendSpace1DBlendModeBlendModeInterpolated AnimationNodeBlendSpace1DBlendMode = 0
  AnimationNodeBlendSpace1DBlendModeBlendModeDiscrete AnimationNodeBlendSpace1DBlendMode = 1
  AnimationNodeBlendSpace1DBlendModeBlendModeDiscreteCarry AnimationNodeBlendSpace1DBlendMode = 2
)

func (me *AnimationNodeBlendSpace1D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeBlendSpace1D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeBlendSpace1D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AnimationNodeBlendSpace1D) AddBlendPoint(node AnimationRootNode, pos float32, at_index int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) SetBlendPointPosition(point int, pos float32, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) GetBlendPointPosition(point int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) SetBlendPointNode(point int, node AnimationRootNode, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) GetBlendPointNode(point int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) RemoveBlendPoint(point int, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) GetBlendPointCount()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) SetMinSpace(min_space float32, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) GetMinSpace()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) SetMaxSpace(max_space float32, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) GetMaxSpace()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) SetSnap(snap float32, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) GetSnap()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) SetValueLabel(text String, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) GetValueLabel()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) SetBlendMode(mode AnimationNodeBlendSpace1DBlendMode, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) GetBlendMode()  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) SetUseSync(enable bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeBlendSpace1D) IsUsingSync()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
