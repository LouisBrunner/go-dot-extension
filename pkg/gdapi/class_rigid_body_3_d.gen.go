// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RigidBody3D struct {
  obj gdc.ObjectPtr
}

func (me *RigidBody3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RigidBody3D) BaseClass() string {
  return "RigidBody3D"
}

type RigidBody3DFreezeMode int
const (
  RigidBody3DFreezeModeFreezeModeStatic RigidBody3DFreezeMode = 0
  RigidBody3DFreezeModeFreezeModeKinematic RigidBody3DFreezeMode = 1
)

type RigidBody3DCenterOfMassMode int
const (
  RigidBody3DCenterOfMassModeCenterOfMassModeAuto RigidBody3DCenterOfMassMode = 0
  RigidBody3DCenterOfMassModeCenterOfMassModeCustom RigidBody3DCenterOfMassMode = 1
)

type RigidBody3DDampMode int
const (
  RigidBody3DDampModeDampModeCombine RigidBody3DDampMode = 0
  RigidBody3DDampModeDampModeReplace RigidBody3DDampMode = 1
)

func  (me *RigidBody3D) XIntegrateForces(state PhysicsDirectBodyState3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetMass(mass float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetMass() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetInertia(inertia Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetInertia() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetCenterOfMassMode(mode RigidBody3DCenterOfMassMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetCenterOfMassMode() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetCenterOfMass(center_of_mass Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetCenterOfMass() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetPhysicsMaterialOverride(physics_material_override PhysicsMaterial, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetPhysicsMaterialOverride() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetLinearVelocity(linear_velocity Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetLinearVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetAngularVelocity(angular_velocity Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetAngularVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetInverseInertiaTensor() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetGravityScale(gravity_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetGravityScale() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetLinearDampMode(linear_damp_mode RigidBody3DDampMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetLinearDampMode() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetAngularDampMode(angular_damp_mode RigidBody3DDampMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetAngularDampMode() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetLinearDamp(linear_damp float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetLinearDamp() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetAngularDamp(angular_damp float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetAngularDamp() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetMaxContactsReported(amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetMaxContactsReported() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetContactCount() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetUseCustomIntegrator(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) IsUsingCustomIntegrator() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetContactMonitor(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) IsContactMonitorEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetUseContinuousCollisionDetection(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) IsUsingContinuousCollisionDetection() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetAxisVelocity(axis_velocity Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) ApplyCentralImpulse(impulse Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) ApplyImpulse(impulse Vector3, position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) ApplyTorqueImpulse(impulse Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) ApplyCentralForce(force Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) ApplyForce(force Vector3, position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) ApplyTorque(torque Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) AddConstantCentralForce(force Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) AddConstantForce(force Vector3, position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) AddConstantTorque(torque Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetConstantForce(force Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetConstantForce() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetConstantTorque(torque Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetConstantTorque() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetSleeping(sleeping bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) IsSleeping() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetCanSleep(able_to_sleep bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) IsAbleToSleep() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetLockRotationEnabled(lock_rotation bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) IsLockRotationEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetFreezeEnabled(freeze_mode bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) IsFreezeEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) SetFreezeMode(freeze_mode RigidBody3DFreezeMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetFreezeMode() { // TODO: return value
  // TODO: implement
}

func  (me *RigidBody3D) GetCollidingBodies() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
