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

type ptrsForNavigationServer3DList struct {
	fnGetMaps                         gdc.MethodBindPtr
	fnMapCreate                       gdc.MethodBindPtr
	fnMapSetActive                    gdc.MethodBindPtr
	fnMapIsActive                     gdc.MethodBindPtr
	fnMapSetUp                        gdc.MethodBindPtr
	fnMapGetUp                        gdc.MethodBindPtr
	fnMapSetCellSize                  gdc.MethodBindPtr
	fnMapGetCellSize                  gdc.MethodBindPtr
	fnMapSetCellHeight                gdc.MethodBindPtr
	fnMapGetCellHeight                gdc.MethodBindPtr
	fnMapSetUseEdgeConnections        gdc.MethodBindPtr
	fnMapGetUseEdgeConnections        gdc.MethodBindPtr
	fnMapSetEdgeConnectionMargin      gdc.MethodBindPtr
	fnMapGetEdgeConnectionMargin      gdc.MethodBindPtr
	fnMapSetLinkConnectionRadius      gdc.MethodBindPtr
	fnMapGetLinkConnectionRadius      gdc.MethodBindPtr
	fnMapGetPath                      gdc.MethodBindPtr
	fnMapGetClosestPointToSegment     gdc.MethodBindPtr
	fnMapGetClosestPoint              gdc.MethodBindPtr
	fnMapGetClosestPointNormal        gdc.MethodBindPtr
	fnMapGetClosestPointOwner         gdc.MethodBindPtr
	fnMapGetLinks                     gdc.MethodBindPtr
	fnMapGetRegions                   gdc.MethodBindPtr
	fnMapGetAgents                    gdc.MethodBindPtr
	fnMapGetObstacles                 gdc.MethodBindPtr
	fnMapForceUpdate                  gdc.MethodBindPtr
	fnQueryPath                       gdc.MethodBindPtr
	fnRegionCreate                    gdc.MethodBindPtr
	fnRegionSetEnabled                gdc.MethodBindPtr
	fnRegionGetEnabled                gdc.MethodBindPtr
	fnRegionSetUseEdgeConnections     gdc.MethodBindPtr
	fnRegionGetUseEdgeConnections     gdc.MethodBindPtr
	fnRegionSetEnterCost              gdc.MethodBindPtr
	fnRegionGetEnterCost              gdc.MethodBindPtr
	fnRegionSetTravelCost             gdc.MethodBindPtr
	fnRegionGetTravelCost             gdc.MethodBindPtr
	fnRegionSetOwnerId                gdc.MethodBindPtr
	fnRegionGetOwnerId                gdc.MethodBindPtr
	fnRegionOwnsPoint                 gdc.MethodBindPtr
	fnRegionSetMap                    gdc.MethodBindPtr
	fnRegionGetMap                    gdc.MethodBindPtr
	fnRegionSetNavigationLayers       gdc.MethodBindPtr
	fnRegionGetNavigationLayers       gdc.MethodBindPtr
	fnRegionSetTransform              gdc.MethodBindPtr
	fnRegionSetNavigationMesh         gdc.MethodBindPtr
	fnRegionBakeNavigationMesh        gdc.MethodBindPtr
	fnRegionGetConnectionsCount       gdc.MethodBindPtr
	fnRegionGetConnectionPathwayStart gdc.MethodBindPtr
	fnRegionGetConnectionPathwayEnd   gdc.MethodBindPtr
	fnLinkCreate                      gdc.MethodBindPtr
	fnLinkSetMap                      gdc.MethodBindPtr
	fnLinkGetMap                      gdc.MethodBindPtr
	fnLinkSetEnabled                  gdc.MethodBindPtr
	fnLinkGetEnabled                  gdc.MethodBindPtr
	fnLinkSetBidirectional            gdc.MethodBindPtr
	fnLinkIsBidirectional             gdc.MethodBindPtr
	fnLinkSetNavigationLayers         gdc.MethodBindPtr
	fnLinkGetNavigationLayers         gdc.MethodBindPtr
	fnLinkSetStartPosition            gdc.MethodBindPtr
	fnLinkGetStartPosition            gdc.MethodBindPtr
	fnLinkSetEndPosition              gdc.MethodBindPtr
	fnLinkGetEndPosition              gdc.MethodBindPtr
	fnLinkSetEnterCost                gdc.MethodBindPtr
	fnLinkGetEnterCost                gdc.MethodBindPtr
	fnLinkSetTravelCost               gdc.MethodBindPtr
	fnLinkGetTravelCost               gdc.MethodBindPtr
	fnLinkSetOwnerId                  gdc.MethodBindPtr
	fnLinkGetOwnerId                  gdc.MethodBindPtr
	fnAgentCreate                     gdc.MethodBindPtr
	fnAgentSetAvoidanceEnabled        gdc.MethodBindPtr
	fnAgentGetAvoidanceEnabled        gdc.MethodBindPtr
	fnAgentSetUse3DAvoidance          gdc.MethodBindPtr
	fnAgentGetUse3DAvoidance          gdc.MethodBindPtr
	fnAgentSetMap                     gdc.MethodBindPtr
	fnAgentGetMap                     gdc.MethodBindPtr
	fnAgentSetPaused                  gdc.MethodBindPtr
	fnAgentGetPaused                  gdc.MethodBindPtr
	fnAgentSetNeighborDistance        gdc.MethodBindPtr
	fnAgentSetMaxNeighbors            gdc.MethodBindPtr
	fnAgentSetTimeHorizonAgents       gdc.MethodBindPtr
	fnAgentSetTimeHorizonObstacles    gdc.MethodBindPtr
	fnAgentSetRadius                  gdc.MethodBindPtr
	fnAgentSetHeight                  gdc.MethodBindPtr
	fnAgentSetMaxSpeed                gdc.MethodBindPtr
	fnAgentSetVelocityForced          gdc.MethodBindPtr
	fnAgentSetVelocity                gdc.MethodBindPtr
	fnAgentSetPosition                gdc.MethodBindPtr
	fnAgentIsMapChanged               gdc.MethodBindPtr
	fnAgentSetAvoidanceCallback       gdc.MethodBindPtr
	fnAgentSetAvoidanceLayers         gdc.MethodBindPtr
	fnAgentSetAvoidanceMask           gdc.MethodBindPtr
	fnAgentSetAvoidancePriority       gdc.MethodBindPtr
	fnObstacleCreate                  gdc.MethodBindPtr
	fnObstacleSetAvoidanceEnabled     gdc.MethodBindPtr
	fnObstacleGetAvoidanceEnabled     gdc.MethodBindPtr
	fnObstacleSetUse3DAvoidance       gdc.MethodBindPtr
	fnObstacleGetUse3DAvoidance       gdc.MethodBindPtr
	fnObstacleSetMap                  gdc.MethodBindPtr
	fnObstacleGetMap                  gdc.MethodBindPtr
	fnObstacleSetPaused               gdc.MethodBindPtr
	fnObstacleGetPaused               gdc.MethodBindPtr
	fnObstacleSetRadius               gdc.MethodBindPtr
	fnObstacleSetHeight               gdc.MethodBindPtr
	fnObstacleSetVelocity             gdc.MethodBindPtr
	fnObstacleSetPosition             gdc.MethodBindPtr
	fnObstacleSetVertices             gdc.MethodBindPtr
	fnObstacleSetAvoidanceLayers      gdc.MethodBindPtr
	fnParseSourceGeometryData         gdc.MethodBindPtr
	fnBakeFromSourceGeometryData      gdc.MethodBindPtr
	fnBakeFromSourceGeometryDataAsync gdc.MethodBindPtr
	fnFreeRid                         gdc.MethodBindPtr
	fnSetActive                       gdc.MethodBindPtr
	fnSetDebugEnabled                 gdc.MethodBindPtr
	fnGetDebugEnabled                 gdc.MethodBindPtr
	fnGetProcessInfo                  gdc.MethodBindPtr
}

