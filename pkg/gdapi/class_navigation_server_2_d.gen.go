// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationServer2D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationServer2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationServer2D) BaseClass() string {
  return "NavigationServer2D"
}



// Enums

func (me *NavigationServer2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationServer2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *NavigationServer2D) GetMaps()  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapCreate()  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapSetActive(map_ RID, active bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapIsActive(map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapSetCellSize(map_ RID, cell_size float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapGetCellSize(map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapSetUseEdgeConnections(map_ RID, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapGetUseEdgeConnections(map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapSetEdgeConnectionMargin(map_ RID, margin float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapGetEdgeConnectionMargin(map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapSetLinkConnectionRadius(map_ RID, radius float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapGetLinkConnectionRadius(map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapGetPath(map_ RID, origin Vector2, destination Vector2, optimize bool, navigation_layers int, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapGetClosestPoint(map_ RID, to_point Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapGetClosestPointOwner(map_ RID, to_point Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapGetLinks(map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapGetRegions(map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapGetAgents(map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapGetObstacles(map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) MapForceUpdate(map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) QueryPath(parameters NavigationPathQueryParameters2D, result NavigationPathQueryResult2D, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionCreate()  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionSetUseEdgeConnections(region RID, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionGetUseEdgeConnections(region RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionSetEnterCost(region RID, enter_cost float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionGetEnterCost(region RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionSetTravelCost(region RID, travel_cost float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionGetTravelCost(region RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionSetOwnerId(region RID, owner_id int, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionGetOwnerId(region RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionOwnsPoint(region RID, point Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionSetMap(region RID, map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionGetMap(region RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionSetNavigationLayers(region RID, navigation_layers int, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionGetNavigationLayers(region RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionSetTransform(region RID, transform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionSetNavigationPolygon(region RID, navigation_polygon NavigationPolygon, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionGetConnectionsCount(region RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionGetConnectionPathwayStart(region RID, connection int, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) RegionGetConnectionPathwayEnd(region RID, connection int, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkCreate()  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkSetMap(link RID, map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkGetMap(link RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkSetBidirectional(link RID, bidirectional bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkIsBidirectional(link RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkSetNavigationLayers(link RID, navigation_layers int, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkGetNavigationLayers(link RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkSetStartPosition(link RID, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkGetStartPosition(link RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkSetEndPosition(link RID, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkGetEndPosition(link RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkSetEnterCost(link RID, enter_cost float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkGetEnterCost(link RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkSetTravelCost(link RID, travel_cost float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkGetTravelCost(link RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkSetOwnerId(link RID, owner_id int, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) LinkGetOwnerId(link RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentCreate()  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetAvoidanceEnabled(agent RID, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentGetAvoidanceEnabled(agent RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetMap(agent RID, map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentGetMap(agent RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetPaused(agent RID, paused bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentGetPaused(agent RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetNeighborDistance(agent RID, distance float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetMaxNeighbors(agent RID, count int, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetTimeHorizonAgents(agent RID, time_horizon float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetTimeHorizonObstacles(agent RID, time_horizon float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetRadius(agent RID, radius float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetMaxSpeed(agent RID, max_speed float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetVelocityForced(agent RID, velocity Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetVelocity(agent RID, velocity Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetPosition(agent RID, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentIsMapChanged(agent RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetAvoidanceCallback(agent RID, callback Callable, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetAvoidanceLayers(agent RID, layers int, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetAvoidanceMask(agent RID, mask int, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) AgentSetAvoidancePriority(agent RID, priority float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) ObstacleCreate()  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) ObstacleSetAvoidanceEnabled(obstacle RID, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) ObstacleGetAvoidanceEnabled(obstacle RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) ObstacleSetMap(obstacle RID, map_ RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) ObstacleGetMap(obstacle RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) ObstacleSetPaused(obstacle RID, paused bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) ObstacleGetPaused(obstacle RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) ObstacleSetRadius(obstacle RID, radius float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) ObstacleSetVelocity(obstacle RID, velocity Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) ObstacleSetPosition(obstacle RID, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) ObstacleSetVertices(obstacle RID, vertices PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) ObstacleSetAvoidanceLayers(obstacle RID, layers int, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) FreeRid(rid RID, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) SetDebugEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationServer2D) GetDebugEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
