// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer3D) BaseClass() string {
  return "PhysicsServer3D"
}

type PhysicsServer3DJointType int
const (
  PhysicsServer3DJointTypePin PhysicsServer3DJointType = 0
  PhysicsServer3DJointTypeHinge PhysicsServer3DJointType = 1
  PhysicsServer3DJointTypeSlider PhysicsServer3DJointType = 2
  PhysicsServer3DJointTypeConeTwist PhysicsServer3DJointType = 3
  PhysicsServer3DJointType6Dof PhysicsServer3DJointType = 4
  PhysicsServer3DJointTypeMax PhysicsServer3DJointType = 5
)

type PhysicsServer3DPinJointParam int
const (
  PhysicsServer3DPinJointBias PhysicsServer3DPinJointParam = 0
  PhysicsServer3DPinJointDamping PhysicsServer3DPinJointParam = 1
  PhysicsServer3DPinJointImpulseClamp PhysicsServer3DPinJointParam = 2
)

type PhysicsServer3DHingeJointParam int
const (
  PhysicsServer3DHingeJointBias PhysicsServer3DHingeJointParam = 0
  PhysicsServer3DHingeJointLimitUpper PhysicsServer3DHingeJointParam = 1
  PhysicsServer3DHingeJointLimitLower PhysicsServer3DHingeJointParam = 2
  PhysicsServer3DHingeJointLimitBias PhysicsServer3DHingeJointParam = 3
  PhysicsServer3DHingeJointLimitSoftness PhysicsServer3DHingeJointParam = 4
  PhysicsServer3DHingeJointLimitRelaxation PhysicsServer3DHingeJointParam = 5
  PhysicsServer3DHingeJointMotorTargetVelocity PhysicsServer3DHingeJointParam = 6
  PhysicsServer3DHingeJointMotorMaxImpulse PhysicsServer3DHingeJointParam = 7
)

type PhysicsServer3DHingeJointFlag int
const (
  PhysicsServer3DHingeJointFlagUseLimit PhysicsServer3DHingeJointFlag = 0
  PhysicsServer3DHingeJointFlagEnableMotor PhysicsServer3DHingeJointFlag = 1
)

type PhysicsServer3DSliderJointParam int
const (
  PhysicsServer3DSliderJointLinearLimitUpper PhysicsServer3DSliderJointParam = 0
  PhysicsServer3DSliderJointLinearLimitLower PhysicsServer3DSliderJointParam = 1
  PhysicsServer3DSliderJointLinearLimitSoftness PhysicsServer3DSliderJointParam = 2
  PhysicsServer3DSliderJointLinearLimitRestitution PhysicsServer3DSliderJointParam = 3
  PhysicsServer3DSliderJointLinearLimitDamping PhysicsServer3DSliderJointParam = 4
  PhysicsServer3DSliderJointLinearMotionSoftness PhysicsServer3DSliderJointParam = 5
  PhysicsServer3DSliderJointLinearMotionRestitution PhysicsServer3DSliderJointParam = 6
  PhysicsServer3DSliderJointLinearMotionDamping PhysicsServer3DSliderJointParam = 7
  PhysicsServer3DSliderJointLinearOrthogonalSoftness PhysicsServer3DSliderJointParam = 8
  PhysicsServer3DSliderJointLinearOrthogonalRestitution PhysicsServer3DSliderJointParam = 9
  PhysicsServer3DSliderJointLinearOrthogonalDamping PhysicsServer3DSliderJointParam = 10
  PhysicsServer3DSliderJointAngularLimitUpper PhysicsServer3DSliderJointParam = 11
  PhysicsServer3DSliderJointAngularLimitLower PhysicsServer3DSliderJointParam = 12
  PhysicsServer3DSliderJointAngularLimitSoftness PhysicsServer3DSliderJointParam = 13
  PhysicsServer3DSliderJointAngularLimitRestitution PhysicsServer3DSliderJointParam = 14
  PhysicsServer3DSliderJointAngularLimitDamping PhysicsServer3DSliderJointParam = 15
  PhysicsServer3DSliderJointAngularMotionSoftness PhysicsServer3DSliderJointParam = 16
  PhysicsServer3DSliderJointAngularMotionRestitution PhysicsServer3DSliderJointParam = 17
  PhysicsServer3DSliderJointAngularMotionDamping PhysicsServer3DSliderJointParam = 18
  PhysicsServer3DSliderJointAngularOrthogonalSoftness PhysicsServer3DSliderJointParam = 19
  PhysicsServer3DSliderJointAngularOrthogonalRestitution PhysicsServer3DSliderJointParam = 20
  PhysicsServer3DSliderJointAngularOrthogonalDamping PhysicsServer3DSliderJointParam = 21
  PhysicsServer3DSliderJointMax PhysicsServer3DSliderJointParam = 22
)

