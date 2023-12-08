// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ArrayMesh struct {
  obj gdc.ObjectPtr
}

func (me *ArrayMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ArrayMesh) BaseClass() string {
  return "ArrayMesh"
}

func  (me *ArrayMesh) AddBlendShape(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) GetBlendShapeCount() { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) GetBlendShapeName(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) SetBlendShapeName(index int, name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) ClearBlendShapes() { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) SetBlendShapeMode(mode MeshBlendShapeMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) GetBlendShapeMode() { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) AddSurfaceFromArrays(primitive MeshPrimitiveType, arrays Array, blend_shapes Array, lods Dictionary, flags MeshArrayFormat, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) ClearSurfaces() { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) SurfaceUpdateVertexRegion(surf_idx int, offset int, data PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) SurfaceUpdateAttributeRegion(surf_idx int, offset int, data PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) SurfaceUpdateSkinRegion(surf_idx int, offset int, data PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) SurfaceGetArrayLen(surf_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) SurfaceGetArrayIndexLen(surf_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) SurfaceGetFormat(surf_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) SurfaceGetPrimitiveType(surf_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) SurfaceFindByName(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) SurfaceSetName(surf_idx int, name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) SurfaceGetName(surf_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) RegenNormalMaps() { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) LightmapUnwrap(transform Transform3D, texel_size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) SetCustomAabb(aabb AABB, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) GetCustomAabb() { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) SetShadowMesh(mesh ArrayMesh, ) { // TODO: return value
  // TODO: implement
}

func  (me *ArrayMesh) GetShadowMesh() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
