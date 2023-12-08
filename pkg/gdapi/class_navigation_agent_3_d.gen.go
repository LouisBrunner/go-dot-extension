// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationAgent3D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationAgent3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationAgent3D) BaseClass() string {
  return "NavigationAgent3D"
}

func  (me *NavigationAgent3D) GetRid() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetAvoidanceEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetAvoidanceEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetPathDesiredDistance(desired_distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetPathDesiredDistance() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetTargetDesiredDistance(desired_distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetTargetDesiredDistance() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetRadius(radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetRadius() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetHeight(height float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetHeight() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetPathHeightOffset(path_height_offset float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetPathHeightOffset() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetUse3DAvoidance(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetUse3DAvoidance() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetNeighborDistance(neighbor_distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetNeighborDistance() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetMaxNeighbors(max_neighbors int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetMaxNeighbors() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetTimeHorizonAgents(time_horizon float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetTimeHorizonAgents() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetTimeHorizonObstacles(time_horizon float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetTimeHorizonObstacles() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetMaxSpeed(max_speed float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetMaxSpeed() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetPathMaxDistance(max_speed float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetPathMaxDistance() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetNavigationLayers(navigation_layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetNavigationLayers() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetNavigationLayerValue(layer_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetNavigationLayerValue(layer_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters3DPathfindingAlgorithm, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetPathfindingAlgorithm() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters3DPathPostProcessing, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetPathPostprocessing() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetPathMetadataFlags(flags NavigationPathQueryParameters3DPathMetadataFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetPathMetadataFlags() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetNavigationMap(navigation_map RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetNavigationMap() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetTargetPosition(position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetTargetPosition() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetNextPathPosition() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetVelocityForced(velocity Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetVelocity(velocity Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) DistanceToTarget() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetCurrentNavigationResult() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetCurrentNavigationPath() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetCurrentNavigationPathIndex() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) IsTargetReached() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) IsTargetReachable() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) IsNavigationFinished() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetFinalPosition() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetAvoidanceLayers(layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetAvoidanceLayers() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetAvoidanceMask(mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetAvoidanceMask() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetAvoidanceLayerValue(layer_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetAvoidanceLayerValue(layer_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetAvoidanceMaskValue(mask_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetAvoidanceMaskValue(mask_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetAvoidancePriority(priority float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetAvoidancePriority() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetDebugEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetDebugEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetDebugUseCustom(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetDebugUseCustom() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetDebugPathCustomColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetDebugPathCustomColor() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) SetDebugPathCustomPointSize(point_size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent3D) GetDebugPathCustomPointSize() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
