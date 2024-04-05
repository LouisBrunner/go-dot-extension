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

type ptrsForNavigationMeshList struct {
  fnSetSamplePartitionType gdc.MethodBindPtr
  fnGetSamplePartitionType gdc.MethodBindPtr
  fnSetParsedGeometryType gdc.MethodBindPtr
  fnGetParsedGeometryType gdc.MethodBindPtr
  fnSetCollisionMask gdc.MethodBindPtr
  fnGetCollisionMask gdc.MethodBindPtr
  fnSetCollisionMaskValue gdc.MethodBindPtr
  fnGetCollisionMaskValue gdc.MethodBindPtr
  fnSetSourceGeometryMode gdc.MethodBindPtr
  fnGetSourceGeometryMode gdc.MethodBindPtr
  fnSetSourceGroupName gdc.MethodBindPtr
  fnGetSourceGroupName gdc.MethodBindPtr
  fnSetCellSize gdc.MethodBindPtr
  fnGetCellSize gdc.MethodBindPtr
  fnSetCellHeight gdc.MethodBindPtr
  fnGetCellHeight gdc.MethodBindPtr
  fnSetAgentHeight gdc.MethodBindPtr
  fnGetAgentHeight gdc.MethodBindPtr
  fnSetAgentRadius gdc.MethodBindPtr
  fnGetAgentRadius gdc.MethodBindPtr
  fnSetAgentMaxClimb gdc.MethodBindPtr
  fnGetAgentMaxClimb gdc.MethodBindPtr
  fnSetAgentMaxSlope gdc.MethodBindPtr
  fnGetAgentMaxSlope gdc.MethodBindPtr
  fnSetRegionMinSize gdc.MethodBindPtr
  fnGetRegionMinSize gdc.MethodBindPtr
  fnSetRegionMergeSize gdc.MethodBindPtr
  fnGetRegionMergeSize gdc.MethodBindPtr
  fnSetEdgeMaxLength gdc.MethodBindPtr
  fnGetEdgeMaxLength gdc.MethodBindPtr
  fnSetEdgeMaxError gdc.MethodBindPtr
  fnGetEdgeMaxError gdc.MethodBindPtr
  fnSetVerticesPerPolygon gdc.MethodBindPtr
  fnGetVerticesPerPolygon gdc.MethodBindPtr
  fnSetDetailSampleDistance gdc.MethodBindPtr
  fnGetDetailSampleDistance gdc.MethodBindPtr
  fnSetDetailSampleMaxError gdc.MethodBindPtr
  fnGetDetailSampleMaxError gdc.MethodBindPtr
  fnSetFilterLowHangingObstacles gdc.MethodBindPtr
  fnGetFilterLowHangingObstacles gdc.MethodBindPtr
  fnSetFilterLedgeSpans gdc.MethodBindPtr
  fnGetFilterLedgeSpans gdc.MethodBindPtr
  fnSetFilterWalkableLowHeightSpans gdc.MethodBindPtr
  fnGetFilterWalkableLowHeightSpans gdc.MethodBindPtr
  fnSetFilterBakingAabb gdc.MethodBindPtr
  fnGetFilterBakingAabb gdc.MethodBindPtr
  fnSetFilterBakingAabbOffset gdc.MethodBindPtr
  fnGetFilterBakingAabbOffset gdc.MethodBindPtr
  fnSetVertices gdc.MethodBindPtr
  fnGetVertices gdc.MethodBindPtr
  fnAddPolygon gdc.MethodBindPtr
  fnGetPolygonCount gdc.MethodBindPtr
  fnGetPolygon gdc.MethodBindPtr
  fnClearPolygons gdc.MethodBindPtr
  fnCreateFromMesh gdc.MethodBindPtr
  fnClear gdc.MethodBindPtr
}

var ptrsForNavigationMesh ptrsForNavigationMeshList

