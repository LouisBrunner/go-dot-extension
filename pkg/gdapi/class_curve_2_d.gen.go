// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Curve2D struct {
  obj gdc.ObjectPtr
}

func (me *Curve2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Curve2D) BaseClass() string {
  return "Curve2D"
}



// Enums

func (me *Curve2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Curve2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Curve2D) GetPointCount()  {
  panic("TODO: implement")
}

func  (me *Curve2D) SetPointCount(count int, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) AddPoint(position Vector2, in Vector2, out Vector2, index int, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) SetPointPosition(idx int, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) GetPointPosition(idx int, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) SetPointIn(idx int, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) GetPointIn(idx int, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) SetPointOut(idx int, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) GetPointOut(idx int, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) RemovePoint(idx int, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) ClearPoints()  {
  panic("TODO: implement")
}

func  (me *Curve2D) Sample(idx int, t float32, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) Samplef(fofs float32, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) SetBakeInterval(distance float32, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) GetBakeInterval()  {
  panic("TODO: implement")
}

func  (me *Curve2D) GetBakedLength()  {
  panic("TODO: implement")
}

func  (me *Curve2D) SampleBaked(offset float32, cubic bool, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) SampleBakedWithRotation(offset float32, cubic bool, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) GetBakedPoints()  {
  panic("TODO: implement")
}

func  (me *Curve2D) GetClosestPoint(to_point Vector2, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) GetClosestOffset(to_point Vector2, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) Tessellate(max_stages int, tolerance_degrees float32, )  {
  panic("TODO: implement")
}

func  (me *Curve2D) TessellateEvenLength(max_stages int, tolerance_length float32, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
