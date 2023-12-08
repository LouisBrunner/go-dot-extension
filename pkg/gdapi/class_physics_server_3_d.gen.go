// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  PhysicsServer3DJointTypeJointTypePin PhysicsServer3DJointType = 0
  PhysicsServer3DJointTypeJointTypeHinge PhysicsServer3DJointType = 1
  PhysicsServer3DJointTypeJointTypeSlider PhysicsServer3DJointType = 2
  PhysicsServer3DJointTypeJointTypeConeTwist PhysicsServer3DJointType = 3
  PhysicsServer3DJointTypeJointType6Dof PhysicsServer3DJointType = 4
  PhysicsServer3DJointTypeJointTypeMax PhysicsServer3DJointType = 5
)

type PhysicsServer3DPinJointParam int
const (
  PhysicsServer3DPinJointParamPinJointBias PhysicsServer3DPinJointParam = 0
  PhysicsServer3DPinJointParamPinJointDamping PhysicsServer3DPinJointParam = 1
  PhysicsServer3DPinJointParamPinJointImpulseClamp PhysicsServer3DPinJointParam = 2
)

type PhysicsServer3DHingeJointParam int
const (
  PhysicsServer3DHingeJointParamHingeJointBias PhysicsServer3DHingeJointParam = 0
  PhysicsServer3DHingeJointParamHingeJointLimitUpper PhysicsServer3DHingeJointParam = 1
  PhysicsServer3DHingeJointParamHingeJointLimitLower PhysicsServer3DHingeJointParam = 2
  PhysicsServer3DHingeJointParamHingeJointLimitBias PhysicsServer3DHingeJointParam = 3
  PhysicsServer3DHingeJointParamHingeJointLimitSoftness PhysicsServer3DHingeJointParam = 4
  PhysicsServer3DHingeJointParamHingeJointLimitRelaxation PhysicsServer3DHingeJointParam = 5
  PhysicsServer3DHingeJointParamHingeJointMotorTargetVelocity PhysicsServer3DHingeJointParam = 6
  PhysicsServer3DHingeJointParamHingeJointMotorMaxImpulse PhysicsServer3DHingeJointParam = 7
)

type PhysicsServer3DHingeJointFlag int
const (
  PhysicsServer3DHingeJointFlagHingeJointFlagUseLimit PhysicsServer3DHingeJointFlag = 0
  PhysicsServer3DHingeJointFlagHingeJointFlagEnableMotor PhysicsServer3DHingeJointFlag = 1
)

