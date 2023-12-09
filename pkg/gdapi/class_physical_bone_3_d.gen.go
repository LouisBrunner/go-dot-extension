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



// Enums

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

func (me *PhysicalBone3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicalBone3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicalBone3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicalBone3D) XIntegrateForces(state PhysicsDirectBodyState3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) ApplyCentralImpulse(impulse Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) ApplyImpulse(impulse Vector3, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetJointType(joint_type PhysicalBone3DJointType, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetJointType()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetJointOffset(offset Transform3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetJointOffset()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetJointRotation(euler Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetJointRotation()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetBodyOffset(offset Transform3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetBodyOffset()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetSimulatePhysics()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) IsSimulatingPhysics()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetBoneId()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetMass(mass float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetMass()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetFriction(friction float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetFriction()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetBounce(bounce float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetBounce()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetGravityScale(gravity_scale float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetGravityScale()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetLinearDampMode(linear_damp_mode PhysicalBone3DDampMode, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetLinearDampMode()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetAngularDampMode(angular_damp_mode PhysicalBone3DDampMode, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetAngularDampMode()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetLinearDamp(linear_damp float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetLinearDamp()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetAngularDamp(angular_damp float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetAngularDamp()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetLinearVelocity(linear_velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetLinearVelocity()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetAngularVelocity(angular_velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) GetAngularVelocity()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetUseCustomIntegrator(enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) IsUsingCustomIntegrator()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) SetCanSleep(able_to_sleep bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone3D) IsAbleToSleep()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
