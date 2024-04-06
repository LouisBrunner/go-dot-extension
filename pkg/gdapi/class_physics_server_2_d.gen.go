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

type ptrsForPhysicsServer2DList struct {
	fnWorldBoundaryShapeCreate                gdc.MethodBindPtr
	fnSeparationRayShapeCreate                gdc.MethodBindPtr
	fnSegmentShapeCreate                      gdc.MethodBindPtr
	fnCircleShapeCreate                       gdc.MethodBindPtr
	fnRectangleShapeCreate                    gdc.MethodBindPtr
	fnCapsuleShapeCreate                      gdc.MethodBindPtr
	fnConvexPolygonShapeCreate                gdc.MethodBindPtr
	fnConcavePolygonShapeCreate               gdc.MethodBindPtr
	fnShapeSetData                            gdc.MethodBindPtr
	fnShapeGetType                            gdc.MethodBindPtr
	fnShapeGetData                            gdc.MethodBindPtr
	fnSpaceCreate                             gdc.MethodBindPtr
	fnSpaceSetActive                          gdc.MethodBindPtr
	fnSpaceIsActive                           gdc.MethodBindPtr
	fnSpaceSetParam                           gdc.MethodBindPtr
	fnSpaceGetParam                           gdc.MethodBindPtr
	fnSpaceGetDirectState                     gdc.MethodBindPtr
	fnAreaCreate                              gdc.MethodBindPtr
	fnAreaSetSpace                            gdc.MethodBindPtr
	fnAreaGetSpace                            gdc.MethodBindPtr
	fnAreaAddShape                            gdc.MethodBindPtr
	fnAreaSetShape                            gdc.MethodBindPtr
	fnAreaSetShapeTransform                   gdc.MethodBindPtr
	fnAreaSetShapeDisabled                    gdc.MethodBindPtr
	fnAreaGetShapeCount                       gdc.MethodBindPtr
	fnAreaGetShape                            gdc.MethodBindPtr
	fnAreaGetShapeTransform                   gdc.MethodBindPtr
	fnAreaRemoveShape                         gdc.MethodBindPtr
	fnAreaClearShapes                         gdc.MethodBindPtr
	fnAreaSetCollisionLayer                   gdc.MethodBindPtr
	fnAreaGetCollisionLayer                   gdc.MethodBindPtr
	fnAreaSetCollisionMask                    gdc.MethodBindPtr
	fnAreaGetCollisionMask                    gdc.MethodBindPtr
	fnAreaSetParam                            gdc.MethodBindPtr
	fnAreaSetTransform                        gdc.MethodBindPtr
	fnAreaGetParam                            gdc.MethodBindPtr
	fnAreaGetTransform                        gdc.MethodBindPtr
	fnAreaAttachObjectInstanceId              gdc.MethodBindPtr
	fnAreaGetObjectInstanceId                 gdc.MethodBindPtr
	fnAreaAttachCanvasInstanceId              gdc.MethodBindPtr
	fnAreaGetCanvasInstanceId                 gdc.MethodBindPtr
	fnAreaSetMonitorCallback                  gdc.MethodBindPtr
	fnAreaSetAreaMonitorCallback              gdc.MethodBindPtr
	fnAreaSetMonitorable                      gdc.MethodBindPtr
	fnBodyCreate                              gdc.MethodBindPtr
	fnBodySetSpace                            gdc.MethodBindPtr
	fnBodyGetSpace                            gdc.MethodBindPtr
	fnBodySetMode                             gdc.MethodBindPtr
	fnBodyGetMode                             gdc.MethodBindPtr
	fnBodyAddShape                            gdc.MethodBindPtr
	fnBodySetShape                            gdc.MethodBindPtr
	fnBodySetShapeTransform                   gdc.MethodBindPtr
	fnBodyGetShapeCount                       gdc.MethodBindPtr
	fnBodyGetShape                            gdc.MethodBindPtr
	fnBodyGetShapeTransform                   gdc.MethodBindPtr
	fnBodyRemoveShape                         gdc.MethodBindPtr
	fnBodyClearShapes                         gdc.MethodBindPtr
	fnBodySetShapeDisabled                    gdc.MethodBindPtr
	fnBodySetShapeAsOneWayCollision           gdc.MethodBindPtr
	fnBodyAttachObjectInstanceId              gdc.MethodBindPtr
	fnBodyGetObjectInstanceId                 gdc.MethodBindPtr
	fnBodyAttachCanvasInstanceId              gdc.MethodBindPtr
	fnBodyGetCanvasInstanceId                 gdc.MethodBindPtr
	fnBodySetContinuousCollisionDetectionMode gdc.MethodBindPtr
	fnBodyGetContinuousCollisionDetectionMode gdc.MethodBindPtr
	fnBodySetCollisionLayer                   gdc.MethodBindPtr
	fnBodyGetCollisionLayer                   gdc.MethodBindPtr
	fnBodySetCollisionMask                    gdc.MethodBindPtr
	fnBodyGetCollisionMask                    gdc.MethodBindPtr
	fnBodySetCollisionPriority                gdc.MethodBindPtr
	fnBodyGetCollisionPriority                gdc.MethodBindPtr
	fnBodySetParam                            gdc.MethodBindPtr
	fnBodyGetParam                            gdc.MethodBindPtr
	fnBodyResetMassProperties                 gdc.MethodBindPtr
	fnBodySetState                            gdc.MethodBindPtr
	fnBodyGetState                            gdc.MethodBindPtr
	fnBodyApplyCentralImpulse                 gdc.MethodBindPtr
	fnBodyApplyTorqueImpulse                  gdc.MethodBindPtr
	fnBodyApplyImpulse                        gdc.MethodBindPtr
	fnBodyApplyCentralForce                   gdc.MethodBindPtr
	fnBodyApplyForce                          gdc.MethodBindPtr
	fnBodyApplyTorque                         gdc.MethodBindPtr
	fnBodyAddConstantCentralForce             gdc.MethodBindPtr
	fnBodyAddConstantForce                    gdc.MethodBindPtr
	fnBodyAddConstantTorque                   gdc.MethodBindPtr
	fnBodySetConstantForce                    gdc.MethodBindPtr
	fnBodyGetConstantForce                    gdc.MethodBindPtr
	fnBodySetConstantTorque                   gdc.MethodBindPtr
	fnBodyGetConstantTorque                   gdc.MethodBindPtr
	fnBodySetAxisVelocity                     gdc.MethodBindPtr
	fnBodyAddCollisionException               gdc.MethodBindPtr
	fnBodyRemoveCollisionException            gdc.MethodBindPtr
	fnBodySetMaxContactsReported              gdc.MethodBindPtr
	fnBodyGetMaxContactsReported              gdc.MethodBindPtr
	fnBodySetOmitForceIntegration             gdc.MethodBindPtr
	fnBodyIsOmittingForceIntegration          gdc.MethodBindPtr
	fnBodySetForceIntegrationCallback         gdc.MethodBindPtr
	fnBodyTestMotion                          gdc.MethodBindPtr
	fnBodyGetDirectState                      gdc.MethodBindPtr
	fnJointCreate                             gdc.MethodBindPtr
	fnJointClear                              gdc.MethodBindPtr
	fnJointSetParam                           gdc.MethodBindPtr
	fnJointGetParam                           gdc.MethodBindPtr
	fnJointDisableCollisionsBetweenBodies     gdc.MethodBindPtr
	fnJointIsDisabledCollisionsBetweenBodies  gdc.MethodBindPtr
	fnJointMakePin                            gdc.MethodBindPtr
	fnJointMakeGroove                         gdc.MethodBindPtr
	fnJointMakeDampedSpring                   gdc.MethodBindPtr
	fnPinJointSetFlag                         gdc.MethodBindPtr
	fnPinJointGetFlag                         gdc.MethodBindPtr
	fnPinJointSetParam                        gdc.MethodBindPtr
	fnPinJointGetParam                        gdc.MethodBindPtr
	fnDampedSpringJointSetParam               gdc.MethodBindPtr
	fnDampedSpringJointGetParam               gdc.MethodBindPtr
	fnJointGetType                            gdc.MethodBindPtr
	fnFreeRid                                 gdc.MethodBindPtr
	fnSetActive                               gdc.MethodBindPtr
	fnGetProcessInfo                          gdc.MethodBindPtr
}

