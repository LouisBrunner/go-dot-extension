// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CollisionPolygon2D struct {
  obj gdc.ObjectPtr
}

func createCollisionPolygon2D(obj gdc.ObjectPtr) *CollisionPolygon2D {
  return &CollisionPolygon2D{
    obj: obj,
  }
}

func (me *CollisionPolygon2D) BaseClass() string {
  return "CollisionPolygon2D"
}
