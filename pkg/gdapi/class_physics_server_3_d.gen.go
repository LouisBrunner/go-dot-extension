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

type ptrsForPhysicsServer3DList struct {
	fnWorldBoundaryShapeCreate                  gdc.MethodBindPtr
	fnSeparationRayShapeCreate                  gdc.MethodBindPtr
	fnSphereShapeCreate                         gdc.MethodBindPtr
	fnBoxShapeCreate                            gdc.MethodBindPtr
	fnCapsuleShapeCreate                        gdc.MethodBindPtr
	fnCylinderShapeCreate                       gdc.MethodBindPtr
	fnConvexPolygonShapeCreate                  gdc.MethodBindPtr
	fnConcavePolygonShapeCreate                 gdc.MethodBindPtr
	fnHeightmapShapeCreate                      gdc.MethodBindPtr
	fnCustomShapeCreate                         gdc.MethodBindPtr
	fnShapeSetData                              gdc.MethodBindPtr
	fnShapeGetType                              gdc.MethodBindPtr
	fnShapeGetData                              gdc.MethodBindPtr
	fnSpaceCreate                               gdc.MethodBindPtr
	fnSpaceSetActive                            gdc.MethodBindPtr
	fnSpaceIsActive                             gdc.MethodBindPtr
	fnSpaceSetParam                             gdc.MethodBindPtr
	fnSpaceGetParam                             gdc.MethodBindPtr
	fnSpaceGetDirectState                       gdc.MethodBindPtr
	fnAreaCreate                                gdc.MethodBindPtr
	fnAreaSetSpace                              gdc.MethodBindPtr
	fnAreaGetSpace                              gdc.MethodBindPtr
	fnAreaAddShape                              gdc.MethodBindPtr
	fnAreaSetShape                              gdc.MethodBindPtr
	fnAreaSetShapeTransform                     gdc.MethodBindPtr
	fnAreaSetShapeDisabled                      gdc.MethodBindPtr
	fnAreaGetShapeCount                         gdc.MethodBindPtr
	fnAreaGetShape                              gdc.MethodBindPtr
	fnAreaGetShapeTransform                     gdc.MethodBindPtr
	fnAreaRemoveShape                           gdc.MethodBindPtr
	fnAreaClearShapes                           gdc.MethodBindPtr
	fnAreaSetCollisionLayer                     gdc.MethodBindPtr
	fnAreaGetCollisionLayer                     gdc.MethodBindPtr
	fnAreaSetCollisionMask                      gdc.MethodBindPtr
	fnAreaGetCollisionMask                      gdc.MethodBindPtr
	fnAreaSetParam                              gdc.MethodBindPtr
	fnAreaSetTransform                          gdc.MethodBindPtr
	fnAreaGetParam                              gdc.MethodBindPtr
	fnAreaGetTransform                          gdc.MethodBindPtr
	fnAreaAttachObjectInstanceId                gdc.MethodBindPtr
	fnAreaGetObjectInstanceId                   gdc.MethodBindPtr
	fnAreaSetMonitorCallback                    gdc.MethodBindPtr
	fnAreaSetAreaMonitorCallback                gdc.MethodBindPtr
	fnAreaSetMonitorable                        gdc.MethodBindPtr
	fnAreaSetRayPickable                        gdc.MethodBindPtr
	fnBodyCreate                                gdc.MethodBindPtr
	fnBodySetSpace                              gdc.MethodBindPtr
	fnBodyGetSpace                              gdc.MethodBindPtr
	fnBodySetMode                               gdc.MethodBindPtr
	fnBodyGetMode                               gdc.MethodBindPtr
	fnBodySetCollisionLayer                     gdc.MethodBindPtr
	fnBodyGetCollisionLayer                     gdc.MethodBindPtr
	fnBodySetCollisionMask                      gdc.MethodBindPtr
	fnBodyGetCollisionMask                      gdc.MethodBindPtr
	fnBodySetCollisionPriority                  gdc.MethodBindPtr
	fnBodyGetCollisionPriority                  gdc.MethodBindPtr
	fnBodyAddShape                              gdc.MethodBindPtr
	fnBodySetShape                              gdc.MethodBindPtr
	fnBodySetShapeTransform                     gdc.MethodBindPtr
	fnBodySetShapeDisabled                      gdc.MethodBindPtr
	fnBodyGetShapeCount                         gdc.MethodBindPtr
	fnBodyGetShape                              gdc.MethodBindPtr
	fnBodyGetShapeTransform                     gdc.MethodBindPtr
	fnBodyRemoveShape                           gdc.MethodBindPtr
	fnBodyClearShapes                           gdc.MethodBindPtr
	fnBodyAttachObjectInstanceId                gdc.MethodBindPtr
	fnBodyGetObjectInstanceId                   gdc.MethodBindPtr
	fnBodySetEnableContinuousCollisionDetection gdc.MethodBindPtr
	fnBodyIsContinuousCollisionDetectionEnabled gdc.MethodBindPtr
	fnBodySetParam                              gdc.MethodBindPtr
	fnBodyGetParam                              gdc.MethodBindPtr
	fnBodyResetMassProperties                   gdc.MethodBindPtr
	fnBodySetState                              gdc.MethodBindPtr
	fnBodyGetState                              gdc.MethodBindPtr
	fnBodyApplyCentralImpulse                   gdc.MethodBindPtr
	fnBodyApplyImpulse                          gdc.MethodBindPtr
	fnBodyApplyTorqueImpulse                    gdc.MethodBindPtr
	fnBodyApplyCentralForce                     gdc.MethodBindPtr
	fnBodyApplyForce                            gdc.MethodBindPtr
	fnBodyApplyTorque                           gdc.MethodBindPtr
	fnBodyAddConstantCentralForce               gdc.MethodBindPtr
	fnBodyAddConstantForce                      gdc.MethodBindPtr
	fnBodyAddConstantTorque                     gdc.MethodBindPtr
	fnBodySetConstantForce                      gdc.MethodBindPtr
	fnBodyGetConstantForce                      gdc.MethodBindPtr
	fnBodySetConstantTorque                     gdc.MethodBindPtr
	fnBodyGetConstantTorque                     gdc.MethodBindPtr
	fnBodySetAxisVelocity                       gdc.MethodBindPtr
	fnBodySetAxisLock                           gdc.MethodBindPtr
	fnBodyIsAxisLocked                          gdc.MethodBindPtr
	fnBodyAddCollisionException                 gdc.MethodBindPtr
	fnBodyRemoveCollisionException              gdc.MethodBindPtr
	fnBodySetMaxContactsReported                gdc.MethodBindPtr
	fnBodyGetMaxContactsReported                gdc.MethodBindPtr
	fnBodySetOmitForceIntegration               gdc.MethodBindPtr
	fnBodyIsOmittingForceIntegration            gdc.MethodBindPtr
	fnBodySetForceIntegrationCallback           gdc.MethodBindPtr
	fnBodySetRayPickable                        gdc.MethodBindPtr
	fnBodyTestMotion                            gdc.MethodBindPtr
	fnBodyGetDirectState                        gdc.MethodBindPtr
	fnSoftBodyGetBounds                         gdc.MethodBindPtr
	fnJointCreate                               gdc.MethodBindPtr
	fnJointClear                                gdc.MethodBindPtr
	fnJointMakePin                              gdc.MethodBindPtr
	fnPinJointSetParam                          gdc.MethodBindPtr
	fnPinJointGetParam                          gdc.MethodBindPtr
	fnPinJointSetLocalA                         gdc.MethodBindPtr
	fnPinJointGetLocalA                         gdc.MethodBindPtr
	fnPinJointSetLocalB                         gdc.MethodBindPtr
	fnPinJointGetLocalB                         gdc.MethodBindPtr
	fnJointMakeHinge                            gdc.MethodBindPtr
	fnHingeJointSetParam                        gdc.MethodBindPtr
	fnHingeJointGetParam                        gdc.MethodBindPtr
	fnHingeJointSetFlag                         gdc.MethodBindPtr
	fnHingeJointGetFlag                         gdc.MethodBindPtr
	fnJointMakeSlider                           gdc.MethodBindPtr
	fnSliderJointSetParam                       gdc.MethodBindPtr
	fnSliderJointGetParam                       gdc.MethodBindPtr
	fnJointMakeConeTwist                        gdc.MethodBindPtr
	fnConeTwistJointSetParam                    gdc.MethodBindPtr
	fnConeTwistJointGetParam                    gdc.MethodBindPtr
	fnJointGetType                              gdc.MethodBindPtr
	fnJointSetSolverPriority                    gdc.MethodBindPtr
	fnJointGetSolverPriority                    gdc.MethodBindPtr
	fnJointDisableCollisionsBetweenBodies       gdc.MethodBindPtr
	fnJointIsDisabledCollisionsBetweenBodies    gdc.MethodBindPtr
	fnJointMakeGeneric6Dof                      gdc.MethodBindPtr
	fnGeneric6DofJointSetParam                  gdc.MethodBindPtr
	fnGeneric6DofJointGetParam                  gdc.MethodBindPtr
	fnGeneric6DofJointSetFlag                   gdc.MethodBindPtr
	fnGeneric6DofJointGetFlag                   gdc.MethodBindPtr
	fnFreeRid                                   gdc.MethodBindPtr
	fnSetActive                                 gdc.MethodBindPtr
	fnGetProcessInfo                            gdc.MethodBindPtr
}

