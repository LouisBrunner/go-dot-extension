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

type ptrsForRigidBody2DList struct {
  fnXIntegrateForces gdc.MethodBindPtr
  fnSetMass gdc.MethodBindPtr
  fnGetMass gdc.MethodBindPtr
  fnGetInertia gdc.MethodBindPtr
  fnSetInertia gdc.MethodBindPtr
  fnSetCenterOfMassMode gdc.MethodBindPtr
  fnGetCenterOfMassMode gdc.MethodBindPtr
  fnSetCenterOfMass gdc.MethodBindPtr
  fnGetCenterOfMass gdc.MethodBindPtr
  fnSetPhysicsMaterialOverride gdc.MethodBindPtr
  fnGetPhysicsMaterialOverride gdc.MethodBindPtr
  fnSetGravityScale gdc.MethodBindPtr
  fnGetGravityScale gdc.MethodBindPtr
  fnSetLinearDampMode gdc.MethodBindPtr
  fnGetLinearDampMode gdc.MethodBindPtr
  fnSetAngularDampMode gdc.MethodBindPtr
  fnGetAngularDampMode gdc.MethodBindPtr
  fnSetLinearDamp gdc.MethodBindPtr
  fnGetLinearDamp gdc.MethodBindPtr
  fnSetAngularDamp gdc.MethodBindPtr
  fnGetAngularDamp gdc.MethodBindPtr
  fnSetLinearVelocity gdc.MethodBindPtr
  fnGetLinearVelocity gdc.MethodBindPtr
  fnSetAngularVelocity gdc.MethodBindPtr
  fnGetAngularVelocity gdc.MethodBindPtr
  fnSetMaxContactsReported gdc.MethodBindPtr
  fnGetMaxContactsReported gdc.MethodBindPtr
  fnGetContactCount gdc.MethodBindPtr
  fnSetUseCustomIntegrator gdc.MethodBindPtr
  fnIsUsingCustomIntegrator gdc.MethodBindPtr
  fnSetContactMonitor gdc.MethodBindPtr
  fnIsContactMonitorEnabled gdc.MethodBindPtr
  fnSetContinuousCollisionDetectionMode gdc.MethodBindPtr
  fnGetContinuousCollisionDetectionMode gdc.MethodBindPtr
  fnSetAxisVelocity gdc.MethodBindPtr
  fnApplyCentralImpulse gdc.MethodBindPtr
  fnApplyImpulse gdc.MethodBindPtr
  fnApplyTorqueImpulse gdc.MethodBindPtr
  fnApplyCentralForce gdc.MethodBindPtr
  fnApplyForce gdc.MethodBindPtr
  fnApplyTorque gdc.MethodBindPtr
  fnAddConstantCentralForce gdc.MethodBindPtr
  fnAddConstantForce gdc.MethodBindPtr
  fnAddConstantTorque gdc.MethodBindPtr
  fnSetConstantForce gdc.MethodBindPtr
  fnGetConstantForce gdc.MethodBindPtr
  fnSetConstantTorque gdc.MethodBindPtr
  fnGetConstantTorque gdc.MethodBindPtr
  fnSetSleeping gdc.MethodBindPtr
  fnIsSleeping gdc.MethodBindPtr
  fnSetCanSleep gdc.MethodBindPtr
  fnIsAbleToSleep gdc.MethodBindPtr
  fnSetLockRotationEnabled gdc.MethodBindPtr
  fnIsLockRotationEnabled gdc.MethodBindPtr
  fnSetFreezeEnabled gdc.MethodBindPtr
  fnIsFreezeEnabled gdc.MethodBindPtr
  fnSetFreezeMode gdc.MethodBindPtr
  fnGetFreezeMode gdc.MethodBindPtr
  fnGetCollidingBodies gdc.MethodBindPtr
}

var ptrsForRigidBody2D ptrsForRigidBody2DList