var ptrsForPhysicsServer2D ptrsForPhysicsServer2DList

func initPhysicsServer2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsServer2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("world_boundary_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnWorldBoundaryShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("separation_ray_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnSeparationRayShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("segment_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnSegmentShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("circle_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnCircleShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("rectangle_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnRectangleShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("capsule_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnCapsuleShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("convex_polygon_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnConvexPolygonShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("concave_polygon_shape_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnConcavePolygonShapeCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("shape_set_data")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnShapeSetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3175752987))
	}
	{
		methodName := StringNameFromStr("shape_get_type")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnShapeGetType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1240598777))
	}
	{
		methodName := StringNameFromStr("shape_get_data")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnShapeGetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4171304767))
	}
	{
		methodName := StringNameFromStr("space_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnSpaceCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("space_set_active")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnSpaceSetActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("space_is_active")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnSpaceIsActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("space_set_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnSpaceSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 949194586))
	}
	{
		methodName := StringNameFromStr("space_get_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnSpaceGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 874111783))
	}
	{
		methodName := StringNameFromStr("space_get_direct_state")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnSpaceGetDirectState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3160173886))
	}
	{
		methodName := StringNameFromStr("area_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("area_set_space")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaSetSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
	}
	{
		methodName := StringNameFromStr("area_get_space")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaGetSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
	}
	{
		methodName := StringNameFromStr("area_add_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaAddShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 339056240))
	}
	{
		methodName := StringNameFromStr("area_set_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaSetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2310537182))
	}
	{
		methodName := StringNameFromStr("area_set_shape_transform")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaSetShapeTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 736082694))
	}
	{
		methodName := StringNameFromStr("area_set_shape_disabled")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaSetShapeDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2658558584))
	}
	{
		methodName := StringNameFromStr("area_get_shape_count")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaGetShapeCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("area_get_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1066463050))
	}
	{
		methodName := StringNameFromStr("area_get_shape_transform")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaGetShapeTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1324854622))
	}
	{
		methodName := StringNameFromStr("area_remove_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaRemoveShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("area_clear_shapes")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaClearShapes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("area_set_collision_layer")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaSetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("area_get_collision_layer")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaGetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("area_set_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("area_get_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("area_set_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1257146028))
	}
	{
		methodName := StringNameFromStr("area_set_transform")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1246044741))
	}
	{
		methodName := StringNameFromStr("area_get_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3047435120))
	}
	{
		methodName := StringNameFromStr("area_get_transform")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 213527486))
	}
	{
		methodName := StringNameFromStr("area_attach_object_instance_id")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaAttachObjectInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("area_get_object_instance_id")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaGetObjectInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("area_attach_canvas_instance_id")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaAttachCanvasInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("area_get_canvas_instance_id")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaGetCanvasInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("area_set_monitor_callback")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaSetMonitorCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3379118538))
	}
	{
		methodName := StringNameFromStr("area_set_area_monitor_callback")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaSetAreaMonitorCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3379118538))
	}
	{
		methodName := StringNameFromStr("area_set_monitorable")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnAreaSetMonitorable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("body_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("body_set_space")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
	}
	{
		methodName := StringNameFromStr("body_get_space")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
	}
	{
		methodName := StringNameFromStr("body_set_mode")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1658067650))
	}
	{
		methodName := StringNameFromStr("body_get_mode")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3261702585))
	}
	{
		methodName := StringNameFromStr("body_add_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyAddShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 339056240))
	}
	{
		methodName := StringNameFromStr("body_set_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2310537182))
	}
	{
		methodName := StringNameFromStr("body_set_shape_transform")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetShapeTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 736082694))
	}
	{
		methodName := StringNameFromStr("body_get_shape_count")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetShapeCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("body_get_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1066463050))
	}
	{
		methodName := StringNameFromStr("body_get_shape_transform")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetShapeTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1324854622))
	}
	{
		methodName := StringNameFromStr("body_remove_shape")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyRemoveShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("body_clear_shapes")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyClearShapes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("body_set_shape_disabled")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetShapeDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2658558584))
	}
	{
		methodName := StringNameFromStr("body_set_shape_as_one_way_collision")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetShapeAsOneWayCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2556489974))
	}
	{
		methodName := StringNameFromStr("body_attach_object_instance_id")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyAttachObjectInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("body_get_object_instance_id")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetObjectInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("body_attach_canvas_instance_id")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyAttachCanvasInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("body_get_canvas_instance_id")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetCanvasInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("body_set_continuous_collision_detection_mode")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetContinuousCollisionDetectionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1882257015))
	}
	{
		methodName := StringNameFromStr("body_get_continuous_collision_detection_mode")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetContinuousCollisionDetectionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2661282217))
	}
	{
		methodName := StringNameFromStr("body_set_collision_layer")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("body_get_collision_layer")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetCollisionLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("body_set_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("body_get_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("body_set_collision_priority")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetCollisionPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("body_get_collision_priority")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetCollisionPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("body_set_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2715630609))
	}
	{
		methodName := StringNameFromStr("body_get_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3208033526))
	}
	{
		methodName := StringNameFromStr("body_reset_mass_properties")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyResetMassProperties = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("body_set_state")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1706355209))
	}
	{
		methodName := StringNameFromStr("body_get_state")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4036367961))
	}
	{
		methodName := StringNameFromStr("body_apply_central_impulse")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyApplyCentralImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3201125042))
	}
	{
		methodName := StringNameFromStr("body_apply_torque_impulse")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyApplyTorqueImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("body_apply_impulse")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyApplyImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 205485391))
	}
	{
		methodName := StringNameFromStr("body_apply_central_force")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyApplyCentralForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3201125042))
	}
	{
		methodName := StringNameFromStr("body_apply_force")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyApplyForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 205485391))
	}
	{
		methodName := StringNameFromStr("body_apply_torque")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyApplyTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("body_add_constant_central_force")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyAddConstantCentralForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3201125042))
	}
	{
		methodName := StringNameFromStr("body_add_constant_force")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyAddConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 205485391))
	}
	{
		methodName := StringNameFromStr("body_add_constant_torque")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyAddConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("body_set_constant_force")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3201125042))
	}
	{
		methodName := StringNameFromStr("body_get_constant_force")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetConstantForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2440833711))
	}
	{
		methodName := StringNameFromStr("body_set_constant_torque")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("body_get_constant_torque")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetConstantTorque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("body_set_axis_velocity")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetAxisVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3201125042))
	}
	{
		methodName := StringNameFromStr("body_add_collision_exception")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyAddCollisionException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
	}
	{
		methodName := StringNameFromStr("body_remove_collision_exception")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyRemoveCollisionException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
	}
	{
		methodName := StringNameFromStr("body_set_max_contacts_reported")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetMaxContactsReported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("body_get_max_contacts_reported")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetMaxContactsReported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("body_set_omit_force_integration")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetOmitForceIntegration = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("body_is_omitting_force_integration")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyIsOmittingForceIntegration = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("body_set_force_integration_callback")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodySetForceIntegrationCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3059434249))
	}
	{
		methodName := StringNameFromStr("body_test_motion")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyTestMotion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1699844009))
	}
	{
		methodName := StringNameFromStr("body_get_direct_state")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnBodyGetDirectState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1191931871))
	}
	{
		methodName := StringNameFromStr("joint_create")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnJointCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("joint_clear")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnJointClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("joint_set_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnJointSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3972556514))
	}
	{
		methodName := StringNameFromStr("joint_get_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnJointGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4016448949))
	}
	{
		methodName := StringNameFromStr("joint_disable_collisions_between_bodies")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnJointDisableCollisionsBetweenBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("joint_is_disabled_collisions_between_bodies")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnJointIsDisabledCollisionsBetweenBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("joint_make_pin")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnJointMakePin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1612646186))
	}
	{
		methodName := StringNameFromStr("joint_make_groove")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnJointMakeGroove = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 481430435))
	}
	{
		methodName := StringNameFromStr("joint_make_damped_spring")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnJointMakeDampedSpring = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1994657646))
	}
	{
		methodName := StringNameFromStr("pin_joint_set_flag")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnPinJointSetFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3520002352))
	}
	{
		methodName := StringNameFromStr("pin_joint_get_flag")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnPinJointGetFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2647867364))
	}
	{
		methodName := StringNameFromStr("pin_joint_set_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnPinJointSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 550574241))
	}
	{
		methodName := StringNameFromStr("pin_joint_get_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnPinJointGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 348281383))
	}
	{
		methodName := StringNameFromStr("damped_spring_joint_set_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnDampedSpringJointSetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 220564071))
	}
	{
		methodName := StringNameFromStr("damped_spring_joint_get_param")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnDampedSpringJointGetParam = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2075871277))
	}
	{
		methodName := StringNameFromStr("joint_get_type")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnJointGetType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4262502231))
	}
	{
		methodName := StringNameFromStr("free_rid")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnFreeRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("set_active")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnSetActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_process_info")
		defer methodName.Destroy()
		ptrsForPhysicsServer2D.fnGetProcessInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 576496006))
	}
}

