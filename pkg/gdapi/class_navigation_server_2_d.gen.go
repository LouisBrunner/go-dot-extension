// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

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
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_maps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[RID](ret)
}

func  (me *NavigationServer2D) MapCreate() RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) MapSetActive(map_ RID, active bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&active) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) MapIsActive(map_ RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) MapSetCellSize(map_ RID, cell_size float64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&cell_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) MapGetCellSize(map_ RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) MapSetUseEdgeConnections(map_ RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_use_edge_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) MapGetUseEdgeConnections(map_ RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_use_edge_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) MapSetEdgeConnectionMargin(map_ RID, margin float64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_edge_connection_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) MapGetEdgeConnectionMargin(map_ RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_edge_connection_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) MapSetLinkConnectionRadius(map_ RID, radius float64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_link_connection_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) MapGetLinkConnectionRadius(map_ RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_link_connection_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) MapGetPath(map_ RID, origin Vector2, destination Vector2, optimize bool, navigation_layers int64, ) PackedVector2Array {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3146466012) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), origin.AsCTypePtr(), destination.AsCTypePtr(), gdc.ConstTypePtr(&optimize) , gdc.ConstTypePtr(&navigation_layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()
  pinner.Pin(&optimize)
  pinner.Pin(&navigation_layers)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) MapGetClosestPoint(map_ RID, to_point Vector2, ) Vector2 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_closest_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1358334418) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), to_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) MapGetClosestPointOwner(map_ RID, to_point Vector2, ) RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_closest_point_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1353467510) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), to_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) MapGetLinks(map_ RID, ) []RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_links")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2684255073) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[RID](ret)
}

func  (me *NavigationServer2D) MapGetRegions(map_ RID, ) []RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_regions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2684255073) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[RID](ret)
}

func  (me *NavigationServer2D) MapGetAgents(map_ RID, ) []RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_agents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2684255073) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[RID](ret)
}

func  (me *NavigationServer2D) MapGetObstacles(map_ RID, ) []RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_obstacles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2684255073) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[RID](ret)
}

func  (me *NavigationServer2D) MapForceUpdate(map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_force_update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) QueryPath(parameters NavigationPathQueryParameters2D, result NavigationPathQueryResult2D, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("query_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3394638789) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), result.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionCreate() RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) RegionSetEnabled(region RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetEnabled(region RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionSetUseEdgeConnections(region RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_use_edge_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetUseEdgeConnections(region RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_use_edge_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionSetEnterCost(region RID, enter_cost float64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&enter_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetEnterCost(region RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionSetTravelCost(region RID, travel_cost float64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&travel_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetTravelCost(region RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionSetOwnerId(region RID, owner_id int64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_owner_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetOwnerId(region RID, ) int64 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_owner_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionOwnsPoint(region RID, point Vector2, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_owns_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 219849798) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionSetMap(region RID, map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetMap(region RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814569979) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) RegionSetNavigationLayers(region RID, navigation_layers int64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&navigation_layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetNavigationLayers(region RID, ) int64 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionSetTransform(region RID, transform Transform2D, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1246044741) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionSetNavigationPolygon(region RID, navigation_polygon NavigationPolygon, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_navigation_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3633623451) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), navigation_polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) RegionGetConnectionsCount(region RID, ) int64 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_connections_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) RegionGetConnectionPathwayStart(region RID, connection int64, ) Vector2 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_connection_pathway_start")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2546185844) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&connection) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&connection)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) RegionGetConnectionPathwayEnd(region RID, connection int64, ) Vector2 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_connection_pathway_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2546185844) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&connection) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&connection)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) LinkCreate() RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) LinkSetMap(link RID, map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetMap(link RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814569979) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) LinkSetEnabled(link RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetEnabled(link RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) LinkSetBidirectional(link RID, bidirectional bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_bidirectional")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&bidirectional) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkIsBidirectional(link RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_is_bidirectional")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) LinkSetNavigationLayers(link RID, navigation_layers int64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&navigation_layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetNavigationLayers(link RID, ) int64 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) LinkSetStartPosition(link RID, position Vector2, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetStartPosition(link RID, ) Vector2 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2440833711) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) LinkSetEndPosition(link RID, position Vector2, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_end_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetEndPosition(link RID, ) Vector2 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_end_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2440833711) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) LinkSetEnterCost(link RID, enter_cost float64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&enter_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetEnterCost(link RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) LinkSetTravelCost(link RID, travel_cost float64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&travel_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetTravelCost(link RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) LinkSetOwnerId(link RID, owner_id int64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_owner_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) LinkGetOwnerId(link RID, ) int64 {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_owner_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) AgentCreate() RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) AgentSetAvoidanceEnabled(agent RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentGetAvoidanceEnabled(agent RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_get_avoidance_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) AgentSetMap(agent RID, map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentGetMap(agent RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_get_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814569979) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) AgentSetPaused(agent RID, paused bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&paused) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentGetPaused(agent RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_get_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) AgentSetNeighborDistance(agent RID, distance float64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_neighbor_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetMaxNeighbors(agent RID, count int64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_max_neighbors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetTimeHorizonAgents(agent RID, time_horizon float64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_time_horizon_agents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&time_horizon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetTimeHorizonObstacles(agent RID, time_horizon float64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_time_horizon_obstacles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&time_horizon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetRadius(agent RID, radius float64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetMaxSpeed(agent RID, max_speed float64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_max_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&max_speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetVelocityForced(agent RID, velocity Vector2, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_velocity_forced")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetVelocity(agent RID, velocity Vector2, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetPosition(agent RID, position Vector2, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentIsMapChanged(agent RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_is_map_changed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) AgentSetAvoidanceCallback(agent RID, callback Callable, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3379118538) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetAvoidanceLayers(agent RID, layers int64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetAvoidanceMask(agent RID, mask int64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) AgentSetAvoidancePriority(agent RID, priority float64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleCreate() RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) ObstacleSetAvoidanceEnabled(obstacle RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_avoidance_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleGetAvoidanceEnabled(obstacle RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_get_avoidance_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) ObstacleSetMap(obstacle RID, map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleGetMap(obstacle RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_get_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814569979) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer2D) ObstacleSetPaused(obstacle RID, paused bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&paused) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleGetPaused(obstacle RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_get_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer2D) ObstacleSetRadius(obstacle RID, radius float64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleSetVelocity(obstacle RID, velocity Vector2, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleSetPosition(obstacle RID, position Vector2, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleSetVertices(obstacle RID, vertices PackedVector2Array, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 29476483) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), vertices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ObstacleSetAvoidanceLayers(obstacle RID, layers int64, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_avoidance_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) ParseSourceGeometryData(navigation_polygon NavigationPolygon, source_geometry_data NavigationMeshSourceGeometryData2D, root_node Node, callback Callable, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("parse_source_geometry_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1176164995) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{navigation_polygon.AsCTypePtr(), source_geometry_data.AsCTypePtr(), root_node.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) BakeFromSourceGeometryData(navigation_polygon NavigationPolygon, source_geometry_data NavigationMeshSourceGeometryData2D, callback Callable, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bake_from_source_geometry_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2909414286) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{navigation_polygon.AsCTypePtr(), source_geometry_data.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) BakeFromSourceGeometryDataAsync(navigation_polygon NavigationPolygon, source_geometry_data NavigationMeshSourceGeometryData2D, callback Callable, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bake_from_source_geometry_data_async")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2909414286) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{navigation_polygon.AsCTypePtr(), source_geometry_data.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) FreeRid(rid RID, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("free_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) SetDebugEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer2D) GetDebugEnabled() bool {
  classNameV := StringNameFromStr("NavigationServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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
