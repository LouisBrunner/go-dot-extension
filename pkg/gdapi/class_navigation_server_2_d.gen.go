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

type ptrsForNavigationServer2DList struct {
  fnGetMaps gdc.MethodBindPtr
  fnMapCreate gdc.MethodBindPtr
  fnMapSetActive gdc.MethodBindPtr
  fnMapIsActive gdc.MethodBindPtr
  fnMapSetCellSize gdc.MethodBindPtr
  fnMapGetCellSize gdc.MethodBindPtr
  fnMapSetUseEdgeConnections gdc.MethodBindPtr
  fnMapGetUseEdgeConnections gdc.MethodBindPtr
  fnMapSetEdgeConnectionMargin gdc.MethodBindPtr
  fnMapGetEdgeConnectionMargin gdc.MethodBindPtr
  fnMapSetLinkConnectionRadius gdc.MethodBindPtr
  fnMapGetLinkConnectionRadius gdc.MethodBindPtr
  fnMapGetPath gdc.MethodBindPtr
  fnMapGetClosestPoint gdc.MethodBindPtr
  fnMapGetClosestPointOwner gdc.MethodBindPtr
  fnMapGetLinks gdc.MethodBindPtr
  fnMapGetRegions gdc.MethodBindPtr
  fnMapGetAgents gdc.MethodBindPtr
  fnMapGetObstacles gdc.MethodBindPtr
  fnMapForceUpdate gdc.MethodBindPtr
  fnQueryPath gdc.MethodBindPtr
  fnRegionCreate gdc.MethodBindPtr
  fnRegionSetEnabled gdc.MethodBindPtr
  fnRegionGetEnabled gdc.MethodBindPtr
  fnRegionSetUseEdgeConnections gdc.MethodBindPtr
  fnRegionGetUseEdgeConnections gdc.MethodBindPtr
  fnRegionSetEnterCost gdc.MethodBindPtr
  fnRegionGetEnterCost gdc.MethodBindPtr
  fnRegionSetTravelCost gdc.MethodBindPtr
  fnRegionGetTravelCost gdc.MethodBindPtr
  fnRegionSetOwnerId gdc.MethodBindPtr
  fnRegionGetOwnerId gdc.MethodBindPtr
  fnRegionOwnsPoint gdc.MethodBindPtr
  fnRegionSetMap gdc.MethodBindPtr
  fnRegionGetMap gdc.MethodBindPtr
  fnRegionSetNavigationLayers gdc.MethodBindPtr
  fnRegionGetNavigationLayers gdc.MethodBindPtr
  fnRegionSetTransform gdc.MethodBindPtr
  fnRegionSetNavigationPolygon gdc.MethodBindPtr
  fnRegionGetConnectionsCount gdc.MethodBindPtr
  fnRegionGetConnectionPathwayStart gdc.MethodBindPtr
  fnRegionGetConnectionPathwayEnd gdc.MethodBindPtr
  fnLinkCreate gdc.MethodBindPtr
  fnLinkSetMap gdc.MethodBindPtr
  fnLinkGetMap gdc.MethodBindPtr
  fnLinkSetEnabled gdc.MethodBindPtr
  fnLinkGetEnabled gdc.MethodBindPtr
  fnLinkSetBidirectional gdc.MethodBindPtr
  fnLinkIsBidirectional gdc.MethodBindPtr
  fnLinkSetNavigationLayers gdc.MethodBindPtr
  fnLinkGetNavigationLayers gdc.MethodBindPtr
  fnLinkSetStartPosition gdc.MethodBindPtr
  fnLinkGetStartPosition gdc.MethodBindPtr
  fnLinkSetEndPosition gdc.MethodBindPtr
  fnLinkGetEndPosition gdc.MethodBindPtr
  fnLinkSetEnterCost gdc.MethodBindPtr
  fnLinkGetEnterCost gdc.MethodBindPtr
  fnLinkSetTravelCost gdc.MethodBindPtr
  fnLinkGetTravelCost gdc.MethodBindPtr
  fnLinkSetOwnerId gdc.MethodBindPtr
  fnLinkGetOwnerId gdc.MethodBindPtr
  fnAgentCreate gdc.MethodBindPtr
  fnAgentSetAvoidanceEnabled gdc.MethodBindPtr
  fnAgentGetAvoidanceEnabled gdc.MethodBindPtr
  fnAgentSetMap gdc.MethodBindPtr
  fnAgentGetMap gdc.MethodBindPtr
  fnAgentSetPaused gdc.MethodBindPtr
  fnAgentGetPaused gdc.MethodBindPtr
  fnAgentSetNeighborDistance gdc.MethodBindPtr
  fnAgentSetMaxNeighbors gdc.MethodBindPtr
  fnAgentSetTimeHorizonAgents gdc.MethodBindPtr
  fnAgentSetTimeHorizonObstacles gdc.MethodBindPtr
  fnAgentSetRadius gdc.MethodBindPtr
  fnAgentSetMaxSpeed gdc.MethodBindPtr
  fnAgentSetVelocityForced gdc.MethodBindPtr
  fnAgentSetVelocity gdc.MethodBindPtr
  fnAgentSetPosition gdc.MethodBindPtr
  fnAgentIsMapChanged gdc.MethodBindPtr
  fnAgentSetAvoidanceCallback gdc.MethodBindPtr
  fnAgentSetAvoidanceLayers gdc.MethodBindPtr
  fnAgentSetAvoidanceMask gdc.MethodBindPtr
  fnAgentSetAvoidancePriority gdc.MethodBindPtr
  fnObstacleCreate gdc.MethodBindPtr
  fnObstacleSetAvoidanceEnabled gdc.MethodBindPtr
  fnObstacleGetAvoidanceEnabled gdc.MethodBindPtr
  fnObstacleSetMap gdc.MethodBindPtr
  fnObstacleGetMap gdc.MethodBindPtr
  fnObstacleSetPaused gdc.MethodBindPtr
  fnObstacleGetPaused gdc.MethodBindPtr
  fnObstacleSetRadius gdc.MethodBindPtr
  fnObstacleSetVelocity gdc.MethodBindPtr
  fnObstacleSetPosition gdc.MethodBindPtr
  fnObstacleSetVertices gdc.MethodBindPtr
  fnObstacleSetAvoidanceLayers gdc.MethodBindPtr
  fnParseSourceGeometryData gdc.MethodBindPtr
  fnBakeFromSourceGeometryData gdc.MethodBindPtr
  fnBakeFromSourceGeometryDataAsync gdc.MethodBindPtr
  fnFreeRid gdc.MethodBindPtr
  fnSetDebugEnabled gdc.MethodBindPtr
  fnGetDebugEnabled gdc.MethodBindPtr
}