type PhysicsServer3DSliderJointParam int
const (
  PhysicsServer3DSliderJointParamSliderJointLinearLimitUpper PhysicsServer3DSliderJointParam = 0
  PhysicsServer3DSliderJointParamSliderJointLinearLimitLower PhysicsServer3DSliderJointParam = 1
  PhysicsServer3DSliderJointParamSliderJointLinearLimitSoftness PhysicsServer3DSliderJointParam = 2
  PhysicsServer3DSliderJointParamSliderJointLinearLimitRestitution PhysicsServer3DSliderJointParam = 3
  PhysicsServer3DSliderJointParamSliderJointLinearLimitDamping PhysicsServer3DSliderJointParam = 4
  PhysicsServer3DSliderJointParamSliderJointLinearMotionSoftness PhysicsServer3DSliderJointParam = 5
  PhysicsServer3DSliderJointParamSliderJointLinearMotionRestitution PhysicsServer3DSliderJointParam = 6
  PhysicsServer3DSliderJointParamSliderJointLinearMotionDamping PhysicsServer3DSliderJointParam = 7
  PhysicsServer3DSliderJointParamSliderJointLinearOrthogonalSoftness PhysicsServer3DSliderJointParam = 8
  PhysicsServer3DSliderJointParamSliderJointLinearOrthogonalRestitution PhysicsServer3DSliderJointParam = 9
  PhysicsServer3DSliderJointParamSliderJointLinearOrthogonalDamping PhysicsServer3DSliderJointParam = 10
  PhysicsServer3DSliderJointParamSliderJointAngularLimitUpper PhysicsServer3DSliderJointParam = 11
  PhysicsServer3DSliderJointParamSliderJointAngularLimitLower PhysicsServer3DSliderJointParam = 12
  PhysicsServer3DSliderJointParamSliderJointAngularLimitSoftness PhysicsServer3DSliderJointParam = 13
  PhysicsServer3DSliderJointParamSliderJointAngularLimitRestitution PhysicsServer3DSliderJointParam = 14
  PhysicsServer3DSliderJointParamSliderJointAngularLimitDamping PhysicsServer3DSliderJointParam = 15
  PhysicsServer3DSliderJointParamSliderJointAngularMotionSoftness PhysicsServer3DSliderJointParam = 16
  PhysicsServer3DSliderJointParamSliderJointAngularMotionRestitution PhysicsServer3DSliderJointParam = 17
  PhysicsServer3DSliderJointParamSliderJointAngularMotionDamping PhysicsServer3DSliderJointParam = 18
  PhysicsServer3DSliderJointParamSliderJointAngularOrthogonalSoftness PhysicsServer3DSliderJointParam = 19
  PhysicsServer3DSliderJointParamSliderJointAngularOrthogonalRestitution PhysicsServer3DSliderJointParam = 20
  PhysicsServer3DSliderJointParamSliderJointAngularOrthogonalDamping PhysicsServer3DSliderJointParam = 21
  PhysicsServer3DSliderJointParamSliderJointMax PhysicsServer3DSliderJointParam = 22
)

type PhysicsServer3DConeTwistJointParam int
const (
  PhysicsServer3DConeTwistJointParamConeTwistJointSwingSpan PhysicsServer3DConeTwistJointParam = 0
  PhysicsServer3DConeTwistJointParamConeTwistJointTwistSpan PhysicsServer3DConeTwistJointParam = 1
  PhysicsServer3DConeTwistJointParamConeTwistJointBias PhysicsServer3DConeTwistJointParam = 2
  PhysicsServer3DConeTwistJointParamConeTwistJointSoftness PhysicsServer3DConeTwistJointParam = 3
  PhysicsServer3DConeTwistJointParamConeTwistJointRelaxation PhysicsServer3DConeTwistJointParam = 4
)

type PhysicsServer3DG6DOFJointAxisParam int
const (
  PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearLowerLimit PhysicsServer3DG6DOFJointAxisParam = 0
  PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearUpperLimit PhysicsServer3DG6DOFJointAxisParam = 1
  PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearLimitSoftness PhysicsServer3DG6DOFJointAxisParam = 2
  PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearRestitution PhysicsServer3DG6DOFJointAxisParam = 3
  PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearDamping PhysicsServer3DG6DOFJointAxisParam = 4
  PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearMotorTargetVelocity PhysicsServer3DG6DOFJointAxisParam = 5
  PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearMotorForceLimit PhysicsServer3DG6DOFJointAxisParam = 6
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularLowerLimit PhysicsServer3DG6DOFJointAxisParam = 10
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularUpperLimit PhysicsServer3DG6DOFJointAxisParam = 11
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularLimitSoftness PhysicsServer3DG6DOFJointAxisParam = 12
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularDamping PhysicsServer3DG6DOFJointAxisParam = 13
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularRestitution PhysicsServer3DG6DOFJointAxisParam = 14
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularForceLimit PhysicsServer3DG6DOFJointAxisParam = 15
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularErp PhysicsServer3DG6DOFJointAxisParam = 16
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularMotorTargetVelocity PhysicsServer3DG6DOFJointAxisParam = 17
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularMotorForceLimit PhysicsServer3DG6DOFJointAxisParam = 18
)

