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



// Enums

func (me *ImporterMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ImporterMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ImporterMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ImporterMesh) AddBlendShape(name String, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetBlendShapeCount()  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetBlendShapeName(blend_shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) SetBlendShapeMode(mode MeshBlendShapeMode, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetBlendShapeMode()  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) AddSurface(primitive MeshPrimitiveType, arrays Array, blend_shapes Array, lods Dictionary, material Material, name String, flags int, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetSurfaceCount()  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetSurfacePrimitiveType(surface_idx int, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetSurfaceName(surface_idx int, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetSurfaceArrays(surface_idx int, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetSurfaceBlendShapeArrays(surface_idx int, blend_shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetSurfaceLodCount(surface_idx int, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetSurfaceLodSize(surface_idx int, lod_idx int, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetSurfaceLodIndices(surface_idx int, lod_idx int, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetSurfaceMaterial(surface_idx int, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetSurfaceFormat(surface_idx int, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) SetSurfaceName(surface_idx int, name String, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) SetSurfaceMaterial(surface_idx int, material Material, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GenerateLods(normal_merge_angle float32, normal_split_angle float32, bone_transform_array Array, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetMesh(base_mesh ArrayMesh, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) Clear()  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) SetLightmapSizeHint(size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *ImporterMesh) GetLightmapSizeHint()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
