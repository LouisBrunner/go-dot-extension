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



// Enums

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

func (me *RigidBody3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RigidBody3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RigidBody3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *RigidBody3D) XIntegrateForces(state PhysicsDirectBodyState3D, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetMass(mass float32, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetMass()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetInertia(inertia Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetInertia()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetCenterOfMassMode(mode RigidBody3DCenterOfMassMode, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetCenterOfMassMode()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetCenterOfMass(center_of_mass Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetCenterOfMass()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetPhysicsMaterialOverride(physics_material_override PhysicsMaterial, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetPhysicsMaterialOverride()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetLinearVelocity(linear_velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetLinearVelocity()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetAngularVelocity(angular_velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetAngularVelocity()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetInverseInertiaTensor()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetGravityScale(gravity_scale float32, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetGravityScale()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetLinearDampMode(linear_damp_mode RigidBody3DDampMode, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetLinearDampMode()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetAngularDampMode(angular_damp_mode RigidBody3DDampMode, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetAngularDampMode()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetLinearDamp(linear_damp float32, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetLinearDamp()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetAngularDamp(angular_damp float32, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetAngularDamp()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetMaxContactsReported(amount int, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetMaxContactsReported()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetContactCount()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetUseCustomIntegrator(enable bool, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) IsUsingCustomIntegrator()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetContactMonitor(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) IsContactMonitorEnabled()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetUseContinuousCollisionDetection(enable bool, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) IsUsingContinuousCollisionDetection()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetAxisVelocity(axis_velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) ApplyCentralImpulse(impulse Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) ApplyImpulse(impulse Vector3, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) ApplyTorqueImpulse(impulse Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) ApplyCentralForce(force Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) ApplyForce(force Vector3, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) ApplyTorque(torque Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) AddConstantCentralForce(force Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) AddConstantForce(force Vector3, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) AddConstantTorque(torque Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetConstantForce(force Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetConstantForce()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetConstantTorque(torque Vector3, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetConstantTorque()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetSleeping(sleeping bool, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) IsSleeping()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetCanSleep(able_to_sleep bool, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) IsAbleToSleep()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetLockRotationEnabled(lock_rotation bool, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) IsLockRotationEnabled()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetFreezeEnabled(freeze_mode bool, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) IsFreezeEnabled()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) SetFreezeMode(freeze_mode RigidBody3DFreezeMode, )  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetFreezeMode()  {
  panic("TODO: implement")
}

func  (me *RigidBody3D) GetCollidingBodies()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
