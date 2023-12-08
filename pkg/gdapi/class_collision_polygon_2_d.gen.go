// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  CollisionPolygon2DBuildModeBuildSolids CollisionPolygon2DBuildMode = 0
  CollisionPolygon2DBuildModeBuildSegments CollisionPolygon2DBuildMode = 1
)

func  (me *CollisionPolygon2D) SetPolygon(polygon PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionPolygon2D) GetPolygon() { // TODO: return value
  // TODO: implement
}

func  (me *CollisionPolygon2D) SetBuildMode(build_mode CollisionPolygon2DBuildMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionPolygon2D) GetBuildMode() { // TODO: return value
  // TODO: implement
}

func  (me *CollisionPolygon2D) SetDisabled(disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionPolygon2D) IsDisabled() { // TODO: return value
  // TODO: implement
}

func  (me *CollisionPolygon2D) SetOneWayCollision(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionPolygon2D) IsOneWayCollisionEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CollisionPolygon2D) SetOneWayCollisionMargin(margin float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CollisionPolygon2D) GetOneWayCollisionMargin() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
