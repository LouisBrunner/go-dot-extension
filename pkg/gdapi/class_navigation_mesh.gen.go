// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetSamplePartitionType()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetParsedGeometryType(geometry_type NavigationMeshParsedGeometryType, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetParsedGeometryType()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetCollisionMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetCollisionMask()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetCollisionMaskValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetCollisionMaskValue(layer_number int, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetSourceGeometryMode(mask NavigationMeshSourceGeometryMode, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetSourceGeometryMode()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetSourceGroupName(mask StringName, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetSourceGroupName()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetCellSize(cell_size float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetCellSize()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetCellHeight(cell_height float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetCellHeight()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetAgentHeight(agent_height float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetAgentHeight()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetAgentRadius(agent_radius float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetAgentRadius()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetAgentMaxClimb(agent_max_climb float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetAgentMaxClimb()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetAgentMaxSlope(agent_max_slope float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetAgentMaxSlope()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetRegionMinSize(region_min_size float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetRegionMinSize()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetRegionMergeSize(region_merge_size float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetRegionMergeSize()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetEdgeMaxLength(edge_max_length float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetEdgeMaxLength()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetEdgeMaxError(edge_max_error float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetEdgeMaxError()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetVerticesPerPolygon(vertices_per_polygon float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetVerticesPerPolygon()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetDetailSampleDistance(detail_sample_dist float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetDetailSampleDistance()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetDetailSampleMaxError(detail_sample_max_error float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetDetailSampleMaxError()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetFilterLowHangingObstacles(filter_low_hanging_obstacles bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetFilterLowHangingObstacles()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetFilterLedgeSpans(filter_ledge_spans bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetFilterLedgeSpans()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetFilterWalkableLowHeightSpans(filter_walkable_low_height_spans bool, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetFilterWalkableLowHeightSpans()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetFilterBakingAabb(baking_aabb AABB, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetFilterBakingAabb()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetFilterBakingAabbOffset(baking_aabb_offset Vector3, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetFilterBakingAabbOffset()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) SetVertices(vertices PackedVector3Array, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetVertices()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) AddPolygon(polygon PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetPolygonCount()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) GetPolygon(idx int, )  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) ClearPolygons()  {
  panic("TODO: implement")
}

func  (me *NavigationMesh) CreateFromMesh(mesh Mesh, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
