// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PointMesh struct {
  obj gdc.ObjectPtr
}

func (me *PointMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PointMesh) BaseClass() string {
  return "PointMesh"
}
