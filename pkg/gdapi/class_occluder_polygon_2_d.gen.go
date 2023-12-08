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

type OccluderPolygon2DCullMode int
const (
  OccluderPolygon2DCullModeCullDisabled OccluderPolygon2DCullMode = 0
  OccluderPolygon2DCullModeCullClockwise OccluderPolygon2DCullMode = 1
  OccluderPolygon2DCullModeCullCounterClockwise OccluderPolygon2DCullMode = 2
)

func  (me *OccluderPolygon2D) SetClosed(closed bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *OccluderPolygon2D) IsClosed() { // TODO: return value
  // TODO: implement
}

func  (me *OccluderPolygon2D) SetCullMode(cull_mode OccluderPolygon2DCullMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *OccluderPolygon2D) GetCullMode() { // TODO: return value
  // TODO: implement
}

func  (me *OccluderPolygon2D) SetPolygon(polygon PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *OccluderPolygon2D) GetPolygon() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
