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

type ptrsForNavigationAgent2DList struct {
	fnGetRid                        gdc.MethodBindPtr
	fnSetAvoidanceEnabled           gdc.MethodBindPtr
	fnGetAvoidanceEnabled           gdc.MethodBindPtr
	fnSetPathDesiredDistance        gdc.MethodBindPtr
	fnGetPathDesiredDistance        gdc.MethodBindPtr
	fnSetTargetDesiredDistance      gdc.MethodBindPtr
	fnGetTargetDesiredDistance      gdc.MethodBindPtr
	fnSetRadius                     gdc.MethodBindPtr
	fnGetRadius                     gdc.MethodBindPtr
	fnSetNeighborDistance           gdc.MethodBindPtr
	fnGetNeighborDistance           gdc.MethodBindPtr
	fnSetMaxNeighbors               gdc.MethodBindPtr
	fnGetMaxNeighbors               gdc.MethodBindPtr
	fnSetTimeHorizonAgents          gdc.MethodBindPtr
	fnGetTimeHorizonAgents          gdc.MethodBindPtr
	fnSetTimeHorizonObstacles       gdc.MethodBindPtr
	fnGetTimeHorizonObstacles       gdc.MethodBindPtr
	fnSetMaxSpeed                   gdc.MethodBindPtr
	fnGetMaxSpeed                   gdc.MethodBindPtr
	fnSetPathMaxDistance            gdc.MethodBindPtr
	fnGetPathMaxDistance            gdc.MethodBindPtr
	fnSetNavigationLayers           gdc.MethodBindPtr
	fnGetNavigationLayers           gdc.MethodBindPtr
	fnSetNavigationLayerValue       gdc.MethodBindPtr
	fnGetNavigationLayerValue       gdc.MethodBindPtr
	fnSetPathfindingAlgorithm       gdc.MethodBindPtr
	fnGetPathfindingAlgorithm       gdc.MethodBindPtr
	fnSetPathPostprocessing         gdc.MethodBindPtr
	fnGetPathPostprocessing         gdc.MethodBindPtr
	fnSetPathMetadataFlags          gdc.MethodBindPtr
	fnGetPathMetadataFlags          gdc.MethodBindPtr
	fnSetNavigationMap              gdc.MethodBindPtr
	fnGetNavigationMap              gdc.MethodBindPtr
	fnSetTargetPosition             gdc.MethodBindPtr
	fnGetTargetPosition             gdc.MethodBindPtr
	fnGetNextPathPosition           gdc.MethodBindPtr
	fnSetVelocityForced             gdc.MethodBindPtr
	fnSetVelocity                   gdc.MethodBindPtr
	fnGetVelocity                   gdc.MethodBindPtr
	fnDistanceToTarget              gdc.MethodBindPtr
	fnGetCurrentNavigationResult    gdc.MethodBindPtr
	fnGetCurrentNavigationPath      gdc.MethodBindPtr
	fnGetCurrentNavigationPathIndex gdc.MethodBindPtr
	fnIsTargetReached               gdc.MethodBindPtr
	fnIsTargetReachable             gdc.MethodBindPtr
	fnIsNavigationFinished          gdc.MethodBindPtr
	fnGetFinalPosition              gdc.MethodBindPtr
	fnSetAvoidanceLayers            gdc.MethodBindPtr
	fnGetAvoidanceLayers            gdc.MethodBindPtr
	fnSetAvoidanceMask              gdc.MethodBindPtr
	fnGetAvoidanceMask              gdc.MethodBindPtr
	fnSetAvoidanceLayerValue        gdc.MethodBindPtr
	fnGetAvoidanceLayerValue        gdc.MethodBindPtr
	fnSetAvoidanceMaskValue         gdc.MethodBindPtr
	fnGetAvoidanceMaskValue         gdc.MethodBindPtr
	fnSetAvoidancePriority          gdc.MethodBindPtr
	fnGetAvoidancePriority          gdc.MethodBindPtr
	fnSetDebugEnabled               gdc.MethodBindPtr
	fnGetDebugEnabled               gdc.MethodBindPtr
	fnSetDebugUseCustom             gdc.MethodBindPtr
	fnGetDebugUseCustom             gdc.MethodBindPtr
	fnSetDebugPathCustomColor       gdc.MethodBindPtr
	fnGetDebugPathCustomColor       gdc.MethodBindPtr
	fnSetDebugPathCustomPointSize   gdc.MethodBindPtr
	fnGetDebugPathCustomPointSize   gdc.MethodBindPtr
	fnSetDebugPathCustomLineWidth   gdc.MethodBindPtr
	fnGetDebugPathCustomLineWidth   gdc.MethodBindPtr
}

