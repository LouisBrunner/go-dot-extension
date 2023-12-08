// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SurfaceTool struct {
  obj gdc.ObjectPtr
}

func (me *SurfaceTool) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SurfaceTool) BaseClass() string {
  return "SurfaceTool"
}

type SurfaceToolCustomFormat int
const (
  SurfaceToolCustomFormatCustomRgba8Unorm SurfaceToolCustomFormat = 0
  SurfaceToolCustomFormatCustomRgba8Snorm SurfaceToolCustomFormat = 1
  SurfaceToolCustomFormatCustomRgHalf SurfaceToolCustomFormat = 2
  SurfaceToolCustomFormatCustomRgbaHalf SurfaceToolCustomFormat = 3
  SurfaceToolCustomFormatCustomRFloat SurfaceToolCustomFormat = 4
  SurfaceToolCustomFormatCustomRgFloat SurfaceToolCustomFormat = 5
  SurfaceToolCustomFormatCustomRgbFloat SurfaceToolCustomFormat = 6
  SurfaceToolCustomFormatCustomRgbaFloat SurfaceToolCustomFormat = 7
  SurfaceToolCustomFormatCustomMax SurfaceToolCustomFormat = 8
)

type SurfaceToolSkinWeightCount int
const (
  SurfaceToolSkinWeightCountSkin4Weights SurfaceToolSkinWeightCount = 0
  SurfaceToolSkinWeightCountSkin8Weights SurfaceToolSkinWeightCount = 1
)

func  (me *SurfaceTool) SetSkinWeightCount(count SurfaceToolSkinWeightCount, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) GetSkinWeightCount() { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) SetCustomFormat(channel_index int, format SurfaceToolCustomFormat, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) GetCustomFormat(channel_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) Begin(primitive MeshPrimitiveType, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) AddVertex(vertex Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) SetColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) SetNormal(normal Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) SetTangent(tangent Plane, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) SetUv(uv Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) SetUv2(uv2 Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) SetBones(bones PackedInt32Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) SetWeights(weights PackedFloat32Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) SetCustom(channel_index int, custom_color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) SetSmoothGroup(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) AddTriangleFan(vertices PackedVector3Array, uvs PackedVector2Array, colors PackedColorArray, uv2s PackedVector2Array, normals PackedVector3Array, tangents Plane, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) AddIndex(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) Index() { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) Deindex() { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) GenerateNormals(flip bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) GenerateTangents() { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) OptimizeIndicesForCache() { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) GetAabb() { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) GenerateLod(nd_threshold float32, target_index_count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) SetMaterial(material Material, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) GetPrimitiveType() { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) CreateFrom(existing Mesh, surface int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) CreateFromBlendShape(existing Mesh, surface int, blend_shape String, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) AppendFrom(existing Mesh, surface int, transform Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) Commit(existing ArrayMesh, flags int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SurfaceTool) CommitToArrays() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