var ptrsForNavigationServer3D ptrsForNavigationServer3DList

func initNavigationServer3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NavigationServer3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_maps")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnGetMaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("map_create")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("map_set_active")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapSetActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("map_is_active")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapIsActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("map_set_up")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapSetUp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("map_get_up")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetUp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 531438156))
	}
	{
		methodName := StringNameFromStr("map_set_cell_size")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapSetCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("map_get_cell_size")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("map_set_cell_height")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapSetCellHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("map_get_cell_height")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetCellHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("map_set_use_edge_connections")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapSetUseEdgeConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("map_get_use_edge_connections")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetUseEdgeConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("map_set_edge_connection_margin")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapSetEdgeConnectionMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("map_get_edge_connection_margin")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetEdgeConnectionMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("map_set_link_connection_radius")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapSetLinkConnectionRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("map_get_link_connection_radius")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetLinkConnectionRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("map_get_path")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1187418690))
	}
	{
		methodName := StringNameFromStr("map_get_closest_point_to_segment")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetClosestPointToSegment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3830095642))
	}
	{
		methodName := StringNameFromStr("map_get_closest_point")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetClosestPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2056183332))
	}
	{
		methodName := StringNameFromStr("map_get_closest_point_normal")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetClosestPointNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2056183332))
	}
	{
		methodName := StringNameFromStr("map_get_closest_point_owner")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetClosestPointOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 553364610))
	}
	{
		methodName := StringNameFromStr("map_get_links")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetLinks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
	}
	{
		methodName := StringNameFromStr("map_get_regions")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetRegions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
	}
	{
		methodName := StringNameFromStr("map_get_agents")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetAgents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
	}
	{
		methodName := StringNameFromStr("map_get_obstacles")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapGetObstacles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2684255073))
	}
	{
		methodName := StringNameFromStr("map_force_update")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnMapForceUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("query_path")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnQueryPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3415008901))
	}
	{
		methodName := StringNameFromStr("region_create")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("region_set_enabled")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("region_get_enabled")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionGetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("region_set_use_edge_connections")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionSetUseEdgeConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("region_get_use_edge_connections")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionGetUseEdgeConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("region_set_enter_cost")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionSetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("region_get_enter_cost")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionGetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("region_set_travel_cost")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionSetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("region_get_travel_cost")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionGetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("region_set_owner_id")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionSetOwnerId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("region_get_owner_id")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionGetOwnerId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("region_owns_point")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionOwnsPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2360011153))
	}
	{
		methodName := StringNameFromStr("region_set_map")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionSetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
	}
	{
		methodName := StringNameFromStr("region_get_map")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionGetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
	}
	{
		methodName := StringNameFromStr("region_set_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionSetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("region_get_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionGetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("region_set_transform")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3935195649))
	}
	{
		methodName := StringNameFromStr("region_set_navigation_mesh")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionSetNavigationMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2764952978))
	}
	{
		methodName := StringNameFromStr("region_bake_navigation_mesh")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionBakeNavigationMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1401173477))
	}
	{
		methodName := StringNameFromStr("region_get_connections_count")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionGetConnectionsCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("region_get_connection_pathway_start")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionGetConnectionPathwayStart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3440143363))
	}
	{
		methodName := StringNameFromStr("region_get_connection_pathway_end")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnRegionGetConnectionPathwayEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3440143363))
	}
	{
		methodName := StringNameFromStr("link_create")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("link_set_map")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkSetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
	}
	{
		methodName := StringNameFromStr("link_get_map")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkGetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
	}
	{
		methodName := StringNameFromStr("link_set_enabled")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("link_get_enabled")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkGetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("link_set_bidirectional")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkSetBidirectional = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("link_is_bidirectional")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkIsBidirectional = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("link_set_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkSetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("link_get_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkGetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("link_set_start_position")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkSetStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("link_get_start_position")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkGetStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 531438156))
	}
	{
		methodName := StringNameFromStr("link_set_end_position")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkSetEndPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("link_get_end_position")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkGetEndPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 531438156))
	}
	{
		methodName := StringNameFromStr("link_set_enter_cost")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkSetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("link_get_enter_cost")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkGetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("link_set_travel_cost")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkSetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("link_get_travel_cost")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkGetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866169185))
	}
	{
		methodName := StringNameFromStr("link_set_owner_id")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkSetOwnerId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("link_get_owner_id")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnLinkGetOwnerId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2198884583))
	}
	{
		methodName := StringNameFromStr("agent_create")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("agent_set_avoidance_enabled")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("agent_get_avoidance_enabled")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentGetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("agent_set_use_3d_avoidance")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetUse3DAvoidance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("agent_get_use_3d_avoidance")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentGetUse3DAvoidance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("agent_set_map")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
	}
	{
		methodName := StringNameFromStr("agent_get_map")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentGetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
	}
	{
		methodName := StringNameFromStr("agent_set_paused")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("agent_get_paused")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentGetPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("agent_set_neighbor_distance")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetNeighborDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("agent_set_max_neighbors")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetMaxNeighbors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("agent_set_time_horizon_agents")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetTimeHorizonAgents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("agent_set_time_horizon_obstacles")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetTimeHorizonObstacles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("agent_set_radius")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("agent_set_height")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("agent_set_max_speed")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetMaxSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("agent_set_velocity_forced")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetVelocityForced = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("agent_set_velocity")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("agent_set_position")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("agent_is_map_changed")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentIsMapChanged = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("agent_set_avoidance_callback")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetAvoidanceCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3379118538))
	}
	{
		methodName := StringNameFromStr("agent_set_avoidance_layers")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetAvoidanceLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("agent_set_avoidance_mask")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetAvoidanceMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("agent_set_avoidance_priority")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnAgentSetAvoidancePriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("obstacle_create")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("obstacle_set_avoidance_enabled")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleSetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("obstacle_get_avoidance_enabled")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleGetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("obstacle_set_use_3d_avoidance")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleSetUse3DAvoidance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("obstacle_get_use_3d_avoidance")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleGetUse3DAvoidance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("obstacle_set_map")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleSetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 395945892))
	}
	{
		methodName := StringNameFromStr("obstacle_get_map")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleGetMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814569979))
	}
	{
		methodName := StringNameFromStr("obstacle_set_paused")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleSetPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265174801))
	}
	{
		methodName := StringNameFromStr("obstacle_get_paused")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleGetPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("obstacle_set_radius")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("obstacle_set_height")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1794382983))
	}
	{
		methodName := StringNameFromStr("obstacle_set_velocity")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleSetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("obstacle_set_position")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3227306858))
	}
	{
		methodName := StringNameFromStr("obstacle_set_vertices")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleSetVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4030257846))
	}
	{
		methodName := StringNameFromStr("obstacle_set_avoidance_layers")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnObstacleSetAvoidanceLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3411492887))
	}
	{
		methodName := StringNameFromStr("parse_source_geometry_data")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnParseSourceGeometryData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 685862123))
	}
	{
		methodName := StringNameFromStr("bake_from_source_geometry_data")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnBakeFromSourceGeometryData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2469318639))
	}
	{
		methodName := StringNameFromStr("bake_from_source_geometry_data_async")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnBakeFromSourceGeometryDataAsync = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2469318639))
	}
	{
		methodName := StringNameFromStr("free_rid")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnFreeRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("set_active")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnSetActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_debug_enabled")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnSetDebugEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_debug_enabled")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnGetDebugEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_process_info")
		defer methodName.Destroy()
		ptrsForNavigationServer3D.fnGetProcessInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1938440894))
	}
}

