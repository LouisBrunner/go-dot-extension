// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports


  "unsafe"




  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer2DExtension struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer2DExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer2DExtension) BaseClass() string {
  return "PhysicsServer2DExtension"
}

func  (me *PhysicsServer2DExtension) XWorldBoundaryShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XSeparationRayShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XSegmentShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XCircleShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XRectangleShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XCapsuleShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XConvexPolygonShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XConcavePolygonShapeCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XShapeSetData(shape RID, data Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XShapeSetCustomSolverBias(shape RID, bias float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XShapeGetType(shape RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XShapeGetData(shape RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XShapeGetCustomSolverBias(shape RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XShapeCollide(shape_A RID, xform_A Transform2D, motion_A Vector2, shape_B RID, xform_B Transform2D, motion_B Vector2, results unsafe.Pointer, result_max int, result_count *int32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XSpaceCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XSpaceSetActive(space RID, active bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XSpaceIsActive(space RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XSpaceSetParam(space RID, param PhysicsServer2DSpaceParameter, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XSpaceGetParam(space RID, param PhysicsServer2DSpaceParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XSpaceGetDirectState(space RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XSpaceSetDebugContacts(space RID, max_contacts int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XSpaceGetContacts(space RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XSpaceGetContactCount(space RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaSetSpace(area RID, space RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaGetSpace(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaAddShape(area RID, shape RID, transform Transform2D, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaSetShape(area RID, shape_idx int, shape RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaSetShapeTransform(area RID, shape_idx int, transform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaSetShapeDisabled(area RID, shape_idx int, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaGetShapeCount(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaGetShape(area RID, shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaGetShapeTransform(area RID, shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaRemoveShape(area RID, shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaClearShapes(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaAttachObjectInstanceId(area RID, id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaGetObjectInstanceId(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaAttachCanvasInstanceId(area RID, id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaGetCanvasInstanceId(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaSetParam(area RID, param PhysicsServer2DAreaParameter, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaSetTransform(area RID, transform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaGetParam(area RID, param PhysicsServer2DAreaParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaGetTransform(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaSetCollisionLayer(area RID, layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaGetCollisionLayer(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaSetCollisionMask(area RID, mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaGetCollisionMask(area RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaSetMonitorable(area RID, monitorable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaSetPickable(area RID, pickable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaSetMonitorCallback(area RID, callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XAreaSetAreaMonitorCallback(area RID, callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetSpace(body RID, space RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetSpace(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetMode(body RID, mode PhysicsServer2DBodyMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetMode(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyAddShape(body RID, shape RID, transform Transform2D, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetShape(body RID, shape_idx int, shape RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetShapeTransform(body RID, shape_idx int, transform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetShapeCount(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetShape(body RID, shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetShapeTransform(body RID, shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetShapeDisabled(body RID, shape_idx int, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetShapeAsOneWayCollision(body RID, shape_idx int, enable bool, margin float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyRemoveShape(body RID, shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyClearShapes(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyAttachObjectInstanceId(body RID, id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetObjectInstanceId(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyAttachCanvasInstanceId(body RID, id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetCanvasInstanceId(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetContinuousCollisionDetectionMode(body RID, mode PhysicsServer2DCCDMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetContinuousCollisionDetectionMode(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetCollisionLayer(body RID, layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetCollisionLayer(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetCollisionMask(body RID, mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetCollisionMask(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetCollisionPriority(body RID, priority float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetCollisionPriority(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetParam(body RID, param PhysicsServer2DBodyParameter, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetParam(body RID, param PhysicsServer2DBodyParameter, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyResetMassProperties(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetState(body RID, state PhysicsServer2DBodyState, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetState(body RID, state PhysicsServer2DBodyState, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyApplyCentralImpulse(body RID, impulse Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyApplyTorqueImpulse(body RID, impulse float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyApplyImpulse(body RID, impulse Vector2, position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyApplyCentralForce(body RID, force Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyApplyForce(body RID, force Vector2, position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyApplyTorque(body RID, torque float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyAddConstantCentralForce(body RID, force Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyAddConstantForce(body RID, force Vector2, position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyAddConstantTorque(body RID, torque float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetConstantForce(body RID, force Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetConstantForce(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetConstantTorque(body RID, torque float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetConstantTorque(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetAxisVelocity(body RID, axis_velocity Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyAddCollisionException(body RID, excepted_body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyRemoveCollisionException(body RID, excepted_body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetCollisionExceptions(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetMaxContactsReported(body RID, amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetMaxContactsReported(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetContactsReportedDepthThreshold(body RID, threshold float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetContactsReportedDepthThreshold(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetOmitForceIntegration(body RID, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyIsOmittingForceIntegration(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetStateSyncCallback(body RID, callable Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetForceIntegrationCallback(body RID, callable Callable, userdata Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyCollideShape(body RID, body_shape int, shape RID, shape_xform Transform2D, motion Vector2, results unsafe.Pointer, result_max int, result_count *int32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodySetPickable(body RID, pickable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyGetDirectState(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XBodyTestMotion(body RID, from Transform2D, motion Vector2, margin float32, collide_separation_ray bool, recovery_as_collision bool, result *PhysicsServer2DExtensionMotionResult, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XJointCreate() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XJointClear(joint RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XJointSetParam(joint RID, param PhysicsServer2DJointParam, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XJointGetParam(joint RID, param PhysicsServer2DJointParam, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XJointDisableCollisionsBetweenBodies(joint RID, disable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XJointIsDisabledCollisionsBetweenBodies(joint RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XJointMakePin(joint RID, anchor Vector2, body_a RID, body_b RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XJointMakeGroove(joint RID, a_groove1 Vector2, a_groove2 Vector2, b_anchor Vector2, body_a RID, body_b RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XJointMakeDampedSpring(joint RID, anchor_a Vector2, anchor_b Vector2, body_a RID, body_b RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XPinJointSetParam(joint RID, param PhysicsServer2DPinJointParam, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XPinJointGetParam(joint RID, param PhysicsServer2DPinJointParam, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XDampedSpringJointSetParam(joint RID, param PhysicsServer2DDampedSpringParam, value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XDampedSpringJointGetParam(joint RID, param PhysicsServer2DDampedSpringParam, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XJointGetType(joint RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XFreeRid(rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XSetActive(active bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XInit() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XStep(step float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XSync() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XFlushQueries() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XEndSync() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XFinish() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XIsFlushingQueries() { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) XGetProcessInfo(process_info PhysicsServer2DProcessInfo, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) BodyTestMotionIsExcludingBody(body RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *PhysicsServer2DExtension) BodyTestMotionIsExcludingObject(object int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
