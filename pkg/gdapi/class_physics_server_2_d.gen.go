// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer2D) BaseClass() string {
  return "PhysicsServer2D"
}

type PhysicsServer2DSpaceParameter int
const (
  PhysicsServer2DSpaceParamContactRecycleRadius PhysicsServer2DSpaceParameter = 0
  PhysicsServer2DSpaceParamContactMaxSeparation PhysicsServer2DSpaceParameter = 1
  PhysicsServer2DSpaceParamContactMaxAllowedPenetration PhysicsServer2DSpaceParameter = 2
  PhysicsServer2DSpaceParamContactDefaultBias PhysicsServer2DSpaceParameter = 3
  PhysicsServer2DSpaceParamBodyLinearVelocitySleepThreshold PhysicsServer2DSpaceParameter = 4
  PhysicsServer2DSpaceParamBodyAngularVelocitySleepThreshold PhysicsServer2DSpaceParameter = 5
  PhysicsServer2DSpaceParamBodyTimeToSleep PhysicsServer2DSpaceParameter = 6
  PhysicsServer2DSpaceParamConstraintDefaultBias PhysicsServer2DSpaceParameter = 7
  PhysicsServer2DSpaceParamSolverIterations PhysicsServer2DSpaceParameter = 8
)

type PhysicsServer2DShapeType int
const (
  PhysicsServer2DShapeWorldBoundary PhysicsServer2DShapeType = 0
  PhysicsServer2DShapeSeparationRay PhysicsServer2DShapeType = 1
  PhysicsServer2DShapeSegment PhysicsServer2DShapeType = 2
  PhysicsServer2DShapeCircle PhysicsServer2DShapeType = 3
  PhysicsServer2DShapeRectangle PhysicsServer2DShapeType = 4
  PhysicsServer2DShapeCapsule PhysicsServer2DShapeType = 5
  PhysicsServer2DShapeConvexPolygon PhysicsServer2DShapeType = 6
  PhysicsServer2DShapeConcavePolygon PhysicsServer2DShapeType = 7
  PhysicsServer2DShapeCustom PhysicsServer2DShapeType = 8
)

type PhysicsServer2DAreaParameter int
const (
  PhysicsServer2DAreaParamGravityOverrideMode PhysicsServer2DAreaParameter = 0
  PhysicsServer2DAreaParamGravity PhysicsServer2DAreaParameter = 1
  PhysicsServer2DAreaParamGravityVector PhysicsServer2DAreaParameter = 2
  PhysicsServer2DAreaParamGravityIsPoint PhysicsServer2DAreaParameter = 3
  PhysicsServer2DAreaParamGravityPointUnitDistance PhysicsServer2DAreaParameter = 4
  PhysicsServer2DAreaParamLinearDampOverrideMode PhysicsServer2DAreaParameter = 5
  PhysicsServer2DAreaParamLinearDamp PhysicsServer2DAreaParameter = 6
  PhysicsServer2DAreaParamAngularDampOverrideMode PhysicsServer2DAreaParameter = 7
  PhysicsServer2DAreaParamAngularDamp PhysicsServer2DAreaParameter = 8
  PhysicsServer2DAreaParamPriority PhysicsServer2DAreaParameter = 9
)

type PhysicsServer2DAreaSpaceOverrideMode int
const (
  PhysicsServer2DAreaSpaceOverrideDisabled PhysicsServer2DAreaSpaceOverrideMode = 0
  PhysicsServer2DAreaSpaceOverrideCombine PhysicsServer2DAreaSpaceOverrideMode = 1
  PhysicsServer2DAreaSpaceOverrideCombineReplace PhysicsServer2DAreaSpaceOverrideMode = 2
  PhysicsServer2DAreaSpaceOverrideReplace PhysicsServer2DAreaSpaceOverrideMode = 3
  PhysicsServer2DAreaSpaceOverrideReplaceCombine PhysicsServer2DAreaSpaceOverrideMode = 4
)

type PhysicsServer2DBodyMode int
const (
  PhysicsServer2DBodyModeStatic PhysicsServer2DBodyMode = 0
  PhysicsServer2DBodyModeKinematic PhysicsServer2DBodyMode = 1
  PhysicsServer2DBodyModeRigid PhysicsServer2DBodyMode = 2
  PhysicsServer2DBodyModeRigidLinear PhysicsServer2DBodyMode = 3
)