type PhysicsServer3DG6DOFJointAxisFlag int
const (
  PhysicsServer3DG6DOFJointAxisFlagG6DofJointFlagEnableLinearLimit PhysicsServer3DG6DOFJointAxisFlag = 0
  PhysicsServer3DG6DOFJointAxisFlagG6DofJointFlagEnableAngularLimit PhysicsServer3DG6DOFJointAxisFlag = 1
  PhysicsServer3DG6DOFJointAxisFlagG6DofJointFlagEnableMotor PhysicsServer3DG6DOFJointAxisFlag = 4
  PhysicsServer3DG6DOFJointAxisFlagG6DofJointFlagEnableLinearMotor PhysicsServer3DG6DOFJointAxisFlag = 5
)

type PhysicsServer3DShapeType int
const (
  PhysicsServer3DShapeTypeShapeWorldBoundary PhysicsServer3DShapeType = 0
  PhysicsServer3DShapeTypeShapeSeparationRay PhysicsServer3DShapeType = 1
  PhysicsServer3DShapeTypeShapeSphere PhysicsServer3DShapeType = 2
  PhysicsServer3DShapeTypeShapeBox PhysicsServer3DShapeType = 3
  PhysicsServer3DShapeTypeShapeCapsule PhysicsServer3DShapeType = 4
  PhysicsServer3DShapeTypeShapeCylinder PhysicsServer3DShapeType = 5
  PhysicsServer3DShapeTypeShapeConvexPolygon PhysicsServer3DShapeType = 6
  PhysicsServer3DShapeTypeShapeConcavePolygon PhysicsServer3DShapeType = 7
  PhysicsServer3DShapeTypeShapeHeightmap PhysicsServer3DShapeType = 8
  PhysicsServer3DShapeTypeShapeSoftBody PhysicsServer3DShapeType = 9
  PhysicsServer3DShapeTypeShapeCustom PhysicsServer3DShapeType = 10
)

type PhysicsServer3DAreaParameter int
const (
  PhysicsServer3DAreaParameterAreaParamGravityOverrideMode PhysicsServer3DAreaParameter = 0
  PhysicsServer3DAreaParameterAreaParamGravity PhysicsServer3DAreaParameter = 1
  PhysicsServer3DAreaParameterAreaParamGravityVector PhysicsServer3DAreaParameter = 2
  PhysicsServer3DAreaParameterAreaParamGravityIsPoint PhysicsServer3DAreaParameter = 3
  PhysicsServer3DAreaParameterAreaParamGravityPointUnitDistance PhysicsServer3DAreaParameter = 4
  PhysicsServer3DAreaParameterAreaParamLinearDampOverrideMode PhysicsServer3DAreaParameter = 5
  PhysicsServer3DAreaParameterAreaParamLinearDamp PhysicsServer3DAreaParameter = 6
  PhysicsServer3DAreaParameterAreaParamAngularDampOverrideMode PhysicsServer3DAreaParameter = 7
  PhysicsServer3DAreaParameterAreaParamAngularDamp PhysicsServer3DAreaParameter = 8
  PhysicsServer3DAreaParameterAreaParamPriority PhysicsServer3DAreaParameter = 9
  PhysicsServer3DAreaParameterAreaParamWindForceMagnitude PhysicsServer3DAreaParameter = 10
  PhysicsServer3DAreaParameterAreaParamWindSource PhysicsServer3DAreaParameter = 11
  PhysicsServer3DAreaParameterAreaParamWindDirection PhysicsServer3DAreaParameter = 12
  PhysicsServer3DAreaParameterAreaParamWindAttenuationFactor PhysicsServer3DAreaParameter = 13
)

type PhysicsServer3DAreaSpaceOverrideMode int
const (
  PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideDisabled PhysicsServer3DAreaSpaceOverrideMode = 0
  PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideCombine PhysicsServer3DAreaSpaceOverrideMode = 1
  PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideCombineReplace PhysicsServer3DAreaSpaceOverrideMode = 2
  PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideReplace PhysicsServer3DAreaSpaceOverrideMode = 3
  PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideReplaceCombine PhysicsServer3DAreaSpaceOverrideMode = 4
)

