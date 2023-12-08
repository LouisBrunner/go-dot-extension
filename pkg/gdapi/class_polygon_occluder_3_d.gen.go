// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PolygonOccluder3D struct {
  obj gdc.ObjectPtr
}

func (me *PolygonOccluder3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PolygonOccluder3D) BaseClass() string {
  return "PolygonOccluder3D"
}