func initNavigationMeshPtrs(iface gdc.Interface) {

  className := StringNameFromStr("NavigationMesh")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_sample_partition_type")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetSamplePartitionType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2472437533))
  }
  {
    methodName := StringNameFromStr("get_sample_partition_type")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetSamplePartitionType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 833513918))
  }
  {
    methodName := StringNameFromStr("set_parsed_geometry_type")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetParsedGeometryType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3064713163))
  }
  {
    methodName := StringNameFromStr("get_parsed_geometry_type")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetParsedGeometryType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3928011953))
  }
  {
    methodName := StringNameFromStr("set_collision_mask")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_collision_mask")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_collision_mask_value")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_collision_mask_value")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetCollisionMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_source_geometry_mode")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetSourceGeometryMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2700825194))
  }
  {
    methodName := StringNameFromStr("get_source_geometry_mode")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetSourceGeometryMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2770484141))
  }
  {
    methodName := StringNameFromStr("set_source_group_name")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetSourceGroupName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("get_source_group_name")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetSourceGroupName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
  }
  {
    methodName := StringNameFromStr("set_cell_size")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_cell_size")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_cell_height")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetCellHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_cell_height")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetCellHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_agent_height")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetAgentHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_agent_height")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetAgentHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_agent_radius")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetAgentRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_agent_radius")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetAgentRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
  {
    methodName := StringNameFromStr("set_agent_max_climb")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetAgentMaxClimb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_agent_max_climb")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetAgentMaxClimb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_agent_max_slope")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetAgentMaxSlope = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_agent_max_slope")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetAgentMaxSlope = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_region_min_size")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetRegionMinSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_region_min_size")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetRegionMinSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_region_merge_size")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetRegionMergeSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_region_merge_size")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetRegionMergeSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_edge_max_length")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetEdgeMaxLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_edge_max_length")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetEdgeMaxLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_edge_max_error")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetEdgeMaxError = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_edge_max_error")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetEdgeMaxError = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_vertices_per_polygon")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetVerticesPerPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_vertices_per_polygon")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetVerticesPerPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_detail_sample_distance")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetDetailSampleDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_detail_sample_distance")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetDetailSampleDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_detail_sample_max_error")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetDetailSampleMaxError = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_detail_sample_max_error")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetDetailSampleMaxError = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_filter_low_hanging_obstacles")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetFilterLowHangingObstacles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_filter_low_hanging_obstacles")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetFilterLowHangingObstacles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_filter_ledge_spans")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetFilterLedgeSpans = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_filter_ledge_spans")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetFilterLedgeSpans = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_filter_walkable_low_height_spans")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetFilterWalkableLowHeightSpans = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_filter_walkable_low_height_spans")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetFilterWalkableLowHeightSpans = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_filter_baking_aabb")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetFilterBakingAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 259215842))
  }
  {
    methodName := StringNameFromStr("get_filter_baking_aabb")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetFilterBakingAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1068685055))
  }
  {
    methodName := StringNameFromStr("set_filter_baking_aabb_offset")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetFilterBakingAabbOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_filter_baking_aabb_offset")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetFilterBakingAabbOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_vertices")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnSetVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 334873810))
  }
  {
    methodName := StringNameFromStr("get_vertices")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 497664490))
  }
  {
    methodName := StringNameFromStr("add_polygon")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnAddPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
  }
  {
    methodName := StringNameFromStr("get_polygon_count")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetPolygonCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_polygon")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnGetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3668444399))
  }
  {
    methodName := StringNameFromStr("clear_polygons")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnClearPolygons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("create_from_mesh")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnCreateFromMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 194775623))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForNavigationMesh.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
}

type NavigationMesh struct {
  Resource
}

func (me *NavigationMesh) BaseClass() string {
  return "NavigationMesh"
}

