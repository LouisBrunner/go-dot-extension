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

type ptrsForPhysicsServer3DExtensionList struct {
	fnXWorldBoundaryShapeCreate                  gdc.MethodBindPtr
	fnXSeparationRayShapeCreate                  gdc.MethodBindPtr
	fnXSphereShapeCreate                         gdc.MethodBindPtr
	fnXBoxShapeCreate                            gdc.MethodBindPtr
	fnXCapsuleShapeCreate                        gdc.MethodBindPtr
	fnXCylinderShapeCreate                       gdc.MethodBindPtr
	fnXConvexPolygonShapeCreate                  gdc.MethodBindPtr
	fnXConcavePolygonShapeCreate                 gdc.MethodBindPtr
	fnXHeightmapShapeCreate                      gdc.MethodBindPtr
	fnXCustomShapeCreate                         gdc.MethodBindPtr
	fnXShapeSetData                              gdc.MethodBindPtr
	fnXShapeSetCustomSolverBias                  gdc.MethodBindPtr
	fnXShapeSetMargin                            gdc.MethodBindPtr
	fnXShapeGetMargin                            gdc.MethodBindPtr
	fnXShapeGetType                              gdc.MethodBindPtr
	fnXShapeGetData                              gdc.MethodBindPtr
	fnXShapeGetCustomSolverBias                  gdc.MethodBindPtr
	fnXSpaceCreate                               gdc.MethodBindPtr
	fnXSpaceSetActive                            gdc.MethodBindPtr
	fnXSpaceIsActive                             gdc.MethodBindPtr
	fnXSpaceSetParam                             gdc.MethodBindPtr
	fnXSpaceGetParam                             gdc.MethodBindPtr
	fnXSpaceGetDirectState                       gdc.MethodBindPtr
	fnXSpaceSetDebugContacts                     gdc.MethodBindPtr
	fnXSpaceGetContacts                          gdc.MethodBindPtr
	fnXSpaceGetContactCount                      gdc.MethodBindPtr
	fnXAreaCreate                                gdc.MethodBindPtr
	fnXAreaSetSpace                              gdc.MethodBindPtr
	fnXAreaGetSpace                              gdc.MethodBindPtr
	fnXAreaAddShape                              gdc.MethodBindPtr
	fnXAreaSetShape                              gdc.MethodBindPtr
	fnXAreaSetShapeTransform                     gdc.MethodBindPtr
	fnXAreaSetShapeDisabled                      gdc.MethodBindPtr
	fnXAreaGetShapeCount                         gdc.MethodBindPtr
	fnXAreaGetShape                              gdc.MethodBindPtr
	fnXAreaGetShapeTransform                     gdc.MethodBindPtr
	fnXAreaRemoveShape                           gdc.MethodBindPtr
	fnXAreaClearShapes                           gdc.MethodBindPtr
	fnXAreaAttachObjectInstanceId                gdc.MethodBindPtr
	fnXAreaGetObjectInstanceId                   gdc.MethodBindPtr
	fnXAreaSetParam                              gdc.MethodBindPtr
	fnXAreaSetTransform                          gdc.MethodBindPtr
	fnXAreaGetParam                              gdc.MethodBindPtr
	fnXAreaGetTransform                          gdc.MethodBindPtr
	fnXAreaSetCollisionLayer                     gdc.MethodBindPtr
	fnXAreaGetCollisionLayer                     gdc.MethodBindPtr
	fnXAreaSetCollisionMask                      gdc.MethodBindPtr
	fnXAreaGetCollisionMask                      gdc.MethodBindPtr
	fnXAreaSetMonitorable                        gdc.MethodBindPtr
	fnXAreaSetRayPickable                        gdc.MethodBindPtr
	fnXAreaSetMonitorCallback                    gdc.MethodBindPtr
	fnXAreaSetAreaMonitorCallback                gdc.MethodBindPtr
	fnXBodyCreate                                gdc.MethodBindPtr
	fnXBodySetSpace                              gdc.MethodBindPtr
	fnXBodyGetSpace                              gdc.MethodBindPtr
	fnXBodySetMode                               gdc.MethodBindPtr
	fnXBodyGetMode                               gdc.MethodBindPtr
	fnXBodyAddShape                              gdc.MethodBindPtr
	fnXBodySetShape                              gdc.MethodBindPtr
	fnXBodySetShapeTransform                     gdc.MethodBindPtr
	fnXBodySetShapeDisabled                      gdc.MethodBindPtr
	fnXBodyGetShapeCount                         gdc.MethodBindPtr
	fnXBodyGetShape                              gdc.MethodBindPtr
	fnXBodyGetShapeTransform                     gdc.MethodBindPtr
	fnXBodyRemoveShape                           gdc.MethodBindPtr
	fnXBodyClearShapes                           gdc.MethodBindPtr
	fnXBodyAttachObjectInstanceId                gdc.MethodBindPtr
	fnXBodyGetObjectInstanceId                   gdc.MethodBindPtr
	fnXBodySetEnableContinuousCollisionDetection gdc.MethodBindPtr
	fnXBodyIsContinuousCollisionDetectionEnabled gdc.MethodBindPtr
	fnXBodySetCollisionLayer                     gdc.MethodBindPtr
	fnXBodyGetCollisionLayer                     gdc.MethodBindPtr
	fnXBodySetCollisionMask                      gdc.MethodBindPtr
	fnXBodyGetCollisionMask                      gdc.MethodBindPtr
	fnXBodySetCollisionPriority                  gdc.MethodBindPtr
	fnXBodyGetCollisionPriority                  gdc.MethodBindPtr
	fnXBodySetUserFlags                          gdc.MethodBindPtr
	fnXBodyGetUserFlags                          gdc.MethodBindPtr
	fnXBodySetParam                              gdc.MethodBindPtr
	fnXBodyGetParam                              gdc.MethodBindPtr
	fnXBodyResetMassProperties                   gdc.MethodBindPtr
	fnXBodySetState                              gdc.MethodBindPtr
	fnXBodyGetState                              gdc.MethodBindPtr
	fnXBodyApplyCentralImpulse                   gdc.MethodBindPtr
	fnXBodyApplyImpulse                          gdc.MethodBindPtr
	fnXBodyApplyTorqueImpulse                    gdc.MethodBindPtr
	fnXBodyApplyCentralForce                     gdc.MethodBindPtr
	fnXBodyApplyForce                            gdc.MethodBindPtr
	fnXBodyApplyTorque                           gdc.MethodBindPtr
	fnXBodyAddConstantCentralForce               gdc.MethodBindPtr
	fnXBodyAddConstantForce                      gdc.MethodBindPtr
	fnXBodyAddConstantTorque                     gdc.MethodBindPtr
	fnXBodySetConstantForce                      gdc.MethodBindPtr
	fnXBodyGetConstantForce                      gdc.MethodBindPtr
	fnXBodySetConstantTorque                     gdc.MethodBindPtr
	fnXBodyGetConstantTorque                     gdc.MethodBindPtr
	fnXBodySetAxisVelocity                       gdc.MethodBindPtr
	fnXBodySetAxisLock                           gdc.MethodBindPtr
	fnXBodyIsAxisLocked                          gdc.MethodBindPtr
	fnXBodyAddCollisionException                 gdc.MethodBindPtr
	fnXBodyRemoveCollisionException              gdc.MethodBindPtr
	fnXBodyGetCollisionExceptions                gdc.MethodBindPtr
	fnXBodySetMaxContactsReported                gdc.MethodBindPtr
	fnXBodyGetMaxContactsReported                gdc.MethodBindPtr
	fnXBodySetContactsReportedDepthThreshold     gdc.MethodBindPtr
	fnXBodyGetContactsReportedDepthThreshold     gdc.MethodBindPtr
	fnXBodySetOmitForceIntegration               gdc.MethodBindPtr
	fnXBodyIsOmittingForceIntegration            gdc.MethodBindPtr
	fnXBodySetStateSyncCallback                  gdc.MethodBindPtr
	fnXBodySetForceIntegrationCallback           gdc.MethodBindPtr
	fnXBodySetRayPickable                        gdc.MethodBindPtr
	fnXBodyTestMotion                            gdc.MethodBindPtr
	fnXBodyGetDirectState                        gdc.MethodBindPtr
	fnXSoftBodyCreate                            gdc.MethodBindPtr
	fnXSoftBodyUpdateRenderingServer             gdc.MethodBindPtr
	fnXSoftBodySetSpace                          gdc.MethodBindPtr
	fnXSoftBodyGetSpace                          gdc.MethodBindPtr
	fnXSoftBodySetRayPickable                    gdc.MethodBindPtr
	fnXSoftBodySetCollisionLayer                 gdc.MethodBindPtr
	fnXSoftBodyGetCollisionLayer                 gdc.MethodBindPtr
	fnXSoftBodySetCollisionMask                  gdc.MethodBindPtr
	fnXSoftBodyGetCollisionMask                  gdc.MethodBindPtr
	fnXSoftBodyAddCollisionException             gdc.MethodBindPtr
	fnXSoftBodyRemoveCollisionException          gdc.MethodBindPtr
	fnXSoftBodyGetCollisionExceptions            gdc.MethodBindPtr
	fnXSoftBodySetState                          gdc.MethodBindPtr
	fnXSoftBodyGetState                          gdc.MethodBindPtr
	fnXSoftBodySetTransform                      gdc.MethodBindPtr
	fnXSoftBodySetSimulationPrecision            gdc.MethodBindPtr
	fnXSoftBodyGetSimulationPrecision            gdc.MethodBindPtr
	fnXSoftBodySetTotalMass                      gdc.MethodBindPtr
	fnXSoftBodyGetTotalMass                      gdc.MethodBindPtr
	fnXSoftBodySetLinearStiffness                gdc.MethodBindPtr
	fnXSoftBodyGetLinearStiffness                gdc.MethodBindPtr
	fnXSoftBodySetPressureCoefficient            gdc.MethodBindPtr
	fnXSoftBodyGetPressureCoefficient            gdc.MethodBindPtr
	fnXSoftBodySetDampingCoefficient             gdc.MethodBindPtr
	fnXSoftBodyGetDampingCoefficient             gdc.MethodBindPtr
	fnXSoftBodySetDragCoefficient                gdc.MethodBindPtr
	fnXSoftBodyGetDragCoefficient                gdc.MethodBindPtr
	fnXSoftBodySetMesh                           gdc.MethodBindPtr
	fnXSoftBodyGetBounds                         gdc.MethodBindPtr
	fnXSoftBodyMovePoint                         gdc.MethodBindPtr
	fnXSoftBodyGetPointGlobalPosition            gdc.MethodBindPtr
	fnXSoftBodyRemoveAllPinnedPoints             gdc.MethodBindPtr
	fnXSoftBodyPinPoint                          gdc.MethodBindPtr
	fnXSoftBodyIsPointPinned                     gdc.MethodBindPtr
	fnXJointCreate                               gdc.MethodBindPtr
	fnXJointClear                                gdc.MethodBindPtr
	fnXJointMakePin                              gdc.MethodBindPtr
	fnXPinJointSetParam                          gdc.MethodBindPtr
	fnXPinJointGetParam                          gdc.MethodBindPtr
	fnXPinJointSetLocalA                         gdc.MethodBindPtr
	fnXPinJointGetLocalA                         gdc.MethodBindPtr
	fnXPinJointSetLocalB                         gdc.MethodBindPtr
	fnXPinJointGetLocalB                         gdc.MethodBindPtr
	fnXJointMakeHinge                            gdc.MethodBindPtr
	fnXJointMakeHingeSimple                      gdc.MethodBindPtr
	fnXHingeJointSetParam                        gdc.MethodBindPtr
	fnXHingeJointGetParam                        gdc.MethodBindPtr
	fnXHingeJointSetFlag                         gdc.MethodBindPtr
	fnXHingeJointGetFlag                         gdc.MethodBindPtr
	fnXJointMakeSlider                           gdc.MethodBindPtr
	fnXSliderJointSetParam                       gdc.MethodBindPtr
	fnXSliderJointGetParam                       gdc.MethodBindPtr
	fnXJointMakeConeTwist                        gdc.MethodBindPtr
	fnXConeTwistJointSetParam                    gdc.MethodBindPtr
	fnXConeTwistJointGetParam                    gdc.MethodBindPtr
	fnXJointMakeGeneric6Dof                      gdc.MethodBindPtr
	fnXGeneric6DofJointSetParam                  gdc.MethodBindPtr
	fnXGeneric6DofJointGetParam                  gdc.MethodBindPtr
	fnXGeneric6DofJointSetFlag                   gdc.MethodBindPtr
	fnXGeneric6DofJointGetFlag                   gdc.MethodBindPtr
	fnXJointGetType                              gdc.MethodBindPtr
	fnXJointSetSolverPriority                    gdc.MethodBindPtr
	fnXJointGetSolverPriority                    gdc.MethodBindPtr
	fnXJointDisableCollisionsBetweenBodies       gdc.MethodBindPtr
	fnXJointIsDisabledCollisionsBetweenBodies    gdc.MethodBindPtr
	fnXFreeRid                                   gdc.MethodBindPtr
	fnXSetActive                                 gdc.MethodBindPtr
	fnXInit                                      gdc.MethodBindPtr
	fnXStep                                      gdc.MethodBindPtr
	fnXSync                                      gdc.MethodBindPtr
	fnXFlushQueries                              gdc.MethodBindPtr
	fnXEndSync                                   gdc.MethodBindPtr
	fnXFinish                                    gdc.MethodBindPtr
	fnXIsFlushingQueries                         gdc.MethodBindPtr
	fnXGetProcessInfo                            gdc.MethodBindPtr
	fnBodyTestMotionIsExcludingBody              gdc.MethodBindPtr
	fnBodyTestMotionIsExcludingObject            gdc.MethodBindPtr
}

var ptrsForPhysicsServer3DExtension ptrsForPhysicsServer3DExtensionList

func initPhysicsServer3DExtensionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsServer3DExtension")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("body_test_motion_is_excluding_body")
		defer methodName.Destroy()
		ptrsForPhysicsServer3DExtension.fnBodyTestMotionIsExcludingBody = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("body_test_motion_is_excluding_object")
		defer methodName.Destroy()
		ptrsForPhysicsServer3DExtension.fnBodyTestMotionIsExcludingObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
}

type PhysicsServer3DExtension struct {
	PhysicsServer3D
}

func (me *PhysicsServer3DExtension) BaseClass() string {
	return "PhysicsServer3DExtension"
}

func NewPhysicsServer3DExtension() *PhysicsServer3DExtension {
	str := StringNameFromStr("PhysicsServer3DExtension") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsServer3DExtension{}
	obj.SetBaseObject(objPtr)
	return obj
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

func (me *PhysicsServer3DExtension) BodyTestMotionIsExcludingBody(body RID) bool {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3DExtension.fnBodyTestMotionIsExcludingBody), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3DExtension) BodyTestMotionIsExcludingObject(object int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&object)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&object)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3DExtension.fnBodyTestMotionIsExcludingObject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
