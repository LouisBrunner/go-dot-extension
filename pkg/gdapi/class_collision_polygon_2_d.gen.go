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



// Enums

type CollisionPolygon2DBuildMode int
const (
  CollisionPolygon2DBuildModeBuildSolids CollisionPolygon2DBuildMode = 0
  CollisionPolygon2DBuildModeBuildSegments CollisionPolygon2DBuildMode = 1
)

func (me *CollisionPolygon2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CollisionPolygon2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CollisionPolygon2D) SetPolygon(polygon PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon2D) GetPolygon()  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon2D) SetBuildMode(build_mode CollisionPolygon2DBuildMode, )  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon2D) GetBuildMode()  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon2D) SetDisabled(disabled bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon2D) IsDisabled()  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon2D) SetOneWayCollision(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon2D) IsOneWayCollisionEnabled()  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon2D) SetOneWayCollisionMargin(margin float32, )  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon2D) GetOneWayCollisionMargin()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