var ptrsForPhysicsServer3D ptrsForPhysicsServer3DList

func initPhysicsServer3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsServer3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("world_boundary_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnWorldBoundaryShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("separation_ray_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnSeparationRayShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("sphere_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnSphereShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("box_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBoxShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("capsule_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnCapsuleShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("cylinder_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnCylinderShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("convex_polygon_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnConvexPolygonShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("concave_polygon_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnConcavePolygonShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("heightmap_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnHeightmapShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("custom_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnCustomShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("shape_set_data")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnShapeSetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3175752987))
	}
	{
		methodName := StringNameFromStr("shape_get_type")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnShapeGetType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3418923367))
	}
	{
		methodName := StringNameFromStr("shape_get_data")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnShapeGetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4171304767))
	}
	{
		methodName := StringNameFromStr("space_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnSpaceCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("space_set_active")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnSpaceSetActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("space_is_active")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnSpaceIsActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("space_set_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnSpaceSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2406017470))
	}
	{
		methodName := StringNameFromStr("space_get_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnSpaceGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1523206731))
	}
	{
		methodName := StringNameFromStr("space_get_direct_state")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnSpaceGetDirectState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2048616813))
	}
	{
		methodName := StringNameFromStr("area_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("area_set_space")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaSetSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
	}
	{
		methodName := StringNameFromStr("area_get_space")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaGetSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
	}
	{
		methodName := StringNameFromStr("area_add_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaAddShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3711419014))
	}
	{
		methodName := StringNameFromStr("area_set_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaSetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2310537182))
	}
	{
		methodName := StringNameFromStr("area_set_shape_transform")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaSetShapeTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 675327471))
	}
	{
		methodName := StringNameFromStr("area_set_shape_disabled")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaSetShapeDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2658558584))
	}
	{
		methodName := StringNameFromStr("area_get_shape_count")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaGetShapeCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("area_get_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1066463050))
	}
	{
		methodName := StringNameFromStr("area_get_shape_transform")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaGetShapeTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1050775521))
	}
	{
		methodName := StringNameFromStr("area_remove_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaRemoveShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("area_clear_shapes")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaClearShapes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("area_set_collision_layer")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaSetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("area_get_collision_layer")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaGetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("area_set_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("area_get_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("area_set_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2980114638))
	}
	{
		methodName := StringNameFromStr("area_set_transform")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3935195649))
	}
	{
		methodName := StringNameFromStr("area_get_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 890056067))
	}
	{
		methodName := StringNameFromStr("area_get_transform")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1128465797))
	}
	{
		methodName := StringNameFromStr("area_attach_object_instance_id")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaAttachObjectInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("area_get_object_instance_id")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaGetObjectInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("area_set_monitor_callback")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaSetMonitorCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3379118538))
	}
	{
		methodName := StringNameFromStr("area_set_area_monitor_callback")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaSetAreaMonitorCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3379118538))
	}
	{
		methodName := StringNameFromStr("area_set_monitorable")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaSetMonitorable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("area_set_ray_pickable")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnAreaSetRayPickable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("body_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("body_set_space")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
	}
	{
		methodName := StringNameFromStr("body_get_space")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
	}
	{
		methodName := StringNameFromStr("body_set_mode")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 606803466))
	}
	{
		methodName := StringNameFromStr("body_get_mode")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2488819728))
	}
	{
		methodName := StringNameFromStr("body_set_collision_layer")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("body_get_collision_layer")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("body_set_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("body_get_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("body_set_collision_priority")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetCollisionPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("body_get_collision_priority")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetCollisionPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("body_add_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyAddShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3711419014))
	}
	{
		methodName := StringNameFromStr("body_set_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2310537182))
	}
	{
		methodName := StringNameFromStr("body_set_shape_transform")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetShapeTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 675327471))
	}
	{
		methodName := StringNameFromStr("body_set_shape_disabled")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetShapeDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2658558584))
	}
	{
		methodName := StringNameFromStr("body_get_shape_count")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetShapeCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("body_get_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1066463050))
	}
	{
		methodName := StringNameFromStr("body_get_shape_transform")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetShapeTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1050775521))
	}
	{
		methodName := StringNameFromStr("body_remove_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyRemoveShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("body_clear_shapes")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyClearShapes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("body_attach_object_instance_id")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyAttachObjectInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("body_get_object_instance_id")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetObjectInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("body_set_enable_continuous_collision_detection")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetEnableContinuousCollisionDetection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("body_is_continuous_collision_detection_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyIsContinuousCollisionDetectionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("body_set_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 910941953))
	}
	{
		methodName := StringNameFromStr("body_get_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3385027841))
	}
	{
		methodName := StringNameFromStr("body_reset_mass_properties")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyResetMassProperties = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("body_set_state")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 599977762))
	}
	{
		methodName := StringNameFromStr("body_get_state")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1850449534))
	}
	{
		methodName := StringNameFromStr("body_apply_central_impulse")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyApplyCentralImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("body_apply_impulse")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyApplyImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 390416203))
	}
	{
		methodName := StringNameFromStr("body_apply_torque_impulse")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyApplyTorqueImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("body_apply_central_force")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyApplyCentralForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("body_apply_force")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyApplyForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 390416203))
	}
	{
		methodName := StringNameFromStr("body_apply_torque")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyApplyTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("body_add_constant_central_force")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyAddConstantCentralForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("body_add_constant_force")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyAddConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 390416203))
	}
	{
		methodName := StringNameFromStr("body_add_constant_torque")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyAddConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("body_set_constant_force")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("body_get_constant_force")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 531438156))
	}
	{
		methodName := StringNameFromStr("body_set_constant_torque")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("body_get_constant_torque")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 531438156))
	}
	{
		methodName := StringNameFromStr("body_set_axis_velocity")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetAxisVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("body_set_axis_lock")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetAxisLock = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2020836892))
	}
	{
		methodName := StringNameFromStr("body_is_axis_locked")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyIsAxisLocked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 587853580))
	}
	{
		methodName := StringNameFromStr("body_add_collision_exception")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyAddCollisionException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
	}
	{
		methodName := StringNameFromStr("body_remove_collision_exception")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyRemoveCollisionException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
	}
	{
		methodName := StringNameFromStr("body_set_max_contacts_reported")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetMaxContactsReported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("body_get_max_contacts_reported")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetMaxContactsReported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("body_set_omit_force_integration")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetOmitForceIntegration = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("body_is_omitting_force_integration")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyIsOmittingForceIntegration = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("body_set_force_integration_callback")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetForceIntegrationCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3059434249))
	}
	{
		methodName := StringNameFromStr("body_set_ray_pickable")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodySetRayPickable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("body_test_motion")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyTestMotion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1944921792))
	}
	{
		methodName := StringNameFromStr("body_get_direct_state")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnBodyGetDirectState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3029727957))
	}
	{
		methodName := StringNameFromStr("soft_body_get_bounds")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnSoftBodyGetBounds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 974181306))
	}
	{
		methodName := StringNameFromStr("joint_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnJointCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("joint_clear")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnJointClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("joint_make_pin")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnJointMakePin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4280171926))
	}
	{
		methodName := StringNameFromStr("pin_joint_set_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnPinJointSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 810685294))
	}
	{
		methodName := StringNameFromStr("pin_joint_get_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnPinJointGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2817972347))
	}
	{
		methodName := StringNameFromStr("pin_joint_set_local_a")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnPinJointSetLocalA = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("pin_joint_get_local_a")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnPinJointGetLocalA = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 531438156))
	}
	{
		methodName := StringNameFromStr("pin_joint_set_local_b")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnPinJointSetLocalB = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("pin_joint_get_local_b")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnPinJointGetLocalB = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 531438156))
	}
	{
		methodName := StringNameFromStr("joint_make_hinge")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnJointMakeHinge = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1684107643))
	}
	{
		methodName := StringNameFromStr("hinge_joint_set_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnHingeJointSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3165502333))
	}
	{
		methodName := StringNameFromStr("hinge_joint_get_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnHingeJointGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2129207581))
	}
	{
		methodName := StringNameFromStr("hinge_joint_set_flag")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnHingeJointSetFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1601626188))
	}
	{
		methodName := StringNameFromStr("hinge_joint_get_flag")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnHingeJointGetFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4165147865))
	}
	{
		methodName := StringNameFromStr("joint_make_slider")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnJointMakeSlider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1684107643))
	}
	{
		methodName := StringNameFromStr("slider_joint_set_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnSliderJointSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2264833593))
	}
	{
		methodName := StringNameFromStr("slider_joint_get_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnSliderJointGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3498644957))
	}
	{
		methodName := StringNameFromStr("joint_make_cone_twist")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnJointMakeConeTwist = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1684107643))
	}
	{
		methodName := StringNameFromStr("cone_twist_joint_set_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnConeTwistJointSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 808587618))
	}
	{
		methodName := StringNameFromStr("cone_twist_joint_get_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnConeTwistJointGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1134789658))
	}
	{
		methodName := StringNameFromStr("joint_get_type")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnJointGetType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4290791900))
	}
	{
		methodName := StringNameFromStr("joint_set_solver_priority")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnJointSetSolverPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("joint_get_solver_priority")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnJointGetSolverPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("joint_disable_collisions_between_bodies")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnJointDisableCollisionsBetweenBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("joint_is_disabled_collisions_between_bodies")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnJointIsDisabledCollisionsBetweenBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("joint_make_generic_6dof")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnJointMakeGeneric6Dof = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1684107643))
	}
	{
		methodName := StringNameFromStr("generic_6dof_joint_set_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnGeneric6DofJointSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2600081391))
	}
	{
		methodName := StringNameFromStr("generic_6dof_joint_get_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnGeneric6DofJointGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 467122058))
	}
	{
		methodName := StringNameFromStr("generic_6dof_joint_set_flag")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnGeneric6DofJointSetFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3570926903))
	}
	{
		methodName := StringNameFromStr("generic_6dof_joint_get_flag")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnGeneric6DofJointGetFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4158090196))
	}
	{
		methodName := StringNameFromStr("free_rid")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnFreeRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("set_active")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnSetActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_process_info")
		defer methodName.Destroy()
		ptrsForPhysicsServer3D.fnGetProcessInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1332958745))
	}
}