type PhysicsServer2DBodyParameter int
const (
  PhysicsServer2DBodyParamBounce PhysicsServer2DBodyParameter = 0
  PhysicsServer2DBodyParamFriction PhysicsServer2DBodyParameter = 1
  PhysicsServer2DBodyParamMass PhysicsServer2DBodyParameter = 2
  PhysicsServer2DBodyParamInertia PhysicsServer2DBodyParameter = 3
  PhysicsServer2DBodyParamCenterOfMass PhysicsServer2DBodyParameter = 4
  PhysicsServer2DBodyParamGravityScale PhysicsServer2DBodyParameter = 5
  PhysicsServer2DBodyParamLinearDampMode PhysicsServer2DBodyParameter = 6
  PhysicsServer2DBodyParamAngularDampMode PhysicsServer2DBodyParameter = 7
  PhysicsServer2DBodyParamLinearDamp PhysicsServer2DBodyParameter = 8
  PhysicsServer2DBodyParamAngularDamp PhysicsServer2DBodyParameter = 9
  PhysicsServer2DBodyParamMax PhysicsServer2DBodyParameter = 10
)

type PhysicsServer2DBodyDampMode int
const (
  PhysicsServer2DBodyDampModeCombine PhysicsServer2DBodyDampMode = 0
  PhysicsServer2DBodyDampModeReplace PhysicsServer2DBodyDampMode = 1
)

type PhysicsServer2DBodyState int
const (
  PhysicsServer2DBodyStateTransform PhysicsServer2DBodyState = 0
  PhysicsServer2DBodyStateLinearVelocity PhysicsServer2DBodyState = 1
  PhysicsServer2DBodyStateAngularVelocity PhysicsServer2DBodyState = 2
  PhysicsServer2DBodyStateSleeping PhysicsServer2DBodyState = 3
  PhysicsServer2DBodyStateCanSleep PhysicsServer2DBodyState = 4
)

type PhysicsServer2DJointType int
const (
  PhysicsServer2DJointTypePin PhysicsServer2DJointType = 0
  PhysicsServer2DJointTypeGroove PhysicsServer2DJointType = 1
  PhysicsServer2DJointTypeDampedSpring PhysicsServer2DJointType = 2
  PhysicsServer2DJointTypeMax PhysicsServer2DJointType = 3
)

type PhysicsServer2DJointParam int
const (
  PhysicsServer2DJointParamBias PhysicsServer2DJointParam = 0
  PhysicsServer2DJointParamMaxBias PhysicsServer2DJointParam = 1
  PhysicsServer2DJointParamMaxForce PhysicsServer2DJointParam = 2
)

type PhysicsServer2DPinJointParam int
const (
  PhysicsServer2DPinJointSoftness PhysicsServer2DPinJointParam = 0
)

type PhysicsServer2DDampedSpringParam int
const (
  PhysicsServer2DDampedSpringRestLength PhysicsServer2DDampedSpringParam = 0
  PhysicsServer2DDampedSpringStiffness PhysicsServer2DDampedSpringParam = 1
  PhysicsServer2DDampedSpringDamping PhysicsServer2DDampedSpringParam = 2
)

type PhysicsServer2DCCDMode int
const (
  PhysicsServer2DCcdModeDisabled PhysicsServer2DCCDMode = 0
  PhysicsServer2DCcdModeCastRay PhysicsServer2DCCDMode = 1
  PhysicsServer2DCcdModeCastShape PhysicsServer2DCCDMode = 2
)

type PhysicsServer2DAreaBodyStatus int
const (
  PhysicsServer2DAreaBodyAdded PhysicsServer2DAreaBodyStatus = 0
  PhysicsServer2DAreaBodyRemoved PhysicsServer2DAreaBodyStatus = 1
)

type PhysicsServer2DProcessInfo int
const (
  PhysicsServer2DInfoActiveObjects PhysicsServer2DProcessInfo = 0
  PhysicsServer2DInfoCollisionPairs PhysicsServer2DProcessInfo = 1
  PhysicsServer2DInfoIslandCount PhysicsServer2DProcessInfo = 2
)
