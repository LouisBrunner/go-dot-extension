// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationAgent2D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationAgent2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationAgent2D) BaseClass() string {
  return "NavigationAgent2D"
}

func  (me *NavigationAgent2D) GetRid() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetAvoidanceEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetAvoidanceEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetPathDesiredDistance(desired_distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetPathDesiredDistance() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetTargetDesiredDistance(desired_distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetTargetDesiredDistance() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetRadius(radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetRadius() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetNeighborDistance(neighbor_distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetNeighborDistance() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetMaxNeighbors(max_neighbors int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetMaxNeighbors() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetTimeHorizonAgents(time_horizon float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetTimeHorizonAgents() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetTimeHorizonObstacles(time_horizon float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetTimeHorizonObstacles() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetMaxSpeed(max_speed float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetMaxSpeed() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetPathMaxDistance(max_speed float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetPathMaxDistance() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetNavigationLayers(navigation_layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetNavigationLayers() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetNavigationLayerValue(layer_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetNavigationLayerValue(layer_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters2DPathfindingAlgorithm, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetPathfindingAlgorithm() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters2DPathPostProcessing, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetPathPostprocessing() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetPathMetadataFlags(flags NavigationPathQueryParameters2DPathMetadataFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetPathMetadataFlags() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetNavigationMap(navigation_map RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetNavigationMap() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetTargetPosition(position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetTargetPosition() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetNextPathPosition() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetVelocityForced(velocity Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetVelocity(velocity Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetVelocity() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) DistanceToTarget() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetCurrentNavigationResult() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetCurrentNavigationPath() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetCurrentNavigationPathIndex() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) IsTargetReached() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) IsTargetReachable() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) IsNavigationFinished() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetFinalPosition() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetAvoidanceLayers(layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetAvoidanceLayers() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetAvoidanceMask(mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetAvoidanceMask() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetAvoidanceLayerValue(layer_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetAvoidanceLayerValue(layer_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetAvoidanceMaskValue(mask_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetAvoidanceMaskValue(mask_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetAvoidancePriority(priority float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetAvoidancePriority() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetDebugEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetDebugEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetDebugUseCustom(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetDebugUseCustom() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetDebugPathCustomColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetDebugPathCustomColor() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetDebugPathCustomPointSize(point_size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetDebugPathCustomPointSize() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) SetDebugPathCustomLineWidth(line_width float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationAgent2D) GetDebugPathCustomLineWidth() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
