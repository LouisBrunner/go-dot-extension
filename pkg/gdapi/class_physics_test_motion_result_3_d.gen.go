// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type PhysicsTestMotionResult3D struct {
  RefCounted
}

func (me *PhysicsTestMotionResult3D) BaseClass() string {
  return "PhysicsTestMotionResult3D"
}

func NewPhysicsTestMotionResult3D() *PhysicsTestMotionResult3D {
  str := StringNameFromStr("PhysicsTestMotionResult3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicsTestMotionResult3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PhysicsTestMotionResult3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsTestMotionResult3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsTestMotionResult3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsTestMotionResult3D) GetTravel() Vector3 {
  classNameV := StringNameFromStr("PhysicsTestMotionResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_travel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionResult3D) GetRemainder() Vector3 {
  classNameV := StringNameFromStr("PhysicsTestMotionResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_remainder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionResult3D) GetCollisionSafeFraction() float64 {
  classNameV := StringNameFromStr("PhysicsTestMotionResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_safe_fraction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionResult3D) GetCollisionUnsafeFraction() float64 {
  classNameV := StringNameFromStr("PhysicsTestMotionResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_unsafe_fraction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionResult3D) GetCollisionCount() int64 {
  classNameV := StringNameFromStr("PhysicsTestMotionResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionResult3D) GetCollisionPoint(collision_index int64, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsTestMotionResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1914908202) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionResult3D) GetCollisionNormal(collision_index int64, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsTestMotionResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1914908202) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionResult3D) GetColliderVelocity(collision_index int64, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsTestMotionResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collider_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1914908202) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionResult3D) GetColliderId(collision_index int64, ) int64 {
  classNameV := StringNameFromStr("PhysicsTestMotionResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collider_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1591665591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionResult3D) GetColliderRid(collision_index int64, ) RID {
  classNameV := StringNameFromStr("PhysicsTestMotionResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collider_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1231817359) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionResult3D) GetCollider(collision_index int64, ) Object {
  classNameV := StringNameFromStr("PhysicsTestMotionResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collider")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2639523548) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewObject()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionResult3D) GetColliderShape(collision_index int64, ) int64 {
  classNameV := StringNameFromStr("PhysicsTestMotionResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collider_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1591665591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionResult3D) GetCollisionLocalShape(collision_index int64, ) int64 {
  classNameV := StringNameFromStr("PhysicsTestMotionResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_local_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1591665591) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionResult3D) GetCollisionDepth(collision_index int64, ) float64 {
  classNameV := StringNameFromStr("PhysicsTestMotionResult3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 218038398) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
