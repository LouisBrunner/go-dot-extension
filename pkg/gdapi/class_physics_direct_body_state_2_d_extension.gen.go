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

type ptrsForPhysicsDirectBodyState2DExtensionList struct {
  fnXGetTotalGravity gdc.MethodBindPtr
  fnXGetTotalLinearDamp gdc.MethodBindPtr
  fnXGetTotalAngularDamp gdc.MethodBindPtr
  fnXGetCenterOfMass gdc.MethodBindPtr
  fnXGetCenterOfMassLocal gdc.MethodBindPtr
  fnXGetInverseMass gdc.MethodBindPtr
  fnXGetInverseInertia gdc.MethodBindPtr
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
  fnXGetContactLocalShape gdc.MethodBindPtr
  fnXGetContactLocalVelocityAtPosition gdc.MethodBindPtr
  fnXGetContactCollider gdc.MethodBindPtr
  fnXGetContactColliderPosition gdc.MethodBindPtr
  fnXGetContactColliderId gdc.MethodBindPtr
  fnXGetContactColliderObject gdc.MethodBindPtr
  fnXGetContactColliderShape gdc.MethodBindPtr
  fnXGetContactColliderVelocityAtPosition gdc.MethodBindPtr
  fnXGetContactImpulse gdc.MethodBindPtr
  fnXGetStep gdc.MethodBindPtr
  fnXIntegrateForces gdc.MethodBindPtr
  fnXGetSpaceState gdc.MethodBindPtr
}

var ptrsForPhysicsDirectBodyState2DExtension ptrsForPhysicsDirectBodyState2DExtensionList

func initPhysicsDirectBodyState2DExtensionPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PhysicsDirectBodyState2DExtension")
  defer className.Destroy()
}

type PhysicsDirectBodyState2DExtension struct {
  PhysicsDirectBodyState2D
}

func (me *PhysicsDirectBodyState2DExtension) BaseClass() string {
  return "PhysicsDirectBodyState2DExtension"
}

func NewPhysicsDirectBodyState2DExtension() *PhysicsDirectBodyState2DExtension {
  str := StringNameFromStr("PhysicsDirectBodyState2DExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicsDirectBodyState2DExtension{}
  obj.SetBaseObject(objPtr)
  return obj
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

// Signals
