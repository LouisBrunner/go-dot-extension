// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationServer3D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationServer3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationServer3D) BaseClass() string {
  return "NavigationServer3D"
}

type NavigationServer3DProcessInfo int
const (
  NavigationServer3DProcessInfoInfoActiveMaps NavigationServer3DProcessInfo = 0
  NavigationServer3DProcessInfoInfoRegionCount NavigationServer3DProcessInfo = 1
  NavigationServer3DProcessInfoInfoAgentCount NavigationServer3DProcessInfo = 2
  NavigationServer3DProcessInfoInfoLinkCount NavigationServer3DProcessInfo = 3
  NavigationServer3DProcessInfoInfoPolygonCount NavigationServer3DProcessInfo = 4
  NavigationServer3DProcessInfoInfoEdgeCount NavigationServer3DProcessInfo = 5
  NavigationServer3DProcessInfoInfoEdgeMergeCount NavigationServer3DProcessInfo = 6
  NavigationServer3DProcessInfoInfoEdgeConnectionCount NavigationServer3DProcessInfo = 7
  NavigationServer3DProcessInfoInfoEdgeFreeCount NavigationServer3DProcessInfo = 8
)

func  (me *NavigationServer3D) GetMaps() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapCreate() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapSetActive(map_ RID, active bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapIsActive(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapSetUp(map_ RID, up Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetUp(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapSetCellSize(map_ RID, cell_size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetCellSize(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapSetCellHeight(map_ RID, cell_height float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetCellHeight(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapSetUseEdgeConnections(map_ RID, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetUseEdgeConnections(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapSetEdgeConnectionMargin(map_ RID, margin float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetEdgeConnectionMargin(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapSetLinkConnectionRadius(map_ RID, radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetLinkConnectionRadius(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetPath(map_ RID, origin Vector3, destination Vector3, optimize bool, navigation_layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetClosestPointToSegment(map_ RID, start Vector3, end Vector3, use_collision bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetClosestPoint(map_ RID, to_point Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetClosestPointNormal(map_ RID, to_point Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetClosestPointOwner(map_ RID, to_point Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetLinks(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetRegions(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetAgents(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapGetObstacles(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) MapForceUpdate(map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) QueryPath(parameters NavigationPathQueryParameters3D, result NavigationPathQueryResult3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionCreate() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionSetUseEdgeConnections(region RID, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionGetUseEdgeConnections(region RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionSetEnterCost(region RID, enter_cost float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionGetEnterCost(region RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionSetTravelCost(region RID, travel_cost float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionGetTravelCost(region RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionSetOwnerId(region RID, owner_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionGetOwnerId(region RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionOwnsPoint(region RID, point Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionSetMap(region RID, map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionGetMap(region RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionSetNavigationLayers(region RID, navigation_layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionGetNavigationLayers(region RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionSetTransform(region RID, transform Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionSetNavigationMesh(region RID, navigation_mesh NavigationMesh, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionBakeNavigationMesh(navigation_mesh NavigationMesh, root_node Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionGetConnectionsCount(region RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionGetConnectionPathwayStart(region RID, connection int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) RegionGetConnectionPathwayEnd(region RID, connection int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkCreate() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkSetMap(link RID, map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkGetMap(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkSetBidirectional(link RID, bidirectional bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkIsBidirectional(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkSetNavigationLayers(link RID, navigation_layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkGetNavigationLayers(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkSetStartPosition(link RID, position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkGetStartPosition(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkSetEndPosition(link RID, position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkGetEndPosition(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkSetEnterCost(link RID, enter_cost float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkGetEnterCost(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkSetTravelCost(link RID, travel_cost float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkGetTravelCost(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkSetOwnerId(link RID, owner_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) LinkGetOwnerId(link RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentCreate() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetAvoidanceEnabled(agent RID, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentGetAvoidanceEnabled(agent RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetUse3DAvoidance(agent RID, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentGetUse3DAvoidance(agent RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetMap(agent RID, map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentGetMap(agent RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetPaused(agent RID, paused bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentGetPaused(agent RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetNeighborDistance(agent RID, distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetMaxNeighbors(agent RID, count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetTimeHorizonAgents(agent RID, time_horizon float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetTimeHorizonObstacles(agent RID, time_horizon float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetRadius(agent RID, radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetHeight(agent RID, height float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetMaxSpeed(agent RID, max_speed float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetVelocityForced(agent RID, velocity Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetVelocity(agent RID, velocity Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetPosition(agent RID, position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentIsMapChanged(agent RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetAvoidanceCallback(agent RID, callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetAvoidanceLayers(agent RID, layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetAvoidanceMask(agent RID, mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) AgentSetAvoidancePriority(agent RID, priority float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleCreate() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleSetAvoidanceEnabled(obstacle RID, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleGetAvoidanceEnabled(obstacle RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleSetUse3DAvoidance(obstacle RID, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleGetUse3DAvoidance(obstacle RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleSetMap(obstacle RID, map_ RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleGetMap(obstacle RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleSetPaused(obstacle RID, paused bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleGetPaused(obstacle RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleSetRadius(obstacle RID, radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleSetHeight(obstacle RID, height float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleSetVelocity(obstacle RID, velocity Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleSetPosition(obstacle RID, position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleSetVertices(obstacle RID, vertices PackedVector3Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ObstacleSetAvoidanceLayers(obstacle RID, layers int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) ParseSourceGeometryData(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, root_node Node, callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) BakeFromSourceGeometryData(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) FreeRid(rid RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) SetActive(active bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) SetDebugEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) GetDebugEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationServer3D) GetProcessInfo(process_info NavigationServer3DProcessInfo, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
