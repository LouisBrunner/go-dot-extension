// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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



// Enums

type PhysicsServer2DSpaceParameter int
const (
  PhysicsServer2DSpaceParameterSpaceParamContactRecycleRadius PhysicsServer2DSpaceParameter = 0
  PhysicsServer2DSpaceParameterSpaceParamContactMaxSeparation PhysicsServer2DSpaceParameter = 1
  PhysicsServer2DSpaceParameterSpaceParamContactMaxAllowedPenetration PhysicsServer2DSpaceParameter = 2
  PhysicsServer2DSpaceParameterSpaceParamContactDefaultBias PhysicsServer2DSpaceParameter = 3
  PhysicsServer2DSpaceParameterSpaceParamBodyLinearVelocitySleepThreshold PhysicsServer2DSpaceParameter = 4
  PhysicsServer2DSpaceParameterSpaceParamBodyAngularVelocitySleepThreshold PhysicsServer2DSpaceParameter = 5
  PhysicsServer2DSpaceParameterSpaceParamBodyTimeToSleep PhysicsServer2DSpaceParameter = 6
  PhysicsServer2DSpaceParameterSpaceParamConstraintDefaultBias PhysicsServer2DSpaceParameter = 7
  PhysicsServer2DSpaceParameterSpaceParamSolverIterations PhysicsServer2DSpaceParameter = 8
)

type PhysicsServer2DShapeType int
const (
  PhysicsServer2DShapeTypeShapeWorldBoundary PhysicsServer2DShapeType = 0
  PhysicsServer2DShapeTypeShapeSeparationRay PhysicsServer2DShapeType = 1
  PhysicsServer2DShapeTypeShapeSegment PhysicsServer2DShapeType = 2
  PhysicsServer2DShapeTypeShapeCircle PhysicsServer2DShapeType = 3
  PhysicsServer2DShapeTypeShapeRectangle PhysicsServer2DShapeType = 4
  PhysicsServer2DShapeTypeShapeCapsule PhysicsServer2DShapeType = 5
  PhysicsServer2DShapeTypeShapeConvexPolygon PhysicsServer2DShapeType = 6
  PhysicsServer2DShapeTypeShapeConcavePolygon PhysicsServer2DShapeType = 7
  PhysicsServer2DShapeTypeShapeCustom PhysicsServer2DShapeType = 8
)

type PhysicsServer2DAreaParameter int
const (
  PhysicsServer2DAreaParameterAreaParamGravityOverrideMode PhysicsServer2DAreaParameter = 0
  PhysicsServer2DAreaParameterAreaParamGravity PhysicsServer2DAreaParameter = 1
  PhysicsServer2DAreaParameterAreaParamGravityVector PhysicsServer2DAreaParameter = 2
  PhysicsServer2DAreaParameterAreaParamGravityIsPoint PhysicsServer2DAreaParameter = 3
  PhysicsServer2DAreaParameterAreaParamGravityPointUnitDistance PhysicsServer2DAreaParameter = 4
  PhysicsServer2DAreaParameterAreaParamLinearDampOverrideMode PhysicsServer2DAreaParameter = 5
  PhysicsServer2DAreaParameterAreaParamLinearDamp PhysicsServer2DAreaParameter = 6
  PhysicsServer2DAreaParameterAreaParamAngularDampOverrideMode PhysicsServer2DAreaParameter = 7
  PhysicsServer2DAreaParameterAreaParamAngularDamp PhysicsServer2DAreaParameter = 8
  PhysicsServer2DAreaParameterAreaParamPriority PhysicsServer2DAreaParameter = 9
)

type PhysicsServer2DAreaSpaceOverrideMode int
const (
  PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideDisabled PhysicsServer2DAreaSpaceOverrideMode = 0
  PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideCombine PhysicsServer2DAreaSpaceOverrideMode = 1
  PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideCombineReplace PhysicsServer2DAreaSpaceOverrideMode = 2
  PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideReplace PhysicsServer2DAreaSpaceOverrideMode = 3
  PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideReplaceCombine PhysicsServer2DAreaSpaceOverrideMode = 4
)

type PhysicsServer2DBodyMode int
const (
  PhysicsServer2DBodyModeBodyModeStatic PhysicsServer2DBodyMode = 0
  PhysicsServer2DBodyModeBodyModeKinematic PhysicsServer2DBodyMode = 1
  PhysicsServer2DBodyModeBodyModeRigid PhysicsServer2DBodyMode = 2
  PhysicsServer2DBodyModeBodyModeRigidLinear PhysicsServer2DBodyMode = 3
)

