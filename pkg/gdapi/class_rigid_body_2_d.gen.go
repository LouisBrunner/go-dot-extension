// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RigidBody2D struct {
  obj gdc.ObjectPtr
}

func (me *RigidBody2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RigidBody2D) BaseClass() string {
  return "RigidBody2D"
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

func  (me *RigidBody2D) XIntegrateForces(state PhysicsDirectBodyState2D, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetMass(mass float32, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetMass()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetInertia()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetInertia(inertia float32, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetCenterOfMassMode(mode RigidBody2DCenterOfMassMode, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetCenterOfMassMode()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetCenterOfMass(center_of_mass Vector2, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetCenterOfMass()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetPhysicsMaterialOverride(physics_material_override PhysicsMaterial, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetPhysicsMaterialOverride()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetGravityScale(gravity_scale float32, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetGravityScale()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetLinearDampMode(linear_damp_mode RigidBody2DDampMode, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetLinearDampMode()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetAngularDampMode(angular_damp_mode RigidBody2DDampMode, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetAngularDampMode()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetLinearDamp(linear_damp float32, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetLinearDamp()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetAngularDamp(angular_damp float32, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetAngularDamp()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetLinearVelocity(linear_velocity Vector2, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetLinearVelocity()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetAngularVelocity(angular_velocity float32, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetAngularVelocity()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetMaxContactsReported(amount int, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetMaxContactsReported()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetContactCount()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetUseCustomIntegrator(enable bool, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) IsUsingCustomIntegrator()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetContactMonitor(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) IsContactMonitorEnabled()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetContinuousCollisionDetectionMode(mode RigidBody2DCCDMode, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetContinuousCollisionDetectionMode()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetAxisVelocity(axis_velocity Vector2, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) ApplyCentralImpulse(impulse Vector2, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) ApplyImpulse(impulse Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) ApplyTorqueImpulse(torque float32, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) ApplyCentralForce(force Vector2, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) ApplyForce(force Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) ApplyTorque(torque float32, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) AddConstantCentralForce(force Vector2, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) AddConstantForce(force Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) AddConstantTorque(torque float32, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetConstantForce(force Vector2, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetConstantForce()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetConstantTorque(torque float32, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetConstantTorque()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetSleeping(sleeping bool, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) IsSleeping()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetCanSleep(able_to_sleep bool, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) IsAbleToSleep()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetLockRotationEnabled(lock_rotation bool, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) IsLockRotationEnabled()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetFreezeEnabled(freeze_mode bool, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) IsFreezeEnabled()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) SetFreezeMode(freeze_mode RigidBody2DFreezeMode, )  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetFreezeMode()  {
  panic("TODO: implement")
}

func  (me *RigidBody2D) GetCollidingBodies()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
