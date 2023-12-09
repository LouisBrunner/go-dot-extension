// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Curve3D struct {
  obj gdc.ObjectPtr
}

func (me *Curve3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Curve3D) BaseClass() string {
  return "Curve3D"
}



// Enums

func (me *Curve3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Curve3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Curve3D) GetPointCount()  {
  panic("TODO: implement")
}

func  (me *Curve3D) SetPointCount(count int, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) AddPoint(position Vector3, in Vector3, out Vector3, index int, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) SetPointPosition(idx int, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) GetPointPosition(idx int, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) SetPointTilt(idx int, tilt float32, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) GetPointTilt(idx int, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) SetPointIn(idx int, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) GetPointIn(idx int, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) SetPointOut(idx int, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) GetPointOut(idx int, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) RemovePoint(idx int, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) ClearPoints()  {
  panic("TODO: implement")
}

func  (me *Curve3D) Sample(idx int, t float32, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) Samplef(fofs float32, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) SetBakeInterval(distance float32, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) GetBakeInterval()  {
  panic("TODO: implement")
}

func  (me *Curve3D) SetUpVectorEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) IsUpVectorEnabled()  {
  panic("TODO: implement")
}

func  (me *Curve3D) GetBakedLength()  {
  panic("TODO: implement")
}

func  (me *Curve3D) SampleBaked(offset float32, cubic bool, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) SampleBakedWithRotation(offset float32, cubic bool, apply_tilt bool, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) SampleBakedUpVector(offset float32, apply_tilt bool, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) GetBakedPoints()  {
  panic("TODO: implement")
}

func  (me *Curve3D) GetBakedTilts()  {
  panic("TODO: implement")
}

func  (me *Curve3D) GetBakedUpVectors()  {
  panic("TODO: implement")
}

func  (me *Curve3D) GetClosestPoint(to_point Vector3, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) GetClosestOffset(to_point Vector3, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) Tessellate(max_stages int, tolerance_degrees float32, )  {
  panic("TODO: implement")
}

func  (me *Curve3D) TessellateEvenLength(max_stages int, tolerance_length float32, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