func initRigidBody2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("RigidBody2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_mass")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_mass")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_inertia")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetInertia = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_inertia")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetInertia = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("set_center_of_mass_mode")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetCenterOfMassMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1757235706))
  }
  {
    methodName := StringNameFromStr("get_center_of_mass_mode")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetCenterOfMassMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3277132817))
  }
  {
    methodName := StringNameFromStr("set_center_of_mass")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetCenterOfMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_center_of_mass")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetCenterOfMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_physics_material_override")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetPhysicsMaterialOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1784508650))
  }
  {
    methodName := StringNameFromStr("get_physics_material_override")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetPhysicsMaterialOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2521850424))
  }
  {
    methodName := StringNameFromStr("set_gravity_scale")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetGravityScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_gravity_scale")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetGravityScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_linear_damp_mode")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetLinearDampMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3406533708))
  }
  {
    methodName := StringNameFromStr("get_linear_damp_mode")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetLinearDampMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2970511462))
  }
  {
    methodName := StringNameFromStr("set_angular_damp_mode")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetAngularDampMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3406533708))
  }
  {
    methodName := StringNameFromStr("get_angular_damp_mode")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetAngularDampMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2970511462))
  }
  {
    methodName := StringNameFromStr("set_linear_damp")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetLinearDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_linear_damp")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetLinearDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_angular_damp")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetAngularDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_angular_damp")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetAngularDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_linear_velocity")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_linear_velocity")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_angular_velocity")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_angular_velocity")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_max_contacts_reported")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetMaxContactsReported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_max_contacts_reported")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetMaxContactsReported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_contact_count")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetContactCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_use_custom_integrator")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetUseCustomIntegrator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_using_custom_integrator")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnIsUsingCustomIntegrator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("set_contact_monitor")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetContactMonitor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_contact_monitor_enabled")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnIsContactMonitorEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_continuous_collision_detection_mode")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetContinuousCollisionDetectionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1000241384))
  }
  {
    methodName := StringNameFromStr("get_continuous_collision_detection_mode")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetContinuousCollisionDetectionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 815214376))
  }
  {
    methodName := StringNameFromStr("set_axis_velocity")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetAxisVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("apply_central_impulse")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnApplyCentralImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3862383994))
  }
  {
    methodName := StringNameFromStr("apply_impulse")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnApplyImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4288681949))
  }
  {
    methodName := StringNameFromStr("apply_torque_impulse")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnApplyTorqueImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("apply_central_force")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnApplyCentralForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("apply_force")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnApplyForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4288681949))
  }
  {
    methodName := StringNameFromStr("apply_torque")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnApplyTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("add_constant_central_force")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnAddConstantCentralForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("add_constant_force")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnAddConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4288681949))
  }
  {
    methodName := StringNameFromStr("add_constant_torque")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnAddConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("set_constant_force")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_constant_force")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_constant_torque")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_constant_torque")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_sleeping")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetSleeping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_sleeping")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnIsSleeping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_can_sleep")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetCanSleep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_able_to_sleep")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnIsAbleToSleep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_lock_rotation_enabled")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetLockRotationEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_lock_rotation_enabled")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnIsLockRotationEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_freeze_enabled")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetFreezeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_freeze_enabled")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnIsFreezeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_freeze_mode")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnSetFreezeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1705112154))
  }
  {
    methodName := StringNameFromStr("get_freeze_mode")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetFreezeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2016872314))
  }
  {
    methodName := StringNameFromStr("get_colliding_bodies")
    defer methodName.Destroy()
    ptrsForRigidBody2D.fnGetCollidingBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
}

type RigidBody2D struct {
  PhysicsBody2D
}

func (me *RigidBody2D) BaseClass() string {
  return "RigidBody2D"
}

