// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectSpaceState3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectSpaceState3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectSpaceState3D) BaseClass() string {
  return "PhysicsDirectSpaceState3D"
}



// Enums

func (me *PhysicsDirectSpaceState3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectSpaceState3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsDirectSpaceState3D) IntersectPoint(parameters PhysicsPointQueryParameters3D, max_results int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState3D) IntersectRay(parameters PhysicsRayQueryParameters3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState3D) IntersectShape(parameters PhysicsShapeQueryParameters3D, max_results int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState3D) CastMotion(parameters PhysicsShapeQueryParameters3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState3D) CollideShape(parameters PhysicsShapeQueryParameters3D, max_results int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState3D) GetRestInfo(parameters PhysicsShapeQueryParameters3D, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
