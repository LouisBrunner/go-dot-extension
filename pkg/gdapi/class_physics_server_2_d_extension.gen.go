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

type ptrsForPhysicsServer2DExtensionList struct {
	fnXWorldBoundaryShapeCreate                gdc.MethodBindPtr
	fnXSeparationRayShapeCreate                gdc.MethodBindPtr
	fnXSegmentShapeCreate                      gdc.MethodBindPtr
	fnXCircleShapeCreate                       gdc.MethodBindPtr
	fnXRectangleShapeCreate                    gdc.MethodBindPtr
	fnXCapsuleShapeCreate                      gdc.MethodBindPtr
	fnXConvexPolygonShapeCreate                gdc.MethodBindPtr
	fnXConcavePolygonShapeCreate               gdc.MethodBindPtr
	fnXShapeSetData                            gdc.MethodBindPtr
	fnXShapeSetCustomSolverBias                gdc.MethodBindPtr
	fnXShapeGetType                            gdc.MethodBindPtr
	fnXShapeGetData                            gdc.MethodBindPtr
	fnXShapeGetCustomSolverBias                gdc.MethodBindPtr
	fnXShapeCollide                            gdc.MethodBindPtr
	fnXSpaceCreate                             gdc.MethodBindPtr
	fnXSpaceSetActive                          gdc.MethodBindPtr
	fnXSpaceIsActive                           gdc.MethodBindPtr
	fnXSpaceSetParam                           gdc.MethodBindPtr
	fnXSpaceGetParam                           gdc.MethodBindPtr
	fnXSpaceGetDirectState                     gdc.MethodBindPtr
	fnXSpaceSetDebugContacts                   gdc.MethodBindPtr
	fnXSpaceGetContacts                        gdc.MethodBindPtr
	fnXSpaceGetContactCount                    gdc.MethodBindPtr
	fnXAreaCreate                              gdc.MethodBindPtr
	fnXAreaSetSpace                            gdc.MethodBindPtr
	fnXAreaGetSpace                            gdc.MethodBindPtr
	fnXAreaAddShape                            gdc.MethodBindPtr
	fnXAreaSetShape                            gdc.MethodBindPtr
	fnXAreaSetShapeTransform                   gdc.MethodBindPtr
	fnXAreaSetShapeDisabled                    gdc.MethodBindPtr
	fnXAreaGetShapeCount                       gdc.MethodBindPtr
	fnXAreaGetShape                            gdc.MethodBindPtr
	fnXAreaGetShapeTransform                   gdc.MethodBindPtr
	fnXAreaRemoveShape                         gdc.MethodBindPtr
	fnXAreaClearShapes                         gdc.MethodBindPtr
	fnXAreaAttachObjectInstanceId              gdc.MethodBindPtr
	fnXAreaGetObjectInstanceId                 gdc.MethodBindPtr
	fnXAreaAttachCanvasInstanceId              gdc.MethodBindPtr
	fnXAreaGetCanvasInstanceId                 gdc.MethodBindPtr
	fnXAreaSetParam                            gdc.MethodBindPtr
	fnXAreaSetTransform                        gdc.MethodBindPtr
	fnXAreaGetParam                            gdc.MethodBindPtr
	fnXAreaGetTransform                        gdc.MethodBindPtr
	fnXAreaSetCollisionLayer                   gdc.MethodBindPtr
	fnXAreaGetCollisionLayer                   gdc.MethodBindPtr
	fnXAreaSetCollisionMask                    gdc.MethodBindPtr
	fnXAreaGetCollisionMask                    gdc.MethodBindPtr
	fnXAreaSetMonitorable                      gdc.MethodBindPtr
	fnXAreaSetPickable                         gdc.MethodBindPtr
	fnXAreaSetMonitorCallback                  gdc.MethodBindPtr
	fnXAreaSetAreaMonitorCallback              gdc.MethodBindPtr
	fnXBodyCreate                              gdc.MethodBindPtr
	fnXBodySetSpace                            gdc.MethodBindPtr
	fnXBodyGetSpace                            gdc.MethodBindPtr
	fnXBodySetMode                             gdc.MethodBindPtr
	fnXBodyGetMode                             gdc.MethodBindPtr
	fnXBodyAddShape                            gdc.MethodBindPtr
	fnXBodySetShape                            gdc.MethodBindPtr
	fnXBodySetShapeTransform                   gdc.MethodBindPtr
	fnXBodyGetShapeCount                       gdc.MethodBindPtr
	fnXBodyGetShape                            gdc.MethodBindPtr
	fnXBodyGetShapeTransform                   gdc.MethodBindPtr
	fnXBodySetShapeDisabled                    gdc.MethodBindPtr
	fnXBodySetShapeAsOneWayCollision           gdc.MethodBindPtr
	fnXBodyRemoveShape                         gdc.MethodBindPtr
	fnXBodyClearShapes                         gdc.MethodBindPtr
	fnXBodyAttachObjectInstanceId              gdc.MethodBindPtr
	fnXBodyGetObjectInstanceId                 gdc.MethodBindPtr
	fnXBodyAttachCanvasInstanceId              gdc.MethodBindPtr
	fnXBodyGetCanvasInstanceId                 gdc.MethodBindPtr
	fnXBodySetContinuousCollisionDetectionMode gdc.MethodBindPtr
	fnXBodyGetContinuousCollisionDetectionMode gdc.MethodBindPtr
	fnXBodySetCollisionLayer                   gdc.MethodBindPtr
	fnXBodyGetCollisionLayer                   gdc.MethodBindPtr
	fnXBodySetCollisionMask                    gdc.MethodBindPtr
	fnXBodyGetCollisionMask                    gdc.MethodBindPtr
	fnXBodySetCollisionPriority                gdc.MethodBindPtr
	fnXBodyGetCollisionPriority                gdc.MethodBindPtr
	fnXBodySetParam                            gdc.MethodBindPtr
	fnXBodyGetParam                            gdc.MethodBindPtr
	fnXBodyResetMassProperties                 gdc.MethodBindPtr
	fnXBodySetState                            gdc.MethodBindPtr
	fnXBodyGetState                            gdc.MethodBindPtr
	fnXBodyApplyCentralImpulse                 gdc.MethodBindPtr
	fnXBodyApplyTorqueImpulse                  gdc.MethodBindPtr
	fnXBodyApplyImpulse                        gdc.MethodBindPtr
	fnXBodyApplyCentralForce                   gdc.MethodBindPtr
	fnXBodyApplyForce                          gdc.MethodBindPtr
	fnXBodyApplyTorque                         gdc.MethodBindPtr
	fnXBodyAddConstantCentralForce             gdc.MethodBindPtr
	fnXBodyAddConstantForce                    gdc.MethodBindPtr
	fnXBodyAddConstantTorque                   gdc.MethodBindPtr
	fnXBodySetConstantForce                    gdc.MethodBindPtr
	fnXBodyGetConstantForce                    gdc.MethodBindPtr
	fnXBodySetConstantTorque                   gdc.MethodBindPtr
	fnXBodyGetConstantTorque                   gdc.MethodBindPtr
	fnXBodySetAxisVelocity                     gdc.MethodBindPtr
	fnXBodyAddCollisionException               gdc.MethodBindPtr
	fnXBodyRemoveCollisionException            gdc.MethodBindPtr
	fnXBodyGetCollisionExceptions              gdc.MethodBindPtr
	fnXBodySetMaxContactsReported              gdc.MethodBindPtr
	fnXBodyGetMaxContactsReported              gdc.MethodBindPtr
	fnXBodySetContactsReportedDepthThreshold   gdc.MethodBindPtr
	fnXBodyGetContactsReportedDepthThreshold   gdc.MethodBindPtr
	fnXBodySetOmitForceIntegration             gdc.MethodBindPtr
	fnXBodyIsOmittingForceIntegration          gdc.MethodBindPtr
	fnXBodySetStateSyncCallback                gdc.MethodBindPtr
	fnXBodySetForceIntegrationCallback         gdc.MethodBindPtr
	fnXBodyCollideShape                        gdc.MethodBindPtr
	fnXBodySetPickable                         gdc.MethodBindPtr
	fnXBodyGetDirectState                      gdc.MethodBindPtr
	fnXBodyTestMotion                          gdc.MethodBindPtr
	fnXJointCreate                             gdc.MethodBindPtr
	fnXJointClear                              gdc.MethodBindPtr
	fnXJointSetParam                           gdc.MethodBindPtr
	fnXJointGetParam                           gdc.MethodBindPtr
	fnXJointDisableCollisionsBetweenBodies     gdc.MethodBindPtr
	fnXJointIsDisabledCollisionsBetweenBodies  gdc.MethodBindPtr
	fnXJointMakePin                            gdc.MethodBindPtr
	fnXJointMakeGroove                         gdc.MethodBindPtr
	fnXJointMakeDampedSpring                   gdc.MethodBindPtr
	fnXPinJointSetFlag                         gdc.MethodBindPtr
	fnXPinJointGetFlag                         gdc.MethodBindPtr
	fnXPinJointSetParam                        gdc.MethodBindPtr
	fnXPinJointGetParam                        gdc.MethodBindPtr
	fnXDampedSpringJointSetParam               gdc.MethodBindPtr
	fnXDampedSpringJointGetParam               gdc.MethodBindPtr
	fnXJointGetType                            gdc.MethodBindPtr
	fnXFreeRid                                 gdc.MethodBindPtr
	fnXSetActive                               gdc.MethodBindPtr
	fnXInit                                    gdc.MethodBindPtr
	fnXStep                                    gdc.MethodBindPtr
	fnXSync                                    gdc.MethodBindPtr
	fnXFlushQueries                            gdc.MethodBindPtr
	fnXEndSync                                 gdc.MethodBindPtr
	fnXFinish                                  gdc.MethodBindPtr
	fnXIsFlushingQueries                       gdc.MethodBindPtr
	fnXGetProcessInfo                          gdc.MethodBindPtr
	fnBodyTestMotionIsExcludingBody            gdc.MethodBindPtr
	fnBodyTestMotionIsExcludingObject          gdc.MethodBindPtr
}

var ptrsForPhysicsServer2DExtension ptrsForPhysicsServer2DExtensionList

func initPhysicsServer2DExtensionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsServer2DExtension")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("body_test_motion_is_excluding_body")
		defer methodName.Destroy()
		ptrsForPhysicsServer2DExtension.fnBodyTestMotionIsExcludingBody = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("body_test_motion_is_excluding_object")
		defer methodName.Destroy()
		ptrsForPhysicsServer2DExtension.fnBodyTestMotionIsExcludingObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
}

type PhysicsServer2DExtension struct {
	PhysicsServer2D
}

func (me *PhysicsServer2DExtension) BaseClass() string {
	return "PhysicsServer2DExtension"
}

func NewPhysicsServer2DExtension() *PhysicsServer2DExtension {
	str := StringNameFromStr("PhysicsServer2DExtension") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsServer2DExtension{}
	obj.SetBaseObject(objPtr)
	return obj
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

func (me *PhysicsServer2DExtension) BodyTestMotionIsExcludingBody(body RID) bool {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2DExtension.fnBodyTestMotionIsExcludingBody), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2DExtension) BodyTestMotionIsExcludingObject(object int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&object)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&object)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2DExtension.fnBodyTestMotionIsExcludingObject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