var ptrsForNavigationAgent2D ptrsForNavigationAgent2DList

func initNavigationAgent2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NavigationAgent2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_rid")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_avoidance_enabled")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_avoidance_enabled")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_path_desired_distance")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetPathDesiredDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_path_desired_distance")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetPathDesiredDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_target_desired_distance")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetTargetDesiredDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_target_desired_distance")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetTargetDesiredDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_radius")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_radius")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_neighbor_distance")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetNeighborDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_neighbor_distance")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetNeighborDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_max_neighbors")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetMaxNeighbors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_max_neighbors")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetMaxNeighbors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_time_horizon_agents")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetTimeHorizonAgents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_time_horizon_agents")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetTimeHorizonAgents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_time_horizon_obstacles")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetTimeHorizonObstacles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_time_horizon_obstacles")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetTimeHorizonObstacles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_max_speed")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetMaxSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_max_speed")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetMaxSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_path_max_distance")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetPathMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_path_max_distance")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetPathMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_navigation_layer_value")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetNavigationLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_navigation_layer_value")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetNavigationLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_pathfinding_algorithm")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetPathfindingAlgorithm = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2783519915))
	}
	{
		methodName := StringNameFromStr("get_pathfinding_algorithm")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetPathfindingAlgorithm = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3000421146))
	}
	{
		methodName := StringNameFromStr("set_path_postprocessing")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetPathPostprocessing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2864409082))
	}
	{
		methodName := StringNameFromStr("get_path_postprocessing")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetPathPostprocessing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3798118993))
	}
	{
		methodName := StringNameFromStr("set_path_metadata_flags")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetPathMetadataFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 24274129))
	}
	{
		methodName := StringNameFromStr("get_path_metadata_flags")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetPathMetadataFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 488152976))
	}
	{
		methodName := StringNameFromStr("set_navigation_map")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("get_navigation_map")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_target_position")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_target_position")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_next_path_position")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetNextPathPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1497962370))
	}
	{
		methodName := StringNameFromStr("set_velocity_forced")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetVelocityForced = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("set_velocity")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_velocity")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1497962370))
	}
	{
		methodName := StringNameFromStr("distance_to_target")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnDistanceToTarget = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_current_navigation_result")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetCurrentNavigationResult = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166799483))
	}
	{
		methodName := StringNameFromStr("get_current_navigation_path")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetCurrentNavigationPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
	}
	{
		methodName := StringNameFromStr("get_current_navigation_path_index")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetCurrentNavigationPathIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("is_target_reached")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnIsTargetReached = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_target_reachable")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnIsTargetReachable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("is_navigation_finished")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnIsNavigationFinished = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("get_final_position")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetFinalPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1497962370))
	}
	{
		methodName := StringNameFromStr("set_avoidance_layers")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetAvoidanceLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_avoidance_layers")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetAvoidanceLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_avoidance_mask")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetAvoidanceMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_avoidance_mask")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetAvoidanceMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_avoidance_layer_value")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetAvoidanceLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_avoidance_layer_value")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetAvoidanceLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_avoidance_mask_value")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetAvoidanceMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_avoidance_mask_value")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetAvoidanceMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_avoidance_priority")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetAvoidancePriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_avoidance_priority")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetAvoidancePriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_debug_enabled")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetDebugEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_debug_enabled")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetDebugEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_debug_use_custom")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetDebugUseCustom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_debug_use_custom")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetDebugUseCustom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_debug_path_custom_color")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetDebugPathCustomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_debug_path_custom_color")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetDebugPathCustomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_debug_path_custom_point_size")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetDebugPathCustomPointSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_debug_path_custom_point_size")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetDebugPathCustomPointSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_debug_path_custom_line_width")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnSetDebugPathCustomLineWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_debug_path_custom_line_width")
		defer methodName.Destroy()
		ptrsForNavigationAgent2D.fnGetDebugPathCustomLineWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type NavigationAgent2D struct {
	Node
}