type PhysicsServer2D struct {
	Object
}

func (me *PhysicsServer2D) BaseClass() string {
	return "PhysicsServer2D"
}

func NewPhysicsServer2D() *PhysicsServer2D {
	str := StringNameFromStr("PhysicsServer2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsServer2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type PhysicsServer2DSpaceParameter int

const (
	PhysicsServer2DSpaceParameterSpaceParamContactRecycleRadius              PhysicsServer2DSpaceParameter = 0
	PhysicsServer2DSpaceParameterSpaceParamContactMaxSeparation              PhysicsServer2DSpaceParameter = 1
	PhysicsServer2DSpaceParameterSpaceParamContactMaxAllowedPenetration      PhysicsServer2DSpaceParameter = 2
	PhysicsServer2DSpaceParameterSpaceParamContactDefaultBias                PhysicsServer2DSpaceParameter = 3
	PhysicsServer2DSpaceParameterSpaceParamBodyLinearVelocitySleepThreshold  PhysicsServer2DSpaceParameter = 4
	PhysicsServer2DSpaceParameterSpaceParamBodyAngularVelocitySleepThreshold PhysicsServer2DSpaceParameter = 5
	PhysicsServer2DSpaceParameterSpaceParamBodyTimeToSleep                   PhysicsServer2DSpaceParameter = 6
	PhysicsServer2DSpaceParameterSpaceParamConstraintDefaultBias             PhysicsServer2DSpaceParameter = 7
	PhysicsServer2DSpaceParameterSpaceParamSolverIterations                  PhysicsServer2DSpaceParameter = 8
)

type PhysicsServer2DShapeType int

const (
	PhysicsServer2DShapeTypeShapeWorldBoundary  PhysicsServer2DShapeType = 0
	PhysicsServer2DShapeTypeShapeSeparationRay  PhysicsServer2DShapeType = 1
	PhysicsServer2DShapeTypeShapeSegment        PhysicsServer2DShapeType = 2
	PhysicsServer2DShapeTypeShapeCircle         PhysicsServer2DShapeType = 3
	PhysicsServer2DShapeTypeShapeRectangle      PhysicsServer2DShapeType = 4
	PhysicsServer2DShapeTypeShapeCapsule        PhysicsServer2DShapeType = 5
	PhysicsServer2DShapeTypeShapeConvexPolygon  PhysicsServer2DShapeType = 6
	PhysicsServer2DShapeTypeShapeConcavePolygon PhysicsServer2DShapeType = 7
	PhysicsServer2DShapeTypeShapeCustom         PhysicsServer2DShapeType = 8
)

type PhysicsServer2DAreaParameter int

const (
	PhysicsServer2DAreaParameterAreaParamGravityOverrideMode      PhysicsServer2DAreaParameter = 0
	PhysicsServer2DAreaParameterAreaParamGravity                  PhysicsServer2DAreaParameter = 1
	PhysicsServer2DAreaParameterAreaParamGravityVector            PhysicsServer2DAreaParameter = 2
	PhysicsServer2DAreaParameterAreaParamGravityIsPoint           PhysicsServer2DAreaParameter = 3
	PhysicsServer2DAreaParameterAreaParamGravityPointUnitDistance PhysicsServer2DAreaParameter = 4
	PhysicsServer2DAreaParameterAreaParamLinearDampOverrideMode   PhysicsServer2DAreaParameter = 5
	PhysicsServer2DAreaParameterAreaParamLinearDamp               PhysicsServer2DAreaParameter = 6
	PhysicsServer2DAreaParameterAreaParamAngularDampOverrideMode  PhysicsServer2DAreaParameter = 7
	PhysicsServer2DAreaParameterAreaParamAngularDamp              PhysicsServer2DAreaParameter = 8
	PhysicsServer2DAreaParameterAreaParamPriority                 PhysicsServer2DAreaParameter = 9
)

type PhysicsServer2DAreaSpaceOverrideMode int

const (
	PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideDisabled       PhysicsServer2DAreaSpaceOverrideMode = 0
	PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideCombine        PhysicsServer2DAreaSpaceOverrideMode = 1
	PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideCombineReplace PhysicsServer2DAreaSpaceOverrideMode = 2
	PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideReplace        PhysicsServer2DAreaSpaceOverrideMode = 3
	PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideReplaceCombine PhysicsServer2DAreaSpaceOverrideMode = 4
)

type PhysicsServer2DBodyMode int

const (
	PhysicsServer2DBodyModeBodyModeStatic      PhysicsServer2DBodyMode = 0
	PhysicsServer2DBodyModeBodyModeKinematic   PhysicsServer2DBodyMode = 1
	PhysicsServer2DBodyModeBodyModeRigid       PhysicsServer2DBodyMode = 2
	PhysicsServer2DBodyModeBodyModeRigidLinear PhysicsServer2DBodyMode = 3
)

type PhysicsServer2DBodyParameter int

const (
	PhysicsServer2DBodyParameterBodyParamBounce          PhysicsServer2DBodyParameter = 0
	PhysicsServer2DBodyParameterBodyParamFriction        PhysicsServer2DBodyParameter = 1
	PhysicsServer2DBodyParameterBodyParamMass            PhysicsServer2DBodyParameter = 2
	PhysicsServer2DBodyParameterBodyParamInertia         PhysicsServer2DBodyParameter = 3
	PhysicsServer2DBodyParameterBodyParamCenterOfMass    PhysicsServer2DBodyParameter = 4
	PhysicsServer2DBodyParameterBodyParamGravityScale    PhysicsServer2DBodyParameter = 5
	PhysicsServer2DBodyParameterBodyParamLinearDampMode  PhysicsServer2DBodyParameter = 6
	PhysicsServer2DBodyParameterBodyParamAngularDampMode PhysicsServer2DBodyParameter = 7
	PhysicsServer2DBodyParameterBodyParamLinearDamp      PhysicsServer2DBodyParameter = 8
	PhysicsServer2DBodyParameterBodyParamAngularDamp     PhysicsServer2DBodyParameter = 9
	PhysicsServer2DBodyParameterBodyParamMax             PhysicsServer2DBodyParameter = 10
)

type PhysicsServer2DBodyDampMode int

const (
	PhysicsServer2DBodyDampModeBodyDampModeCombine PhysicsServer2DBodyDampMode = 0
	PhysicsServer2DBodyDampModeBodyDampModeReplace PhysicsServer2DBodyDampMode = 1
)

type PhysicsServer2DBodyState int

const (
	PhysicsServer2DBodyStateBodyStateTransform       PhysicsServer2DBodyState = 0
	PhysicsServer2DBodyStateBodyStateLinearVelocity  PhysicsServer2DBodyState = 1
	PhysicsServer2DBodyStateBodyStateAngularVelocity PhysicsServer2DBodyState = 2
	PhysicsServer2DBodyStateBodyStateSleeping        PhysicsServer2DBodyState = 3
	PhysicsServer2DBodyStateBodyStateCanSleep        PhysicsServer2DBodyState = 4
)

type PhysicsServer2DJointType int

const (
	PhysicsServer2DJointTypeJointTypePin          PhysicsServer2DJointType = 0
	PhysicsServer2DJointTypeJointTypeGroove       PhysicsServer2DJointType = 1
	PhysicsServer2DJointTypeJointTypeDampedSpring PhysicsServer2DJointType = 2
	PhysicsServer2DJointTypeJointTypeMax          PhysicsServer2DJointType = 3
)

type PhysicsServer2DJointParam int

const (
	PhysicsServer2DJointParamJointParamBias     PhysicsServer2DJointParam = 0
	PhysicsServer2DJointParamJointParamMaxBias  PhysicsServer2DJointParam = 1
	PhysicsServer2DJointParamJointParamMaxForce PhysicsServer2DJointParam = 2
)

type PhysicsServer2DPinJointParam int

const (
	PhysicsServer2DPinJointParamPinJointSoftness            PhysicsServer2DPinJointParam = 0
	PhysicsServer2DPinJointParamPinJointLimitUpper          PhysicsServer2DPinJointParam = 1
	PhysicsServer2DPinJointParamPinJointLimitLower          PhysicsServer2DPinJointParam = 2
	PhysicsServer2DPinJointParamPinJointMotorTargetVelocity PhysicsServer2DPinJointParam = 3
)

type PhysicsServer2DPinJointFlag int

const (
	PhysicsServer2DPinJointFlagPinJointFlagAngularLimitEnabled PhysicsServer2DPinJointFlag = 0
	PhysicsServer2DPinJointFlagPinJointFlagMotorEnabled        PhysicsServer2DPinJointFlag = 1
)

type PhysicsServer2DDampedSpringParam int

const (
	PhysicsServer2DDampedSpringParamDampedSpringRestLength PhysicsServer2DDampedSpringParam = 0
	PhysicsServer2DDampedSpringParamDampedSpringStiffness  PhysicsServer2DDampedSpringParam = 1
	PhysicsServer2DDampedSpringParamDampedSpringDamping    PhysicsServer2DDampedSpringParam = 2
)

type PhysicsServer2DCCDMode int

const (
	PhysicsServer2DCCDModeCcdModeDisabled  PhysicsServer2DCCDMode = 0
	PhysicsServer2DCCDModeCcdModeCastRay   PhysicsServer2DCCDMode = 1
	PhysicsServer2DCCDModeCcdModeCastShape PhysicsServer2DCCDMode = 2
)

type PhysicsServer2DAreaBodyStatus int

const (
	PhysicsServer2DAreaBodyStatusAreaBodyAdded   PhysicsServer2DAreaBodyStatus = 0
	PhysicsServer2DAreaBodyStatusAreaBodyRemoved PhysicsServer2DAreaBodyStatus = 1
)

type PhysicsServer2DProcessInfo int

const (
	PhysicsServer2DProcessInfoInfoActiveObjects  PhysicsServer2DProcessInfo = 0
	PhysicsServer2DProcessInfoInfoCollisionPairs PhysicsServer2DProcessInfo = 1
	PhysicsServer2DProcessInfoInfoIslandCount    PhysicsServer2DProcessInfo = 2
)

func (me *PhysicsServer2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsServer2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsServer2D) WorldBoundaryShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnWorldBoundaryShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) SeparationRayShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnSeparationRayShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) SegmentShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnSegmentShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) CircleShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnCircleShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) RectangleShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnRectangleShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) CapsuleShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnCapsuleShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) ConvexPolygonShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnConvexPolygonShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) ConcavePolygonShapeCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnConcavePolygonShapeCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) ShapeSetData(shape RID, data Variant) {
	cargs := []gdc.ConstTypePtr{shape.AsCTypePtr(), data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnShapeSetData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) ShapeGetType(shape RID) PhysicsServer2DShapeType {
	cargs := []gdc.ConstTypePtr{shape.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret PhysicsServer2DShapeType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnShapeGetType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PhysicsServer2D) ShapeGetData(shape RID) Variant {
	cargs := []gdc.ConstTypePtr{shape.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnShapeGetData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) SpaceCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnSpaceCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) SpaceSetActive(space RID, active bool) {
	cargs := []gdc.ConstTypePtr{space.AsCTypePtr(), gdc.ConstTypePtr(&active)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnSpaceSetActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) SpaceIsActive(space RID) bool {
	cargs := []gdc.ConstTypePtr{space.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnSpaceIsActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) SpaceSetParam(space RID, param PhysicsServer2DSpaceParameter, value float64) {
	cargs := []gdc.ConstTypePtr{space.AsCTypePtr(), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnSpaceSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) SpaceGetParam(space RID, param PhysicsServer2DSpaceParameter) float64 {
	cargs := []gdc.ConstTypePtr{space.AsCTypePtr(), gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnSpaceGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) SpaceGetDirectState(space RID) PhysicsDirectSpaceState2D {
	cargs := []gdc.ConstTypePtr{space.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsDirectSpaceState2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnSpaceGetDirectState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) AreaCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) AreaSetSpace(area RID, space RID) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), space.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaSetSpace), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaGetSpace(area RID) RID {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaGetSpace), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) AreaAddShape(area RID, shape RID, transform Transform2D, disabled bool) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), shape.AsCTypePtr(), transform.AsCTypePtr(), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaAddShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaSetShape(area RID, shape_idx int64, shape RID) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx), shape.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaSetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaSetShapeTransform(area RID, shape_idx int64, transform Transform2D) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaSetShapeTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaSetShapeDisabled(area RID, shape_idx int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaSetShapeDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaGetShapeCount(area RID) int64 {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaGetShapeCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) AreaGetShape(area RID, shape_idx int64) RID {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&shape_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaGetShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) AreaGetShapeTransform(area RID, shape_idx int64) Transform2D {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()
	pinner.Pin(&shape_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaGetShapeTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) AreaRemoveShape(area RID, shape_idx int64) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaRemoveShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaClearShapes(area RID) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaClearShapes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaSetCollisionLayer(area RID, layer int64) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaSetCollisionLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaGetCollisionLayer(area RID) int64 {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaGetCollisionLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) AreaSetCollisionMask(area RID, mask int64) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaGetCollisionMask(area RID) int64 {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) AreaSetParam(area RID, param PhysicsServer2DAreaParameter, value Variant) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&param), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaSetTransform(area RID, transform Transform2D) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaGetParam(area RID, param PhysicsServer2DAreaParameter) Variant {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) AreaGetTransform(area RID) Transform2D {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) AreaAttachObjectInstanceId(area RID, id int64) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaAttachObjectInstanceId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaGetObjectInstanceId(area RID) int64 {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaGetObjectInstanceId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) AreaAttachCanvasInstanceId(area RID, id int64) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaAttachCanvasInstanceId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaGetCanvasInstanceId(area RID) int64 {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaGetCanvasInstanceId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) AreaSetMonitorCallback(area RID, callback Callable) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaSetMonitorCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaSetAreaMonitorCallback(area RID, callback Callable) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaSetAreaMonitorCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) AreaSetMonitorable(area RID, monitorable bool) {
	cargs := []gdc.ConstTypePtr{area.AsCTypePtr(), gdc.ConstTypePtr(&monitorable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnAreaSetMonitorable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) BodySetSpace(body RID, space RID) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), space.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetSpace), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyGetSpace(body RID) RID {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetSpace), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) BodySetMode(body RID, mode PhysicsServer2DBodyMode) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyGetMode(body RID) PhysicsServer2DBodyMode {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret PhysicsServer2DBodyMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PhysicsServer2D) BodyAddShape(body RID, shape RID, transform Transform2D, disabled bool) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), shape.AsCTypePtr(), transform.AsCTypePtr(), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyAddShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodySetShape(body RID, shape_idx int64, shape RID) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx), shape.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodySetShapeTransform(body RID, shape_idx int64, transform Transform2D) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetShapeTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyGetShapeCount(body RID) int64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetShapeCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) BodyGetShape(body RID, shape_idx int64) RID {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&shape_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) BodyGetShapeTransform(body RID, shape_idx int64) Transform2D {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()
	pinner.Pin(&shape_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetShapeTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) BodyRemoveShape(body RID, shape_idx int64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyRemoveShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyClearShapes(body RID) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyClearShapes), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodySetShapeDisabled(body RID, shape_idx int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetShapeDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodySetShapeAsOneWayCollision(body RID, shape_idx int64, enable bool, margin float64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(&enable), gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetShapeAsOneWayCollision), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyAttachObjectInstanceId(body RID, id int64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyAttachObjectInstanceId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyGetObjectInstanceId(body RID) int64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetObjectInstanceId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) BodyAttachCanvasInstanceId(body RID, id int64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyAttachCanvasInstanceId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyGetCanvasInstanceId(body RID) int64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetCanvasInstanceId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) BodySetContinuousCollisionDetectionMode(body RID, mode PhysicsServer2DCCDMode) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetContinuousCollisionDetectionMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyGetContinuousCollisionDetectionMode(body RID) PhysicsServer2DCCDMode {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret PhysicsServer2DCCDMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetContinuousCollisionDetectionMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PhysicsServer2D) BodySetCollisionLayer(body RID, layer int64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetCollisionLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyGetCollisionLayer(body RID) int64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetCollisionLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) BodySetCollisionMask(body RID, mask int64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyGetCollisionMask(body RID) int64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) BodySetCollisionPriority(body RID, priority float64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetCollisionPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyGetCollisionPriority(body RID) float64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetCollisionPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) BodySetParam(body RID, param PhysicsServer2DBodyParameter, value Variant) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&param), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyGetParam(body RID, param PhysicsServer2DBodyParameter) Variant {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) BodyResetMassProperties(body RID) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyResetMassProperties), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodySetState(body RID, state PhysicsServer2DBodyState, value Variant) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&state), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetState), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyGetState(body RID, state PhysicsServer2DBodyState) Variant {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&state)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&state)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) BodyApplyCentralImpulse(body RID, impulse Vector2) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), impulse.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyApplyCentralImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyApplyTorqueImpulse(body RID, impulse float64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&impulse)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyApplyTorqueImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyApplyImpulse(body RID, impulse Vector2, position Vector2) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), impulse.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyApplyImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyApplyCentralForce(body RID, force Vector2) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyApplyCentralForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyApplyForce(body RID, force Vector2, position Vector2) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), force.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyApplyForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyApplyTorque(body RID, torque float64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&torque)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyApplyTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyAddConstantCentralForce(body RID, force Vector2) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyAddConstantCentralForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyAddConstantForce(body RID, force Vector2, position Vector2) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), force.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyAddConstantForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyAddConstantTorque(body RID, torque float64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&torque)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyAddConstantTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodySetConstantForce(body RID, force Vector2) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), force.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetConstantForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyGetConstantForce(body RID) Vector2 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetConstantForce), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) BodySetConstantTorque(body RID, torque float64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&torque)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetConstantTorque), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyGetConstantTorque(body RID) float64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetConstantTorque), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) BodySetAxisVelocity(body RID, axis_velocity Vector2) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), axis_velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetAxisVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyAddCollisionException(body RID, excepted_body RID) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), excepted_body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyAddCollisionException), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyRemoveCollisionException(body RID, excepted_body RID) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), excepted_body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyRemoveCollisionException), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodySetMaxContactsReported(body RID, amount int64) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetMaxContactsReported), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyGetMaxContactsReported(body RID) int64 {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetMaxContactsReported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) BodySetOmitForceIntegration(body RID, enable bool) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetOmitForceIntegration), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyIsOmittingForceIntegration(body RID) bool {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyIsOmittingForceIntegration), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) BodySetForceIntegrationCallback(body RID, callable Callable, userdata Variant) {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), callable.AsCTypePtr(), userdata.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodySetForceIntegrationCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) BodyTestMotion(body RID, parameters PhysicsTestMotionParameters2D, result PhysicsTestMotionResult2D) bool {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr(), parameters.AsCTypePtr(), result.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyTestMotion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) BodyGetDirectState(body RID) PhysicsDirectBodyState2D {
	cargs := []gdc.ConstTypePtr{body.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsDirectBodyState2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnBodyGetDirectState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) JointCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnJointCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsServer2D) JointClear(joint RID) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnJointClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) JointSetParam(joint RID, param PhysicsServer2DJointParam, value float64) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnJointSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) JointGetParam(joint RID, param PhysicsServer2DJointParam) float64 {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnJointGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) JointDisableCollisionsBetweenBodies(joint RID, disable bool) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&disable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnJointDisableCollisionsBetweenBodies), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) JointIsDisabledCollisionsBetweenBodies(joint RID) bool {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnJointIsDisabledCollisionsBetweenBodies), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) JointMakePin(joint RID, anchor Vector2, body_a RID, body_b RID) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), anchor.AsCTypePtr(), body_a.AsCTypePtr(), body_b.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnJointMakePin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) JointMakeGroove(joint RID, groove1_a Vector2, groove2_a Vector2, anchor_b Vector2, body_a RID, body_b RID) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), groove1_a.AsCTypePtr(), groove2_a.AsCTypePtr(), anchor_b.AsCTypePtr(), body_a.AsCTypePtr(), body_b.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnJointMakeGroove), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) JointMakeDampedSpring(joint RID, anchor_a Vector2, anchor_b Vector2, body_a RID, body_b RID) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), anchor_a.AsCTypePtr(), anchor_b.AsCTypePtr(), body_a.AsCTypePtr(), body_b.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnJointMakeDampedSpring), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) PinJointSetFlag(joint RID, flag PhysicsServer2DPinJointFlag, enabled bool) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnPinJointSetFlag), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) PinJointGetFlag(joint RID, flag PhysicsServer2DPinJointFlag) bool {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&flag)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&flag)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnPinJointGetFlag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) PinJointSetParam(joint RID, param PhysicsServer2DPinJointParam, value float64) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnPinJointSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) PinJointGetParam(joint RID, param PhysicsServer2DPinJointParam) float64 {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnPinJointGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) DampedSpringJointSetParam(joint RID, param PhysicsServer2DDampedSpringParam, value float64) {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnDampedSpringJointSetParam), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) DampedSpringJointGetParam(joint RID, param PhysicsServer2DDampedSpringParam) float64 {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr(), gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnDampedSpringJointGetParam), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsServer2D) JointGetType(joint RID) PhysicsServer2DJointType {
	cargs := []gdc.ConstTypePtr{joint.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret PhysicsServer2DJointType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnJointGetType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PhysicsServer2D) FreeRid(rid RID) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnFreeRid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) SetActive(active bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnSetActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsServer2D) GetProcessInfo(process_info PhysicsServer2DProcessInfo) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&process_info)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&process_info)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsServer2D.fnGetProcessInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
