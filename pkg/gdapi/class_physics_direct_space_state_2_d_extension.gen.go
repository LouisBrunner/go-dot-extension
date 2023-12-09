// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports



  "unsafe"



  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectSpaceState2DExtension struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectSpaceState2DExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectSpaceState2DExtension) BaseClass() string {
  return "PhysicsDirectSpaceState2DExtension"
}



// Enums

func (me *PhysicsDirectSpaceState2DExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectSpaceState2DExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsDirectSpaceState2DExtension) XIntersectRay(from Vector2, to Vector2, collision_mask int, collide_with_bodies bool, collide_with_areas bool, hit_from_inside bool, result *PhysicsServer2DExtensionRayResult, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState2DExtension) XIntersectPoint(position Vector2, canvas_instance_id int, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results *PhysicsServer2DExtensionShapeResult, max_results int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState2DExtension) XIntersectShape(shape_rid RID, transform Transform2D, motion Vector2, margin float32, collision_mask int, collide_with_bodies bool, collide_with_areas bool, result *PhysicsServer2DExtensionShapeResult, max_results int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState2DExtension) XCastMotion(shape_rid RID, transform Transform2D, motion Vector2, margin float32, collision_mask int, collide_with_bodies bool, collide_with_areas bool, closest_safe *float32, closest_unsafe *float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState2DExtension) XCollideShape(shape_rid RID, transform Transform2D, motion Vector2, margin float32, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results unsafe.Pointer, max_results int, result_count *int32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState2DExtension) XRestInfo(shape_rid RID, transform Transform2D, motion Vector2, margin float32, collision_mask int, collide_with_bodies bool, collide_with_areas bool, rest_info *PhysicsServer2DExtensionShapeRestInfo, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState2DExtension) IsBodyExcludedFromQuery(body RID, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