var ptrsForNavigationServer2D ptrsForNavigationServer2DList

func initNavigationServer2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("NavigationServer2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_maps")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnGetMaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("map_create")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("map_set_active")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapSetActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("map_is_active")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapIsActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
  }
  {
    methodName := StringNameFromStr("map_set_cell_size")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapSetCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("map_get_cell_size")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapGetCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
  }
  {
    methodName := StringNameFromStr("map_set_use_edge_connections")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapSetUseEdgeConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("map_get_use_edge_connections")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapGetUseEdgeConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
  }
  {
    methodName := StringNameFromStr("map_set_edge_connection_margin")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapSetEdgeConnectionMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("map_get_edge_connection_margin")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapGetEdgeConnectionMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
  }
  {
    methodName := StringNameFromStr("map_set_link_connection_radius")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapSetLinkConnectionRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("map_get_link_connection_radius")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapGetLinkConnectionRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
  }
  {
    methodName := StringNameFromStr("map_get_path")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapGetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3146466012))
  }
  {
    methodName := StringNameFromStr("map_get_closest_point")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapGetClosestPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1358334418))
  }
  {
    methodName := StringNameFromStr("map_get_closest_point_owner")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapGetClosestPointOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1353467510))
  }
  {
    methodName := StringNameFromStr("map_get_links")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapGetLinks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
  }
  {
    methodName := StringNameFromStr("map_get_regions")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapGetRegions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
  }
  {
    methodName := StringNameFromStr("map_get_agents")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapGetAgents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
  }
  {
    methodName := StringNameFromStr("map_get_obstacles")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapGetObstacles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
  }
  {
    methodName := StringNameFromStr("map_force_update")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnMapForceUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("query_path")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnQueryPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3394638789))
  }
  {
    methodName := StringNameFromStr("region_create")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("region_set_enabled")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("region_get_enabled")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionGetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
  }
  {
    methodName := StringNameFromStr("region_set_use_edge_connections")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionSetUseEdgeConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("region_get_use_edge_connections")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionGetUseEdgeConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
  }
  {
    methodName := StringNameFromStr("region_set_enter_cost")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionSetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("region_get_enter_cost")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionGetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
  }
  {
    methodName := StringNameFromStr("region_set_travel_cost")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionSetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("region_get_travel_cost")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionGetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
  }
  {
    methodName := StringNameFromStr("region_set_owner_id")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionSetOwnerId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("region_get_owner_id")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionGetOwnerId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
  }
  {
    methodName := StringNameFromStr("region_owns_point")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionOwnsPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 219849798))
  }
  {
    methodName := StringNameFromStr("region_set_map")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionSetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("region_get_map")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionGetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
  }
  {
    methodName := StringNameFromStr("region_set_navigation_layers")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionSetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("region_get_navigation_layers")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionGetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
  }
  {
    methodName := StringNameFromStr("region_set_transform")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1246044741))
  }
  {
    methodName := StringNameFromStr("region_set_navigation_polygon")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionSetNavigationPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3633623451))
  }
  {
    methodName := StringNameFromStr("region_get_connections_count")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionGetConnectionsCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
  }
  {
    methodName := StringNameFromStr("region_get_connection_pathway_start")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionGetConnectionPathwayStart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2546185844))
  }
  {
    methodName := StringNameFromStr("region_get_connection_pathway_end")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnRegionGetConnectionPathwayEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2546185844))
  }
  {
    methodName := StringNameFromStr("link_create")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("link_set_map")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkSetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("link_get_map")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkGetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
  }
  {
    methodName := StringNameFromStr("link_set_enabled")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("link_get_enabled")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkGetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
  }
  {
    methodName := StringNameFromStr("link_set_bidirectional")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkSetBidirectional = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("link_is_bidirectional")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkIsBidirectional = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
  }
  {
    methodName := StringNameFromStr("link_set_navigation_layers")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkSetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("link_get_navigation_layers")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkGetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
  }
  {
    methodName := StringNameFromStr("link_set_start_position")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkSetStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3201125042))
  }
  {
    methodName := StringNameFromStr("link_get_start_position")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkGetStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2440833711))
  }
  {
    methodName := StringNameFromStr("link_set_end_position")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkSetEndPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3201125042))
  }
  {
    methodName := StringNameFromStr("link_get_end_position")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkGetEndPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2440833711))
  }
  {
    methodName := StringNameFromStr("link_set_enter_cost")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkSetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("link_get_enter_cost")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkGetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
  }
  {
    methodName := StringNameFromStr("link_set_travel_cost")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkSetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("link_get_travel_cost")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkGetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
  }
  {
    methodName := StringNameFromStr("link_set_owner_id")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkSetOwnerId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("link_get_owner_id")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnLinkGetOwnerId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
  }
  {
    methodName := StringNameFromStr("agent_create")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("agent_set_avoidance_enabled")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("agent_get_avoidance_enabled")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentGetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
  }
  {
    methodName := StringNameFromStr("agent_set_map")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("agent_get_map")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentGetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
  }
  {
    methodName := StringNameFromStr("agent_set_paused")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("agent_get_paused")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentGetPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
  }
  {
    methodName := StringNameFromStr("agent_set_neighbor_distance")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetNeighborDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("agent_set_max_neighbors")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetMaxNeighbors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("agent_set_time_horizon_agents")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetTimeHorizonAgents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("agent_set_time_horizon_obstacles")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetTimeHorizonObstacles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("agent_set_radius")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("agent_set_max_speed")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetMaxSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("agent_set_velocity_forced")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetVelocityForced = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3201125042))
  }
  {
    methodName := StringNameFromStr("agent_set_velocity")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3201125042))
  }
  {
    methodName := StringNameFromStr("agent_set_position")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3201125042))
  }
  {
    methodName := StringNameFromStr("agent_is_map_changed")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentIsMapChanged = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
  }
  {
    methodName := StringNameFromStr("agent_set_avoidance_callback")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetAvoidanceCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3379118538))
  }
  {
    methodName := StringNameFromStr("agent_set_avoidance_layers")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetAvoidanceLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("agent_set_avoidance_mask")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetAvoidanceMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("agent_set_avoidance_priority")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnAgentSetAvoidancePriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("obstacle_create")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnObstacleCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
  }
  {
    methodName := StringNameFromStr("obstacle_set_avoidance_enabled")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnObstacleSetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("obstacle_get_avoidance_enabled")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnObstacleGetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
  }
  {
    methodName := StringNameFromStr("obstacle_set_map")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnObstacleSetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
  }
  {
    methodName := StringNameFromStr("obstacle_get_map")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnObstacleGetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
  }
  {
    methodName := StringNameFromStr("obstacle_set_paused")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnObstacleSetPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
  }
  {
    methodName := StringNameFromStr("obstacle_get_paused")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnObstacleGetPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
  }
  {
    methodName := StringNameFromStr("obstacle_set_radius")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnObstacleSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
  }
  {
    methodName := StringNameFromStr("obstacle_set_velocity")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnObstacleSetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3201125042))
  }
  {
    methodName := StringNameFromStr("obstacle_set_position")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnObstacleSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3201125042))
  }
  {
    methodName := StringNameFromStr("obstacle_set_vertices")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnObstacleSetVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 29476483))
  }
  {
    methodName := StringNameFromStr("obstacle_set_avoidance_layers")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnObstacleSetAvoidanceLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
  }
  {
    methodName := StringNameFromStr("parse_source_geometry_data")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnParseSourceGeometryData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1176164995))
  }
  {
    methodName := StringNameFromStr("bake_from_source_geometry_data")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnBakeFromSourceGeometryData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2909414286))
  }
  {
    methodName := StringNameFromStr("bake_from_source_geometry_data_async")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnBakeFromSourceGeometryDataAsync = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2909414286))
  }
  {
    methodName := StringNameFromStr("free_rid")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnFreeRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("set_debug_enabled")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnSetDebugEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_debug_enabled")
    defer methodName.Destroy()
    ptrsForNavigationServer2D.fnGetDebugEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type NavigationServer2D struct {
  Object
}

func (me *NavigationServer2D) BaseClass() string {
  return "NavigationServer2D"
}

func NewNavigationServer2D() *NavigationServer2D {
  str := StringNameFromStr("NavigationServer2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NavigationServer2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *NavigationServer2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationServer2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationServer2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationServer2D) GetMaps() []RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnGetMaps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[RID](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *NavigationServer2D) MapCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) MapSetActive(map_ RID, active bool, )  {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&active) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapSetActive), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) MapIsActive(map_ RID, ) bool {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapIsActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) MapSetCellSize(map_ RID, cell_size float64, )  {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&cell_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapSetCellSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) MapGetCellSize(map_ RID, ) float64 {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapGetCellSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) MapSetUseEdgeConnections(map_ RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapSetUseEdgeConnections), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) MapGetUseEdgeConnections(map_ RID, ) bool {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapGetUseEdgeConnections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) MapSetEdgeConnectionMargin(map_ RID, margin float64, )  {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapSetEdgeConnectionMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) MapGetEdgeConnectionMargin(map_ RID, ) float64 {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapGetEdgeConnectionMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) MapSetLinkConnectionRadius(map_ RID, radius float64, )  {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapSetLinkConnectionRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) MapGetLinkConnectionRadius(map_ RID, ) float64 {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapGetLinkConnectionRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) MapGetPath(map_ RID, origin Vector2, destination Vector2, optimize bool, navigation_layers int64, ) PackedVector2Array {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), origin.AsCTypePtr(), destination.AsCTypePtr(), gdc.ConstTypePtr(&optimize) , gdc.ConstTypePtr(&navigation_layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()
  pinner.Pin(&optimize)
  pinner.Pin(&navigation_layers)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapGetPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) MapGetClosestPoint(map_ RID, to_point Vector2, ) Vector2 {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), to_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapGetClosestPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) MapGetClosestPointOwner(map_ RID, to_point Vector2, ) RID {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), to_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapGetClosestPointOwner), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) MapGetLinks(map_ RID, ) []RID {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapGetLinks), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[RID](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *NavigationServer2D) MapGetRegions(map_ RID, ) []RID {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapGetRegions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[RID](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *NavigationServer2D) MapGetAgents(map_ RID, ) []RID {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapGetAgents), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[RID](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *NavigationServer2D) MapGetObstacles(map_ RID, ) []RID {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapGetObstacles), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[RID](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *NavigationServer2D) MapForceUpdate(map_ RID, )  {
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnMapForceUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) QueryPath(parameters NavigationPathQueryParameters2D, result NavigationPathQueryResult2D, )  {
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), result.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnQueryPath), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) RegionSetEnabled(region RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetEnabled(region RID, ) bool {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionGetEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionSetUseEdgeConnections(region RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionSetUseEdgeConnections), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetUseEdgeConnections(region RID, ) bool {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionGetUseEdgeConnections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionSetEnterCost(region RID, enter_cost float64, )  {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&enter_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionSetEnterCost), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetEnterCost(region RID, ) float64 {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionGetEnterCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionSetTravelCost(region RID, travel_cost float64, )  {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&travel_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionSetTravelCost), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetTravelCost(region RID, ) float64 {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionGetTravelCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionSetOwnerId(region RID, owner_id int64, )  {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionSetOwnerId), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetOwnerId(region RID, ) int64 {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionGetOwnerId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionOwnsPoint(region RID, point Vector2, ) bool {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionOwnsPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionSetMap(region RID, map_ RID, )  {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionSetMap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetMap(region RID, ) RID {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionGetMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) RegionSetNavigationLayers(region RID, navigation_layers int64, )  {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&navigation_layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionSetNavigationLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetNavigationLayers(region RID, ) int64 {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionGetNavigationLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionSetTransform(region RID, transform Transform2D, )  {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionSetNavigationPolygon(region RID, navigation_polygon NavigationPolygon, )  {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), navigation_polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionSetNavigationPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetConnectionsCount(region RID, ) int64 {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionGetConnectionsCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionGetConnectionPathwayStart(region RID, connection int64, ) Vector2 {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&connection) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&connection)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionGetConnectionPathwayStart), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) RegionGetConnectionPathwayEnd(region RID, connection int64, ) Vector2 {
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&connection) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&connection)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnRegionGetConnectionPathwayEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) LinkCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) LinkSetMap(link RID, map_ RID, )  {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkSetMap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetMap(link RID, ) RID {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkGetMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) LinkSetEnabled(link RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetEnabled(link RID, ) bool {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkGetEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) LinkSetBidirectional(link RID, bidirectional bool, )  {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&bidirectional) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkSetBidirectional), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkIsBidirectional(link RID, ) bool {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkIsBidirectional), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) LinkSetNavigationLayers(link RID, navigation_layers int64, )  {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&navigation_layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkSetNavigationLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetNavigationLayers(link RID, ) int64 {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkGetNavigationLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) LinkSetStartPosition(link RID, position Vector2, )  {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkSetStartPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetStartPosition(link RID, ) Vector2 {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkGetStartPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) LinkSetEndPosition(link RID, position Vector2, )  {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkSetEndPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetEndPosition(link RID, ) Vector2 {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkGetEndPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) LinkSetEnterCost(link RID, enter_cost float64, )  {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&enter_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkSetEnterCost), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetEnterCost(link RID, ) float64 {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkGetEnterCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) LinkSetTravelCost(link RID, travel_cost float64, )  {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&travel_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkSetTravelCost), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetTravelCost(link RID, ) float64 {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkGetTravelCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) LinkSetOwnerId(link RID, owner_id int64, )  {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkSetOwnerId), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetOwnerId(link RID, ) int64 {
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnLinkGetOwnerId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) AgentCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) AgentSetAvoidanceEnabled(agent RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentGetAvoidanceEnabled(agent RID, ) bool {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentGetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) AgentSetMap(agent RID, map_ RID, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetMap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentGetMap(agent RID, ) RID {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentGetMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) AgentSetPaused(agent RID, paused bool, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&paused) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetPaused), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentGetPaused(agent RID, ) bool {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentGetPaused), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) AgentSetNeighborDistance(agent RID, distance float64, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetNeighborDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetMaxNeighbors(agent RID, count int64, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetMaxNeighbors), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetTimeHorizonAgents(agent RID, time_horizon float64, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&time_horizon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetTimeHorizonAgents), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetTimeHorizonObstacles(agent RID, time_horizon float64, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&time_horizon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetTimeHorizonObstacles), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetRadius(agent RID, radius float64, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetMaxSpeed(agent RID, max_speed float64, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&max_speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetMaxSpeed), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetVelocityForced(agent RID, velocity Vector2, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetVelocityForced), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetVelocity(agent RID, velocity Vector2, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetPosition(agent RID, position Vector2, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentIsMapChanged(agent RID, ) bool {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentIsMapChanged), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) AgentSetAvoidanceCallback(agent RID, callback Callable, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetAvoidanceCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetAvoidanceLayers(agent RID, layers int64, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetAvoidanceLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetAvoidanceMask(agent RID, mask int64, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetAvoidanceMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetAvoidancePriority(agent RID, priority float64, )  {
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnAgentSetAvoidancePriority), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleCreate() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnObstacleCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) ObstacleSetAvoidanceEnabled(obstacle RID, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnObstacleSetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleGetAvoidanceEnabled(obstacle RID, ) bool {
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnObstacleGetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) ObstacleSetMap(obstacle RID, map_ RID, )  {
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnObstacleSetMap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleGetMap(obstacle RID, ) RID {
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnObstacleGetMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) ObstacleSetPaused(obstacle RID, paused bool, )  {
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&paused) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnObstacleSetPaused), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleGetPaused(obstacle RID, ) bool {
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnObstacleGetPaused), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) ObstacleSetRadius(obstacle RID, radius float64, )  {
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnObstacleSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleSetVelocity(obstacle RID, velocity Vector2, )  {
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnObstacleSetVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleSetPosition(obstacle RID, position Vector2, )  {
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnObstacleSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleSetVertices(obstacle RID, vertices PackedVector2Array, )  {
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), vertices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnObstacleSetVertices), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleSetAvoidanceLayers(obstacle RID, layers int64, )  {
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnObstacleSetAvoidanceLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ParseSourceGeometryData(navigation_polygon NavigationPolygon, source_geometry_data NavigationMeshSourceGeometryData2D, root_node Node, callback Callable, )  {
  cargs := []gdc.ConstTypePtr{navigation_polygon.AsCTypePtr(), source_geometry_data.AsCTypePtr(), root_node.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnParseSourceGeometryData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) BakeFromSourceGeometryData(navigation_polygon NavigationPolygon, source_geometry_data NavigationMeshSourceGeometryData2D, callback Callable, )  {
  cargs := []gdc.ConstTypePtr{navigation_polygon.AsCTypePtr(), source_geometry_data.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnBakeFromSourceGeometryData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) BakeFromSourceGeometryDataAsync(navigation_polygon NavigationPolygon, source_geometry_data NavigationMeshSourceGeometryData2D, callback Callable, )  {
  cargs := []gdc.ConstTypePtr{navigation_polygon.AsCTypePtr(), source_geometry_data.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnBakeFromSourceGeometryDataAsync), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) FreeRid(rid RID, )  {
  cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnFreeRid), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) SetDebugEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnSetDebugEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) GetDebugEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer2D.fnGetDebugEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals

type NavigationServer2DMapChangedSignalFn func(map_ RID, )

func (me *NavigationServer2D) ConnectMapChanged(subs SignalSubscribers, fn NavigationServer2DMapChangedSignalFn) {
  sig := StringNameFromStr("map_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationServer2D) DisconnectMapChanged(subs SignalSubscribers, fn NavigationServer2DMapChangedSignalFn) {
  sig := StringNameFromStr("map_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationServer2DNavigationDebugChangedSignalFn func()

func (me *NavigationServer2D) ConnectNavigationDebugChanged(subs SignalSubscribers, fn NavigationServer2DNavigationDebugChangedSignalFn) {
  sig := StringNameFromStr("navigation_debug_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationServer2D) DisconnectNavigationDebugChanged(subs SignalSubscribers, fn NavigationServer2DNavigationDebugChangedSignalFn) {
  sig := StringNameFromStr("navigation_debug_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
