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

func  (me *NavigationMesh) SetSamplePartitionType(sample_partition_type NavigationMeshSamplePartitionType, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetSamplePartitionType() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetParsedGeometryType(geometry_type NavigationMeshParsedGeometryType, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetParsedGeometryType() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetCollisionMask(mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetCollisionMask() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetCollisionMaskValue(layer_number int, value bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetCollisionMaskValue(layer_number int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetSourceGeometryMode(mask NavigationMeshSourceGeometryMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetSourceGeometryMode() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetSourceGroupName(mask StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetSourceGroupName() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetCellSize(cell_size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetCellSize() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetCellHeight(cell_height float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetCellHeight() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetAgentHeight(agent_height float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetAgentHeight() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetAgentRadius(agent_radius float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetAgentRadius() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetAgentMaxClimb(agent_max_climb float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetAgentMaxClimb() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetAgentMaxSlope(agent_max_slope float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetAgentMaxSlope() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetRegionMinSize(region_min_size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetRegionMinSize() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetRegionMergeSize(region_merge_size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetRegionMergeSize() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetEdgeMaxLength(edge_max_length float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetEdgeMaxLength() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetEdgeMaxError(edge_max_error float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetEdgeMaxError() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetVerticesPerPolygon(vertices_per_polygon float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetVerticesPerPolygon() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetDetailSampleDistance(detail_sample_dist float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetDetailSampleDistance() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetDetailSampleMaxError(detail_sample_max_error float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetDetailSampleMaxError() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetFilterLowHangingObstacles(filter_low_hanging_obstacles bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetFilterLowHangingObstacles() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetFilterLedgeSpans(filter_ledge_spans bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetFilterLedgeSpans() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetFilterWalkableLowHeightSpans(filter_walkable_low_height_spans bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetFilterWalkableLowHeightSpans() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetFilterBakingAabb(baking_aabb AABB, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetFilterBakingAabb() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetFilterBakingAabbOffset(baking_aabb_offset Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetFilterBakingAabbOffset() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) SetVertices(vertices PackedVector3Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetVertices() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) AddPolygon(polygon PackedInt32Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetPolygonCount() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) GetPolygon(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) ClearPolygons() { // TODO: return value
  // TODO: implement
}

func  (me *NavigationMesh) CreateFromMesh(mesh Mesh, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