type PhysicsServer3DBodyMode int
const (
  PhysicsServer3DBodyModeBodyModeStatic PhysicsServer3DBodyMode = 0
  PhysicsServer3DBodyModeBodyModeKinematic PhysicsServer3DBodyMode = 1
  PhysicsServer3DBodyModeBodyModeRigid PhysicsServer3DBodyMode = 2
  PhysicsServer3DBodyModeBodyModeRigidLinear PhysicsServer3DBodyMode = 3
)

type PhysicsServer3DBodyParameter int
const (
  PhysicsServer3DBodyParameterBodyParamBounce PhysicsServer3DBodyParameter = 0
  PhysicsServer3DBodyParameterBodyParamFriction PhysicsServer3DBodyParameter = 1
  PhysicsServer3DBodyParameterBodyParamMass PhysicsServer3DBodyParameter = 2
  PhysicsServer3DBodyParameterBodyParamInertia PhysicsServer3DBodyParameter = 3
  PhysicsServer3DBodyParameterBodyParamCenterOfMass PhysicsServer3DBodyParameter = 4
  PhysicsServer3DBodyParameterBodyParamGravityScale PhysicsServer3DBodyParameter = 5
  PhysicsServer3DBodyParameterBodyParamLinearDampMode PhysicsServer3DBodyParameter = 6
  PhysicsServer3DBodyParameterBodyParamAngularDampMode PhysicsServer3DBodyParameter = 7
  PhysicsServer3DBodyParameterBodyParamLinearDamp PhysicsServer3DBodyParameter = 8
  PhysicsServer3DBodyParameterBodyParamAngularDamp PhysicsServer3DBodyParameter = 9
  PhysicsServer3DBodyParameterBodyParamMax PhysicsServer3DBodyParameter = 10
)

type PhysicsServer3DBodyDampMode int
const (
  PhysicsServer3DBodyDampModeBodyDampModeCombine PhysicsServer3DBodyDampMode = 0
  PhysicsServer3DBodyDampModeBodyDampModeReplace PhysicsServer3DBodyDampMode = 1
)

type PhysicsServer3DBodyState int
const (
  PhysicsServer3DBodyStateBodyStateTransform PhysicsServer3DBodyState = 0
  PhysicsServer3DBodyStateBodyStateLinearVelocity PhysicsServer3DBodyState = 1
  PhysicsServer3DBodyStateBodyStateAngularVelocity PhysicsServer3DBodyState = 2
  PhysicsServer3DBodyStateBodyStateSleeping PhysicsServer3DBodyState = 3
  PhysicsServer3DBodyStateBodyStateCanSleep PhysicsServer3DBodyState = 4
)

type PhysicsServer3DAreaBodyStatus int
const (
  PhysicsServer3DAreaBodyStatusAreaBodyAdded PhysicsServer3DAreaBodyStatus = 0
  PhysicsServer3DAreaBodyStatusAreaBodyRemoved PhysicsServer3DAreaBodyStatus = 1
)

type PhysicsServer3DProcessInfo int
const (
  PhysicsServer3DProcessInfoInfoActiveObjects PhysicsServer3DProcessInfo = 0
  PhysicsServer3DProcessInfoInfoCollisionPairs PhysicsServer3DProcessInfo = 1
  PhysicsServer3DProcessInfoInfoIslandCount PhysicsServer3DProcessInfo = 2
)

type PhysicsServer3DSpaceParameter int
const (
  PhysicsServer3DSpaceParameterSpaceParamContactRecycleRadius PhysicsServer3DSpaceParameter = 0
  PhysicsServer3DSpaceParameterSpaceParamContactMaxSeparation PhysicsServer3DSpaceParameter = 1
  PhysicsServer3DSpaceParameterSpaceParamContactMaxAllowedPenetration PhysicsServer3DSpaceParameter = 2
  PhysicsServer3DSpaceParameterSpaceParamContactDefaultBias PhysicsServer3DSpaceParameter = 3
  PhysicsServer3DSpaceParameterSpaceParamBodyLinearVelocitySleepThreshold PhysicsServer3DSpaceParameter = 4
  PhysicsServer3DSpaceParameterSpaceParamBodyAngularVelocitySleepThreshold PhysicsServer3DSpaceParameter = 5
  PhysicsServer3DSpaceParameterSpaceParamBodyTimeToSleep PhysicsServer3DSpaceParameter = 6
  PhysicsServer3DSpaceParameterSpaceParamSolverIterations PhysicsServer3DSpaceParameter = 7
)

