// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CollisionPolygon3D struct {
  obj gdc.ObjectPtr
}

func (me *CollisionPolygon3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CollisionPolygon3D) BaseClass() string {
  return "CollisionPolygon3D"
}
