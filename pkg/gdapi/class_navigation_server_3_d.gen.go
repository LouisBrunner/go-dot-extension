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

func  (me *NavigationServer3D) GetMaps() []RID {
  classNameV := StringNameFromStr("NavigationServer3D")
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
  sliceRet, err := ConvertArrayToSlice[RID](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *NavigationServer3D) MapCreate() RID {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) MapSetActive(map_ RID, active bool, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&active) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) MapIsActive(map_ RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) MapSetUp(map_ RID, up Vector3, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_up")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), up.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) MapGetUp(map_ RID, ) Vector3 {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_up")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 531438156) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer3D) MapSetCellSize(map_ RID, cell_size float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&cell_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) MapGetCellSize(map_ RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) MapSetCellHeight(map_ RID, cell_height float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_cell_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&cell_height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) MapGetCellHeight(map_ RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_cell_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer3D) MapSetUseEdgeConnections(map_ RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_use_edge_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) MapGetUseEdgeConnections(map_ RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) MapSetEdgeConnectionMargin(map_ RID, margin float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_edge_connection_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) MapGetEdgeConnectionMargin(map_ RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) MapSetLinkConnectionRadius(map_ RID, radius float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_set_link_connection_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) MapGetLinkConnectionRadius(map_ RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) MapGetPath(map_ RID, origin Vector3, destination Vector3, optimize bool, navigation_layers int64, ) PackedVector3Array {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1187418690) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), origin.AsCTypePtr(), destination.AsCTypePtr(), gdc.ConstTypePtr(&optimize) , gdc.ConstTypePtr(&navigation_layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()
  pinner.Pin(&optimize)
  pinner.Pin(&navigation_layers)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer3D) MapGetClosestPointToSegment(map_ RID, start Vector3, end Vector3, use_collision bool, ) Vector3 {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_closest_point_to_segment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3830095642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), start.AsCTypePtr(), end.AsCTypePtr(), gdc.ConstTypePtr(&use_collision) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&use_collision)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer3D) MapGetClosestPoint(map_ RID, to_point Vector3, ) Vector3 {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_closest_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2056183332) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), to_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer3D) MapGetClosestPointNormal(map_ RID, to_point Vector3, ) Vector3 {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_closest_point_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2056183332) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), to_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer3D) MapGetClosestPointOwner(map_ RID, to_point Vector3, ) RID {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_get_closest_point_owner")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 553364610) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), to_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer3D) MapGetLinks(map_ RID, ) []RID {
  classNameV := StringNameFromStr("NavigationServer3D")
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
  sliceRet, err := ConvertArrayToSlice[RID](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *NavigationServer3D) MapGetRegions(map_ RID, ) []RID {
  classNameV := StringNameFromStr("NavigationServer3D")
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
  sliceRet, err := ConvertArrayToSlice[RID](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *NavigationServer3D) MapGetAgents(map_ RID, ) []RID {
  classNameV := StringNameFromStr("NavigationServer3D")
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
  sliceRet, err := ConvertArrayToSlice[RID](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *NavigationServer3D) MapGetObstacles(map_ RID, ) []RID {
  classNameV := StringNameFromStr("NavigationServer3D")
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
  sliceRet, err := ConvertArrayToSlice[RID](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *NavigationServer3D) MapForceUpdate(map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("map_force_update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) QueryPath(parameters NavigationPathQueryParameters3D, result NavigationPathQueryResult3D, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("query_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3415008901) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), result.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) RegionCreate() RID {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) RegionSetEnabled(region RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) RegionGetEnabled(region RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) RegionSetUseEdgeConnections(region RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_use_edge_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) RegionGetUseEdgeConnections(region RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) RegionSetEnterCost(region RID, enter_cost float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&enter_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) RegionGetEnterCost(region RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) RegionSetTravelCost(region RID, travel_cost float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&travel_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) RegionGetTravelCost(region RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) RegionSetOwnerId(region RID, owner_id int64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_owner_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) RegionGetOwnerId(region RID, ) int64 {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) RegionOwnsPoint(region RID, point Vector3, ) bool {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_owns_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2360011153) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer3D) RegionSetMap(region RID, map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) RegionGetMap(region RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) RegionSetNavigationLayers(region RID, navigation_layers int64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&navigation_layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) RegionGetNavigationLayers(region RID, ) int64 {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) RegionSetTransform(region RID, transform Transform3D, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3935195649) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) RegionSetNavigationMesh(region RID, navigation_mesh NavigationMesh, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_set_navigation_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2764952978) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), navigation_mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) RegionBakeNavigationMesh(navigation_mesh NavigationMesh, root_node Node, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_bake_navigation_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1401173477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), root_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) RegionGetConnectionsCount(region RID, ) int64 {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) RegionGetConnectionPathwayStart(region RID, connection int64, ) Vector3 {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_connection_pathway_start")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3440143363) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&connection) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&connection)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer3D) RegionGetConnectionPathwayEnd(region RID, connection int64, ) Vector3 {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("region_get_connection_pathway_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3440143363) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&connection) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&connection)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer3D) LinkCreate() RID {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) LinkSetMap(link RID, map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) LinkGetMap(link RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) LinkSetEnabled(link RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) LinkGetEnabled(link RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) LinkSetBidirectional(link RID, bidirectional bool, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_bidirectional")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&bidirectional) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) LinkIsBidirectional(link RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) LinkSetNavigationLayers(link RID, navigation_layers int64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&navigation_layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) LinkGetNavigationLayers(link RID, ) int64 {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) LinkSetStartPosition(link RID, position Vector3, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) LinkGetStartPosition(link RID, ) Vector3 {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_start_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 531438156) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer3D) LinkSetEndPosition(link RID, position Vector3, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_end_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) LinkGetEndPosition(link RID, ) Vector3 {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_get_end_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 531438156) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationServer3D) LinkSetEnterCost(link RID, enter_cost float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&enter_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) LinkGetEnterCost(link RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) LinkSetTravelCost(link RID, travel_cost float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&travel_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) LinkGetTravelCost(link RID, ) float64 {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) LinkSetOwnerId(link RID, owner_id int64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("link_set_owner_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{link.AsCTypePtr(), gdc.ConstTypePtr(&owner_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) LinkGetOwnerId(link RID, ) int64 {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) AgentCreate() RID {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) AgentSetAvoidanceEnabled(agent RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentGetAvoidanceEnabled(agent RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) AgentSetUse3DAvoidance(agent RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_use_3d_avoidance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentGetUse3DAvoidance(agent RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_get_use_3d_avoidance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer3D) AgentSetMap(agent RID, map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentGetMap(agent RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) AgentSetPaused(agent RID, paused bool, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&paused) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentGetPaused(agent RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) AgentSetNeighborDistance(agent RID, distance float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_neighbor_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentSetMaxNeighbors(agent RID, count int64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_max_neighbors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentSetTimeHorizonAgents(agent RID, time_horizon float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_time_horizon_agents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&time_horizon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentSetTimeHorizonObstacles(agent RID, time_horizon float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_time_horizon_obstacles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&time_horizon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentSetRadius(agent RID, radius float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentSetHeight(agent RID, height float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentSetMaxSpeed(agent RID, max_speed float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_max_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&max_speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentSetVelocityForced(agent RID, velocity Vector3, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_velocity_forced")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentSetVelocity(agent RID, velocity Vector3, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentSetPosition(agent RID, position Vector3, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentIsMapChanged(agent RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) AgentSetAvoidanceCallback(agent RID, callback Callable, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3379118538) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentSetAvoidanceLayers(agent RID, layers int64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentSetAvoidanceMask(agent RID, mask int64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) AgentSetAvoidancePriority(agent RID, priority float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("agent_set_avoidance_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{agent.AsCTypePtr(), gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) ObstacleCreate() RID {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) ObstacleSetAvoidanceEnabled(obstacle RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_avoidance_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) ObstacleGetAvoidanceEnabled(obstacle RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) ObstacleSetUse3DAvoidance(obstacle RID, enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_use_3d_avoidance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) ObstacleGetUse3DAvoidance(obstacle RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_get_use_3d_avoidance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationServer3D) ObstacleSetMap(obstacle RID, map_ RID, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), map_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) ObstacleGetMap(obstacle RID, ) RID {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) ObstacleSetPaused(obstacle RID, paused bool, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&paused) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) ObstacleGetPaused(obstacle RID, ) bool {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) ObstacleSetRadius(obstacle RID, radius float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) ObstacleSetHeight(obstacle RID, height float64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) ObstacleSetVelocity(obstacle RID, velocity Vector3, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) ObstacleSetPosition(obstacle RID, position Vector3, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) ObstacleSetVertices(obstacle RID, vertices PackedVector3Array, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4030257846) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), vertices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) ObstacleSetAvoidanceLayers(obstacle RID, layers int64, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("obstacle_set_avoidance_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{obstacle.AsCTypePtr(), gdc.ConstTypePtr(&layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) ParseSourceGeometryData(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, root_node Node, callback Callable, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("parse_source_geometry_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 685862123) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), source_geometry_data.AsCTypePtr(), root_node.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) BakeFromSourceGeometryData(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, callback Callable, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bake_from_source_geometry_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2469318639) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), source_geometry_data.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) BakeFromSourceGeometryDataAsync(navigation_mesh NavigationMesh, source_geometry_data NavigationMeshSourceGeometryData3D, callback Callable, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bake_from_source_geometry_data_async")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2469318639) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), source_geometry_data.AsCTypePtr(), callback.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) FreeRid(rid RID, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("free_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{rid.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) SetActive(active bool, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) SetDebugEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationServer3D) GetDebugEnabled() bool {
  classNameV := StringNameFromStr("NavigationServer3D")
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

func  (me *NavigationServer3D) GetProcessInfo(process_info NavigationServer3DProcessInfo, ) int64 {
  classNameV := StringNameFromStr("NavigationServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1938440894) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&process_info) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&process_info)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals

type NavigationServer3DMapChangedSignalFn func(map_ RID, )

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
