// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectSpaceState2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectSpaceState2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectSpaceState2D) BaseClass() string {
  return "PhysicsDirectSpaceState2D"
}



// Enums

func (me *PhysicsDirectSpaceState2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsDirectSpaceState2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectSpaceState2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsDirectSpaceState2D) IntersectPoint(parameters PhysicsPointQueryParameters2D, max_results int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState2D) IntersectRay(parameters PhysicsRayQueryParameters2D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState2D) IntersectShape(parameters PhysicsShapeQueryParameters2D, max_results int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState2D) CastMotion(parameters PhysicsShapeQueryParameters2D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState2D) CollideShape(parameters PhysicsShapeQueryParameters2D, max_results int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState2D) GetRestInfo(parameters PhysicsShapeQueryParameters2D, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
