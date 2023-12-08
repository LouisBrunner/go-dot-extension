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

func  (me *RigidBody2D) XIntegrateForces(state PhysicsDirectBodyState2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetMass(mass float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetMass() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetInertia() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetInertia(inertia float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetCenterOfMassMode(mode RigidBody2DCenterOfMassMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetCenterOfMassMode() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetCenterOfMass(center_of_mass Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetCenterOfMass() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetPhysicsMaterialOverride(physics_material_override PhysicsMaterial, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetPhysicsMaterialOverride() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetGravityScale(gravity_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetGravityScale() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetLinearDampMode(linear_damp_mode RigidBody2DDampMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetLinearDampMode() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetAngularDampMode(angular_damp_mode RigidBody2DDampMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetAngularDampMode() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetLinearDamp(linear_damp float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetLinearDamp() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetAngularDamp(angular_damp float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetAngularDamp() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetLinearVelocity(linear_velocity Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetLinearVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetAngularVelocity(angular_velocity float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetAngularVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetMaxContactsReported(amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetMaxContactsReported() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetContactCount() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetUseCustomIntegrator(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) IsUsingCustomIntegrator() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetContactMonitor(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) IsContactMonitorEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetContinuousCollisionDetectionMode(mode RigidBody2DCCDMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetContinuousCollisionDetectionMode() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetAxisVelocity(axis_velocity Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) ApplyCentralImpulse(impulse Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) ApplyImpulse(impulse Vector2, position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) ApplyTorqueImpulse(torque float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) ApplyCentralForce(force Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) ApplyForce(force Vector2, position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) ApplyTorque(torque float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) AddConstantCentralForce(force Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) AddConstantForce(force Vector2, position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) AddConstantTorque(torque float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetConstantForce(force Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetConstantForce() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetConstantTorque(torque float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetConstantTorque() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetSleeping(sleeping bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) IsSleeping() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetCanSleep(able_to_sleep bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) IsAbleToSleep() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetLockRotationEnabled(lock_rotation bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) IsLockRotationEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetFreezeEnabled(freeze_mode bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) IsFreezeEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) SetFreezeMode(freeze_mode RigidBody2DFreezeMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetFreezeMode() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody2D) GetCollidingBodies() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
