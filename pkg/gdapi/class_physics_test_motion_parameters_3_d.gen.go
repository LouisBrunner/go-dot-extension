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

type ptrsForPhysicsTestMotionParameters3DList struct {
  fnGetFrom gdc.MethodBindPtr
  fnSetFrom gdc.MethodBindPtr
  fnGetMotion gdc.MethodBindPtr
  fnSetMotion gdc.MethodBindPtr
  fnGetMargin gdc.MethodBindPtr
  fnSetMargin gdc.MethodBindPtr
  fnGetMaxCollisions gdc.MethodBindPtr
  fnSetMaxCollisions gdc.MethodBindPtr
  fnIsCollideSeparationRayEnabled gdc.MethodBindPtr
  fnSetCollideSeparationRayEnabled gdc.MethodBindPtr
  fnGetExcludeBodies gdc.MethodBindPtr
  fnSetExcludeBodies gdc.MethodBindPtr
  fnGetExcludeObjects gdc.MethodBindPtr
  fnSetExcludeObjects gdc.MethodBindPtr
  fnIsRecoveryAsCollisionEnabled gdc.MethodBindPtr
  fnSetRecoveryAsCollisionEnabled gdc.MethodBindPtr
}

var ptrsForPhysicsTestMotionParameters3D ptrsForPhysicsTestMotionParameters3DList

func initPhysicsTestMotionParameters3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_from")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnGetFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
  }
  {
    methodName := StringNameFromStr("set_from")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnSetFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2952846383))
  }
  {
    methodName := StringNameFromStr("get_motion")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnGetMotion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_motion")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnSetMotion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_margin")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnGetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_margin")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnSetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_max_collisions")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnGetMaxCollisions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_max_collisions")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnSetMaxCollisions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("is_collide_separation_ray_enabled")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnIsCollideSeparationRayEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_collide_separation_ray_enabled")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnSetCollideSeparationRayEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_exclude_bodies")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnGetExcludeBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("set_exclude_bodies")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnSetExcludeBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_exclude_objects")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnGetExcludeObjects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("set_exclude_objects")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnSetExcludeObjects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("is_recovery_as_collision_enabled")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnIsRecoveryAsCollisionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_recovery_as_collision_enabled")
    defer methodName.Destroy()
    ptrsForPhysicsTestMotionParameters3D.fnSetRecoveryAsCollisionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
}

type PhysicsTestMotionParameters3D struct {
  RefCounted
}

func (me *PhysicsTestMotionParameters3D) BaseClass() string {
  return "PhysicsTestMotionParameters3D"
}

func NewPhysicsTestMotionParameters3D() *PhysicsTestMotionParameters3D {
  str := StringNameFromStr("PhysicsTestMotionParameters3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicsTestMotionParameters3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PhysicsTestMotionParameters3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsTestMotionParameters3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsTestMotionParameters3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsTestMotionParameters3D) GetFrom() Transform3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnGetFrom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionParameters3D) SetFrom(from Transform3D, )  {
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnSetFrom), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsTestMotionParameters3D) GetMotion() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnGetMotion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsTestMotionParameters3D) SetMotion(motion Vector3, )  {
  cargs := []gdc.ConstTypePtr{motion.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnSetMotion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsTestMotionParameters3D) GetMargin() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnGetMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionParameters3D) SetMargin(margin float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnSetMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsTestMotionParameters3D) GetMaxCollisions() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnGetMaxCollisions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionParameters3D) SetMaxCollisions(max_collisions int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_collisions) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnSetMaxCollisions), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsTestMotionParameters3D) IsCollideSeparationRayEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnIsCollideSeparationRayEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionParameters3D) SetCollideSeparationRayEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnSetCollideSeparationRayEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsTestMotionParameters3D) GetExcludeBodies() []RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnGetExcludeBodies), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[RID](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *PhysicsTestMotionParameters3D) SetExcludeBodies(exclude_list []RID, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_list) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnSetExcludeBodies), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsTestMotionParameters3D) GetExcludeObjects() []int {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnGetExcludeObjects), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[int](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *PhysicsTestMotionParameters3D) SetExcludeObjects(exclude_list []int, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_list) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnSetExcludeObjects), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsTestMotionParameters3D) IsRecoveryAsCollisionEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnIsRecoveryAsCollisionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsTestMotionParameters3D) SetRecoveryAsCollisionEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters3D.fnSetRecoveryAsCollisionEnabled), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