type NavigationServer3D struct {
	Object
}

func (me *NavigationServer3D) BaseClass() string {
	return "NavigationServer3D"
}

func NewNavigationServer3D() *NavigationServer3D {
	str := StringNameFromStr("NavigationServer3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &NavigationServer3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type NavigationServer3DProcessInfo int

const (
	NavigationServer3DProcessInfoInfoActiveMaps          NavigationServer3DProcessInfo = 0
	NavigationServer3DProcessInfoInfoRegionCount         NavigationServer3DProcessInfo = 1
	NavigationServer3DProcessInfoInfoAgentCount          NavigationServer3DProcessInfo = 2
	NavigationServer3DProcessInfoInfoLinkCount           NavigationServer3DProcessInfo = 3
	NavigationServer3DProcessInfoInfoPolygonCount        NavigationServer3DProcessInfo = 4
	NavigationServer3DProcessInfoInfoEdgeCount           NavigationServer3DProcessInfo = 5
	NavigationServer3DProcessInfoInfoEdgeMergeCount      NavigationServer3DProcessInfo = 6
	NavigationServer3DProcessInfoInfoEdgeConnectionCount NavigationServer3DProcessInfo = 7
	NavigationServer3DProcessInfoInfoEdgeFreeCount       NavigationServer3DProcessInfo = 8
)

func (me *NavigationServer3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *NavigationServer3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *NavigationServer3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *NavigationServer3D) GetMaps() []RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnGetMaps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[RID](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *NavigationServer3D) MapCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) MapSetActive(map_ RID, active bool) {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&active)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapSetActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) MapIsActive(map_ RID) bool {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapIsActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) MapSetUp(map_ RID, up Vector3) {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), up.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapSetUp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) MapGetUp(map_ RID) Vector3 {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetUp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) MapSetCellSize(map_ RID, cell_size float64) {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&cell_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapSetCellSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) MapGetCellSize(map_ RID) float64 {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetCellSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) MapSetCellHeight(map_ RID, cell_height float64) {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&cell_height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapSetCellHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) MapGetCellHeight(map_ RID) float64 {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetCellHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) MapSetUseEdgeConnections(map_ RID, enabled bool) {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapSetUseEdgeConnections), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) MapGetUseEdgeConnections(map_ RID) bool {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetUseEdgeConnections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) MapSetEdgeConnectionMargin(map_ RID, margin float64) {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapSetEdgeConnectionMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) MapGetEdgeConnectionMargin(map_ RID) float64 {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetEdgeConnectionMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) MapSetLinkConnectionRadius(map_ RID, radius float64) {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapSetLinkConnectionRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) MapGetLinkConnectionRadius(map_ RID) float64 {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetLinkConnectionRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) MapGetPath(map_ RID, origin Vector3, destination Vector3, optimize bool, navigation_layers int64) PackedVector3Array {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), origin.AsCTypePtr(), destination.AsCTypePtr(), gdc.ConstTypePtr(&optimize), gdc.ConstTypePtr(&navigation_layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector3Array()
	pinner.Pin(&optimize)
	pinner.Pin(&navigation_layers)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) MapGetClosestPointToSegment(map_ RID, start Vector3, end Vector3, use_collision bool) Vector3 {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), start.AsCTypePtr(), end.AsCTypePtr(), gdc.ConstTypePtr(&use_collision)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&use_collision)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetClosestPointToSegment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) MapGetClosestPoint(map_ RID, to_point Vector3) Vector3 {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), to_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetClosestPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) MapGetClosestPointNormal(map_ RID, to_point Vector3) Vector3 {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), to_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetClosestPointNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) MapGetClosestPointOwner(map_ RID, to_point Vector3) RID {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), to_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetClosestPointOwner), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) MapGetLinks(map_ RID) []RID {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetLinks), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[RID](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *NavigationServer3D) MapGetRegions(map_ RID) []RID {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetRegions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[RID](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *NavigationServer3D) MapGetAgents(map_ RID) []RID {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetAgents), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[RID](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *NavigationServer3D) MapGetObstacles(map_ RID) []RID {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapGetObstacles), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[RID](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *NavigationServer3D) MapForceUpdate(map_ RID) {
	cargs := []gdc.ConstTypePtr{map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnMapForceUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) QueryPath(parameters NavigationPathQueryParameters3D, result NavigationPathQueryResult3D) {
	cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), result.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnQueryPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) RegionCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) RegionSetEnabled(region RID, enabled bool) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) RegionGetEnabled(region RID) bool {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionGetEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) RegionSetUseEdgeConnections(region RID, enabled bool) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionSetUseEdgeConnections), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) RegionGetUseEdgeConnections(region RID) bool {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionGetUseEdgeConnections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) RegionSetEnterCost(region RID, enter_cost float64) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&enter_cost)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionSetEnterCost), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) RegionGetEnterCost(region RID) float64 {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionGetEnterCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) RegionSetTravelCost(region RID, travel_cost float64) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&travel_cost)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionSetTravelCost), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) RegionGetTravelCost(region RID) float64 {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionGetTravelCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) RegionSetOwnerId(region RID, owner_id int64) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionSetOwnerId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) RegionGetOwnerId(region RID) int64 {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionGetOwnerId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) RegionOwnsPoint(region RID, point Vector3) bool {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionOwnsPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) RegionSetMap(region RID, map_ RID) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionSetMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) RegionGetMap(region RID) RID {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionGetMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) RegionSetNavigationLayers(region RID, navigation_layers int64) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&navigation_layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionSetNavigationLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) RegionGetNavigationLayers(region RID) int64 {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionGetNavigationLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) RegionSetTransform(region RID, transform Transform3D) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), transform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) RegionSetNavigationMesh(region RID, navigation_mesh NavigationMesh) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), navigation_mesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionSetNavigationMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) RegionBakeNavigationMesh(navigation_mesh NavigationMesh, root_node Node) {
	cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), root_node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionBakeNavigationMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) RegionGetConnectionsCount(region RID) int64 {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionGetConnectionsCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) RegionGetConnectionPathwayStart(region RID, connection int64) Vector3 {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&connection)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&connection)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionGetConnectionPathwayStart), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) RegionGetConnectionPathwayEnd(region RID, connection int64) Vector3 {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&connection)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&connection)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnRegionGetConnectionPathwayEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) LinkCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) LinkSetMap(link RID, map_ RID) {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkSetMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) LinkGetMap(link RID) RID {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkGetMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) LinkSetEnabled(link RID, enabled bool) {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) LinkGetEnabled(link RID) bool {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkGetEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) LinkSetBidirectional(link RID, bidirectional bool) {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&bidirectional)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkSetBidirectional), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) LinkIsBidirectional(link RID) bool {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkIsBidirectional), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) LinkSetNavigationLayers(link RID, navigation_layers int64) {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&navigation_layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkSetNavigationLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) LinkGetNavigationLayers(link RID) int64 {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkGetNavigationLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) LinkSetStartPosition(link RID, position Vector3) {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkSetStartPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) LinkGetStartPosition(link RID) Vector3 {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkGetStartPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) LinkSetEndPosition(link RID, position Vector3) {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkSetEndPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) LinkGetEndPosition(link RID) Vector3 {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkGetEndPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) LinkSetEnterCost(link RID, enter_cost float64) {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&enter_cost)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkSetEnterCost), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) LinkGetEnterCost(link RID) float64 {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkGetEnterCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) LinkSetTravelCost(link RID, travel_cost float64) {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&travel_cost)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkSetTravelCost), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) LinkGetTravelCost(link RID) float64 {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkGetTravelCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) LinkSetOwnerId(link RID, owner_id int64) {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&owner_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkSetOwnerId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) LinkGetOwnerId(link RID) int64 {
	cargs := []gdc.ConstTypePtr{link.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnLinkGetOwnerId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) AgentCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) AgentSetAvoidanceEnabled(agent RID, enabled bool) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentGetAvoidanceEnabled(agent RID) bool {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentGetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) AgentSetUse3DAvoidance(agent RID, enabled bool) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetUse3DAvoidance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentGetUse3DAvoidance(agent RID) bool {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentGetUse3DAvoidance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) AgentSetMap(agent RID, map_ RID) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentGetMap(agent RID) RID {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentGetMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) AgentSetPaused(agent RID, paused bool) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&paused)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetPaused), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentGetPaused(agent RID) bool {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentGetPaused), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) AgentSetNeighborDistance(agent RID, distance float64) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetNeighborDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentSetMaxNeighbors(agent RID, count int64) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetMaxNeighbors), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentSetTimeHorizonAgents(agent RID, time_horizon float64) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&time_horizon)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetTimeHorizonAgents), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentSetTimeHorizonObstacles(agent RID, time_horizon float64) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&time_horizon)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetTimeHorizonObstacles), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentSetRadius(agent RID, radius float64) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentSetHeight(agent RID, height float64) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentSetMaxSpeed(agent RID, max_speed float64) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&max_speed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetMaxSpeed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentSetVelocityForced(agent RID, velocity Vector3) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetVelocityForced), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentSetVelocity(agent RID, velocity Vector3) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentSetPosition(agent RID, position Vector3) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentIsMapChanged(agent RID) bool {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentIsMapChanged), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) AgentSetAvoidanceCallback(agent RID, callback Callable) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetAvoidanceCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentSetAvoidanceLayers(agent RID, layers int64) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetAvoidanceLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentSetAvoidanceMask(agent RID, mask int64) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetAvoidanceMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) AgentSetAvoidancePriority(agent RID, priority float64) {
	cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnAgentSetAvoidancePriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) ObstacleCreate() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) ObstacleSetAvoidanceEnabled(obstacle RID, enabled bool) {
	cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleSetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) ObstacleGetAvoidanceEnabled(obstacle RID) bool {
	cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleGetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) ObstacleSetUse3DAvoidance(obstacle RID, enabled bool) {
	cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleSetUse3DAvoidance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) ObstacleGetUse3DAvoidance(obstacle RID) bool {
	cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleGetUse3DAvoidance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) ObstacleSetMap(obstacle RID, map_ RID) {
	cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), map_.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleSetMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) ObstacleGetMap(obstacle RID) RID {
	cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleGetMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationServer3D) ObstacleSetPaused(obstacle RID, paused bool) {
	cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&paused)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleSetPaused), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) ObstacleGetPaused(obstacle RID) bool {
	cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleGetPaused), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) ObstacleSetRadius(obstacle RID, radius float64) {
	cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) ObstacleSetHeight(obstacle RID, height float64) {
	cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) ObstacleSetVelocity(obstacle RID, velocity Vector3) {
	cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleSetVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) ObstacleSetPosition(obstacle RID, position Vector3) {
	cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) ObstacleSetVertices(obstacle RID, vertices PackedVector3Array) {
	cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), vertices.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleSetVertices), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) ObstacleSetAvoidanceLayers(obstacle RID, layers int64) {
	cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnObstacleSetAvoidanceLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) ParseSourceGeometryData(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, root_node Node, callback Callable) {
	cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), source_geometry_data.AsCTypePtr(), root_node.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnParseSourceGeometryData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) BakeFromSourceGeometryData(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, callback Callable) {
	cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), source_geometry_data.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnBakeFromSourceGeometryData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) BakeFromSourceGeometryDataAsync(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, callback Callable) {
	cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), source_geometry_data.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnBakeFromSourceGeometryDataAsync), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) FreeRid(rid RID) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnFreeRid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) SetActive(active bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnSetActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) SetDebugEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnSetDebugEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationServer3D) GetDebugEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnGetDebugEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationServer3D) GetProcessInfo(process_info NavigationServer3DProcessInfo) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&process_info)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&process_info)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationServer3D.fnGetProcessInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals

