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

func  (me *NavigationServer2D) GetMaps() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapCreate() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapSetActive(map_ RID, active bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapIsActive(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapSetCellSize(map_ RID, cell_size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapGetCellSize(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapSetUseEdgeConnections(map_ RID, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapGetUseEdgeConnections(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapSetEdgeConnectionMargin(map_ RID, margin float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapGetEdgeConnectionMargin(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapSetLinkConnectionRadius(map_ RID, radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapGetLinkConnectionRadius(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapGetPath(map_ RID, origin Vector2, destination Vector2, optimize bool, navigation_layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapGetClosestPoint(map_ RID, to_point Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapGetClosestPointOwner(map_ RID, to_point Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapGetLinks(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapGetRegions(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapGetAgents(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapGetObstacles(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) MapForceUpdate(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) QueryPath(parameters NavigationPathQueryParameters2D, result NavigationPathQueryResult2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionCreate() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionSetUseEdgeConnections(region RID, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionGetUseEdgeConnections(region RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionSetEnterCost(region RID, enter_cost float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionGetEnterCost(region RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionSetTravelCost(region RID, travel_cost float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionGetTravelCost(region RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionSetOwnerId(region RID, owner_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionGetOwnerId(region RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionOwnsPoint(region RID, point Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionSetMap(region RID, map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionGetMap(region RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionSetNavigationLayers(region RID, navigation_layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionGetNavigationLayers(region RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionSetTransform(region RID, transform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionSetNavigationPolygon(region RID, navigation_polygon NavigationPolygon, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionGetConnectionsCount(region RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionGetConnectionPathwayStart(region RID, connection int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) RegionGetConnectionPathwayEnd(region RID, connection int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkCreate() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkSetMap(link RID, map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkGetMap(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkSetBidirectional(link RID, bidirectional bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkIsBidirectional(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkSetNavigationLayers(link RID, navigation_layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkGetNavigationLayers(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkSetStartPosition(link RID, position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkGetStartPosition(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkSetEndPosition(link RID, position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkGetEndPosition(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkSetEnterCost(link RID, enter_cost float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkGetEnterCost(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkSetTravelCost(link RID, travel_cost float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkGetTravelCost(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkSetOwnerId(link RID, owner_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) LinkGetOwnerId(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentCreate() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetAvoidanceEnabled(agent RID, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentGetAvoidanceEnabled(agent RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetMap(agent RID, map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentGetMap(agent RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetPaused(agent RID, paused bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentGetPaused(agent RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetNeighborDistance(agent RID, distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetMaxNeighbors(agent RID, count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetTimeHorizonAgents(agent RID, time_horizon float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetTimeHorizonObstacles(agent RID, time_horizon float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetRadius(agent RID, radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetMaxSpeed(agent RID, max_speed float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetVelocityForced(agent RID, velocity Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetVelocity(agent RID, velocity Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetPosition(agent RID, position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentIsMapChanged(agent RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetAvoidanceCallback(agent RID, callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetAvoidanceLayers(agent RID, layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetAvoidanceMask(agent RID, mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) AgentSetAvoidancePriority(agent RID, priority float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) ObstacleCreate() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) ObstacleSetAvoidanceEnabled(obstacle RID, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) ObstacleGetAvoidanceEnabled(obstacle RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) ObstacleSetMap(obstacle RID, map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) ObstacleGetMap(obstacle RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) ObstacleSetPaused(obstacle RID, paused bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) ObstacleGetPaused(obstacle RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) ObstacleSetRadius(obstacle RID, radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) ObstacleSetVelocity(obstacle RID, velocity Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) ObstacleSetPosition(obstacle RID, position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) ObstacleSetVertices(obstacle RID, vertices PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) ObstacleSetAvoidanceLayers(obstacle RID, layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) FreeRid(rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) SetDebugEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer2D) GetDebugEnabled() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
