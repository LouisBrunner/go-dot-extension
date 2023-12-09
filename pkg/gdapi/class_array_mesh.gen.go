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



// Enums

func (me *ArrayMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ArrayMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ArrayMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ArrayMesh) AddBlendShape(name StringName, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) GetBlendShapeCount()  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) GetBlendShapeName(index int, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) SetBlendShapeName(index int, name StringName, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) ClearBlendShapes()  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) SetBlendShapeMode(mode MeshBlendShapeMode, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) GetBlendShapeMode()  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) AddSurfaceFromArrays(primitive MeshPrimitiveType, arrays Array, blend_shapes Array, lods Dictionary, flags MeshArrayFormat, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) ClearSurfaces()  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) SurfaceUpdateVertexRegion(surf_idx int, offset int, data PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) SurfaceUpdateAttributeRegion(surf_idx int, offset int, data PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) SurfaceUpdateSkinRegion(surf_idx int, offset int, data PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) SurfaceGetArrayLen(surf_idx int, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) SurfaceGetArrayIndexLen(surf_idx int, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) SurfaceGetFormat(surf_idx int, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) SurfaceGetPrimitiveType(surf_idx int, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) SurfaceFindByName(name String, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) SurfaceSetName(surf_idx int, name String, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) SurfaceGetName(surf_idx int, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) RegenNormalMaps()  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) LightmapUnwrap(transform Transform3D, texel_size float32, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) SetCustomAabb(aabb AABB, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) GetCustomAabb()  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) SetShadowMesh(mesh ArrayMesh, )  {
  panic("TODO: implement")
}

func  (me *ArrayMesh) GetShadowMesh()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
