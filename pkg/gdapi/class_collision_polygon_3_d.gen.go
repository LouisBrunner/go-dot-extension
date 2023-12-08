// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CollisionPolygon3D struct {
  obj gdc.ObjectPtr
}

func createCollisionPolygon3D(obj gdc.ObjectPtr) *CollisionPolygon3D {
  return &CollisionPolygon3D{
    obj: obj,
  }
}

func (me *CollisionPolygon3D) BaseClass() string {
  return "CollisionPolygon3D"
}