type PhysicsServer3D struct {
	Object
}

func (me *PhysicsServer3D) BaseClass() string {
	return "PhysicsServer3D"
}

func NewPhysicsServer3D() *PhysicsServer3D {
	str := StringNameFromStr("PhysicsServer3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsServer3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type PhysicsServer3DJointType int

const (
	PhysicsServer3DJointTypeJointTypePin       PhysicsServer3DJointType = 0
	PhysicsServer3DJointTypeJointTypeHinge     PhysicsServer3DJointType = 1
	PhysicsServer3DJointTypeJointTypeSlider    PhysicsServer3DJointType = 2
	PhysicsServer3DJointTypeJointTypeConeTwist PhysicsServer3DJointType = 3
	PhysicsServer3DJointTypeJointType6Dof      PhysicsServer3DJointType = 4
	PhysicsServer3DJointTypeJointTypeMax       PhysicsServer3DJointType = 5
)

type PhysicsServer3DPinJointParam int

const (
	PhysicsServer3DPinJointParamPinJointBias         PhysicsServer3DPinJointParam = 0
	PhysicsServer3DPinJointParamPinJointDamping      PhysicsServer3DPinJointParam = 1
	PhysicsServer3DPinJointParamPinJointImpulseClamp PhysicsServer3DPinJointParam = 2
)

type PhysicsServer3DHingeJointParam int

const (
	PhysicsServer3DHingeJointParamHingeJointBias                PhysicsServer3DHingeJointParam = 0
	PhysicsServer3DHingeJointParamHingeJointLimitUpper          PhysicsServer3DHingeJointParam = 1
	PhysicsServer3DHingeJointParamHingeJointLimitLower          PhysicsServer3DHingeJointParam = 2
	PhysicsServer3DHingeJointParamHingeJointLimitBias           PhysicsServer3DHingeJointParam = 3
	PhysicsServer3DHingeJointParamHingeJointLimitSoftness       PhysicsServer3DHingeJointParam = 4
	PhysicsServer3DHingeJointParamHingeJointLimitRelaxation     PhysicsServer3DHingeJointParam = 5
	PhysicsServer3DHingeJointParamHingeJointMotorTargetVelocity PhysicsServer3DHingeJointParam = 6
	PhysicsServer3DHingeJointParamHingeJointMotorMaxImpulse     PhysicsServer3DHingeJointParam = 7
)

type PhysicsServer3DHingeJointFlag int

const (
	PhysicsServer3DHingeJointFlagHingeJointFlagUseLimit    PhysicsServer3DHingeJointFlag = 0
	PhysicsServer3DHingeJointFlagHingeJointFlagEnableMotor PhysicsServer3DHingeJointFlag = 1
)

type PhysicsServer3DSliderJointParam int

const (
	PhysicsServer3DSliderJointParamSliderJointLinearLimitUpper             PhysicsServer3DSliderJointParam = 0
	PhysicsServer3DSliderJointParamSliderJointLinearLimitLower             PhysicsServer3DSliderJointParam = 1
	PhysicsServer3DSliderJointParamSliderJointLinearLimitSoftness          PhysicsServer3DSliderJointParam = 2
	PhysicsServer3DSliderJointParamSliderJointLinearLimitRestitution       PhysicsServer3DSliderJointParam = 3
	PhysicsServer3DSliderJointParamSliderJointLinearLimitDamping           PhysicsServer3DSliderJointParam = 4
	PhysicsServer3DSliderJointParamSliderJointLinearMotionSoftness         PhysicsServer3DSliderJointParam = 5
	PhysicsServer3DSliderJointParamSliderJointLinearMotionRestitution      PhysicsServer3DSliderJointParam = 6
	PhysicsServer3DSliderJointParamSliderJointLinearMotionDamping          PhysicsServer3DSliderJointParam = 7
	PhysicsServer3DSliderJointParamSliderJointLinearOrthogonalSoftness     PhysicsServer3DSliderJointParam = 8
	PhysicsServer3DSliderJointParamSliderJointLinearOrthogonalRestitution  PhysicsServer3DSliderJointParam = 9
	PhysicsServer3DSliderJointParamSliderJointLinearOrthogonalDamping      PhysicsServer3DSliderJointParam = 10
	PhysicsServer3DSliderJointParamSliderJointAngularLimitUpper            PhysicsServer3DSliderJointParam = 11
	PhysicsServer3DSliderJointParamSliderJointAngularLimitLower            PhysicsServer3DSliderJointParam = 12
	PhysicsServer3DSliderJointParamSliderJointAngularLimitSoftness         PhysicsServer3DSliderJointParam = 13
	PhysicsServer3DSliderJointParamSliderJointAngularLimitRestitution      PhysicsServer3DSliderJointParam = 14
	PhysicsServer3DSliderJointParamSliderJointAngularLimitDamping          PhysicsServer3DSliderJointParam = 15
	PhysicsServer3DSliderJointParamSliderJointAngularMotionSoftness        PhysicsServer3DSliderJointParam = 16
	PhysicsServer3DSliderJointParamSliderJointAngularMotionRestitution     PhysicsServer3DSliderJointParam = 17
	PhysicsServer3DSliderJointParamSliderJointAngularMotionDamping         PhysicsServer3DSliderJointParam = 18
	PhysicsServer3DSliderJointParamSliderJointAngularOrthogonalSoftness    PhysicsServer3DSliderJointParam = 19
	PhysicsServer3DSliderJointParamSliderJointAngularOrthogonalRestitution PhysicsServer3DSliderJointParam = 20
	PhysicsServer3DSliderJointParamSliderJointAngularOrthogonalDamping     PhysicsServer3DSliderJointParam = 21
	PhysicsServer3DSliderJointParamSliderJointMax                          PhysicsServer3DSliderJointParam = 22
)

type PhysicsServer3DConeTwistJointParam int

const (
	PhysicsServer3DConeTwistJointParamConeTwistJointSwingSpan  PhysicsServer3DConeTwistJointParam = 0
	PhysicsServer3DConeTwistJointParamConeTwistJointTwistSpan  PhysicsServer3DConeTwistJointParam = 1
	PhysicsServer3DConeTwistJointParamConeTwistJointBias       PhysicsServer3DConeTwistJointParam = 2
	PhysicsServer3DConeTwistJointParamConeTwistJointSoftness   PhysicsServer3DConeTwistJointParam = 3
	PhysicsServer3DConeTwistJointParamConeTwistJointRelaxation PhysicsServer3DConeTwistJointParam = 4
)

type PhysicsServer3DG6DOFJointAxisParam int

const (
	PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearLowerLimit           PhysicsServer3DG6DOFJointAxisParam = 0
	PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearUpperLimit           PhysicsServer3DG6DOFJointAxisParam = 1
	PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearLimitSoftness        PhysicsServer3DG6DOFJointAxisParam = 2
	PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearRestitution          PhysicsServer3DG6DOFJointAxisParam = 3
	PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearDamping              PhysicsServer3DG6DOFJointAxisParam = 4
	PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearMotorTargetVelocity  PhysicsServer3DG6DOFJointAxisParam = 5
	PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearMotorForceLimit      PhysicsServer3DG6DOFJointAxisParam = 6
	PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularLowerLimit          PhysicsServer3DG6DOFJointAxisParam = 10
	PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularUpperLimit          PhysicsServer3DG6DOFJointAxisParam = 11
	PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularLimitSoftness       PhysicsServer3DG6DOFJointAxisParam = 12
	PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularDamping             PhysicsServer3DG6DOFJointAxisParam = 13
	PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularRestitution         PhysicsServer3DG6DOFJointAxisParam = 14
	PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularForceLimit          PhysicsServer3DG6DOFJointAxisParam = 15
	PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularErp                 PhysicsServer3DG6DOFJointAxisParam = 16
	PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularMotorTargetVelocity PhysicsServer3DG6DOFJointAxisParam = 17
	PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularMotorForceLimit     PhysicsServer3DG6DOFJointAxisParam = 18
)

type PhysicsServer3DG6DOFJointAxisFlag int

const (
	PhysicsServer3DG6DOFJointAxisFlagG6DofJointFlagEnableLinearLimit  PhysicsServer3DG6DOFJointAxisFlag = 0
	PhysicsServer3DG6DOFJointAxisFlagG6DofJointFlagEnableAngularLimit PhysicsServer3DG6DOFJointAxisFlag = 1
	PhysicsServer3DG6DOFJointAxisFlagG6DofJointFlagEnableMotor        PhysicsServer3DG6DOFJointAxisFlag = 4
	PhysicsServer3DG6DOFJointAxisFlagG6DofJointFlagEnableLinearMotor  PhysicsServer3DG6DOFJointAxisFlag = 5
)

type PhysicsServer3DShapeType int

const (
	PhysicsServer3DShapeTypeShapeWorldBoundary  PhysicsServer3DShapeType = 0
	PhysicsServer3DShapeTypeShapeSeparationRay  PhysicsServer3DShapeType = 1
	PhysicsServer3DShapeTypeShapeSphere         PhysicsServer3DShapeType = 2
	PhysicsServer3DShapeTypeShapeBox            PhysicsServer3DShapeType = 3
	PhysicsServer3DShapeTypeShapeCapsule        PhysicsServer3DShapeType = 4
	PhysicsServer3DShapeTypeShapeCylinder       PhysicsServer3DShapeType = 5
	PhysicsServer3DShapeTypeShapeConvexPolygon  PhysicsServer3DShapeType = 6
	PhysicsServer3DShapeTypeShapeConcavePolygon PhysicsServer3DShapeType = 7
	PhysicsServer3DShapeTypeShapeHeightmap      PhysicsServer3DShapeType = 8
	PhysicsServer3DShapeTypeShapeSoftBody       PhysicsServer3DShapeType = 9
	PhysicsServer3DShapeTypeShapeCustom         PhysicsServer3DShapeType = 10
)

type PhysicsServer3DAreaParameter int

const (
	PhysicsServer3DAreaParameterAreaParamGravityOverrideMode      PhysicsServer3DAreaParameter = 0
	PhysicsServer3DAreaParameterAreaParamGravity                  PhysicsServer3DAreaParameter = 1
	PhysicsServer3DAreaParameterAreaParamGravityVector            PhysicsServer3DAreaParameter = 2
	PhysicsServer3DAreaParameterAreaParamGravityIsPoint           PhysicsServer3DAreaParameter = 3
	PhysicsServer3DAreaParameterAreaParamGravityPointUnitDistance PhysicsServer3DAreaParameter = 4
	PhysicsServer3DAreaParameterAreaParamLinearDampOverrideMode   PhysicsServer3DAreaParameter = 5
	PhysicsServer3DAreaParameterAreaParamLinearDamp               PhysicsServer3DAreaParameter = 6
	PhysicsServer3DAreaParameterAreaParamAngularDampOverrideMode  PhysicsServer3DAreaParameter = 7
	PhysicsServer3DAreaParameterAreaParamAngularDamp              PhysicsServer3DAreaParameter = 8
	PhysicsServer3DAreaParameterAreaParamPriority                 PhysicsServer3DAreaParameter = 9
	PhysicsServer3DAreaParameterAreaParamWindForceMagnitude       PhysicsServer3DAreaParameter = 10
	PhysicsServer3DAreaParameterAreaParamWindSource               PhysicsServer3DAreaParameter = 11
	PhysicsServer3DAreaParameterAreaParamWindDirection            PhysicsServer3DAreaParameter = 12
	PhysicsServer3DAreaParameterAreaParamWindAttenuationFactor    PhysicsServer3DAreaParameter = 13
)

type PhysicsServer3DAreaSpaceOverrideMode int

const (
	PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideDisabled       PhysicsServer3DAreaSpaceOverrideMode = 0
	PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideCombine        PhysicsServer3DAreaSpaceOverrideMode = 1
	PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideCombineReplace PhysicsServer3DAreaSpaceOverrideMode = 2
	PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideReplace        PhysicsServer3DAreaSpaceOverrideMode = 3
	PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideReplaceCombine PhysicsServer3DAreaSpaceOverrideMode = 4
)

type PhysicsServer3DBodyMode int

const (
	PhysicsServer3DBodyModeBodyModeStatic      PhysicsServer3DBodyMode = 0
	PhysicsServer3DBodyModeBodyModeKinematic   PhysicsServer3DBodyMode = 1
	PhysicsServer3DBodyModeBodyModeRigid       PhysicsServer3DBodyMode = 2
	PhysicsServer3DBodyModeBodyModeRigidLinear PhysicsServer3DBodyMode = 3
)

type PhysicsServer3DBodyParameter int

const (
	PhysicsServer3DBodyParameterBodyParamBounce          PhysicsServer3DBodyParameter = 0
	PhysicsServer3DBodyParameterBodyParamFriction        PhysicsServer3DBodyParameter = 1
	PhysicsServer3DBodyParameterBodyParamMass            PhysicsServer3DBodyParameter = 2
	PhysicsServer3DBodyParameterBodyParamInertia         PhysicsServer3DBodyParameter = 3
	PhysicsServer3DBodyParameterBodyParamCenterOfMass    PhysicsServer3DBodyParameter = 4
	PhysicsServer3DBodyParameterBodyParamGravityScale    PhysicsServer3DBodyParameter = 5
	PhysicsServer3DBodyParameterBodyParamLinearDampMode  PhysicsServer3DBodyParameter = 6
	PhysicsServer3DBodyParameterBodyParamAngularDampMode PhysicsServer3DBodyParameter = 7
	PhysicsServer3DBodyParameterBodyParamLinearDamp      PhysicsServer3DBodyParameter = 8
	PhysicsServer3DBodyParameterBodyParamAngularDamp     PhysicsServer3DBodyParameter = 9
	PhysicsServer3DBodyParameterBodyParamMax             PhysicsServer3DBodyParameter = 10
)

type PhysicsServer3DBodyDampMode int

const (
	PhysicsServer3DBodyDampModeBodyDampModeCombine PhysicsServer3DBodyDampMode = 0
	PhysicsServer3DBodyDampModeBodyDampModeReplace PhysicsServer3DBodyDampMode = 1
)

type PhysicsServer3DBodyState int

const (
	PhysicsServer3DBodyStateBodyStateTransform       PhysicsServer3DBodyState = 0
	PhysicsServer3DBodyStateBodyStateLinearVelocity  PhysicsServer3DBodyState = 1
	PhysicsServer3DBodyStateBodyStateAngularVelocity PhysicsServer3DBodyState = 2
	PhysicsServer3DBodyStateBodyStateSleeping        PhysicsServer3DBodyState = 3
	PhysicsServer3DBodyStateBodyStateCanSleep        PhysicsServer3DBodyState = 4
)

type PhysicsServer3DAreaBodyStatus int

const (
	PhysicsServer3DAreaBodyStatusAreaBodyAdded   PhysicsServer3DAreaBodyStatus = 0
	PhysicsServer3DAreaBodyStatusAreaBodyRemoved PhysicsServer3DAreaBodyStatus = 1
)

type PhysicsServer3DProcessInfo int

const (
	PhysicsServer3DProcessInfoInfoActiveObjects  PhysicsServer3DProcessInfo = 0
	PhysicsServer3DProcessInfoInfoCollisionPairs PhysicsServer3DProcessInfo = 1
	PhysicsServer3DProcessInfoInfoIslandCount    PhysicsServer3DProcessInfo = 2
)

type PhysicsServer3DSpaceParameter int

const (
	PhysicsServer3DSpaceParameterSpaceParamContactRecycleRadius              PhysicsServer3DSpaceParameter = 0
	PhysicsServer3DSpaceParameterSpaceParamContactMaxSeparation              PhysicsServer3DSpaceParameter = 1
	PhysicsServer3DSpaceParameterSpaceParamContactMaxAllowedPenetration      PhysicsServer3DSpaceParameter = 2
	PhysicsServer3DSpaceParameterSpaceParamContactDefaultBias                PhysicsServer3DSpaceParameter = 3
	PhysicsServer3DSpaceParameterSpaceParamBodyLinearVelocitySleepThreshold  PhysicsServer3DSpaceParameter = 4
	PhysicsServer3DSpaceParameterSpaceParamBodyAngularVelocitySleepThreshold PhysicsServer3DSpaceParameter = 5
	PhysicsServer3DSpaceParameterSpaceParamBodyTimeToSleep                   PhysicsServer3DSpaceParameter = 6
	PhysicsServer3DSpaceParameterSpaceParamSolverIterations                  PhysicsServer3DSpaceParameter = 7
)

type PhysicsServer3DBodyAxis int

const (
	PhysicsServer3DBodyAxisBodyAxisLinearX  PhysicsServer3DBodyAxis = 1
	PhysicsServer3DBodyAxisBodyAxisLinearY  PhysicsServer3DBodyAxis = 2
	PhysicsServer3DBodyAxisBodyAxisLinearZ  PhysicsServer3DBodyAxis = 4
	PhysicsServer3DBodyAxisBodyAxisAngularX PhysicsServer3DBodyAxis = 8
	PhysicsServer3DBodyAxisBodyAxisAngularY PhysicsServer3DBodyAxis = 16
	PhysicsServer3DBodyAxisBodyAxisAngularZ PhysicsServer3DBodyAxis = 32
)

func (me *PhysicsServer3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsServer3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsServer3D) WorldBoundaryShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnWorldBoundaryShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) SeparationRayShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnSeparationRayShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) SphereShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnSphereShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) BoxShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBoxShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) CapsuleShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnCapsuleShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) CylinderShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnCylinderShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) ConvexPolygonShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnConvexPolygonShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) ConcavePolygonShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnConcavePolygonShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) HeightmapShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnHeightmapShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) CustomShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnCustomShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) ShapeSetData(shape RID, data Variant) {
	cargs := []gdc.ConstTypePtr{shape.AsCTypePtr(), data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnShapeSetData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) ShapeGetType(shape RID) PhysicsServer3DShapeType {
	cargs := []gdc.ConstTypePtr{shape.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret PhysicsServer3DShapeType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnShapeGetType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PhysicsServer3D) ShapeGetData(shape RID) Variant {
	cargs := []gdc.ConstTypePtr{shape.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnShapeGetData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) SpaceCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnSpaceCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) SpaceSetActive(space RID, active bool) {
	cargs := []gdc.ConstTypePtr{space.AsCTypePtr(), gdc.ConstTypePtr(&active)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnSpaceSetActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) SpaceIsActive(space RID) bool {
	cargs := []gdc.ConstTypePtr{space.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnSpaceIsActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) SpaceSetParam(space RID, param PhysicsServer3DSpaceParameter, value float64) {
	cargs := []gdc.ConstTypePtr{space.AsCTypePtr(), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnSpaceSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) SpaceGetParam(space RID, param PhysicsServer3DSpaceParameter) float64 {
	cargs := []gdc.ConstTypePtr{space.AsCTypePtr(), gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnSpaceGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) SpaceGetDirectState(space RID) PhysicsDirectSpaceState3D {
	cargs := []gdc.ConstTypePtr{space.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsDirectSpaceState3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnSpaceGetDirectState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) AreaCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) AreaSetSpace(area RID, space RID) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), space.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaSetSpace), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaGetSpace(area RID) RID {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaGetSpace), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) AreaAddShape(area RID, shape RID, transform Transform3D, disabled bool) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), shape.AsCTypePtr(), transform.AsCTypePtr(), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaAddShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaSetShape(area RID, shape_idx int64, shape RID) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx), shape.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaSetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaSetShapeTransform(area RID, shape_idx int64, transform Transform3D) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaSetShapeTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaSetShapeDisabled(area RID, shape_idx int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaSetShapeDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaGetShapeCount(area RID) int64 {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaGetShapeCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) AreaGetShape(area RID, shape_idx int64) RID {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&shape_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaGetShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) AreaGetShapeTransform(area RID, shape_idx int64) Transform3D {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&shape_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaGetShapeTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) AreaRemoveShape(area RID, shape_idx int64) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaRemoveShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaClearShapes(area RID) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaClearShapes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaSetCollisionLayer(area RID, layer int64) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaSetCollisionLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaGetCollisionLayer(area RID) int64 {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaGetCollisionLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) AreaSetCollisionMask(area RID, mask int64) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaGetCollisionMask(area RID) int64 {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) AreaSetParam(area RID, param PhysicsServer3DAreaParameter, value Variant) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&param), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaSetTransform(area RID, transform Transform3D) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaGetParam(area RID, param PhysicsServer3DAreaParameter) Variant {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) AreaGetTransform(area RID) Transform3D {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) AreaAttachObjectInstanceId(area RID, id int64) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaAttachObjectInstanceId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaGetObjectInstanceId(area RID) int64 {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaGetObjectInstanceId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) AreaSetMonitorCallback(area RID, callback Callable) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaSetMonitorCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaSetAreaMonitorCallback(area RID, callback Callable) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaSetAreaMonitorCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaSetMonitorable(area RID, monitorable bool) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&monitorable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaSetMonitorable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) AreaSetRayPickable(area RID, enable bool) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnAreaSetRayPickable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) BodySetSpace(body RID, space RID) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), space.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetSpace), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyGetSpace(body RID) RID {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetSpace), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) BodySetMode(body RID, mode PhysicsServer3DBodyMode) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyGetMode(body RID) PhysicsServer3DBodyMode {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret PhysicsServer3DBodyMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PhysicsServer3D) BodySetCollisionLayer(body RID, layer int64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetCollisionLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyGetCollisionLayer(body RID) int64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetCollisionLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) BodySetCollisionMask(body RID, mask int64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyGetCollisionMask(body RID) int64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) BodySetCollisionPriority(body RID, priority float64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetCollisionPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyGetCollisionPriority(body RID) float64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetCollisionPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) BodyAddShape(body RID, shape RID, transform Transform3D, disabled bool) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), shape.AsCTypePtr(), transform.AsCTypePtr(), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyAddShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodySetShape(body RID, shape_idx int64, shape RID) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx), shape.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodySetShapeTransform(body RID, shape_idx int64, transform Transform3D) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetShapeTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodySetShapeDisabled(body RID, shape_idx int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetShapeDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyGetShapeCount(body RID) int64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetShapeCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) BodyGetShape(body RID, shape_idx int64) RID {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&shape_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) BodyGetShapeTransform(body RID, shape_idx int64) Transform3D {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&shape_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetShapeTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) BodyRemoveShape(body RID, shape_idx int64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyRemoveShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyClearShapes(body RID) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyClearShapes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyAttachObjectInstanceId(body RID, id int64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyAttachObjectInstanceId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyGetObjectInstanceId(body RID) int64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetObjectInstanceId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) BodySetEnableContinuousCollisionDetection(body RID, enable bool) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetEnableContinuousCollisionDetection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyIsContinuousCollisionDetectionEnabled(body RID) bool {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyIsContinuousCollisionDetectionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) BodySetParam(body RID, param PhysicsServer3DBodyParameter, value Variant) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&param), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyGetParam(body RID, param PhysicsServer3DBodyParameter) Variant {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) BodyResetMassProperties(body RID) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyResetMassProperties), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodySetState(body RID, state PhysicsServer3DBodyState, value Variant) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&state), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetState), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyGetState(body RID, state PhysicsServer3DBodyState) Variant {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&state)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&state)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) BodyApplyCentralImpulse(body RID, impulse Vector3) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), impulse.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyApplyCentralImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyApplyImpulse(body RID, impulse Vector3, position Vector3) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), impulse.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyApplyImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyApplyTorqueImpulse(body RID, impulse Vector3) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), impulse.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyApplyTorqueImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyApplyCentralForce(body RID, force Vector3) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyApplyCentralForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyApplyForce(body RID, force Vector3, position Vector3) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), force.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyApplyForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyApplyTorque(body RID, torque Vector3) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), torque.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyApplyTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyAddConstantCentralForce(body RID, force Vector3) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyAddConstantCentralForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyAddConstantForce(body RID, force Vector3, position Vector3) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), force.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyAddConstantForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyAddConstantTorque(body RID, torque Vector3) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), torque.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyAddConstantTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodySetConstantForce(body RID, force Vector3) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetConstantForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyGetConstantForce(body RID) Vector3 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetConstantForce), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) BodySetConstantTorque(body RID, torque Vector3) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), torque.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetConstantTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyGetConstantTorque(body RID) Vector3 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetConstantTorque), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) BodySetAxisVelocity(body RID, axis_velocity Vector3) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), axis_velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetAxisVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodySetAxisLock(body RID, axis PhysicsServer3DBodyAxis, lock bool) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&axis), gdc.ConstTypePtr(&lock)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetAxisLock), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyIsAxisLocked(body RID, axis PhysicsServer3DBodyAxis) bool {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&axis)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&axis)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyIsAxisLocked), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) BodyAddCollisionException(body RID, excepted_body RID) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), excepted_body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyAddCollisionException), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyRemoveCollisionException(body RID, excepted_body RID) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), excepted_body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyRemoveCollisionException), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodySetMaxContactsReported(body RID, amount int64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetMaxContactsReported), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyGetMaxContactsReported(body RID) int64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetMaxContactsReported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) BodySetOmitForceIntegration(body RID, enable bool) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetOmitForceIntegration), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyIsOmittingForceIntegration(body RID) bool {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyIsOmittingForceIntegration), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) BodySetForceIntegrationCallback(body RID, callable Callable, userdata Variant) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), callable.AsCTypePtr(), userdata.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetForceIntegrationCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodySetRayPickable(body RID, enable bool) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodySetRayPickable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) BodyTestMotion(body RID, parameters PhysicsTestMotionParameters3D, result PhysicsTestMotionResult3D) bool {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), parameters.AsCTypePtr(), result.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyTestMotion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) BodyGetDirectState(body RID) PhysicsDirectBodyState3D {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsDirectBodyState3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnBodyGetDirectState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) SoftBodyGetBounds(body RID) AABB {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAABB()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnSoftBodyGetBounds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) JointCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnJointCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) JointClear(joint RID) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnJointClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) JointMakePin(joint RID, body_A RID, local_A Vector3, body_B RID, local_B Vector3) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), body_A.AsCTypePtr(), local_A.AsCTypePtr(), body_B.AsCTypePtr(), local_B.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnJointMakePin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) PinJointSetParam(joint RID, param PhysicsServer3DPinJointParam, value float64) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnPinJointSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) PinJointGetParam(joint RID, param PhysicsServer3DPinJointParam) float64 {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnPinJointGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) PinJointSetLocalA(joint RID, local_A Vector3) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), local_A.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnPinJointSetLocalA), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) PinJointGetLocalA(joint RID) Vector3 {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnPinJointGetLocalA), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) PinJointSetLocalB(joint RID, local_B Vector3) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), local_B.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnPinJointSetLocalB), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) PinJointGetLocalB(joint RID) Vector3 {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnPinJointGetLocalB), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer3D) JointMakeHinge(joint RID, body_A RID, hinge_A Transform3D, body_B RID, hinge_B Transform3D) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), body_A.AsCTypePtr(), hinge_A.AsCTypePtr(), body_B.AsCTypePtr(), hinge_B.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnJointMakeHinge), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) HingeJointSetParam(joint RID, param PhysicsServer3DHingeJointParam, value float64) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnHingeJointSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) HingeJointGetParam(joint RID, param PhysicsServer3DHingeJointParam) float64 {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnHingeJointGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) HingeJointSetFlag(joint RID, flag PhysicsServer3DHingeJointFlag, enabled bool) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnHingeJointSetFlag), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) HingeJointGetFlag(joint RID, flag PhysicsServer3DHingeJointFlag) bool {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&flag)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&flag)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnHingeJointGetFlag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) JointMakeSlider(joint RID, body_A RID, local_ref_A Transform3D, body_B RID, local_ref_B Transform3D) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), body_A.AsCTypePtr(), local_ref_A.AsCTypePtr(), body_B.AsCTypePtr(), local_ref_B.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnJointMakeSlider), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) SliderJointSetParam(joint RID, param PhysicsServer3DSliderJointParam, value float64) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnSliderJointSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) SliderJointGetParam(joint RID, param PhysicsServer3DSliderJointParam) float64 {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnSliderJointGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) JointMakeConeTwist(joint RID, body_A RID, local_ref_A Transform3D, body_B RID, local_ref_B Transform3D) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), body_A.AsCTypePtr(), local_ref_A.AsCTypePtr(), body_B.AsCTypePtr(), local_ref_B.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnJointMakeConeTwist), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) ConeTwistJointSetParam(joint RID, param PhysicsServer3DConeTwistJointParam, value float64) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnConeTwistJointSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) ConeTwistJointGetParam(joint RID, param PhysicsServer3DConeTwistJointParam) float64 {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnConeTwistJointGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) JointGetType(joint RID) PhysicsServer3DJointType {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret PhysicsServer3DJointType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnJointGetType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PhysicsServer3D) JointSetSolverPriority(joint RID, priority int64) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnJointSetSolverPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) JointGetSolverPriority(joint RID) int64 {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnJointGetSolverPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) JointDisableCollisionsBetweenBodies(joint RID, disable bool) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&disable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnJointDisableCollisionsBetweenBodies), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) JointIsDisabledCollisionsBetweenBodies(joint RID) bool {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnJointIsDisabledCollisionsBetweenBodies), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) JointMakeGeneric6Dof(joint RID, body_A RID, local_ref_A Transform3D, body_B RID, local_ref_B Transform3D) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), body_A.AsCTypePtr(), local_ref_A.AsCTypePtr(), body_B.AsCTypePtr(), local_ref_B.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnJointMakeGeneric6Dof), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) Generic6DofJointSetParam(joint RID, axis Vector3Axis, param PhysicsServer3DG6DOFJointAxisParam, value float64) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&axis), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnGeneric6DofJointSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) Generic6DofJointGetParam(joint RID, axis Vector3Axis, param PhysicsServer3DG6DOFJointAxisParam) float64 {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&axis), gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&axis)
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnGeneric6DofJointGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) Generic6DofJointSetFlag(joint RID, axis Vector3Axis, flag PhysicsServer3DG6DOFJointAxisFlag, enable bool) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&axis), gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnGeneric6DofJointSetFlag), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) Generic6DofJointGetFlag(joint RID, axis Vector3Axis, flag PhysicsServer3DG6DOFJointAxisFlag) bool {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&axis), gdc.ConstTypePtr(&flag)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&axis)
	pinner.Pin(&flag)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnGeneric6DofJointGetFlag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer3D) FreeRid(rid RID) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnFreeRid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) SetActive(active bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnSetActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer3D) GetProcessInfo(process_info PhysicsServer3DProcessInfo) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&process_info)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&process_info)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer3D.fnGetProcessInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
