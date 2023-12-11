// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsTestMotionParameters3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsTestMotionParameters3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsTestMotionParameters3D) BaseClass() string {
  return "PhysicsTestMotionParameters3D"
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
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_from")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsTestMotionParameters3D) SetFrom(from Transform3D, )  {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_from")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2952846383) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsTestMotionParameters3D) GetMotion() Vector3 {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsTestMotionParameters3D) SetMotion(motion Vector3, )  {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(motion.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsTestMotionParameters3D) GetMargin() float32 {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsTestMotionParameters3D) SetMargin(margin float32, )  {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsTestMotionParameters3D) GetMaxCollisions() int {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_collisions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsTestMotionParameters3D) SetMaxCollisions(max_collisions int, )  {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_collisions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_collisions), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsTestMotionParameters3D) IsCollideSeparationRayEnabled() bool {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_collide_separation_ray_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsTestMotionParameters3D) SetCollideSeparationRayEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collide_separation_ray_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsTestMotionParameters3D) GetExcludeBodies() RID {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_exclude_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsTestMotionParameters3D) SetExcludeBodies(exclude_list RID, )  {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exclude_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(exclude_list.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsTestMotionParameters3D) GetExcludeObjects() int {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_exclude_objects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsTestMotionParameters3D) SetExcludeObjects(exclude_list int, )  {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_exclude_objects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_list), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsTestMotionParameters3D) IsRecoveryAsCollisionEnabled() bool {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_recovery_as_collision_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsTestMotionParameters3D) SetRecoveryAsCollisionEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("PhysicsTestMotionParameters3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_recovery_as_collision_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