type PhysicsServer3DConeTwistJointParam int
const (
  PhysicsServer3DConeTwistJointSwingSpan PhysicsServer3DConeTwistJointParam = 0
  PhysicsServer3DConeTwistJointTwistSpan PhysicsServer3DConeTwistJointParam = 1
  PhysicsServer3DConeTwistJointBias PhysicsServer3DConeTwistJointParam = 2
  PhysicsServer3DConeTwistJointSoftness PhysicsServer3DConeTwistJointParam = 3
  PhysicsServer3DConeTwistJointRelaxation PhysicsServer3DConeTwistJointParam = 4
)

type PhysicsServer3DG6DOFJointAxisParam int
const (
  PhysicsServer3DG6DofJointLinearLowerLimit PhysicsServer3DG6DOFJointAxisParam = 0
  PhysicsServer3DG6DofJointLinearUpperLimit PhysicsServer3DG6DOFJointAxisParam = 1
  PhysicsServer3DG6DofJointLinearLimitSoftness PhysicsServer3DG6DOFJointAxisParam = 2
  PhysicsServer3DG6DofJointLinearRestitution PhysicsServer3DG6DOFJointAxisParam = 3
  PhysicsServer3DG6DofJointLinearDamping PhysicsServer3DG6DOFJointAxisParam = 4
  PhysicsServer3DG6DofJointLinearMotorTargetVelocity PhysicsServer3DG6DOFJointAxisParam = 5
  PhysicsServer3DG6DofJointLinearMotorForceLimit PhysicsServer3DG6DOFJointAxisParam = 6
  PhysicsServer3DG6DofJointAngularLowerLimit PhysicsServer3DG6DOFJointAxisParam = 10
  PhysicsServer3DG6DofJointAngularUpperLimit PhysicsServer3DG6DOFJointAxisParam = 11
  PhysicsServer3DG6DofJointAngularLimitSoftness PhysicsServer3DG6DOFJointAxisParam = 12
  PhysicsServer3DG6DofJointAngularDamping PhysicsServer3DG6DOFJointAxisParam = 13
  PhysicsServer3DG6DofJointAngularRestitution PhysicsServer3DG6DOFJointAxisParam = 14
  PhysicsServer3DG6DofJointAngularForceLimit PhysicsServer3DG6DOFJointAxisParam = 15
  PhysicsServer3DG6DofJointAngularErp PhysicsServer3DG6DOFJointAxisParam = 16
  PhysicsServer3DG6DofJointAngularMotorTargetVelocity PhysicsServer3DG6DOFJointAxisParam = 17
  PhysicsServer3DG6DofJointAngularMotorForceLimit PhysicsServer3DG6DOFJointAxisParam = 18
)

type PhysicsServer3DG6DOFJointAxisFlag int
const (
  PhysicsServer3DG6DofJointFlagEnableLinearLimit PhysicsServer3DG6DOFJointAxisFlag = 0
  PhysicsServer3DG6DofJointFlagEnableAngularLimit PhysicsServer3DG6DOFJointAxisFlag = 1
  PhysicsServer3DG6DofJointFlagEnableMotor PhysicsServer3DG6DOFJointAxisFlag = 4
  PhysicsServer3DG6DofJointFlagEnableLinearMotor PhysicsServer3DG6DOFJointAxisFlag = 5
)