type NavigationServer3DMapChangedSignalFn func(map_ RID)

func (me *NavigationServer3D) ConnectMapChanged(subs SignalSubscribers, fn NavigationServer3DMapChangedSignalFn) {
	sig := StringNameFromStr("map_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationServer3D) DisconnectMapChanged(subs SignalSubscribers, fn NavigationServer3DMapChangedSignalFn) {
	sig := StringNameFromStr("map_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationServer3DNavigationDebugChangedSignalFn func()

func (me *NavigationServer3D) ConnectNavigationDebugChanged(subs SignalSubscribers, fn NavigationServer3DNavigationDebugChangedSignalFn) {
	sig := StringNameFromStr("navigation_debug_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationServer3D) DisconnectNavigationDebugChanged(subs SignalSubscribers, fn NavigationServer3DNavigationDebugChangedSignalFn) {
	sig := StringNameFromStr("navigation_debug_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationServer3DAvoidanceDebugChangedSignalFn func()

func (me *NavigationServer3D) ConnectAvoidanceDebugChanged(subs SignalSubscribers, fn NavigationServer3DAvoidanceDebugChangedSignalFn) {
	sig := StringNameFromStr("avoidance_debug_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationServer3D) DisconnectAvoidanceDebugChanged(subs SignalSubscribers, fn NavigationServer3DAvoidanceDebugChangedSignalFn) {
	sig := StringNameFromStr("avoidance_debug_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