func NewRigidBody2D() *RigidBody2D {
  str := StringNameFromStr("RigidBody2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RigidBody2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type RigidBody2DFreezeMode int
const (
  RigidBody2DFreezeModeFreezeModeStatic RigidBody2DFreezeMode = 0
  RigidBody2DFreezeModeFreezeModeKinematic RigidBody2DFreezeMode = 1
)

type RigidBody2DCenterOfMassMode int
const (
  RigidBody2DCenterOfMassModeCenterOfMassModeAuto RigidBody2DCenterOfMassMode = 0
  RigidBody2DCenterOfMassModeCenterOfMassModeCustom RigidBody2DCenterOfMassMode = 1
)

type RigidBody2DDampMode int
const (
  RigidBody2DDampModeDampModeCombine RigidBody2DDampMode = 0
  RigidBody2DDampModeDampModeReplace RigidBody2DDampMode = 1
)

type RigidBody2DCCDMode int
const (
  RigidBody2DCCDModeCcdModeDisabled RigidBody2DCCDMode = 0
  RigidBody2DCCDModeCcdModeCastRay RigidBody2DCCDMode = 1
  RigidBody2DCCDModeCcdModeCastShape RigidBody2DCCDMode = 2
)

func (me *RigidBody2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RigidBody2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RigidBody2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RigidBody2D) SetMass(mass float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mass) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetMass), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetMass() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetMass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) GetInertia() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetInertia), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) SetInertia(inertia float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&inertia) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetInertia), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) SetCenterOfMassMode(mode RigidBody2DCenterOfMassMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetCenterOfMassMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetCenterOfMassMode() RigidBody2DCenterOfMassMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RigidBody2DCenterOfMassMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetCenterOfMassMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RigidBody2D) SetCenterOfMass(center_of_mass Vector2, )  {
  cargs := []gdc.ConstTypePtr{center_of_mass.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetCenterOfMass), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetCenterOfMass() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetCenterOfMass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RigidBody2D) SetPhysicsMaterialOverride(physics_material_override PhysicsMaterial, )  {
  cargs := []gdc.ConstTypePtr{physics_material_override.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetPhysicsMaterialOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetPhysicsMaterialOverride() PhysicsMaterial {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPhysicsMaterial()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetPhysicsMaterialOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RigidBody2D) SetGravityScale(gravity_scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gravity_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetGravityScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetGravityScale() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetGravityScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) SetLinearDampMode(linear_damp_mode RigidBody2DDampMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetLinearDampMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetLinearDampMode() RigidBody2DDampMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RigidBody2DDampMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetLinearDampMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RigidBody2D) SetAngularDampMode(angular_damp_mode RigidBody2DDampMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetAngularDampMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetAngularDampMode() RigidBody2DDampMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RigidBody2DDampMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetAngularDampMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RigidBody2D) SetLinearDamp(linear_damp float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetLinearDamp), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetLinearDamp() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetLinearDamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) SetAngularDamp(angular_damp float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetAngularDamp), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetAngularDamp() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetAngularDamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) SetLinearVelocity(linear_velocity Vector2, )  {
  cargs := []gdc.ConstTypePtr{linear_velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetLinearVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetLinearVelocity() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetLinearVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RigidBody2D) SetAngularVelocity(angular_velocity float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_velocity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetAngularVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetAngularVelocity() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetAngularVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) SetMaxContactsReported(amount int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetMaxContactsReported), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetMaxContactsReported() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetMaxContactsReported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) GetContactCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetContactCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) SetUseCustomIntegrator(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetUseCustomIntegrator), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) IsUsingCustomIntegrator() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnIsUsingCustomIntegrator), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) SetContactMonitor(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetContactMonitor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) IsContactMonitorEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnIsContactMonitorEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) SetContinuousCollisionDetectionMode(mode RigidBody2DCCDMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetContinuousCollisionDetectionMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetContinuousCollisionDetectionMode() RigidBody2DCCDMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RigidBody2DCCDMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetContinuousCollisionDetectionMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RigidBody2D) SetAxisVelocity(axis_velocity Vector2, )  {
  cargs := []gdc.ConstTypePtr{axis_velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetAxisVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) ApplyCentralImpulse(impulse Vector2, )  {
  cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnApplyCentralImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) ApplyImpulse(impulse Vector2, position Vector2, )  {
  cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnApplyImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) ApplyTorqueImpulse(torque float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&torque) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnApplyTorqueImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) ApplyCentralForce(force Vector2, )  {
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnApplyCentralForce), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) ApplyForce(force Vector2, position Vector2, )  {
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnApplyForce), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) ApplyTorque(torque float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&torque) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnApplyTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) AddConstantCentralForce(force Vector2, )  {
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnAddConstantCentralForce), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) AddConstantForce(force Vector2, position Vector2, )  {
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnAddConstantForce), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) AddConstantTorque(torque float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&torque) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnAddConstantTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) SetConstantForce(force Vector2, )  {
  cargs := []gdc.ConstTypePtr{force.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetConstantForce), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetConstantForce() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetConstantForce), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RigidBody2D) SetConstantTorque(torque float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&torque) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetConstantTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetConstantTorque() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetConstantTorque), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) SetSleeping(sleeping bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sleeping) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetSleeping), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) IsSleeping() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnIsSleeping), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) SetCanSleep(able_to_sleep bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&able_to_sleep) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetCanSleep), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) IsAbleToSleep() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnIsAbleToSleep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) SetLockRotationEnabled(lock_rotation bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lock_rotation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetLockRotationEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) IsLockRotationEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnIsLockRotationEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) SetFreezeEnabled(freeze_mode bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&freeze_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetFreezeEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) IsFreezeEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnIsFreezeEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RigidBody2D) SetFreezeMode(freeze_mode RigidBody2DFreezeMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&freeze_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnSetFreezeMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RigidBody2D) GetFreezeMode() RigidBody2DFreezeMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RigidBody2DFreezeMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetFreezeMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RigidBody2D) GetCollidingBodies() []Node2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRigidBody2D.fnGetCollidingBodies), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Node2D](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type RigidBody2DBodyShapeEnteredSignalFn func(body_rid RID, body Node, body_shape_index int, local_shape_index int, )

func (me *RigidBody2D) ConnectBodyShapeEntered(subs SignalSubscribers, fn RigidBody2DBodyShapeEnteredSignalFn) {
  sig := StringNameFromStr("body_shape_entered")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *RigidBody2D) DisconnectBodyShapeEntered(subs SignalSubscribers, fn RigidBody2DBodyShapeEnteredSignalFn) {
  sig := StringNameFromStr("body_shape_entered")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type RigidBody2DBodyShapeExitedSignalFn func(body_rid RID, body Node, body_shape_index int, local_shape_index int, )

func (me *RigidBody2D) ConnectBodyShapeExited(subs SignalSubscribers, fn RigidBody2DBodyShapeExitedSignalFn) {
  sig := StringNameFromStr("body_shape_exited")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *RigidBody2D) DisconnectBodyShapeExited(subs SignalSubscribers, fn RigidBody2DBodyShapeExitedSignalFn) {
  sig := StringNameFromStr("body_shape_exited")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type RigidBody2DBodyEnteredSignalFn func(body Node, )

func (me *RigidBody2D) ConnectBodyEntered(subs SignalSubscribers, fn RigidBody2DBodyEnteredSignalFn) {
  sig := StringNameFromStr("body_entered")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *RigidBody2D) DisconnectBodyEntered(subs SignalSubscribers, fn RigidBody2DBodyEnteredSignalFn) {
  sig := StringNameFromStr("body_entered")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type RigidBody2DBodyExitedSignalFn func(body Node, )

func (me *RigidBody2D) ConnectBodyExited(subs SignalSubscribers, fn RigidBody2DBodyExitedSignalFn) {
  sig := StringNameFromStr("body_exited")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *RigidBody2D) DisconnectBodyExited(subs SignalSubscribers, fn RigidBody2DBodyExitedSignalFn) {
  sig := StringNameFromStr("body_exited")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type RigidBody2DSleepingStateChangedSignalFn func()

func (me *RigidBody2D) ConnectSleepingStateChanged(subs SignalSubscribers, fn RigidBody2DSleepingStateChangedSignalFn) {
  sig := StringNameFromStr("sleeping_state_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *RigidBody2D) DisconnectSleepingStateChanged(subs SignalSubscribers, fn RigidBody2DSleepingStateChangedSignalFn) {
  sig := StringNameFromStr("sleeping_state_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