func (me *NavigationAgent2D) BaseClass() string {
	return "NavigationAgent2D"
}

func NewNavigationAgent2D() *NavigationAgent2D {
	str := StringNameFromStr("NavigationAgent2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &NavigationAgent2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *NavigationAgent2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *NavigationAgent2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *NavigationAgent2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *NavigationAgent2D) GetRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationAgent2D) SetAvoidanceEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetAvoidanceEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetPathDesiredDistance(desired_distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&desired_distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetPathDesiredDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetPathDesiredDistance() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetPathDesiredDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetTargetDesiredDistance(desired_distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&desired_distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetTargetDesiredDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetTargetDesiredDistance() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetTargetDesiredDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetNeighborDistance(neighbor_distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&neighbor_distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetNeighborDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetNeighborDistance() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetNeighborDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetMaxNeighbors(max_neighbors int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_neighbors)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetMaxNeighbors), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetMaxNeighbors() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetMaxNeighbors), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetTimeHorizonAgents(time_horizon float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_horizon)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetTimeHorizonAgents), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetTimeHorizonAgents() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetTimeHorizonAgents), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetTimeHorizonObstacles(time_horizon float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_horizon)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetTimeHorizonObstacles), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetTimeHorizonObstacles() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetTimeHorizonObstacles), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetMaxSpeed(max_speed float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_speed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetMaxSpeed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetMaxSpeed() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetMaxSpeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetPathMaxDistance(max_speed float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_speed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetPathMaxDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetPathMaxDistance() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetPathMaxDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetNavigationLayers(navigation_layers int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetNavigationLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetNavigationLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetNavigationLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetNavigationLayerValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetNavigationLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetNavigationLayerValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetNavigationLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters2DPathfindingAlgorithm) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pathfinding_algorithm)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetPathfindingAlgorithm), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetPathfindingAlgorithm() NavigationPathQueryParameters2DPathfindingAlgorithm {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret NavigationPathQueryParameters2DPathfindingAlgorithm

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetPathfindingAlgorithm), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *NavigationAgent2D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters2DPathPostProcessing) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_postprocessing)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetPathPostprocessing), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetPathPostprocessing() NavigationPathQueryParameters2DPathPostProcessing {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret NavigationPathQueryParameters2DPathPostProcessing

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetPathPostprocessing), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *NavigationAgent2D) SetPathMetadataFlags(flags NavigationPathQueryParameters2DPathMetadataFlags) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetPathMetadataFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetPathMetadataFlags() NavigationPathQueryParameters2DPathMetadataFlags {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret NavigationPathQueryParameters2DPathMetadataFlags

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetPathMetadataFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *NavigationAgent2D) SetNavigationMap(navigation_map RID) {
	cargs := []gdc.ConstTypePtr{navigation_map.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetNavigationMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetNavigationMap() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetNavigationMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationAgent2D) SetTargetPosition(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetTargetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetTargetPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetTargetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationAgent2D) GetNextPathPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetNextPathPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationAgent2D) SetVelocityForced(velocity Vector2) {
	cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetVelocityForced), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) SetVelocity(velocity Vector2) {
	cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetVelocity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationAgent2D) DistanceToTarget() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnDistanceToTarget), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) GetCurrentNavigationResult() NavigationPathQueryResult2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNavigationPathQueryResult2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetCurrentNavigationResult), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationAgent2D) GetCurrentNavigationPath() PackedVector2Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetCurrentNavigationPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationAgent2D) GetCurrentNavigationPathIndex() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetCurrentNavigationPathIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) IsTargetReached() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnIsTargetReached), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) IsTargetReachable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnIsTargetReachable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) IsNavigationFinished() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnIsNavigationFinished), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) GetFinalPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetFinalPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationAgent2D) SetAvoidanceLayers(layers int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetAvoidanceLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetAvoidanceLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetAvoidanceLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetAvoidanceMask(mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetAvoidanceMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetAvoidanceMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetAvoidanceMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetAvoidanceLayerValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetAvoidanceLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetAvoidanceLayerValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetAvoidanceLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetAvoidanceMaskValue(mask_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetAvoidanceMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetAvoidanceMaskValue(mask_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&mask_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetAvoidanceMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetAvoidancePriority(priority float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetAvoidancePriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetAvoidancePriority() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetAvoidancePriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetDebugEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetDebugEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetDebugEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetDebugEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetDebugUseCustom(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetDebugUseCustom), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetDebugUseCustom() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetDebugUseCustom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetDebugPathCustomColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetDebugPathCustomColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetDebugPathCustomColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetDebugPathCustomColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationAgent2D) SetDebugPathCustomPointSize(point_size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetDebugPathCustomPointSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetDebugPathCustomPointSize() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetDebugPathCustomPointSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationAgent2D) SetDebugPathCustomLineWidth(line_width float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line_width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnSetDebugPathCustomLineWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationAgent2D) GetDebugPathCustomLineWidth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent2D.fnGetDebugPathCustomLineWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type NavigationAgent2DPathChangedSignalFn func()

func (me *NavigationAgent2D) ConnectPathChanged(subs SignalSubscribers, fn NavigationAgent2DPathChangedSignalFn) {
	sig := StringNameFromStr("path_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent2D) DisconnectPathChanged(subs SignalSubscribers, fn NavigationAgent2DPathChangedSignalFn) {
	sig := StringNameFromStr("path_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent2DTargetReachedSignalFn func()

func (me *NavigationAgent2D) ConnectTargetReached(subs SignalSubscribers, fn NavigationAgent2DTargetReachedSignalFn) {
	sig := StringNameFromStr("target_reached")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent2D) DisconnectTargetReached(subs SignalSubscribers, fn NavigationAgent2DTargetReachedSignalFn) {
	sig := StringNameFromStr("target_reached")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent2DWaypointReachedSignalFn func(details Dictionary)

func (me *NavigationAgent2D) ConnectWaypointReached(subs SignalSubscribers, fn NavigationAgent2DWaypointReachedSignalFn) {
	sig := StringNameFromStr("waypoint_reached")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent2D) DisconnectWaypointReached(subs SignalSubscribers, fn NavigationAgent2DWaypointReachedSignalFn) {
	sig := StringNameFromStr("waypoint_reached")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent2DLinkReachedSignalFn func(details Dictionary)

func (me *NavigationAgent2D) ConnectLinkReached(subs SignalSubscribers, fn NavigationAgent2DLinkReachedSignalFn) {
	sig := StringNameFromStr("link_reached")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent2D) DisconnectLinkReached(subs SignalSubscribers, fn NavigationAgent2DLinkReachedSignalFn) {
	sig := StringNameFromStr("link_reached")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent2DNavigationFinishedSignalFn func()

func (me *NavigationAgent2D) ConnectNavigationFinished(subs SignalSubscribers, fn NavigationAgent2DNavigationFinishedSignalFn) {
	sig := StringNameFromStr("navigation_finished")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent2D) DisconnectNavigationFinished(subs SignalSubscribers, fn NavigationAgent2DNavigationFinishedSignalFn) {
	sig := StringNameFromStr("navigation_finished")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent2DVelocityComputedSignalFn func(safe_velocity Vector2)

func (me *NavigationAgent2D) ConnectVelocityComputed(subs SignalSubscribers, fn NavigationAgent2DVelocityComputedSignalFn) {
	sig := StringNameFromStr("velocity_computed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent2D) DisconnectVelocityComputed(subs SignalSubscribers, fn NavigationAgent2DVelocityComputedSignalFn) {
	sig := StringNameFromStr("velocity_computed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