func NewNavigationMesh() *NavigationMesh {
  str := StringNameFromStr("NavigationMesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NavigationMesh{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type NavigationMeshSamplePartitionType int
const (
  NavigationMeshSamplePartitionTypeSamplePartitionWatershed NavigationMeshSamplePartitionType = 0
  NavigationMeshSamplePartitionTypeSamplePartitionMonotone NavigationMeshSamplePartitionType = 1
  NavigationMeshSamplePartitionTypeSamplePartitionLayers NavigationMeshSamplePartitionType = 2
  NavigationMeshSamplePartitionTypeSamplePartitionMax NavigationMeshSamplePartitionType = 3
)

type NavigationMeshParsedGeometryType int
const (
  NavigationMeshParsedGeometryTypeParsedGeometryMeshInstances NavigationMeshParsedGeometryType = 0
  NavigationMeshParsedGeometryTypeParsedGeometryStaticColliders NavigationMeshParsedGeometryType = 1
  NavigationMeshParsedGeometryTypeParsedGeometryBoth NavigationMeshParsedGeometryType = 2
  NavigationMeshParsedGeometryTypeParsedGeometryMax NavigationMeshParsedGeometryType = 3
)

type NavigationMeshSourceGeometryMode int
const (
  NavigationMeshSourceGeometryModeSourceGeometryRootNodeChildren NavigationMeshSourceGeometryMode = 0
  NavigationMeshSourceGeometryModeSourceGeometryGroupsWithChildren NavigationMeshSourceGeometryMode = 1
  NavigationMeshSourceGeometryModeSourceGeometryGroupsExplicit NavigationMeshSourceGeometryMode = 2
  NavigationMeshSourceGeometryModeSourceGeometryMax NavigationMeshSourceGeometryMode = 3
)

func (me *NavigationMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationMesh) SetSamplePartitionType(sample_partition_type NavigationMeshSamplePartitionType, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sample_partition_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetSamplePartitionType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetSamplePartitionType() NavigationMeshSamplePartitionType {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationMeshSamplePartitionType

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetSamplePartitionType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationMesh) SetParsedGeometryType(geometry_type NavigationMeshParsedGeometryType, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&geometry_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetParsedGeometryType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetParsedGeometryType() NavigationMeshParsedGeometryType {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationMeshParsedGeometryType

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetParsedGeometryType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationMesh) SetCollisionMask(mask int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetCollisionMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetCollisionMaskValue(layer_number int64, value bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetCollisionMaskValue(layer_number int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetCollisionMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetSourceGeometryMode(mask NavigationMeshSourceGeometryMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetSourceGeometryMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetSourceGeometryMode() NavigationMeshSourceGeometryMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationMeshSourceGeometryMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetSourceGeometryMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationMesh) SetSourceGroupName(mask StringName, )  {
  cargs := []gdc.ConstTypePtr{mask.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetSourceGroupName), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetSourceGroupName() StringName {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetSourceGroupName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationMesh) SetCellSize(cell_size float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cell_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetCellSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetCellSize() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetCellSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetCellHeight(cell_height float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cell_height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetCellHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetCellHeight() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetCellHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetAgentHeight(agent_height float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&agent_height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetAgentHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetAgentHeight() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetAgentHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetAgentRadius(agent_radius float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&agent_radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetAgentRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetAgentRadius() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetAgentRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetAgentMaxClimb(agent_max_climb float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&agent_max_climb) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetAgentMaxClimb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetAgentMaxClimb() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetAgentMaxClimb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetAgentMaxSlope(agent_max_slope float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&agent_max_slope) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetAgentMaxSlope), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetAgentMaxSlope() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetAgentMaxSlope), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetRegionMinSize(region_min_size float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&region_min_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetRegionMinSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetRegionMinSize() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetRegionMinSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetRegionMergeSize(region_merge_size float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&region_merge_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetRegionMergeSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetRegionMergeSize() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetRegionMergeSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetEdgeMaxLength(edge_max_length float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edge_max_length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetEdgeMaxLength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetEdgeMaxLength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetEdgeMaxLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetEdgeMaxError(edge_max_error float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edge_max_error) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetEdgeMaxError), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetEdgeMaxError() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetEdgeMaxError), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetVerticesPerPolygon(vertices_per_polygon float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertices_per_polygon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetVerticesPerPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetVerticesPerPolygon() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetVerticesPerPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetDetailSampleDistance(detail_sample_dist float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail_sample_dist) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetDetailSampleDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetDetailSampleDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetDetailSampleDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetDetailSampleMaxError(detail_sample_max_error float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail_sample_max_error) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetDetailSampleMaxError), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetDetailSampleMaxError() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetDetailSampleMaxError), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetFilterLowHangingObstacles(filter_low_hanging_obstacles bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter_low_hanging_obstacles) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetFilterLowHangingObstacles), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetFilterLowHangingObstacles() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetFilterLowHangingObstacles), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetFilterLedgeSpans(filter_ledge_spans bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter_ledge_spans) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetFilterLedgeSpans), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetFilterLedgeSpans() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetFilterLedgeSpans), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetFilterWalkableLowHeightSpans(filter_walkable_low_height_spans bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter_walkable_low_height_spans) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetFilterWalkableLowHeightSpans), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetFilterWalkableLowHeightSpans() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetFilterWalkableLowHeightSpans), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) SetFilterBakingAabb(baking_aabb AABB, )  {
  cargs := []gdc.ConstTypePtr{baking_aabb.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetFilterBakingAabb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetFilterBakingAabb() AABB {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetFilterBakingAabb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationMesh) SetFilterBakingAabbOffset(baking_aabb_offset Vector3, )  {
  cargs := []gdc.ConstTypePtr{baking_aabb_offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetFilterBakingAabbOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetFilterBakingAabbOffset() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetFilterBakingAabbOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationMesh) SetVertices(vertices PackedVector3Array, )  {
  cargs := []gdc.ConstTypePtr{vertices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnSetVertices), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetVertices() PackedVector3Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetVertices), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationMesh) AddPolygon(polygon PackedInt32Array, )  {
  cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnAddPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) GetPolygonCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetPolygonCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationMesh) GetPolygon(idx int64, ) PackedInt32Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnGetPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationMesh) ClearPolygons()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnClearPolygons), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) CreateFromMesh(mesh Mesh, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnCreateFromMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationMesh) Clear()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationMesh.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
