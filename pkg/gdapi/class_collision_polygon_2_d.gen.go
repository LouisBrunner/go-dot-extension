// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CollisionPolygon2D struct {
  obj gdc.ObjectPtr
}

func (me *CollisionPolygon2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CollisionPolygon2D) BaseClass() string {
  return "CollisionPolygon2D"
}

type CollisionPolygon2DBuildMode int
const (
  CollisionPolygon2DBuildSolids CollisionPolygon2DBuildMode = 0
  CollisionPolygon2DBuildSegments CollisionPolygon2DBuildMode = 1
)
