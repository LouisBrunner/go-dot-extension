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

func  (me *NavigationAgent2D) GetRid()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetAvoidanceEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetAvoidanceEnabled()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetPathDesiredDistance(desired_distance float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetPathDesiredDistance()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetTargetDesiredDistance(desired_distance float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetTargetDesiredDistance()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetRadius()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetNeighborDistance(neighbor_distance float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetNeighborDistance()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetMaxNeighbors(max_neighbors int, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetMaxNeighbors()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetTimeHorizonAgents(time_horizon float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetTimeHorizonAgents()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetTimeHorizonObstacles(time_horizon float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetTimeHorizonObstacles()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetMaxSpeed(max_speed float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetMaxSpeed()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetPathMaxDistance(max_speed float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetPathMaxDistance()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetNavigationLayers(navigation_layers int, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetNavigationLayers()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetNavigationLayerValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetNavigationLayerValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters2DPathfindingAlgorithm, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetPathfindingAlgorithm()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters2DPathPostProcessing, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetPathPostprocessing()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetPathMetadataFlags(flags NavigationPathQueryParameters2DPathMetadataFlags, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetPathMetadataFlags()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetNavigationMap(navigation_map RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetNavigationMap()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetTargetPosition(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetTargetPosition()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetNextPathPosition()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetVelocityForced(velocity Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetVelocity(velocity Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetVelocity()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) DistanceToTarget()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetCurrentNavigationResult()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetCurrentNavigationPath()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetCurrentNavigationPathIndex()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) IsTargetReached()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) IsTargetReachable()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) IsNavigationFinished()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetFinalPosition()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetAvoidanceLayers(layers int, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetAvoidanceLayers()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetAvoidanceMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetAvoidanceMask()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetAvoidanceLayerValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetAvoidanceLayerValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetAvoidanceMaskValue(mask_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetAvoidanceMaskValue(mask_number int, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetAvoidancePriority(priority float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetAvoidancePriority()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetDebugEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetDebugEnabled()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetDebugUseCustom(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetDebugUseCustom()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetDebugPathCustomColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetDebugPathCustomColor()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetDebugPathCustomPointSize(point_size float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetDebugPathCustomPointSize()  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) SetDebugPathCustomLineWidth(line_width float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationAgent2D) GetDebugPathCustomLineWidth()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
