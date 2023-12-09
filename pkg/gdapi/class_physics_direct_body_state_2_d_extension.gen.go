// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectBodyState2DExtension struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectBodyState2DExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectBodyState2DExtension) BaseClass() string {
  return "PhysicsDirectBodyState2DExtension"
}



// Enums

func (me *PhysicsDirectBodyState2DExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsDirectBodyState2DExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectBodyState2DExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsDirectBodyState2DExtension) XGetTotalGravity()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetTotalLinearDamp()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetTotalAngularDamp()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetCenterOfMass()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetCenterOfMassLocal()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetInverseMass()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetInverseInertia()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XSetLinearVelocity(velocity Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetLinearVelocity()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XSetAngularVelocity(velocity float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetAngularVelocity()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XSetTransform(transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetTransform()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetVelocityAtLocalPosition(local_position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XApplyCentralImpulse(impulse Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XApplyImpulse(impulse Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XApplyTorqueImpulse(impulse float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XApplyCentralForce(force Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XApplyForce(force Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XApplyTorque(torque float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XAddConstantCentralForce(force Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XAddConstantForce(force Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XAddConstantTorque(torque float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XSetConstantForce(force Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetConstantForce()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XSetConstantTorque(torque float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetConstantTorque()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XSetSleepState(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XIsSleeping()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetContactCount()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetContactLocalPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetContactLocalNormal(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetContactLocalShape(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetContactLocalVelocityAtPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetContactCollider(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetContactColliderPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetContactColliderId(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetContactColliderObject(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetContactColliderShape(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetContactColliderVelocityAtPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetContactImpulse(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetStep()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XIntegrateForces()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState2DExtension) XGetSpaceState()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
