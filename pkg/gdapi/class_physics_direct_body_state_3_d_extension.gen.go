// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsDirectBodyState3DExtension struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectBodyState3DExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectBodyState3DExtension) BaseClass() string {
  return "PhysicsDirectBodyState3DExtension"
}



// Enums

func (me *PhysicsDirectBodyState3DExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectBodyState3DExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsDirectBodyState3DExtension) XGetTotalGravity()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetTotalLinearDamp()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetTotalAngularDamp()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetCenterOfMass()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetCenterOfMassLocal()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetPrincipalInertiaAxes()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetInverseMass()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetInverseInertia()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetInverseInertiaTensor()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XSetLinearVelocity(velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetLinearVelocity()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XSetAngularVelocity(velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetAngularVelocity()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XSetTransform(transform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetTransform()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetVelocityAtLocalPosition(local_position Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XApplyCentralImpulse(impulse Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XApplyImpulse(impulse Vector3, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XApplyTorqueImpulse(impulse Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XApplyCentralForce(force Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XApplyForce(force Vector3, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XApplyTorque(torque Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XAddConstantCentralForce(force Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XAddConstantForce(force Vector3, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XAddConstantTorque(torque Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XSetConstantForce(force Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetConstantForce()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XSetConstantTorque(torque Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetConstantTorque()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XSetSleepState(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XIsSleeping()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetContactCount()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetContactLocalPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetContactLocalNormal(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetContactImpulse(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetContactLocalShape(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetContactLocalVelocityAtPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetContactCollider(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetContactColliderPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetContactColliderId(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetContactColliderObject(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetContactColliderShape(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetContactColliderVelocityAtPosition(contact_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetStep()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XIntegrateForces()  {
  panic("TODO: implement")
}

func  (me *PhysicsDirectBodyState3DExtension) XGetSpaceState()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