type PhysicsServer2DBodyParameter int
const (
  PhysicsServer2DBodyParameterBodyParamBounce PhysicsServer2DBodyParameter = 0
  PhysicsServer2DBodyParameterBodyParamFriction PhysicsServer2DBodyParameter = 1
  PhysicsServer2DBodyParameterBodyParamMass PhysicsServer2DBodyParameter = 2
  PhysicsServer2DBodyParameterBodyParamInertia PhysicsServer2DBodyParameter = 3
  PhysicsServer2DBodyParameterBodyParamCenterOfMass PhysicsServer2DBodyParameter = 4
  PhysicsServer2DBodyParameterBodyParamGravityScale PhysicsServer2DBodyParameter = 5
  PhysicsServer2DBodyParameterBodyParamLinearDampMode PhysicsServer2DBodyParameter = 6
  PhysicsServer2DBodyParameterBodyParamAngularDampMode PhysicsServer2DBodyParameter = 7
  PhysicsServer2DBodyParameterBodyParamLinearDamp PhysicsServer2DBodyParameter = 8
  PhysicsServer2DBodyParameterBodyParamAngularDamp PhysicsServer2DBodyParameter = 9
  PhysicsServer2DBodyParameterBodyParamMax PhysicsServer2DBodyParameter = 10
)

type PhysicsServer2DBodyDampMode int
const (
  PhysicsServer2DBodyDampModeBodyDampModeCombine PhysicsServer2DBodyDampMode = 0
  PhysicsServer2DBodyDampModeBodyDampModeReplace PhysicsServer2DBodyDampMode = 1
)

type PhysicsServer2DBodyState int
const (
  PhysicsServer2DBodyStateBodyStateTransform PhysicsServer2DBodyState = 0
  PhysicsServer2DBodyStateBodyStateLinearVelocity PhysicsServer2DBodyState = 1
  PhysicsServer2DBodyStateBodyStateAngularVelocity PhysicsServer2DBodyState = 2
  PhysicsServer2DBodyStateBodyStateSleeping PhysicsServer2DBodyState = 3
  PhysicsServer2DBodyStateBodyStateCanSleep PhysicsServer2DBodyState = 4
)

type PhysicsServer2DJointType int
const (
  PhysicsServer2DJointTypeJointTypePin PhysicsServer2DJointType = 0
  PhysicsServer2DJointTypeJointTypeGroove PhysicsServer2DJointType = 1
  PhysicsServer2DJointTypeJointTypeDampedSpring PhysicsServer2DJointType = 2
  PhysicsServer2DJointTypeJointTypeMax PhysicsServer2DJointType = 3
)

type PhysicsServer2DJointParam int
const (
  PhysicsServer2DJointParamJointParamBias PhysicsServer2DJointParam = 0
  PhysicsServer2DJointParamJointParamMaxBias PhysicsServer2DJointParam = 1
  PhysicsServer2DJointParamJointParamMaxForce PhysicsServer2DJointParam = 2
)

type PhysicsServer2DPinJointParam int
const (
  PhysicsServer2DPinJointParamPinJointSoftness PhysicsServer2DPinJointParam = 0
)

type PhysicsServer2DDampedSpringParam int
const (
  PhysicsServer2DDampedSpringParamDampedSpringRestLength PhysicsServer2DDampedSpringParam = 0
  PhysicsServer2DDampedSpringParamDampedSpringStiffness PhysicsServer2DDampedSpringParam = 1
  PhysicsServer2DDampedSpringParamDampedSpringDamping PhysicsServer2DDampedSpringParam = 2
)

type PhysicsServer2DCCDMode int
const (
  PhysicsServer2DCCDModeCcdModeDisabled PhysicsServer2DCCDMode = 0
  PhysicsServer2DCCDModeCcdModeCastRay PhysicsServer2DCCDMode = 1
  PhysicsServer2DCCDModeCcdModeCastShape PhysicsServer2DCCDMode = 2
)

type PhysicsServer2DAreaBodyStatus int
const (
  PhysicsServer2DAreaBodyStatusAreaBodyAdded PhysicsServer2DAreaBodyStatus = 0
  PhysicsServer2DAreaBodyStatusAreaBodyRemoved PhysicsServer2DAreaBodyStatus = 1
)

type PhysicsServer2DProcessInfo int
const (
  PhysicsServer2DProcessInfoInfoActiveObjects PhysicsServer2DProcessInfo = 0
  PhysicsServer2DProcessInfoInfoCollisionPairs PhysicsServer2DProcessInfo = 1
  PhysicsServer2DProcessInfoInfoIslandCount PhysicsServer2DProcessInfo = 2
)