type PhysicsServer3DShapeType int
const (
  PhysicsServer3DShapeWorldBoundary PhysicsServer3DShapeType = 0
  PhysicsServer3DShapeSeparationRay PhysicsServer3DShapeType = 1
  PhysicsServer3DShapeSphere PhysicsServer3DShapeType = 2
  PhysicsServer3DShapeBox PhysicsServer3DShapeType = 3
  PhysicsServer3DShapeCapsule PhysicsServer3DShapeType = 4
  PhysicsServer3DShapeCylinder PhysicsServer3DShapeType = 5
  PhysicsServer3DShapeConvexPolygon PhysicsServer3DShapeType = 6
  PhysicsServer3DShapeConcavePolygon PhysicsServer3DShapeType = 7
  PhysicsServer3DShapeHeightmap PhysicsServer3DShapeType = 8
  PhysicsServer3DShapeSoftBody PhysicsServer3DShapeType = 9
  PhysicsServer3DShapeCustom PhysicsServer3DShapeType = 10
)

type PhysicsServer3DAreaParameter int
const (
  PhysicsServer3DAreaParamGravityOverrideMode PhysicsServer3DAreaParameter = 0
  PhysicsServer3DAreaParamGravity PhysicsServer3DAreaParameter = 1
  PhysicsServer3DAreaParamGravityVector PhysicsServer3DAreaParameter = 2
  PhysicsServer3DAreaParamGravityIsPoint PhysicsServer3DAreaParameter = 3
  PhysicsServer3DAreaParamGravityPointUnitDistance PhysicsServer3DAreaParameter = 4
  PhysicsServer3DAreaParamLinearDampOverrideMode PhysicsServer3DAreaParameter = 5
  PhysicsServer3DAreaParamLinearDamp PhysicsServer3DAreaParameter = 6
  PhysicsServer3DAreaParamAngularDampOverrideMode PhysicsServer3DAreaParameter = 7
  PhysicsServer3DAreaParamAngularDamp PhysicsServer3DAreaParameter = 8
  PhysicsServer3DAreaParamPriority PhysicsServer3DAreaParameter = 9
  PhysicsServer3DAreaParamWindForceMagnitude PhysicsServer3DAreaParameter = 10
  PhysicsServer3DAreaParamWindSource PhysicsServer3DAreaParameter = 11
  PhysicsServer3DAreaParamWindDirection PhysicsServer3DAreaParameter = 12
  PhysicsServer3DAreaParamWindAttenuationFactor PhysicsServer3DAreaParameter = 13
)

type PhysicsServer3DAreaSpaceOverrideMode int
const (
  PhysicsServer3DAreaSpaceOverrideDisabled PhysicsServer3DAreaSpaceOverrideMode = 0
  PhysicsServer3DAreaSpaceOverrideCombine PhysicsServer3DAreaSpaceOverrideMode = 1
  PhysicsServer3DAreaSpaceOverrideCombineReplace PhysicsServer3DAreaSpaceOverrideMode = 2
  PhysicsServer3DAreaSpaceOverrideReplace PhysicsServer3DAreaSpaceOverrideMode = 3
  PhysicsServer3DAreaSpaceOverrideReplaceCombine PhysicsServer3DAreaSpaceOverrideMode = 4
)

type PhysicsServer3DBodyMode int
const (
  PhysicsServer3DBodyModeStatic PhysicsServer3DBodyMode = 0
  PhysicsServer3DBodyModeKinematic PhysicsServer3DBodyMode = 1
  PhysicsServer3DBodyModeRigid PhysicsServer3DBodyMode = 2
  PhysicsServer3DBodyModeRigidLinear PhysicsServer3DBodyMode = 3
)

