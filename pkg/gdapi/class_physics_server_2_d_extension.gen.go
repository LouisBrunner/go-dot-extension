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



// Enums

func (me *PhysicsServer2DExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsServer2DExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer2DExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsServer2DExtension) XWorldBoundaryShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XSeparationRayShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XSegmentShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XCircleShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XRectangleShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XCapsuleShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XConvexPolygonShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XConcavePolygonShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XShapeSetData(shape RID, data Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XShapeSetCustomSolverBias(shape RID, bias float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XShapeGetType(shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XShapeGetData(shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XShapeGetCustomSolverBias(shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XShapeCollide(shape_A RID, xform_A Transform2D, motion_A Vector2, shape_B RID, xform_B Transform2D, motion_B Vector2, results unsafe.Pointer, result_max int, result_count *int32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XSpaceCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XSpaceSetActive(space RID, active bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XSpaceIsActive(space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XSpaceSetParam(space RID, param PhysicsServer2DSpaceParameter, value float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XSpaceGetParam(space RID, param PhysicsServer2DSpaceParameter, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XSpaceGetDirectState(space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XSpaceSetDebugContacts(space RID, max_contacts int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XSpaceGetContacts(space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XSpaceGetContactCount(space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaSetSpace(area RID, space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaGetSpace(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaAddShape(area RID, shape RID, transform Transform2D, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaSetShape(area RID, shape_idx int, shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaSetShapeTransform(area RID, shape_idx int, transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaSetShapeDisabled(area RID, shape_idx int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaGetShapeCount(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaGetShape(area RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaGetShapeTransform(area RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaRemoveShape(area RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaClearShapes(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaAttachObjectInstanceId(area RID, id int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaGetObjectInstanceId(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaAttachCanvasInstanceId(area RID, id int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaGetCanvasInstanceId(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaSetParam(area RID, param PhysicsServer2DAreaParameter, value Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaSetTransform(area RID, transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaGetParam(area RID, param PhysicsServer2DAreaParameter, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaGetTransform(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaSetCollisionLayer(area RID, layer int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaGetCollisionLayer(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaSetCollisionMask(area RID, mask int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaGetCollisionMask(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaSetMonitorable(area RID, monitorable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaSetPickable(area RID, pickable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaSetMonitorCallback(area RID, callback Callable, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XAreaSetAreaMonitorCallback(area RID, callback Callable, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetSpace(body RID, space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetSpace(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetMode(body RID, mode PhysicsServer2DBodyMode, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetMode(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyAddShape(body RID, shape RID, transform Transform2D, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetShape(body RID, shape_idx int, shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetShapeTransform(body RID, shape_idx int, transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetShapeCount(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetShape(body RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetShapeTransform(body RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetShapeDisabled(body RID, shape_idx int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetShapeAsOneWayCollision(body RID, shape_idx int, enable bool, margin float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyRemoveShape(body RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyClearShapes(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyAttachObjectInstanceId(body RID, id int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetObjectInstanceId(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyAttachCanvasInstanceId(body RID, id int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetCanvasInstanceId(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetContinuousCollisionDetectionMode(body RID, mode PhysicsServer2DCCDMode, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetContinuousCollisionDetectionMode(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetCollisionLayer(body RID, layer int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetCollisionLayer(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetCollisionMask(body RID, mask int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetCollisionMask(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetCollisionPriority(body RID, priority float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetCollisionPriority(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetParam(body RID, param PhysicsServer2DBodyParameter, value Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetParam(body RID, param PhysicsServer2DBodyParameter, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyResetMassProperties(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetState(body RID, state PhysicsServer2DBodyState, value Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetState(body RID, state PhysicsServer2DBodyState, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyApplyCentralImpulse(body RID, impulse Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyApplyTorqueImpulse(body RID, impulse float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyApplyImpulse(body RID, impulse Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyApplyCentralForce(body RID, force Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyApplyForce(body RID, force Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyApplyTorque(body RID, torque float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyAddConstantCentralForce(body RID, force Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyAddConstantForce(body RID, force Vector2, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyAddConstantTorque(body RID, torque float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetConstantForce(body RID, force Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetConstantForce(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetConstantTorque(body RID, torque float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetConstantTorque(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetAxisVelocity(body RID, axis_velocity Vector2, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyAddCollisionException(body RID, excepted_body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyRemoveCollisionException(body RID, excepted_body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetCollisionExceptions(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetMaxContactsReported(body RID, amount int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetMaxContactsReported(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetContactsReportedDepthThreshold(body RID, threshold float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetContactsReportedDepthThreshold(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetOmitForceIntegration(body RID, enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyIsOmittingForceIntegration(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetStateSyncCallback(body RID, callable Callable, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetForceIntegrationCallback(body RID, callable Callable, userdata Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyCollideShape(body RID, body_shape int, shape RID, shape_xform Transform2D, motion Vector2, results unsafe.Pointer, result_max int, result_count *int32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodySetPickable(body RID, pickable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyGetDirectState(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XBodyTestMotion(body RID, from Transform2D, motion Vector2, margin float32, collide_separation_ray bool, recovery_as_collision bool, result *PhysicsServer2DExtensionMotionResult, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XJointCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XJointClear(joint RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XJointSetParam(joint RID, param PhysicsServer2DJointParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XJointGetParam(joint RID, param PhysicsServer2DJointParam, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XJointDisableCollisionsBetweenBodies(joint RID, disable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XJointIsDisabledCollisionsBetweenBodies(joint RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XJointMakePin(joint RID, anchor Vector2, body_a RID, body_b RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XJointMakeGroove(joint RID, a_groove1 Vector2, a_groove2 Vector2, b_anchor Vector2, body_a RID, body_b RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XJointMakeDampedSpring(joint RID, anchor_a Vector2, anchor_b Vector2, body_a RID, body_b RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XPinJointSetParam(joint RID, param PhysicsServer2DPinJointParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XPinJointGetParam(joint RID, param PhysicsServer2DPinJointParam, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XDampedSpringJointSetParam(joint RID, param PhysicsServer2DDampedSpringParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XDampedSpringJointGetParam(joint RID, param PhysicsServer2DDampedSpringParam, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XJointGetType(joint RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XFreeRid(rid RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XSetActive(active bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XInit()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XStep(step float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XSync()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XFlushQueries()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XEndSync()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XFinish()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XIsFlushingQueries()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) XGetProcessInfo(process_info PhysicsServer2DProcessInfo, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) BodyTestMotionIsExcludingBody(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer2DExtension) BodyTestMotionIsExcludingObject(object int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
