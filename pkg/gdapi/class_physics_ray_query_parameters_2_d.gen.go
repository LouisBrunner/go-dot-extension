// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsRayQueryParameters2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsRayQueryParameters2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsRayQueryParameters2D) BaseClass() string {
  return "PhysicsRayQueryParameters2D"
}



// Enums

func (me *PhysicsRayQueryParameters2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsRayQueryParameters2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsRayQueryParameters2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  PhysicsRayQueryParameters2DCreate(from Vector2, to Vector2, collision_mask int, exclude RID, ) PhysicsRayQueryParameters2D {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1118143851) // FIXME: should cache?
  var ret PhysicsRayQueryParameters2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(to.AsCTypePtr()), gdc.ConstTypePtr(&collision_mask), gdc.ConstTypePtr(exclude.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsRayQueryParameters2D) SetFrom(from Vector2, )  {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_from")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsRayQueryParameters2D) GetFrom() Vector2 {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_from")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsRayQueryParameters2D) SetTo(to Vector2, )  {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_to")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(to.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsRayQueryParameters2D) GetTo() Vector2 {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_to")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsRayQueryParameters2D) SetCollisionMask(collision_mask int, )  {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsRayQueryParameters2D) GetCollisionMask() int {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsRayQueryParameters2D) SetExclude(exclude RID, )  {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exclude")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(exclude.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsRayQueryParameters2D) GetExclude() RID {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_exclude")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsRayQueryParameters2D) SetCollideWithBodies(enable bool, )  {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collide_with_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsRayQueryParameters2D) IsCollideWithBodiesEnabled() bool {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collide_with_bodies_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsRayQueryParameters2D) SetCollideWithAreas(enable bool, )  {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collide_with_areas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsRayQueryParameters2D) IsCollideWithAreasEnabled() bool {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collide_with_areas_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsRayQueryParameters2D) SetHitFromInside(enable bool, )  {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hit_from_inside")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsRayQueryParameters2D) IsHitFromInsideEnabled() bool {
  classNameV := StringNameFromStr("PhysicsRayQueryParameters2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_hit_from_inside_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
