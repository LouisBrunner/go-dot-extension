// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGPolygon3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGPolygon3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGPolygon3D) BaseClass() string {
  return "CSGPolygon3D"
}

type CSGPolygon3DMode int
const (
  CSGPolygon3DModeDepth CSGPolygon3DMode = 0
  CSGPolygon3DModeSpin CSGPolygon3DMode = 1
  CSGPolygon3DModePath CSGPolygon3DMode = 2
)

type CSGPolygon3DPathRotation int
const (
  CSGPolygon3DPathRotationPolygon CSGPolygon3DPathRotation = 0
  CSGPolygon3DPathRotationPath CSGPolygon3DPathRotation = 1
  CSGPolygon3DPathRotationPathFollow CSGPolygon3DPathRotation = 2
)

type CSGPolygon3DPathIntervalType int
const (
  CSGPolygon3DPathIntervalDistance CSGPolygon3DPathIntervalType = 0
  CSGPolygon3DPathIntervalSubdivide CSGPolygon3DPathIntervalType = 1
)
