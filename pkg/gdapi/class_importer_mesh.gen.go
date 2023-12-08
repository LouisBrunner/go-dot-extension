// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImporterMesh struct {
  obj gdc.ObjectPtr
}

func (me *ImporterMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ImporterMesh) BaseClass() string {
  return "ImporterMesh"
}

func  (me *ImporterMesh) AddBlendShape(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetBlendShapeCount() { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetBlendShapeName(blend_shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) SetBlendShapeMode(mode MeshBlendShapeMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetBlendShapeMode() { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) AddSurface(primitive MeshPrimitiveType, arrays Array, blend_shapes Array, lods Dictionary, material Material, name String, flags int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetSurfaceCount() { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetSurfacePrimitiveType(surface_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetSurfaceName(surface_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetSurfaceArrays(surface_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetSurfaceBlendShapeArrays(surface_idx int, blend_shape_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetSurfaceLodCount(surface_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetSurfaceLodSize(surface_idx int, lod_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetSurfaceLodIndices(surface_idx int, lod_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetSurfaceMaterial(surface_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetSurfaceFormat(surface_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) SetSurfaceName(surface_idx int, name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) SetSurfaceMaterial(surface_idx int, material Material, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GenerateLods(normal_merge_angle float32, normal_split_angle float32, bone_transform_array Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetMesh(base_mesh ArrayMesh, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) SetLightmapSizeHint(size Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *ImporterMesh) GetLightmapSizeHint() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