type PhysicsServer3DBodyAxis int
const (
  PhysicsServer3DBodyAxisBodyAxisLinearX PhysicsServer3DBodyAxis = 1
  PhysicsServer3DBodyAxisBodyAxisLinearY PhysicsServer3DBodyAxis = 2
  PhysicsServer3DBodyAxisBodyAxisLinearZ PhysicsServer3DBodyAxis = 4
  PhysicsServer3DBodyAxisBodyAxisAngularX PhysicsServer3DBodyAxis = 8
  PhysicsServer3DBodyAxisBodyAxisAngularY PhysicsServer3DBodyAxis = 16
  PhysicsServer3DBodyAxisBodyAxisAngularZ PhysicsServer3DBodyAxis = 32
)

func  (me *PhysicsServer3D) WorldBoundaryShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) SeparationRayShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) SphereShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BoxShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) CapsuleShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) CylinderShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) ConvexPolygonShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) ConcavePolygonShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) HeightmapShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) CustomShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) ShapeSetData(shape RID, data Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) ShapeGetType(shape RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) ShapeGetData(shape RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) SpaceCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) SpaceSetActive(space RID, active bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) SpaceIsActive(space RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) SpaceSetParam(space RID, param PhysicsServer3DSpaceParameter, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) SpaceGetParam(space RID, param PhysicsServer3DSpaceParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) SpaceGetDirectState(space RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaSetSpace(area RID, space RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaGetSpace(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaAddShape(area RID, shape RID, transform Transform3D, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaSetShape(area RID, shape_idx int, shape RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaSetShapeTransform(area RID, shape_idx int, transform Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaSetShapeDisabled(area RID, shape_idx int, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaGetShapeCount(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaGetShape(area RID, shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaGetShapeTransform(area RID, shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaRemoveShape(area RID, shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaClearShapes(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaSetCollisionLayer(area RID, layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaGetCollisionLayer(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaSetCollisionMask(area RID, mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaGetCollisionMask(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaSetParam(area RID, param PhysicsServer3DAreaParameter, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaSetTransform(area RID, transform Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaGetParam(area RID, param PhysicsServer3DAreaParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaGetTransform(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaAttachObjectInstanceId(area RID, id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaGetObjectInstanceId(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaSetMonitorCallback(area RID, callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaSetAreaMonitorCallback(area RID, callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaSetMonitorable(area RID, monitorable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) AreaSetRayPickable(area RID, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetSpace(body RID, space RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetSpace(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetMode(body RID, mode PhysicsServer3DBodyMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetMode(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetCollisionLayer(body RID, layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetCollisionLayer(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetCollisionMask(body RID, mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetCollisionMask(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetCollisionPriority(body RID, priority float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetCollisionPriority(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyAddShape(body RID, shape RID, transform Transform3D, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetShape(body RID, shape_idx int, shape RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetShapeTransform(body RID, shape_idx int, transform Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetShapeDisabled(body RID, shape_idx int, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetShapeCount(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetShape(body RID, shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetShapeTransform(body RID, shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyRemoveShape(body RID, shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyClearShapes(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyAttachObjectInstanceId(body RID, id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetObjectInstanceId(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetEnableContinuousCollisionDetection(body RID, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyIsContinuousCollisionDetectionEnabled(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetParam(body RID, param PhysicsServer3DBodyParameter, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetParam(body RID, param PhysicsServer3DBodyParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyResetMassProperties(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetState(body RID, state PhysicsServer3DBodyState, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetState(body RID, state PhysicsServer3DBodyState, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyApplyCentralImpulse(body RID, impulse Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyApplyImpulse(body RID, impulse Vector3, position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyApplyTorqueImpulse(body RID, impulse Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyApplyCentralForce(body RID, force Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyApplyForce(body RID, force Vector3, position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyApplyTorque(body RID, torque Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyAddConstantCentralForce(body RID, force Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyAddConstantForce(body RID, force Vector3, position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyAddConstantTorque(body RID, torque Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetConstantForce(body RID, force Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetConstantForce(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetConstantTorque(body RID, torque Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetConstantTorque(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetAxisVelocity(body RID, axis_velocity Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetAxisLock(body RID, axis PhysicsServer3DBodyAxis, lock bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyIsAxisLocked(body RID, axis PhysicsServer3DBodyAxis, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyAddCollisionException(body RID, excepted_body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyRemoveCollisionException(body RID, excepted_body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetMaxContactsReported(body RID, amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetMaxContactsReported(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetOmitForceIntegration(body RID, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyIsOmittingForceIntegration(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetForceIntegrationCallback(body RID, callable Callable, userdata Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodySetRayPickable(body RID, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyTestMotion(body RID, parameters PhysicsTestMotionParameters3D, result PhysicsTestMotionResult3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) BodyGetDirectState(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) SoftBodyGetBounds(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) JointCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) JointClear(joint RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) JointMakePin(joint RID, body_A RID, local_A Vector3, body_B RID, local_B Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) PinJointSetParam(joint RID, param PhysicsServer3DPinJointParam, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) PinJointGetParam(joint RID, param PhysicsServer3DPinJointParam, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) PinJointSetLocalA(joint RID, local_A Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) PinJointGetLocalA(joint RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) PinJointSetLocalB(joint RID, local_B Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) PinJointGetLocalB(joint RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) JointMakeHinge(joint RID, body_A RID, hinge_A Transform3D, body_B RID, hinge_B Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) HingeJointSetParam(joint RID, param PhysicsServer3DHingeJointParam, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) HingeJointGetParam(joint RID, param PhysicsServer3DHingeJointParam, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) HingeJointSetFlag(joint RID, flag PhysicsServer3DHingeJointFlag, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) HingeJointGetFlag(joint RID, flag PhysicsServer3DHingeJointFlag, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) JointMakeSlider(joint RID, body_A RID, local_ref_A Transform3D, body_B RID, local_ref_B Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) SliderJointSetParam(joint RID, param PhysicsServer3DSliderJointParam, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) SliderJointGetParam(joint RID, param PhysicsServer3DSliderJointParam, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) JointMakeConeTwist(joint RID, body_A RID, local_ref_A Transform3D, body_B RID, local_ref_B Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) ConeTwistJointSetParam(joint RID, param PhysicsServer3DConeTwistJointParam, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) ConeTwistJointGetParam(joint RID, param PhysicsServer3DConeTwistJointParam, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) JointGetType(joint RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) JointSetSolverPriority(joint RID, priority int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) JointGetSolverPriority(joint RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) JointDisableCollisionsBetweenBodies(joint RID, disable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) JointIsDisabledCollisionsBetweenBodies(joint RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) JointMakeGeneric6Dof(joint RID, body_A RID, local_ref_A Transform3D, body_B RID, local_ref_B Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) Generic6DofJointSetParam(joint RID, axis Vector3Axis, param PhysicsServer3DG6DOFJointAxisParam, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) Generic6DofJointGetParam(joint RID, axis Vector3Axis, param PhysicsServer3DG6DOFJointAxisParam, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) Generic6DofJointSetFlag(joint RID, axis Vector3Axis, flag PhysicsServer3DG6DOFJointAxisFlag, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) Generic6DofJointGetFlag(joint RID, axis Vector3Axis, flag PhysicsServer3DG6DOFJointAxisFlag, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) FreeRid(rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) SetActive(active bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer3D) GetProcessInfo(process_info PhysicsServer3DProcessInfo, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
