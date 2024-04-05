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

type ptrsForNavigationAgent3DList struct {
  fnGetRid gdc.MethodBindPtr
  fnSetAvoidanceEnabled gdc.MethodBindPtr
  fnGetAvoidanceEnabled gdc.MethodBindPtr
  fnSetPathDesiredDistance gdc.MethodBindPtr
  fnGetPathDesiredDistance gdc.MethodBindPtr
  fnSetTargetDesiredDistance gdc.MethodBindPtr
  fnGetTargetDesiredDistance gdc.MethodBindPtr
  fnSetRadius gdc.MethodBindPtr
  fnGetRadius gdc.MethodBindPtr
  fnSetHeight gdc.MethodBindPtr
  fnGetHeight gdc.MethodBindPtr
  fnSetPathHeightOffset gdc.MethodBindPtr
  fnGetPathHeightOffset gdc.MethodBindPtr
  fnSetUse3DAvoidance gdc.MethodBindPtr
  fnGetUse3DAvoidance gdc.MethodBindPtr
  fnSetKeepYVelocity gdc.MethodBindPtr
  fnGetKeepYVelocity gdc.MethodBindPtr
  fnSetNeighborDistance gdc.MethodBindPtr
  fnGetNeighborDistance gdc.MethodBindPtr
  fnSetMaxNeighbors gdc.MethodBindPtr
  fnGetMaxNeighbors gdc.MethodBindPtr
  fnSetTimeHorizonAgents gdc.MethodBindPtr
  fnGetTimeHorizonAgents gdc.MethodBindPtr
  fnSetTimeHorizonObstacles gdc.MethodBindPtr
  fnGetTimeHorizonObstacles gdc.MethodBindPtr
  fnSetMaxSpeed gdc.MethodBindPtr
  fnGetMaxSpeed gdc.MethodBindPtr
  fnSetPathMaxDistance gdc.MethodBindPtr
  fnGetPathMaxDistance gdc.MethodBindPtr
  fnSetNavigationLayers gdc.MethodBindPtr
  fnGetNavigationLayers gdc.MethodBindPtr
  fnSetNavigationLayerValue gdc.MethodBindPtr
  fnGetNavigationLayerValue gdc.MethodBindPtr
  fnSetPathfindingAlgorithm gdc.MethodBindPtr
  fnGetPathfindingAlgorithm gdc.MethodBindPtr
  fnSetPathPostprocessing gdc.MethodBindPtr
  fnGetPathPostprocessing gdc.MethodBindPtr
  fnSetPathMetadataFlags gdc.MethodBindPtr
  fnGetPathMetadataFlags gdc.MethodBindPtr
  fnSetNavigationMap gdc.MethodBindPtr
  fnGetNavigationMap gdc.MethodBindPtr
  fnSetTargetPosition gdc.MethodBindPtr
  fnGetTargetPosition gdc.MethodBindPtr
  fnGetNextPathPosition gdc.MethodBindPtr
  fnSetVelocityForced gdc.MethodBindPtr
  fnSetVelocity gdc.MethodBindPtr
  fnGetVelocity gdc.MethodBindPtr
  fnDistanceToTarget gdc.MethodBindPtr
  fnGetCurrentNavigationResult gdc.MethodBindPtr
  fnGetCurrentNavigationPath gdc.MethodBindPtr
  fnGetCurrentNavigationPathIndex gdc.MethodBindPtr
  fnIsTargetReached gdc.MethodBindPtr
  fnIsTargetReachable gdc.MethodBindPtr
  fnIsNavigationFinished gdc.MethodBindPtr
  fnGetFinalPosition gdc.MethodBindPtr
  fnSetAvoidanceLayers gdc.MethodBindPtr
  fnGetAvoidanceLayers gdc.MethodBindPtr
  fnSetAvoidanceMask gdc.MethodBindPtr
  fnGetAvoidanceMask gdc.MethodBindPtr
  fnSetAvoidanceLayerValue gdc.MethodBindPtr
  fnGetAvoidanceLayerValue gdc.MethodBindPtr
  fnSetAvoidanceMaskValue gdc.MethodBindPtr
  fnGetAvoidanceMaskValue gdc.MethodBindPtr
  fnSetAvoidancePriority gdc.MethodBindPtr
  fnGetAvoidancePriority gdc.MethodBindPtr
  fnSetDebugEnabled gdc.MethodBindPtr
  fnGetDebugEnabled gdc.MethodBindPtr
  fnSetDebugUseCustom gdc.MethodBindPtr
  fnGetDebugUseCustom gdc.MethodBindPtr
  fnSetDebugPathCustomColor gdc.MethodBindPtr
  fnGetDebugPathCustomColor gdc.MethodBindPtr
  fnSetDebugPathCustomPointSize gdc.MethodBindPtr
  fnGetDebugPathCustomPointSize gdc.MethodBindPtr
}

