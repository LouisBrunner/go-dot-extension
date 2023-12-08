// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConcavePolygonShape3D struct {
  obj gdc.ObjectPtr
}

func (me *ConcavePolygonShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConcavePolygonShape3D) BaseClass() string {
  return "ConcavePolygonShape3D"
}
