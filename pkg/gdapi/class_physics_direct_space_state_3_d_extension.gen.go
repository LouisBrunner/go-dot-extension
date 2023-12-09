// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports




  "unsafe"


  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectSpaceState3DExtension struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectSpaceState3DExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectSpaceState3DExtension) BaseClass() string {
  return "PhysicsDirectSpaceState3DExtension"
}



// Enums

func (me *PhysicsDirectSpaceState3DExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsDirectSpaceState3DExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectSpaceState3DExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsDirectSpaceState3DExtension) XIntersectRay(from Vector3, to Vector3, collision_mask int, collide_with_bodies bool, collide_with_areas bool, hit_from_inside bool, hit_back_faces bool, pick_ray bool, result *PhysicsServer3DExtensionRayResult, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState3DExtension) XIntersectPoint(position Vector3, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results *PhysicsServer3DExtensionShapeResult, max_results int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState3DExtension) XIntersectShape(shape_rid RID, transform Transform3D, motion Vector3, margin float32, collision_mask int, collide_with_bodies bool, collide_with_areas bool, result_count *PhysicsServer3DExtensionShapeResult, max_results int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState3DExtension) XCastMotion(shape_rid RID, transform Transform3D, motion Vector3, margin float32, collision_mask int, collide_with_bodies bool, collide_with_areas bool, closest_safe *float32, closest_unsafe *float32, info *PhysicsServer3DExtensionShapeRestInfo, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState3DExtension) XCollideShape(shape_rid RID, transform Transform3D, motion Vector3, margin float32, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results unsafe.Pointer, max_results int, result_count *int32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState3DExtension) XRestInfo(shape_rid RID, transform Transform3D, motion Vector3, margin float32, collision_mask int, collide_with_bodies bool, collide_with_areas bool, rest_info *PhysicsServer3DExtensionShapeRestInfo, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState3DExtension) XGetClosestPointToObjectVolume(object RID, point Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectSpaceState3DExtension) IsBodyExcludedFromQuery(body RID, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
