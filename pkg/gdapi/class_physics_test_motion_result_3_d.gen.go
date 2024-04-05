// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForPhysicsTestMotionResult3DList struct {
  fnGetTravel gdc.MethodBindPtr
  fnGetRemainder gdc.MethodBindPtr
  fnGetCollisionSafeFraction gdc.MethodBindPtr
  fnGetCollisionUnsafeFraction gdc.MethodBindPtr
  fnGetCollisionCount gdc.MethodBindPtr
  fnGetCollisionPoint gdc.MethodBindPtr
  fnGetCollisionNormal gdc.MethodBindPtr
  fnGetColliderVelocity gdc.MethodBindPtr
  fnGetColliderId gdc.MethodBindPtr
  fnGetColliderRid gdc.MethodBindPtr
  fnGetCollider gdc.MethodBindPtr
  fnGetColliderShape gdc.MethodBindPtr
  fnGetCollisionLocalShape gdc.MethodBindPtr
  fnGetCollisionDepth gdc.MethodBindPtr
}

var ptrsForPhysicsTestMotionResult3D ptrsForPhysicsTestMotionResult3DList

func initPhysicsTestMotionResult3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PhysicsTestMotionResult3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_travel")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionResult3D.fnGetTravel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("get_remainder")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionResult3D.fnGetRemainder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("get_collision_safe_fraction")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionResult3D.fnGetCollisionSafeFraction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_collision_unsafe_fraction")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionResult3D.fnGetCollisionUnsafeFraction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_collision_count")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionResult3D.fnGetCollisionCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_collision_point")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionResult3D.fnGetCollisionPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1914908202))
  }
  {
    methodName := StringNameFromStr("get_collision_normal")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionResult3D.fnGetCollisionNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1914908202))
  }
  {
    methodName := StringNameFromStr("get_collider_velocity")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionResult3D.fnGetColliderVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1914908202))
  }
  {
    methodName := StringNameFromStr("get_collider_id")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionResult3D.fnGetColliderId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
  }
  {
    methodName := StringNameFromStr("get_collider_rid")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionResult3D.fnGetColliderRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1231817359))
  }
  {
    methodName := StringNameFromStr("get_collider")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionResult3D.fnGetCollider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2639523548))
  }
  {
    methodName := StringNameFromStr("get_collider_shape")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionResult3D.fnGetColliderShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
  }
  {
    methodName := StringNameFromStr("get_collision_local_shape")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionResult3D.fnGetCollisionLocalShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
  }
  {
    methodName := StringNameFromStr("get_collision_depth")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionResult3D.fnGetCollisionDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 218038398))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult3D.fnGetTravel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionResult3D) GetRemainder() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult3D.fnGetRemainder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionResult3D) GetCollisionSafeFraction() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult3D.fnGetCollisionSafeFraction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionResult3D) GetCollisionUnsafeFraction() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult3D.fnGetCollisionUnsafeFraction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionResult3D) GetCollisionCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult3D.fnGetCollisionCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionResult3D) GetCollisionPoint(collision_index int64, ) Vector3 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult3D.fnGetCollisionPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionResult3D) GetCollisionNormal(collision_index int64, ) Vector3 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult3D.fnGetCollisionNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionResult3D) GetColliderVelocity(collision_index int64, ) Vector3 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult3D.fnGetColliderVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionResult3D) GetColliderId(collision_index int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult3D.fnGetColliderId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionResult3D) GetColliderRid(collision_index int64, ) RID {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult3D.fnGetColliderRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionResult3D) GetCollider(collision_index int64, ) Object {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewObject()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult3D.fnGetCollider), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionResult3D) GetColliderShape(collision_index int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult3D.fnGetColliderShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionResult3D) GetCollisionLocalShape(collision_index int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult3D.fnGetCollisionLocalShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionResult3D) GetCollisionDepth(collision_index int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&collision_index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult3D.fnGetCollisionDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