func (me *PhysicsServer2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsServer2D) WorldBoundaryShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) SeparationRayShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) SegmentShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) CircleShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) RectangleShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) CapsuleShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) ConvexPolygonShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) ConcavePolygonShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) ShapeSetData(shape RID, data Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) ShapeGetType(shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) ShapeGetData(shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) SpaceCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) SpaceSetActive(space RID, active bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) SpaceIsActive(space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) SpaceSetParam(space RID, param PhysicsServer2DSpaceParameter, value float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) SpaceGetParam(space RID, param PhysicsServer2DSpaceParameter, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) SpaceGetDirectState(space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaSetSpace(area RID, space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaGetSpace(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaAddShape(area RID, shape RID, transform Transform2D, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaSetShape(area RID, shape_idx int, shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaSetShapeTransform(area RID, shape_idx int, transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaSetShapeDisabled(area RID, shape_idx int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaGetShapeCount(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaGetShape(area RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaGetShapeTransform(area RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaRemoveShape(area RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaClearShapes(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaSetCollisionLayer(area RID, layer int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaGetCollisionLayer(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaSetCollisionMask(area RID, mask int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaGetCollisionMask(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaSetParam(area RID, param PhysicsServer2DAreaParameter, value Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaSetTransform(area RID, transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaGetParam(area RID, param PhysicsServer2DAreaParameter, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaGetTransform(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaAttachObjectInstanceId(area RID, id int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaGetObjectInstanceId(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaAttachCanvasInstanceId(area RID, id int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaGetCanvasInstanceId(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaSetMonitorCallback(area RID, callback Callable, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaSetAreaMonitorCallback(area RID, callback Callable, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) AreaSetMonitorable(area RID, monitorable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetSpace(body RID, space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetSpace(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetMode(body RID, mode PhysicsServer2DBodyMode, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetMode(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyAddShape(body RID, shape RID, transform Transform2D, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetShape(body RID, shape_idx int, shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetShapeTransform(body RID, shape_idx int, transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetShapeCount(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetShape(body RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetShapeTransform(body RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyRemoveShape(body RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyClearShapes(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetShapeDisabled(body RID, shape_idx int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetShapeAsOneWayCollision(body RID, shape_idx int, enable bool, margin float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyAttachObjectInstanceId(body RID, id int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetObjectInstanceId(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyAttachCanvasInstanceId(body RID, id int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetCanvasInstanceId(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetContinuousCollisionDetectionMode(body RID, mode PhysicsServer2DCCDMode, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetContinuousCollisionDetectionMode(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetCollisionLayer(body RID, layer int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetCollisionLayer(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetCollisionMask(body RID, mask int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetCollisionMask(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetCollisionPriority(body RID, priority float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetCollisionPriority(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetParam(body RID, param PhysicsServer2DBodyParameter, value Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetParam(body RID, param PhysicsServer2DBodyParameter, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyResetMassProperties(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetState(body RID, state PhysicsServer2DBodyState, value Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetState(body RID, state PhysicsServer2DBodyState, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyApplyCentralImpulse(body RID, impulse Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyApplyTorqueImpulse(body RID, impulse float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyApplyImpulse(body RID, impulse Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyApplyCentralForce(body RID, force Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyApplyForce(body RID, force Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyApplyTorque(body RID, torque float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyAddConstantCentralForce(body RID, force Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyAddConstantForce(body RID, force Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyAddConstantTorque(body RID, torque float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetConstantForce(body RID, force Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetConstantForce(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetConstantTorque(body RID, torque float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetConstantTorque(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetAxisVelocity(body RID, axis_velocity Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyAddCollisionException(body RID, excepted_body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyRemoveCollisionException(body RID, excepted_body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetMaxContactsReported(body RID, amount int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetMaxContactsReported(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetOmitForceIntegration(body RID, enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyIsOmittingForceIntegration(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodySetForceIntegrationCallback(body RID, callable Callable, userdata Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyTestMotion(body RID, parameters PhysicsTestMotionParameters2D, result PhysicsTestMotionResult2D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) BodyGetDirectState(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) JointCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) JointClear(joint RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) JointSetParam(joint RID, param PhysicsServer2DJointParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) JointGetParam(joint RID, param PhysicsServer2DJointParam, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) JointDisableCollisionsBetweenBodies(joint RID, disable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) JointIsDisabledCollisionsBetweenBodies(joint RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) JointMakePin(joint RID, anchor Vector2, body_a RID, body_b RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) JointMakeGroove(joint RID, groove1_a Vector2, groove2_a Vector2, anchor_b Vector2, body_a RID, body_b RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) JointMakeDampedSpring(joint RID, anchor_a Vector2, anchor_b Vector2, body_a RID, body_b RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) PinJointSetParam(joint RID, param PhysicsServer2DPinJointParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) PinJointGetParam(joint RID, param PhysicsServer2DPinJointParam, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) DampedSpringJointSetParam(joint RID, param PhysicsServer2DDampedSpringParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) DampedSpringJointGetParam(joint RID, param PhysicsServer2DDampedSpringParam, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) JointGetType(joint RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) FreeRid(rid RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) SetActive(active bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2D) GetProcessInfo(process_info PhysicsServer2DProcessInfo, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
