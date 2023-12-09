// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer3DExtension struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer3DExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer3DExtension) BaseClass() string {
  return "PhysicsServer3DExtension"
}



// Enums

func (me *PhysicsServer3DExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsServer3DExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer3DExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsServer3DExtension) XWorldBoundaryShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSeparationRayShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSphereShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBoxShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XCapsuleShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XCylinderShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XConvexPolygonShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XConcavePolygonShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XHeightmapShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XCustomShapeCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XShapeSetData(shape RID, data Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XShapeSetCustomSolverBias(shape RID, bias float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XShapeSetMargin(shape RID, margin float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XShapeGetMargin(shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XShapeGetType(shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XShapeGetData(shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XShapeGetCustomSolverBias(shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSpaceCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSpaceSetActive(space RID, active bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSpaceIsActive(space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSpaceSetParam(space RID, param PhysicsServer3DSpaceParameter, value float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSpaceGetParam(space RID, param PhysicsServer3DSpaceParameter, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSpaceGetDirectState(space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSpaceSetDebugContacts(space RID, max_contacts int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSpaceGetContacts(space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSpaceGetContactCount(space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaSetSpace(area RID, space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaGetSpace(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaAddShape(area RID, shape RID, transform Transform3D, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaSetShape(area RID, shape_idx int, shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaSetShapeTransform(area RID, shape_idx int, transform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaSetShapeDisabled(area RID, shape_idx int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaGetShapeCount(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaGetShape(area RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaGetShapeTransform(area RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaRemoveShape(area RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaClearShapes(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaAttachObjectInstanceId(area RID, id int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaGetObjectInstanceId(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaSetParam(area RID, param PhysicsServer3DAreaParameter, value Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaSetTransform(area RID, transform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaGetParam(area RID, param PhysicsServer3DAreaParameter, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaGetTransform(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaSetCollisionLayer(area RID, layer int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaGetCollisionLayer(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaSetCollisionMask(area RID, mask int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaGetCollisionMask(area RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaSetMonitorable(area RID, monitorable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaSetRayPickable(area RID, enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaSetMonitorCallback(area RID, callback Callable, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XAreaSetAreaMonitorCallback(area RID, callback Callable, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetSpace(body RID, space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetSpace(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetMode(body RID, mode PhysicsServer3DBodyMode, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetMode(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyAddShape(body RID, shape RID, transform Transform3D, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetShape(body RID, shape_idx int, shape RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetShapeTransform(body RID, shape_idx int, transform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetShapeDisabled(body RID, shape_idx int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetShapeCount(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetShape(body RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetShapeTransform(body RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyRemoveShape(body RID, shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyClearShapes(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyAttachObjectInstanceId(body RID, id int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetObjectInstanceId(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetEnableContinuousCollisionDetection(body RID, enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyIsContinuousCollisionDetectionEnabled(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetCollisionLayer(body RID, layer int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetCollisionLayer(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetCollisionMask(body RID, mask int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetCollisionMask(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetCollisionPriority(body RID, priority float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetCollisionPriority(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetUserFlags(body RID, flags int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetUserFlags(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetParam(body RID, param PhysicsServer3DBodyParameter, value Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetParam(body RID, param PhysicsServer3DBodyParameter, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyResetMassProperties(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetState(body RID, state PhysicsServer3DBodyState, value Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetState(body RID, state PhysicsServer3DBodyState, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyApplyCentralImpulse(body RID, impulse Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyApplyImpulse(body RID, impulse Vector3, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyApplyTorqueImpulse(body RID, impulse Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyApplyCentralForce(body RID, force Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyApplyForce(body RID, force Vector3, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyApplyTorque(body RID, torque Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyAddConstantCentralForce(body RID, force Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyAddConstantForce(body RID, force Vector3, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyAddConstantTorque(body RID, torque Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetConstantForce(body RID, force Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetConstantForce(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetConstantTorque(body RID, torque Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetConstantTorque(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetAxisVelocity(body RID, axis_velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetAxisLock(body RID, axis PhysicsServer3DBodyAxis, lock bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyIsAxisLocked(body RID, axis PhysicsServer3DBodyAxis, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyAddCollisionException(body RID, excepted_body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyRemoveCollisionException(body RID, excepted_body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetCollisionExceptions(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetMaxContactsReported(body RID, amount int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetMaxContactsReported(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetContactsReportedDepthThreshold(body RID, threshold float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetContactsReportedDepthThreshold(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetOmitForceIntegration(body RID, enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyIsOmittingForceIntegration(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetStateSyncCallback(body RID, callable Callable, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetForceIntegrationCallback(body RID, callable Callable, userdata Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodySetRayPickable(body RID, enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyTestMotion(body RID, from Transform3D, motion Vector3, margin float32, max_collisions int, collide_separation_ray bool, recovery_as_collision bool, result *PhysicsServer3DExtensionMotionResult, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XBodyGetDirectState(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyUpdateRenderingServer(body RID, rendering_server_handler PhysicsServer3DRenderingServerHandler, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodySetSpace(body RID, space RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyGetSpace(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodySetRayPickable(body RID, enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodySetCollisionLayer(body RID, layer int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyGetCollisionLayer(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodySetCollisionMask(body RID, mask int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyGetCollisionMask(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyAddCollisionException(body RID, body_b RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyRemoveCollisionException(body RID, body_b RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyGetCollisionExceptions(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodySetState(body RID, state PhysicsServer3DBodyState, variant Variant, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyGetState(body RID, state PhysicsServer3DBodyState, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodySetTransform(body RID, transform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodySetSimulationPrecision(body RID, simulation_precision int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyGetSimulationPrecision(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodySetTotalMass(body RID, total_mass float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyGetTotalMass(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodySetLinearStiffness(body RID, linear_stiffness float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyGetLinearStiffness(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodySetPressureCoefficient(body RID, pressure_coefficient float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyGetPressureCoefficient(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodySetDampingCoefficient(body RID, damping_coefficient float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyGetDampingCoefficient(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodySetDragCoefficient(body RID, drag_coefficient float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyGetDragCoefficient(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodySetMesh(body RID, mesh RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyGetBounds(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyMovePoint(body RID, point_index int, global_position Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyGetPointGlobalPosition(body RID, point_index int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyRemoveAllPinnedPoints(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyPinPoint(body RID, point_index int, pin bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSoftBodyIsPointPinned(body RID, point_index int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XJointCreate()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XJointClear(joint RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XJointMakePin(joint RID, body_A RID, local_A Vector3, body_B RID, local_B Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XPinJointSetParam(joint RID, param PhysicsServer3DPinJointParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XPinJointGetParam(joint RID, param PhysicsServer3DPinJointParam, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XPinJointSetLocalA(joint RID, local_A Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XPinJointGetLocalA(joint RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XPinJointSetLocalB(joint RID, local_B Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XPinJointGetLocalB(joint RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XJointMakeHinge(joint RID, body_A RID, hinge_A Transform3D, body_B RID, hinge_B Transform3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XJointMakeHingeSimple(joint RID, body_A RID, pivot_A Vector3, axis_A Vector3, body_B RID, pivot_B Vector3, axis_B Vector3, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XHingeJointSetParam(joint RID, param PhysicsServer3DHingeJointParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XHingeJointGetParam(joint RID, param PhysicsServer3DHingeJointParam, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XHingeJointSetFlag(joint RID, flag PhysicsServer3DHingeJointFlag, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XHingeJointGetFlag(joint RID, flag PhysicsServer3DHingeJointFlag, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XJointMakeSlider(joint RID, body_A RID, local_ref_A Transform3D, body_B RID, local_ref_B Transform3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSliderJointSetParam(joint RID, param PhysicsServer3DSliderJointParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSliderJointGetParam(joint RID, param PhysicsServer3DSliderJointParam, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XJointMakeConeTwist(joint RID, body_A RID, local_ref_A Transform3D, body_B RID, local_ref_B Transform3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XConeTwistJointSetParam(joint RID, param PhysicsServer3DConeTwistJointParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XConeTwistJointGetParam(joint RID, param PhysicsServer3DConeTwistJointParam, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XJointMakeGeneric6Dof(joint RID, body_A RID, local_ref_A Transform3D, body_B RID, local_ref_B Transform3D, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XGeneric6DofJointSetParam(joint RID, axis Vector3Axis, param PhysicsServer3DG6DOFJointAxisParam, value float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XGeneric6DofJointGetParam(joint RID, axis Vector3Axis, param PhysicsServer3DG6DOFJointAxisParam, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XGeneric6DofJointSetFlag(joint RID, axis Vector3Axis, flag PhysicsServer3DG6DOFJointAxisFlag, enable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XGeneric6DofJointGetFlag(joint RID, axis Vector3Axis, flag PhysicsServer3DG6DOFJointAxisFlag, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XJointGetType(joint RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XJointSetSolverPriority(joint RID, priority int, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XJointGetSolverPriority(joint RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XJointDisableCollisionsBetweenBodies(joint RID, disable bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XJointIsDisabledCollisionsBetweenBodies(joint RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XFreeRid(rid RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSetActive(active bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XInit()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XStep(step float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XSync()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XFlushQueries()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XEndSync()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XFinish()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XIsFlushingQueries()  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) XGetProcessInfo(process_info PhysicsServer3DProcessInfo, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) BodyTestMotionIsExcludingBody(body RID, )  {
  panic("TODO: implement")
}

func  (me *PhysicsServer3DExtension) BodyTestMotionIsExcludingObject(object int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
