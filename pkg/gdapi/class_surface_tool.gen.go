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



// Enums

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

func (me *SurfaceTool) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SurfaceTool) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SurfaceTool) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SurfaceTool) SetSkinWeightCount(count SurfaceToolSkinWeightCount, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) GetSkinWeightCount()  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) SetCustomFormat(channel_index int, format SurfaceToolCustomFormat, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) GetCustomFormat(channel_index int, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) Begin(primitive MeshPrimitiveType, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) AddVertex(vertex Vector3, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) SetColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) SetNormal(normal Vector3, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) SetTangent(tangent Plane, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) SetUv(uv Vector2, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) SetUv2(uv2 Vector2, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) SetBones(bones PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) SetWeights(weights PackedFloat32Array, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) SetCustom(channel_index int, custom_color Color, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) SetSmoothGroup(index int, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) AddTriangleFan(vertices PackedVector3Array, uvs PackedVector2Array, colors PackedColorArray, uv2s PackedVector2Array, normals PackedVector3Array, tangents Plane, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) AddIndex(index int, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) Index()  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) Deindex()  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) GenerateNormals(flip bool, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) GenerateTangents()  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) OptimizeIndicesForCache()  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) GetAabb()  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) GenerateLod(nd_threshold float32, target_index_count int, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) SetMaterial(material Material, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) GetPrimitiveType()  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) Clear()  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) CreateFrom(existing Mesh, surface int, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) CreateFromBlendShape(existing Mesh, surface int, blend_shape String, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) AppendFrom(existing Mesh, surface int, transform Transform3D, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) Commit(existing ArrayMesh, flags int, )  {
  panic("TODO: implement")
}

func  (me *SurfaceTool) CommitToArrays()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