var ptrsForNavigationAgent3D ptrsForNavigationAgent3DList

func initNavigationAgent3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("NavigationAgent3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_rid")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
  }
  {
    methodName := StringNameFromStr("set_avoidance_enabled")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_avoidance_enabled")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_path_desired_distance")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetPathDesiredDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_path_desired_distance")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetPathDesiredDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_target_desired_distance")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetTargetDesiredDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_target_desired_distance")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetTargetDesiredDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_radius")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_radius")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_height")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_height")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_path_height_offset")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetPathHeightOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_path_height_offset")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetPathHeightOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_use_3d_avoidance")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetUse3DAvoidance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_use_3d_avoidance")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetUse3DAvoidance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_keep_y_velocity")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetKeepYVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_keep_y_velocity")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetKeepYVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_neighbor_distance")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetNeighborDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_neighbor_distance")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetNeighborDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_max_neighbors")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetMaxNeighbors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_max_neighbors")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetMaxNeighbors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_time_horizon_agents")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetTimeHorizonAgents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_time_horizon_agents")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetTimeHorizonAgents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_time_horizon_obstacles")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetTimeHorizonObstacles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_time_horizon_obstacles")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetTimeHorizonObstacles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_max_speed")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetMaxSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_max_speed")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetMaxSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_path_max_distance")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetPathMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_path_max_distance")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetPathMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
  {
    methodName := StringNameFromStr("set_navigation_layers")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_navigation_layers")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_navigation_layer_value")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetNavigationLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_navigation_layer_value")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetNavigationLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_pathfinding_algorithm")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetPathfindingAlgorithm = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 394560454))
  }
  {
    methodName := StringNameFromStr("get_pathfinding_algorithm")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetPathfindingAlgorithm = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3398491350))
  }
  {
    methodName := StringNameFromStr("set_path_postprocessing")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetPathPostprocessing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2267362344))
  }
  {
    methodName := StringNameFromStr("get_path_postprocessing")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetPathPostprocessing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3883858360))
  }
  {
    methodName := StringNameFromStr("set_path_metadata_flags")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetPathMetadataFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2713846708))
  }
  {
    methodName := StringNameFromStr("get_path_metadata_flags")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetPathMetadataFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1582332802))
  }
  {
    methodName := StringNameFromStr("set_navigation_map")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("get_navigation_map")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
  }
  {
    methodName := StringNameFromStr("set_target_position")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_target_position")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetTargetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("get_next_path_position")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetNextPathPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3783033775))
  }
  {
    methodName := StringNameFromStr("set_velocity_forced")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetVelocityForced = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("set_velocity")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_velocity")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3783033775))
  }
  {
    methodName := StringNameFromStr("distance_to_target")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnDistanceToTarget = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_current_navigation_result")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetCurrentNavigationResult = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 728825684))
  }
  {
    methodName := StringNameFromStr("get_current_navigation_path")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetCurrentNavigationPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 497664490))
  }
  {
    methodName := StringNameFromStr("get_current_navigation_path_index")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetCurrentNavigationPathIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("is_target_reached")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnIsTargetReached = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("is_target_reachable")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnIsTargetReachable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("is_navigation_finished")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnIsNavigationFinished = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("get_final_position")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetFinalPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3783033775))
  }
  {
    methodName := StringNameFromStr("set_avoidance_layers")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetAvoidanceLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_avoidance_layers")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetAvoidanceLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_avoidance_mask")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetAvoidanceMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_avoidance_mask")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetAvoidanceMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_avoidance_layer_value")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetAvoidanceLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_avoidance_layer_value")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetAvoidanceLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_avoidance_mask_value")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetAvoidanceMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_avoidance_mask_value")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetAvoidanceMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_avoidance_priority")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetAvoidancePriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_avoidance_priority")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetAvoidancePriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_debug_enabled")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetDebugEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_debug_enabled")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetDebugEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_debug_use_custom")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetDebugUseCustom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_debug_use_custom")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetDebugUseCustom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_debug_path_custom_color")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetDebugPathCustomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_debug_path_custom_color")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetDebugPathCustomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_debug_path_custom_point_size")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnSetDebugPathCustomPointSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_debug_path_custom_point_size")
    defer methodName.Destroy()
    ptrsForNavigationAgent3D.fnGetDebugPathCustomPointSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

