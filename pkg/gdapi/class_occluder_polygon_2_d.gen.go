// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  OccluderPolygon2DCullDisabled OccluderPolygon2DCullMode = 0
  OccluderPolygon2DCullClockwise OccluderPolygon2DCullMode = 1
  OccluderPolygon2DCullCounterClockwise OccluderPolygon2DCullMode = 2
)
