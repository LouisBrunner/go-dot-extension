// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForPhysicsDirectBodyState3DExtensionList struct {
  fnXGetTotalGravity gdc.MethodBindPtr
  fnXGetTotalLinearDamp gdc.MethodBindPtr
  fnXGetTotalAngularDamp gdc.MethodBindPtr
  fnXGetCenterOfMass gdc.MethodBindPtr
  fnXGetCenterOfMassLocal gdc.MethodBindPtr
  fnXGetPrincipalInertiaAxes gdc.MethodBindPtr
  fnXGetInverseMass gdc.MethodBindPtr
  fnXGetInverseInertia gdc.MethodBindPtr
  fnXGetInverseInertiaTensor gdc.MethodBindPtr
  fnXSetLinearVelocity gdc.MethodBindPtr
  fnXGetLinearVelocity gdc.MethodBindPtr
  fnXSetAngularVelocity gdc.MethodBindPtr
  fnXGetAngularVelocity gdc.MethodBindPtr
  fnXSetTransform gdc.MethodBindPtr
  fnXGetTransform gdc.MethodBindPtr
  fnXGetVelocityAtLocalPosition gdc.MethodBindPtr
  fnXApplyCentralImpulse gdc.MethodBindPtr
  fnXApplyImpulse gdc.MethodBindPtr
  fnXApplyTorqueImpulse gdc.MethodBindPtr
  fnXApplyCentralForce gdc.MethodBindPtr
  fnXApplyForce gdc.MethodBindPtr
  fnXApplyTorque gdc.MethodBindPtr
  fnXAddConstantCentralForce gdc.MethodBindPtr
  fnXAddConstantForce gdc.MethodBindPtr
  fnXAddConstantTorque gdc.MethodBindPtr
  fnXSetConstantForce gdc.MethodBindPtr
  fnXGetConstantForce gdc.MethodBindPtr
  fnXSetConstantTorque gdc.MethodBindPtr
  fnXGetConstantTorque gdc.MethodBindPtr
  fnXSetSleepState gdc.MethodBindPtr
  fnXIsSleeping gdc.MethodBindPtr
  fnXGetContactCount gdc.MethodBindPtr
  fnXGetContactLocalPosition gdc.MethodBindPtr
  fnXGetContactLocalNormal gdc.MethodBindPtr
  fnXGetContactImpulse gdc.MethodBindPtr
  fnXGetContactLocalShape gdc.MethodBindPtr
  fnXGetContactLocalVelocityAtPosition gdc.MethodBindPtr
  fnXGetContactCollider gdc.MethodBindPtr
  fnXGetContactColliderPosition gdc.MethodBindPtr
  fnXGetContactColliderId gdc.MethodBindPtr
  fnXGetContactColliderObject gdc.MethodBindPtr
  fnXGetContactColliderShape gdc.MethodBindPtr
  fnXGetContactColliderVelocityAtPosition gdc.MethodBindPtr
  fnXGetStep gdc.MethodBindPtr
  fnXIntegrateForces gdc.MethodBindPtr
  fnXGetSpaceState gdc.MethodBindPtr
}

var ptrsForPhysicsDirectBodyState3DExtension ptrsForPhysicsDirectBodyState3DExtensionList

func initPhysicsDirectBodyState3DExtensionPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PhysicsDirectBodyState3DExtension")
  defer className.Destroy()
}

type PhysicsDirectBodyState3DExtension struct {
  PhysicsDirectBodyState3D
}

func (me *PhysicsDirectBodyState3DExtension) BaseClass() string {
  return "PhysicsDirectBodyState3DExtension"
}

func NewPhysicsDirectBodyState3DExtension() *PhysicsDirectBodyState3DExtension {
  str := StringNameFromStr("PhysicsDirectBodyState3DExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicsDirectBodyState3DExtension{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PhysicsDirectBodyState3DExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsDirectBodyState3DExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectBodyState3DExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
