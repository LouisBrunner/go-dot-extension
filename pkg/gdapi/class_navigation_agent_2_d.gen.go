// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type NavigationAgent2D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationAgent2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationAgent2D) BaseClass() string {
  return "NavigationAgent2D"
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

func  (me *NavigationAgent2D) GetRid() RID {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetAvoidanceEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_avoidance_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetAvoidanceEnabled() bool {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_avoidance_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetPathDesiredDistance(desired_distance float32, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_desired_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&desired_distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetPathDesiredDistance() float32 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_desired_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetTargetDesiredDistance(desired_distance float32, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_desired_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&desired_distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetTargetDesiredDistance() float32 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_desired_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetRadius(radius float32, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetRadius() float32 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetNeighborDistance(neighbor_distance float32, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_neighbor_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&neighbor_distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetNeighborDistance() float32 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_neighbor_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetMaxNeighbors(max_neighbors int, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_neighbors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_neighbors), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetMaxNeighbors() int {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_neighbors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetTimeHorizonAgents(time_horizon float32, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_time_horizon_agents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_horizon), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetTimeHorizonAgents() float32 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_horizon_agents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetTimeHorizonObstacles(time_horizon float32, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_time_horizon_obstacles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_horizon), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetTimeHorizonObstacles() float32 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_horizon_obstacles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetMaxSpeed(max_speed float32, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_speed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetMaxSpeed() float32 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetPathMaxDistance(max_speed float32, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_speed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetPathMaxDistance() float32 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetNavigationLayers(navigation_layers int, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetNavigationLayers() int {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetNavigationLayerValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetNavigationLayerValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters2DPathfindingAlgorithm, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pathfinding_algorithm")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2783519915) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pathfinding_algorithm), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetPathfindingAlgorithm() NavigationPathQueryParameters2DPathfindingAlgorithm {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pathfinding_algorithm")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3000421146) // FIXME: should cache?
  var ret NavigationPathQueryParameters2DPathfindingAlgorithm
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters2DPathPostProcessing, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_postprocessing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2864409082) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_postprocessing), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetPathPostprocessing() NavigationPathQueryParameters2DPathPostProcessing {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_postprocessing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3798118993) // FIXME: should cache?
  var ret NavigationPathQueryParameters2DPathPostProcessing
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetPathMetadataFlags(flags NavigationPathQueryParameters2DPathMetadataFlags, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_metadata_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 24274129) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetPathMetadataFlags() NavigationPathQueryParameters2DPathMetadataFlags {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_metadata_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 488152976) // FIXME: should cache?
  var ret NavigationPathQueryParameters2DPathMetadataFlags
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetNavigationMap(navigation_map RID, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(navigation_map.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetNavigationMap() RID {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetTargetPosition(position Vector2, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetTargetPosition() Vector2 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) GetNextPathPosition() Vector2 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_path_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1497962370) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetVelocityForced(velocity Vector2, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_velocity_forced")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) SetVelocity(velocity Vector2, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetVelocity() Vector2 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1497962370) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) DistanceToTarget() float32 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("distance_to_target")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) GetCurrentNavigationResult() NavigationPathQueryResult2D {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_navigation_result")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166799483) // FIXME: should cache?
  var ret NavigationPathQueryResult2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) GetCurrentNavigationPath() PackedVector2Array {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_navigation_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) GetCurrentNavigationPathIndex() int {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_navigation_path_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) IsTargetReached() bool {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_target_reached")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) IsTargetReachable() bool {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_target_reachable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) IsNavigationFinished() bool {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_navigation_finished")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) GetFinalPosition() Vector2 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_final_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1497962370) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetAvoidanceLayers(layers int, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_avoidance_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetAvoidanceLayers() int {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_avoidance_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetAvoidanceMask(mask int, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_avoidance_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetAvoidanceMask() int {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_avoidance_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetAvoidanceLayerValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_avoidance_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetAvoidanceLayerValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_avoidance_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetAvoidanceMaskValue(mask_number int, value bool, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_avoidance_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetAvoidanceMaskValue(mask_number int, ) bool {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_avoidance_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetAvoidancePriority(priority float32, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_avoidance_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetAvoidancePriority() float32 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_avoidance_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetDebugEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetDebugEnabled() bool {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetDebugUseCustom(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_use_custom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetDebugUseCustom() bool {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_use_custom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetDebugPathCustomColor(color Color, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_path_custom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetDebugPathCustomColor() Color {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_path_custom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetDebugPathCustomPointSize(point_size float32, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_path_custom_point_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetDebugPathCustomPointSize() float32 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_path_custom_point_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationAgent2D) SetDebugPathCustomLineWidth(line_width float32, )  {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_path_custom_line_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line_width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationAgent2D) GetDebugPathCustomLineWidth() float32 {
  classNameV := StringNameFromStr("NavigationAgent2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_path_custom_line_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *NavigationAgent2D) GetPropTargetPosition() Vector2 {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropTargetPosition(value Vector2) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropPathDesiredDistance() float32 {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropPathDesiredDistance(value float32) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropTargetDesiredDistance() float32 {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropTargetDesiredDistance(value float32) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropPathMaxDistance() float32 {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropPathMaxDistance(value float32) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropNavigationLayers() int {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropNavigationLayers(value int) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropPathfindingAlgorithm() int {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropPathfindingAlgorithm(value int) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropPathPostprocessing() int {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropPathPostprocessing(value int) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropPathMetadataFlags() int {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropPathMetadataFlags(value int) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropAvoidanceEnabled() bool {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropAvoidanceEnabled(value bool) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropVelocity() Vector2 {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropVelocity(value Vector2) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropRadius() float32 {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropRadius(value float32) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropNeighborDistance() float32 {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropNeighborDistance(value float32) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropMaxNeighbors() int {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropMaxNeighbors(value int) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropTimeHorizonAgents() float32 {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropTimeHorizonAgents(value float32) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropTimeHorizonObstacles() float32 {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropTimeHorizonObstacles(value float32) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropMaxSpeed() float32 {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropMaxSpeed(value float32) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropAvoidanceLayers() int {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropAvoidanceLayers(value int) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropAvoidanceMask() int {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropAvoidanceMask(value int) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropAvoidancePriority() float32 {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropAvoidancePriority(value float32) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropDebugEnabled() bool {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropDebugEnabled(value bool) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropDebugUseCustom() bool {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropDebugUseCustom(value bool) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropDebugPathCustomColor() Color {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropDebugPathCustomColor(value Color) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropDebugPathCustomPointSize() float32 {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropDebugPathCustomPointSize(value float32) {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) GetPropDebugPathCustomLineWidth() float32 {
  panic("TODO: implement")
}

func (me *NavigationAgent2D) SetPropDebugPathCustomLineWidth(value float32) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API