// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *NavigationServer2D) GetMaps() RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_maps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) MapCreate() RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) MapSetActive(map_ RID, active bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), gdc.ConstTypePtr(&active), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) MapIsActive(map_ RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) MapSetCellSize(map_ RID, cell_size float32, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), gdc.ConstTypePtr(&cell_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) MapGetCellSize(map_ RID, ) float32 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) MapSetUseEdgeConnections(map_ RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_use_edge_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) MapGetUseEdgeConnections(map_ RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_use_edge_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) MapSetEdgeConnectionMargin(map_ RID, margin float32, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_edge_connection_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) MapGetEdgeConnectionMargin(map_ RID, ) float32 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_edge_connection_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) MapSetLinkConnectionRadius(map_ RID, radius float32, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_link_connection_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) MapGetLinkConnectionRadius(map_ RID, ) float32 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_link_connection_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) MapGetPath(map_ RID, origin Vector2, destination Vector2, optimize bool, navigation_layers int, ) PackedVector2Array {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 56240621) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), gdc.ConstTypePtr(origin.AsCTypePtr()), gdc.ConstTypePtr(destination.AsCTypePtr()), gdc.ConstTypePtr(&optimize), gdc.ConstTypePtr(&navigation_layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) MapGetClosestPoint(map_ RID, to_point Vector2, ) Vector2 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_closest_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1358334418) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), gdc.ConstTypePtr(to_point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) MapGetClosestPointOwner(map_ RID, to_point Vector2, ) RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_closest_point_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1353467510) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), gdc.ConstTypePtr(to_point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) MapGetLinks(map_ RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_links")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2684255073) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) MapGetRegions(map_ RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_regions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2684255073) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) MapGetAgents(map_ RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_agents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2684255073) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) MapGetObstacles(map_ RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_obstacles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2684255073) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) MapForceUpdate(map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_force_update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) QueryPath(parameters NavigationPathQueryParameters2D, result NavigationPathQueryResult2D, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("query_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3394638789) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parameters.AsCTypePtr()), gdc.ConstTypePtr(result.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) RegionCreate() RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) RegionSetUseEdgeConnections(region RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_use_edge_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) RegionGetUseEdgeConnections(region RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_use_edge_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) RegionSetEnterCost(region RID, enter_cost float32, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), gdc.ConstTypePtr(&enter_cost), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) RegionGetEnterCost(region RID, ) float32 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) RegionSetTravelCost(region RID, travel_cost float32, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), gdc.ConstTypePtr(&travel_cost), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) RegionGetTravelCost(region RID, ) float32 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) RegionSetOwnerId(region RID, owner_id int, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_owner_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) RegionGetOwnerId(region RID, ) int {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_owner_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) RegionOwnsPoint(region RID, point Vector2, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_owns_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 219849798) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), gdc.ConstTypePtr(point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) RegionSetMap(region RID, map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) RegionGetMap(region RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814569979) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) RegionSetNavigationLayers(region RID, navigation_layers int, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), gdc.ConstTypePtr(&navigation_layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) RegionGetNavigationLayers(region RID, ) int {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) RegionSetTransform(region RID, transform Transform2D, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1246044741) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) RegionSetNavigationPolygon(region RID, navigation_polygon NavigationPolygon, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_navigation_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3633623451) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), gdc.ConstTypePtr(navigation_polygon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) RegionGetConnectionsCount(region RID, ) int {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_connections_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) RegionGetConnectionPathwayStart(region RID, connection int, ) Vector2 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_connection_pathway_start")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2546185844) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), gdc.ConstTypePtr(&connection), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) RegionGetConnectionPathwayEnd(region RID, connection int, ) Vector2 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_connection_pathway_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2546185844) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), gdc.ConstTypePtr(&connection), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) LinkCreate() RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) LinkSetMap(link RID, map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) LinkGetMap(link RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814569979) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) LinkSetBidirectional(link RID, bidirectional bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_bidirectional")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), gdc.ConstTypePtr(&bidirectional), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) LinkIsBidirectional(link RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_is_bidirectional")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) LinkSetNavigationLayers(link RID, navigation_layers int, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), gdc.ConstTypePtr(&navigation_layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) LinkGetNavigationLayers(link RID, ) int {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) LinkSetStartPosition(link RID, position Vector2, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) LinkGetStartPosition(link RID, ) Vector2 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2440833711) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) LinkSetEndPosition(link RID, position Vector2, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_end_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) LinkGetEndPosition(link RID, ) Vector2 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_end_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2440833711) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) LinkSetEnterCost(link RID, enter_cost float32, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), gdc.ConstTypePtr(&enter_cost), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) LinkGetEnterCost(link RID, ) float32 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) LinkSetTravelCost(link RID, travel_cost float32, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), gdc.ConstTypePtr(&travel_cost), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) LinkGetTravelCost(link RID, ) float32 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) LinkSetOwnerId(link RID, owner_id int, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_owner_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), gdc.ConstTypePtr(&owner_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) LinkGetOwnerId(link RID, ) int {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_owner_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(link.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) AgentCreate() RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) AgentSetAvoidanceEnabled(agent RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentGetAvoidanceEnabled(agent RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_get_avoidance_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) AgentSetMap(agent RID, map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentGetMap(agent RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_get_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814569979) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) AgentSetPaused(agent RID, paused bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(&paused), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentGetPaused(agent RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_get_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) AgentSetNeighborDistance(agent RID, distance float32, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_neighbor_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(&distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentSetMaxNeighbors(agent RID, count int, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_max_neighbors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(&count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentSetTimeHorizonAgents(agent RID, time_horizon float32, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_time_horizon_agents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(&time_horizon), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentSetTimeHorizonObstacles(agent RID, time_horizon float32, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_time_horizon_obstacles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(&time_horizon), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentSetRadius(agent RID, radius float32, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentSetMaxSpeed(agent RID, max_speed float32, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_max_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(&max_speed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentSetVelocityForced(agent RID, velocity Vector2, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_velocity_forced")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentSetVelocity(agent RID, velocity Vector2, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentSetPosition(agent RID, position Vector2, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentIsMapChanged(agent RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_is_map_changed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) AgentSetAvoidanceCallback(agent RID, callback Callable, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3379118538) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentSetAvoidanceLayers(agent RID, layers int, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(&layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentSetAvoidanceMask(agent RID, mask int, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) AgentSetAvoidancePriority(agent RID, priority float32, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(agent.AsCTypePtr()), gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) ObstacleCreate() RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) ObstacleSetAvoidanceEnabled(obstacle RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_avoidance_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(obstacle.AsCTypePtr()), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) ObstacleGetAvoidanceEnabled(obstacle RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_get_avoidance_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(obstacle.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) ObstacleSetMap(obstacle RID, map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(obstacle.AsCTypePtr()), gdc.ConstTypePtr(map_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) ObstacleGetMap(obstacle RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_get_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814569979) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(obstacle.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) ObstacleSetPaused(obstacle RID, paused bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(obstacle.AsCTypePtr()), gdc.ConstTypePtr(&paused), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) ObstacleGetPaused(obstacle RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_get_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(obstacle.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationServer2D) ObstacleSetRadius(obstacle RID, radius float32, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(obstacle.AsCTypePtr()), gdc.ConstTypePtr(&radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) ObstacleSetVelocity(obstacle RID, velocity Vector2, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(obstacle.AsCTypePtr()), gdc.ConstTypePtr(velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) ObstacleSetPosition(obstacle RID, position Vector2, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(obstacle.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) ObstacleSetVertices(obstacle RID, vertices PackedVector2Array, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 29476483) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(obstacle.AsCTypePtr()), gdc.ConstTypePtr(vertices.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) ObstacleSetAvoidanceLayers(obstacle RID, layers int, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_avoidance_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(obstacle.AsCTypePtr()), gdc.ConstTypePtr(&layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) FreeRid(rid RID, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("free_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) SetDebugEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationServer2D) GetDebugEnabled() bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API