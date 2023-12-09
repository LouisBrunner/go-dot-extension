// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectBodyState3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectBodyState3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectBodyState3D) BaseClass() string {
  return "PhysicsDirectBodyState3D"
}



// Enums

func (me *PhysicsDirectBodyState3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectBodyState3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsDirectBodyState3D) GetTotalGravity()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetTotalLinearDamp()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetTotalAngularDamp()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetCenterOfMass()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetCenterOfMassLocal()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetPrincipalInertiaAxes()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetInverseMass()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetInverseInertia()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetInverseInertiaTensor()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) SetLinearVelocity(velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetLinearVelocity()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) SetAngularVelocity(velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetAngularVelocity()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) SetTransform(transform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetTransform()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetVelocityAtLocalPosition(local_position Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) ApplyCentralImpulse(impulse Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) ApplyImpulse(impulse Vector3, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) ApplyTorqueImpulse(impulse Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) ApplyCentralForce(force Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) ApplyForce(force Vector3, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) ApplyTorque(torque Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) AddConstantCentralForce(force Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) AddConstantForce(force Vector3, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) AddConstantTorque(torque Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) SetConstantForce(force Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetConstantForce()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) SetConstantTorque(torque Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetConstantTorque()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) SetSleepState(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) IsSleeping()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetContactCount()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetContactLocalPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetContactLocalNormal(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetContactImpulse(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetContactLocalShape(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetContactLocalVelocityAtPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetContactCollider(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetContactColliderPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetContactColliderId(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetContactColliderObject(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetContactColliderShape(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetContactColliderVelocityAtPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetStep()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) IntegrateForces()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3D) GetSpaceState()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
