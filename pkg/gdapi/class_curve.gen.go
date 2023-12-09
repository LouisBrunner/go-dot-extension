// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Curve struct {
  obj gdc.ObjectPtr
}

func (me *Curve) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Curve) BaseClass() string {
  return "Curve"
}



// Enums

type CurveTangentMode int
const (
  CurveTangentModeTangentFree CurveTangentMode = 0
  CurveTangentModeTangentLinear CurveTangentMode = 1
  CurveTangentModeTangentModeCount CurveTangentMode = 2
)

func (me *Curve) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Curve) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Curve) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Curve) GetPointCount()  {
  panic("TODO: implement")
}

func  (me *Curve) SetPointCount(count int, )  {
  panic("TODO: implement")
}

func  (me *Curve) AddPoint(position Vector2, left_tangent float32, right_tangent float32, left_mode CurveTangentMode, right_mode CurveTangentMode, )  {
  panic("TODO: implement")
}

func  (me *Curve) RemovePoint(index int, )  {
  panic("TODO: implement")
}

func  (me *Curve) ClearPoints()  {
  panic("TODO: implement")
}

func  (me *Curve) GetPointPosition(index int, )  {
  panic("TODO: implement")
}

func  (me *Curve) SetPointValue(index int, y float32, )  {
  panic("TODO: implement")
}

func  (me *Curve) SetPointOffset(index int, offset float32, )  {
  panic("TODO: implement")
}

func  (me *Curve) Sample(offset float32, )  {
  panic("TODO: implement")
}

func  (me *Curve) SampleBaked(offset float32, )  {
  panic("TODO: implement")
}

func  (me *Curve) GetPointLeftTangent(index int, )  {
  panic("TODO: implement")
}

func  (me *Curve) GetPointRightTangent(index int, )  {
  panic("TODO: implement")
}

func  (me *Curve) GetPointLeftMode(index int, )  {
  panic("TODO: implement")
}

func  (me *Curve) GetPointRightMode(index int, )  {
  panic("TODO: implement")
}

func  (me *Curve) SetPointLeftTangent(index int, tangent float32, )  {
  panic("TODO: implement")
}

func  (me *Curve) SetPointRightTangent(index int, tangent float32, )  {
  panic("TODO: implement")
}

func  (me *Curve) SetPointLeftMode(index int, mode CurveTangentMode, )  {
  panic("TODO: implement")
}

func  (me *Curve) SetPointRightMode(index int, mode CurveTangentMode, )  {
  panic("TODO: implement")
}

func  (me *Curve) GetMinValue()  {
  panic("TODO: implement")
}

func  (me *Curve) SetMinValue(min float32, )  {
  panic("TODO: implement")
}

func  (me *Curve) GetMaxValue()  {
  panic("TODO: implement")
}

func  (me *Curve) SetMaxValue(max float32, )  {
  panic("TODO: implement")
}

func  (me *Curve) CleanDupes()  {
  panic("TODO: implement")
}

func  (me *Curve) Bake()  {
  panic("TODO: implement")
}

func  (me *Curve) GetBakeResolution()  {
  panic("TODO: implement")
}

func  (me *Curve) SetBakeResolution(resolution int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