type NavigationAgent3D struct {
  Node
}

func (me *NavigationAgent3D) BaseClass() string {
  return "NavigationAgent3D"
}

func NewNavigationAgent3D() *NavigationAgent3D {
  str := StringNameFromStr("NavigationAgent3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NavigationAgent3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *NavigationAgent3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationAgent3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationAgent3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationAgent3D) GetRid() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) SetAvoidanceEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetAvoidanceEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetPathDesiredDistance(desired_distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&desired_distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetPathDesiredDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetPathDesiredDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetPathDesiredDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetTargetDesiredDistance(desired_distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&desired_distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetTargetDesiredDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetTargetDesiredDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetTargetDesiredDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetRadius(radius float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetRadius() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetHeight(height float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetHeight() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetPathHeightOffset(path_height_offset float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_height_offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetPathHeightOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetPathHeightOffset() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetPathHeightOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetUse3DAvoidance(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetUse3DAvoidance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetUse3DAvoidance() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetUse3DAvoidance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetKeepYVelocity(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetKeepYVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetKeepYVelocity() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetKeepYVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetNeighborDistance(neighbor_distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&neighbor_distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetNeighborDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetNeighborDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetNeighborDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetMaxNeighbors(max_neighbors int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_neighbors) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetMaxNeighbors), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetMaxNeighbors() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetMaxNeighbors), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetTimeHorizonAgents(time_horizon float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_horizon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetTimeHorizonAgents), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetTimeHorizonAgents() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetTimeHorizonAgents), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetTimeHorizonObstacles(time_horizon float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_horizon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetTimeHorizonObstacles), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetTimeHorizonObstacles() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetTimeHorizonObstacles), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetMaxSpeed(max_speed float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetMaxSpeed), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetMaxSpeed() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetMaxSpeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetPathMaxDistance(max_speed float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetPathMaxDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetPathMaxDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetPathMaxDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetNavigationLayers(navigation_layers int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetNavigationLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetNavigationLayers() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetNavigationLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetNavigationLayerValue(layer_number int64, value bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetNavigationLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetNavigationLayerValue(layer_number int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetNavigationLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters3DPathfindingAlgorithm, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pathfinding_algorithm) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetPathfindingAlgorithm), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetPathfindingAlgorithm() NavigationPathQueryParameters3DPathfindingAlgorithm {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationPathQueryParameters3DPathfindingAlgorithm

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetPathfindingAlgorithm), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationAgent3D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters3DPathPostProcessing, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_postprocessing) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetPathPostprocessing), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetPathPostprocessing() NavigationPathQueryParameters3DPathPostProcessing {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationPathQueryParameters3DPathPostProcessing

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetPathPostprocessing), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationAgent3D) SetPathMetadataFlags(flags NavigationPathQueryParameters3DPathMetadataFlags, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetPathMetadataFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetPathMetadataFlags() NavigationPathQueryParameters3DPathMetadataFlags {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationPathQueryParameters3DPathMetadataFlags

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetPathMetadataFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationAgent3D) SetNavigationMap(navigation_map RID, )  {
  cargs := []gdc.ConstTypePtr{navigation_map.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetNavigationMap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetNavigationMap() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetNavigationMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) SetTargetPosition(position Vector3, )  {
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetTargetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetTargetPosition() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetTargetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) GetNextPathPosition() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetNextPathPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) SetVelocityForced(velocity Vector3, )  {
  cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetVelocityForced), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) SetVelocity(velocity Vector3, )  {
  cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetVelocity() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) DistanceToTarget() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnDistanceToTarget), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) GetCurrentNavigationResult() NavigationPathQueryResult3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNavigationPathQueryResult3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetCurrentNavigationResult), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) GetCurrentNavigationPath() PackedVector3Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetCurrentNavigationPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) GetCurrentNavigationPathIndex() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetCurrentNavigationPathIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) IsTargetReached() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnIsTargetReached), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) IsTargetReachable() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnIsTargetReachable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) IsNavigationFinished() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnIsNavigationFinished), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) GetFinalPosition() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetFinalPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) SetAvoidanceLayers(layers int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetAvoidanceLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetAvoidanceLayers() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetAvoidanceLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetAvoidanceMask(mask int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetAvoidanceMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetAvoidanceMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetAvoidanceMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetAvoidanceLayerValue(layer_number int64, value bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetAvoidanceLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetAvoidanceLayerValue(layer_number int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetAvoidanceLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetAvoidanceMaskValue(mask_number int64, value bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetAvoidanceMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetAvoidanceMaskValue(mask_number int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&mask_number)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetAvoidanceMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetAvoidancePriority(priority float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetAvoidancePriority), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetAvoidancePriority() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetAvoidancePriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetDebugEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetDebugEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetDebugEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetDebugEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetDebugUseCustom(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetDebugUseCustom), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetDebugUseCustom() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetDebugUseCustom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetDebugPathCustomColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetDebugPathCustomColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetDebugPathCustomColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetDebugPathCustomColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) SetDebugPathCustomPointSize(point_size float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnSetDebugPathCustomPointSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetDebugPathCustomPointSize() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationAgent3D.fnGetDebugPathCustomPointSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type NavigationAgent3DPathChangedSignalFn func()

func (me *NavigationAgent3D) ConnectPathChanged(subs SignalSubscribers, fn NavigationAgent3DPathChangedSignalFn) {
  sig := StringNameFromStr("path_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent3D) DisconnectPathChanged(subs SignalSubscribers, fn NavigationAgent3DPathChangedSignalFn) {
  sig := StringNameFromStr("path_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent3DTargetReachedSignalFn func()

func (me *NavigationAgent3D) ConnectTargetReached(subs SignalSubscribers, fn NavigationAgent3DTargetReachedSignalFn) {
  sig := StringNameFromStr("target_reached")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent3D) DisconnectTargetReached(subs SignalSubscribers, fn NavigationAgent3DTargetReachedSignalFn) {
  sig := StringNameFromStr("target_reached")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent3DWaypointReachedSignalFn func(details Dictionary, )

func (me *NavigationAgent3D) ConnectWaypointReached(subs SignalSubscribers, fn NavigationAgent3DWaypointReachedSignalFn) {
  sig := StringNameFromStr("waypoint_reached")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent3D) DisconnectWaypointReached(subs SignalSubscribers, fn NavigationAgent3DWaypointReachedSignalFn) {
  sig := StringNameFromStr("waypoint_reached")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent3DLinkReachedSignalFn func(details Dictionary, )

func (me *NavigationAgent3D) ConnectLinkReached(subs SignalSubscribers, fn NavigationAgent3DLinkReachedSignalFn) {
  sig := StringNameFromStr("link_reached")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent3D) DisconnectLinkReached(subs SignalSubscribers, fn NavigationAgent3DLinkReachedSignalFn) {
  sig := StringNameFromStr("link_reached")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent3DNavigationFinishedSignalFn func()

func (me *NavigationAgent3D) ConnectNavigationFinished(subs SignalSubscribers, fn NavigationAgent3DNavigationFinishedSignalFn) {
  sig := StringNameFromStr("navigation_finished")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent3D) DisconnectNavigationFinished(subs SignalSubscribers, fn NavigationAgent3DNavigationFinishedSignalFn) {
  sig := StringNameFromStr("navigation_finished")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent3DVelocityComputedSignalFn func(safe_velocity Vector3, )

func (me *NavigationAgent3D) ConnectVelocityComputed(subs SignalSubscribers, fn NavigationAgent3DVelocityComputedSignalFn) {
  sig := StringNameFromStr("velocity_computed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent3D) DisconnectVelocityComputed(subs SignalSubscribers, fn NavigationAgent3DVelocityComputedSignalFn) {
  sig := StringNameFromStr("velocity_computed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
