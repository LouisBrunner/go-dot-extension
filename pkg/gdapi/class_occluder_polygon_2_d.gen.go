// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OccluderPolygon2D struct {
  obj gdc.ObjectPtr
}

func (me *OccluderPolygon2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OccluderPolygon2D) BaseClass() string {
  return "OccluderPolygon2D"
}



// Enums

type OccluderPolygon2DCullMode int
const (
  OccluderPolygon2DCullModeCullDisabled OccluderPolygon2DCullMode = 0
  OccluderPolygon2DCullModeCullClockwise OccluderPolygon2DCullMode = 1
  OccluderPolygon2DCullModeCullCounterClockwise OccluderPolygon2DCullMode = 2
)

func (me *OccluderPolygon2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OccluderPolygon2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OccluderPolygon2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *OccluderPolygon2D) SetClosed(closed bool, )  {
  panic("TODO: implement")
}

func  (me *OccluderPolygon2D) IsClosed()  {
  panic("TODO: implement")
}

func  (me *OccluderPolygon2D) SetCullMode(cull_mode OccluderPolygon2DCullMode, )  {
  panic("TODO: implement")
}

func  (me *OccluderPolygon2D) GetCullMode()  {
  panic("TODO: implement")
}

func  (me *OccluderPolygon2D) SetPolygon(polygon PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *OccluderPolygon2D) GetPolygon()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
