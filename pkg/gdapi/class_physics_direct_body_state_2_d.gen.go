// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectBodyState2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectBodyState2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectBodyState2D) BaseClass() string {
  return "PhysicsDirectBodyState2D"
}



// Enums

func (me *PhysicsDirectBodyState2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectBodyState2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsDirectBodyState2D) GetTotalGravity()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetTotalLinearDamp()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetTotalAngularDamp()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetCenterOfMass()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetCenterOfMassLocal()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetInverseMass()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetInverseInertia()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) SetLinearVelocity(velocity Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetLinearVelocity()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) SetAngularVelocity(velocity float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetAngularVelocity()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) SetTransform(transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetTransform()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetVelocityAtLocalPosition(local_position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) ApplyCentralImpulse(impulse Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) ApplyTorqueImpulse(impulse float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) ApplyImpulse(impulse Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) ApplyCentralForce(force Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) ApplyForce(force Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) ApplyTorque(torque float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) AddConstantCentralForce(force Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) AddConstantForce(force Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) AddConstantTorque(torque float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) SetConstantForce(force Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetConstantForce()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) SetConstantTorque(torque float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetConstantTorque()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) SetSleepState(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) IsSleeping()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetContactCount()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetContactLocalPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetContactLocalNormal(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetContactLocalShape(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetContactLocalVelocityAtPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetContactCollider(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetContactColliderPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetContactColliderId(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetContactColliderObject(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetContactColliderShape(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetContactColliderVelocityAtPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetContactImpulse(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetStep()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) IntegrateForces()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2D) GetSpaceState()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
