// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicalBone3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicalBone3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicalBone3D) BaseClass() string {
  return "PhysicalBone3D"
}

type PhysicalBone3DDampMode int
const (
  PhysicalBone3DDampModeDampModeCombine PhysicalBone3DDampMode = 0
  PhysicalBone3DDampModeDampModeReplace PhysicalBone3DDampMode = 1
)

type PhysicalBone3DJointType int
const (
  PhysicalBone3DJointTypeJointTypeNone PhysicalBone3DJointType = 0
  PhysicalBone3DJointTypeJointTypePin PhysicalBone3DJointType = 1
  PhysicalBone3DJointTypeJointTypeCone PhysicalBone3DJointType = 2
  PhysicalBone3DJointTypeJointTypeHinge PhysicalBone3DJointType = 3
  PhysicalBone3DJointTypeJointTypeSlider PhysicalBone3DJointType = 4
  PhysicalBone3DJointTypeJointType6Dof PhysicalBone3DJointType = 5
)

func  (me *PhysicalBone3D) XIntegrateForces(state PhysicsDirectBodyState3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) ApplyCentralImpulse(impulse Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) ApplyImpulse(impulse Vector3, position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetJointType(joint_type PhysicalBone3DJointType, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetJointType() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetJointOffset(offset Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetJointOffset() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetJointRotation(euler Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetJointRotation() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetBodyOffset(offset Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetBodyOffset() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetSimulatePhysics() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) IsSimulatingPhysics() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetBoneId() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetMass(mass float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetMass() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetFriction(friction float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetFriction() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetBounce(bounce float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetBounce() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetGravityScale(gravity_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetGravityScale() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetLinearDampMode(linear_damp_mode PhysicalBone3DDampMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetLinearDampMode() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetAngularDampMode(angular_damp_mode PhysicalBone3DDampMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetAngularDampMode() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetLinearDamp(linear_damp float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetLinearDamp() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetAngularDamp(angular_damp float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetAngularDamp() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetLinearVelocity(linear_velocity Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetLinearVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetAngularVelocity(angular_velocity Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) GetAngularVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetUseCustomIntegrator(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) IsUsingCustomIntegrator() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) SetCanSleep(able_to_sleep bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicalBone3D) IsAbleToSleep() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
