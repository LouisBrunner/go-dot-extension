// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type NavigationMesh struct {
  obj gdc.ObjectPtr
}

func (me *NavigationMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationMesh) BaseClass() string {
  return "NavigationMesh"
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
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sample_partition_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2472437533) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sample_partition_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetSamplePartitionType() NavigationMeshSamplePartitionType {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sample_partition_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 833513918) // FIXME: should cache?
  var ret NavigationMeshSamplePartitionType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetParsedGeometryType(geometry_type NavigationMeshParsedGeometryType, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_parsed_geometry_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3064713163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&geometry_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetParsedGeometryType() NavigationMeshParsedGeometryType {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parsed_geometry_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3928011953) // FIXME: should cache?
  var ret NavigationMeshParsedGeometryType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetCollisionMask(mask int, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetCollisionMask() int {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetCollisionMaskValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetCollisionMaskValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetSourceGeometryMode(mask NavigationMeshSourceGeometryMode, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_source_geometry_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2700825194) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetSourceGeometryMode() NavigationMeshSourceGeometryMode {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_source_geometry_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2770484141) // FIXME: should cache?
  var ret NavigationMeshSourceGeometryMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetSourceGroupName(mask StringName, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_source_group_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mask.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetSourceGroupName() StringName {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_source_group_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetCellSize(cell_size float32, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cell_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetCellSize() float32 {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetCellHeight(cell_height float32, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cell_height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetCellHeight() float32 {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cell_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetAgentHeight(agent_height float32, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_agent_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&agent_height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetAgentHeight() float32 {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_agent_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetAgentRadius(agent_radius float32, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_agent_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&agent_radius), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetAgentRadius() float32 {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_agent_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetAgentMaxClimb(agent_max_climb float32, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_agent_max_climb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&agent_max_climb), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetAgentMaxClimb() float32 {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_agent_max_climb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetAgentMaxSlope(agent_max_slope float32, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_agent_max_slope")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&agent_max_slope), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetAgentMaxSlope() float32 {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_agent_max_slope")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetRegionMinSize(region_min_size float32, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_region_min_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&region_min_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetRegionMinSize() float32 {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_region_min_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetRegionMergeSize(region_merge_size float32, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_region_merge_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&region_merge_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetRegionMergeSize() float32 {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_region_merge_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetEdgeMaxLength(edge_max_length float32, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_edge_max_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edge_max_length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetEdgeMaxLength() float32 {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edge_max_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetEdgeMaxError(edge_max_error float32, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_edge_max_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edge_max_error), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetEdgeMaxError() float32 {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edge_max_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetVerticesPerPolygon(vertices_per_polygon float32, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertices_per_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertices_per_polygon), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetVerticesPerPolygon() float32 {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertices_per_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetDetailSampleDistance(detail_sample_dist float32, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_detail_sample_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail_sample_dist), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetDetailSampleDistance() float32 {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_detail_sample_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetDetailSampleMaxError(detail_sample_max_error float32, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_detail_sample_max_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&detail_sample_max_error), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetDetailSampleMaxError() float32 {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_detail_sample_max_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetFilterLowHangingObstacles(filter_low_hanging_obstacles bool, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filter_low_hanging_obstacles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter_low_hanging_obstacles), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetFilterLowHangingObstacles() bool {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_filter_low_hanging_obstacles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetFilterLedgeSpans(filter_ledge_spans bool, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filter_ledge_spans")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter_ledge_spans), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetFilterLedgeSpans() bool {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_filter_ledge_spans")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetFilterWalkableLowHeightSpans(filter_walkable_low_height_spans bool, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filter_walkable_low_height_spans")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter_walkable_low_height_spans), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetFilterWalkableLowHeightSpans() bool {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_filter_walkable_low_height_spans")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetFilterBakingAabb(baking_aabb AABB, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filter_baking_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 259215842) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(baking_aabb.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetFilterBakingAabb() AABB {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_filter_baking_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  var ret AABB
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetFilterBakingAabbOffset(baking_aabb_offset Vector3, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_filter_baking_aabb_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(baking_aabb_offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetFilterBakingAabbOffset() Vector3 {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_filter_baking_aabb_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) SetVertices(vertices PackedVector3Array, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 334873810) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(vertices.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetVertices() PackedVector3Array {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  var ret PackedVector3Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) AddPolygon(polygon PackedInt32Array, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(polygon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) GetPolygonCount() int {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polygon_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) GetPolygon(idx int, ) PackedInt32Array {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3668444399) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *NavigationMesh) ClearPolygons()  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_polygons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *NavigationMesh) CreateFromMesh(mesh Mesh, )  {
  classNameV := StringNameFromStr("NavigationMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_from_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 194775623) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mesh.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
