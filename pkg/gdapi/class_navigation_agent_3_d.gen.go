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



// Enums

func (me *NavigationAgent3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationAgent3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *NavigationAgent3D) GetRid()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetAvoidanceEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetAvoidanceEnabled()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetPathDesiredDistance(desired_distance float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetPathDesiredDistance()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetTargetDesiredDistance(desired_distance float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetTargetDesiredDistance()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetRadius()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetHeight(height float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetHeight()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetPathHeightOffset(path_height_offset float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetPathHeightOffset()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetUse3DAvoidance(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetUse3DAvoidance()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetNeighborDistance(neighbor_distance float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetNeighborDistance()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetMaxNeighbors(max_neighbors int, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetMaxNeighbors()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetTimeHorizonAgents(time_horizon float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetTimeHorizonAgents()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetTimeHorizonObstacles(time_horizon float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetTimeHorizonObstacles()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetMaxSpeed(max_speed float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetMaxSpeed()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetPathMaxDistance(max_speed float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetPathMaxDistance()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetNavigationLayers(navigation_layers int, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetNavigationLayers()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetNavigationLayerValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetNavigationLayerValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters3DPathfindingAlgorithm, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetPathfindingAlgorithm()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters3DPathPostProcessing, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetPathPostprocessing()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetPathMetadataFlags(flags NavigationPathQueryParameters3DPathMetadataFlags, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetPathMetadataFlags()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetNavigationMap(navigation_map RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetNavigationMap()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetTargetPosition(position Vector3, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetTargetPosition()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetNextPathPosition()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetVelocityForced(velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetVelocity(velocity Vector3, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetVelocity()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) DistanceToTarget()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetCurrentNavigationResult()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetCurrentNavigationPath()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetCurrentNavigationPathIndex()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) IsTargetReached()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) IsTargetReachable()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) IsNavigationFinished()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetFinalPosition()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetAvoidanceLayers(layers int, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetAvoidanceLayers()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetAvoidanceMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetAvoidanceMask()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetAvoidanceLayerValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetAvoidanceLayerValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetAvoidanceMaskValue(mask_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetAvoidanceMaskValue(mask_number int, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetAvoidancePriority(priority float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetAvoidancePriority()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetDebugEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetDebugEnabled()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetDebugUseCustom(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetDebugUseCustom()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetDebugPathCustomColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetDebugPathCustomColor()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) SetDebugPathCustomPointSize(point_size float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent3D) GetDebugPathCustomPointSize()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