type PhysicsServer3DBodyParameter int
const (
  PhysicsServer3DBodyParamBounce PhysicsServer3DBodyParameter = 0
  PhysicsServer3DBodyParamFriction PhysicsServer3DBodyParameter = 1
  PhysicsServer3DBodyParamMass PhysicsServer3DBodyParameter = 2
  PhysicsServer3DBodyParamInertia PhysicsServer3DBodyParameter = 3
  PhysicsServer3DBodyParamCenterOfMass PhysicsServer3DBodyParameter = 4
  PhysicsServer3DBodyParamGravityScale PhysicsServer3DBodyParameter = 5
  PhysicsServer3DBodyParamLinearDampMode PhysicsServer3DBodyParameter = 6
  PhysicsServer3DBodyParamAngularDampMode PhysicsServer3DBodyParameter = 7
  PhysicsServer3DBodyParamLinearDamp PhysicsServer3DBodyParameter = 8
  PhysicsServer3DBodyParamAngularDamp PhysicsServer3DBodyParameter = 9
  PhysicsServer3DBodyParamMax PhysicsServer3DBodyParameter = 10
)

type PhysicsServer3DBodyDampMode int
const (
  PhysicsServer3DBodyDampModeCombine PhysicsServer3DBodyDampMode = 0
  PhysicsServer3DBodyDampModeReplace PhysicsServer3DBodyDampMode = 1
)

type PhysicsServer3DBodyState int
const (
  PhysicsServer3DBodyStateTransform PhysicsServer3DBodyState = 0
  PhysicsServer3DBodyStateLinearVelocity PhysicsServer3DBodyState = 1
  PhysicsServer3DBodyStateAngularVelocity PhysicsServer3DBodyState = 2
  PhysicsServer3DBodyStateSleeping PhysicsServer3DBodyState = 3
  PhysicsServer3DBodyStateCanSleep PhysicsServer3DBodyState = 4
)

type PhysicsServer3DAreaBodyStatus int
const (
  PhysicsServer3DAreaBodyAdded PhysicsServer3DAreaBodyStatus = 0
  PhysicsServer3DAreaBodyRemoved PhysicsServer3DAreaBodyStatus = 1
)

type PhysicsServer3DProcessInfo int
const (
  PhysicsServer3DInfoActiveObjects PhysicsServer3DProcessInfo = 0
  PhysicsServer3DInfoCollisionPairs PhysicsServer3DProcessInfo = 1
  PhysicsServer3DInfoIslandCount PhysicsServer3DProcessInfo = 2
)

type PhysicsServer3DSpaceParameter int
const (
  PhysicsServer3DSpaceParamContactRecycleRadius PhysicsServer3DSpaceParameter = 0
  PhysicsServer3DSpaceParamContactMaxSeparation PhysicsServer3DSpaceParameter = 1
  PhysicsServer3DSpaceParamContactMaxAllowedPenetration PhysicsServer3DSpaceParameter = 2
  PhysicsServer3DSpaceParamContactDefaultBias PhysicsServer3DSpaceParameter = 3
  PhysicsServer3DSpaceParamBodyLinearVelocitySleepThreshold PhysicsServer3DSpaceParameter = 4
  PhysicsServer3DSpaceParamBodyAngularVelocitySleepThreshold PhysicsServer3DSpaceParameter = 5
  PhysicsServer3DSpaceParamBodyTimeToSleep PhysicsServer3DSpaceParameter = 6
  PhysicsServer3DSpaceParamSolverIterations PhysicsServer3DSpaceParameter = 7
)

type PhysicsServer3DBodyAxis int
const (
  PhysicsServer3DBodyAxisLinearX PhysicsServer3DBodyAxis = 1
  PhysicsServer3DBodyAxisLinearY PhysicsServer3DBodyAxis = 2
  PhysicsServer3DBodyAxisLinearZ PhysicsServer3DBodyAxis = 4
  PhysicsServer3DBodyAxisAngularX PhysicsServer3DBodyAxis = 8
  PhysicsServer3DBodyAxisAngularY PhysicsServer3DBodyAxis = 16
  PhysicsServer3DBodyAxisAngularZ PhysicsServer3DBodyAxis = 32
)
