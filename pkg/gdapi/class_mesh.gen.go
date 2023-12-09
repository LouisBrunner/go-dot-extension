// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Mesh struct {
  obj gdc.ObjectPtr
}

func (me *Mesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Mesh) BaseClass() string {
  return "Mesh"
}



// Enums

type MeshPrimitiveType int
const (
  MeshPrimitiveTypePrimitivePoints MeshPrimitiveType = 0
  MeshPrimitiveTypePrimitiveLines MeshPrimitiveType = 1
  MeshPrimitiveTypePrimitiveLineStrip MeshPrimitiveType = 2
  MeshPrimitiveTypePrimitiveTriangles MeshPrimitiveType = 3
  MeshPrimitiveTypePrimitiveTriangleStrip MeshPrimitiveType = 4
)

type MeshArrayType int
const (
  MeshArrayTypeArrayVertex MeshArrayType = 0
  MeshArrayTypeArrayNormal MeshArrayType = 1
  MeshArrayTypeArrayTangent MeshArrayType = 2
  MeshArrayTypeArrayColor MeshArrayType = 3
  MeshArrayTypeArrayTexUv MeshArrayType = 4
  MeshArrayTypeArrayTexUv2 MeshArrayType = 5
  MeshArrayTypeArrayCustom0 MeshArrayType = 6
  MeshArrayTypeArrayCustom1 MeshArrayType = 7
  MeshArrayTypeArrayCustom2 MeshArrayType = 8
  MeshArrayTypeArrayCustom3 MeshArrayType = 9
  MeshArrayTypeArrayBones MeshArrayType = 10
  MeshArrayTypeArrayWeights MeshArrayType = 11
  MeshArrayTypeArrayIndex MeshArrayType = 12
  MeshArrayTypeArrayMax MeshArrayType = 13
)

type MeshArrayCustomFormat int
const (
  MeshArrayCustomFormatArrayCustomRgba8Unorm MeshArrayCustomFormat = 0
  MeshArrayCustomFormatArrayCustomRgba8Snorm MeshArrayCustomFormat = 1
  MeshArrayCustomFormatArrayCustomRgHalf MeshArrayCustomFormat = 2
  MeshArrayCustomFormatArrayCustomRgbaHalf MeshArrayCustomFormat = 3
  MeshArrayCustomFormatArrayCustomRFloat MeshArrayCustomFormat = 4
  MeshArrayCustomFormatArrayCustomRgFloat MeshArrayCustomFormat = 5
  MeshArrayCustomFormatArrayCustomRgbFloat MeshArrayCustomFormat = 6
  MeshArrayCustomFormatArrayCustomRgbaFloat MeshArrayCustomFormat = 7
  MeshArrayCustomFormatArrayCustomMax MeshArrayCustomFormat = 8
)

type MeshArrayFormat int
const (
  MeshArrayFormatArrayFormatVertex MeshArrayFormat = 1
  MeshArrayFormatArrayFormatNormal MeshArrayFormat = 2
  MeshArrayFormatArrayFormatTangent MeshArrayFormat = 4
  MeshArrayFormatArrayFormatColor MeshArrayFormat = 8
  MeshArrayFormatArrayFormatTexUv MeshArrayFormat = 16
  MeshArrayFormatArrayFormatTexUv2 MeshArrayFormat = 32
  MeshArrayFormatArrayFormatCustom0 MeshArrayFormat = 64
  MeshArrayFormatArrayFormatCustom1 MeshArrayFormat = 128
  MeshArrayFormatArrayFormatCustom2 MeshArrayFormat = 256
  MeshArrayFormatArrayFormatCustom3 MeshArrayFormat = 512
  MeshArrayFormatArrayFormatBones MeshArrayFormat = 1024
  MeshArrayFormatArrayFormatWeights MeshArrayFormat = 2048
  MeshArrayFormatArrayFormatIndex MeshArrayFormat = 4096
  MeshArrayFormatArrayFormatBlendShapeMask MeshArrayFormat = 7
  MeshArrayFormatArrayFormatCustomBase MeshArrayFormat = 13
  MeshArrayFormatArrayFormatCustomBits MeshArrayFormat = 3
  MeshArrayFormatArrayFormatCustom0Shift MeshArrayFormat = 13
  MeshArrayFormatArrayFormatCustom1Shift MeshArrayFormat = 16
  MeshArrayFormatArrayFormatCustom2Shift MeshArrayFormat = 19
  MeshArrayFormatArrayFormatCustom3Shift MeshArrayFormat = 22
  MeshArrayFormatArrayFormatCustomMask MeshArrayFormat = 7
  MeshArrayFormatArrayCompressFlagsBase MeshArrayFormat = 25
  MeshArrayFormatArrayFlagUse2DVertices MeshArrayFormat = 33554432
  MeshArrayFormatArrayFlagUseDynamicUpdate MeshArrayFormat = 67108864
  MeshArrayFormatArrayFlagUse8BoneWeights MeshArrayFormat = 134217728
  MeshArrayFormatArrayFlagUsesEmptyVertexArray MeshArrayFormat = 268435456
)

type MeshBlendShapeMode int
const (
  MeshBlendShapeModeBlendShapeModeNormalized MeshBlendShapeMode = 0
  MeshBlendShapeModeBlendShapeModeRelative MeshBlendShapeMode = 1
)

func (me *Mesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Mesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Mesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Mesh) XGetSurfaceCount()  {
  panic("TODO: implement")
}

func  (me *Mesh) XSurfaceGetArrayLen(index int, )  {
  panic("TODO: implement")
}

func  (me *Mesh) XSurfaceGetArrayIndexLen(index int, )  {
  panic("TODO: implement")
}

func  (me *Mesh) XSurfaceGetArrays(index int, )  {
  panic("TODO: implement")
}

func  (me *Mesh) XSurfaceGetBlendShapeArrays(index int, )  {
  panic("TODO: implement")
}

func  (me *Mesh) XSurfaceGetLods(index int, )  {
  panic("TODO: implement")
}

func  (me *Mesh) XSurfaceGetFormat(index int, )  {
  panic("TODO: implement")
}

func  (me *Mesh) XSurfaceGetPrimitiveType(index int, )  {
  panic("TODO: implement")
}

func  (me *Mesh) XSurfaceSetMaterial(index int, material Material, )  {
  panic("TODO: implement")
}

func  (me *Mesh) XSurfaceGetMaterial(index int, )  {
  panic("TODO: implement")
}

func  (me *Mesh) XGetBlendShapeCount()  {
  panic("TODO: implement")
}

func  (me *Mesh) XGetBlendShapeName(index int, )  {
  panic("TODO: implement")
}

func  (me *Mesh) XSetBlendShapeName(index int, name StringName, )  {
  panic("TODO: implement")
}

func  (me *Mesh) XGetAabb()  {
  panic("TODO: implement")
}

func  (me *Mesh) SetLightmapSizeHint(size Vector2i, )  {
  panic("TODO: implement")
}

func  (me *Mesh) GetLightmapSizeHint()  {
  panic("TODO: implement")
}

func  (me *Mesh) GetAabb()  {
  panic("TODO: implement")
}

func  (me *Mesh) GetFaces()  {
  panic("TODO: implement")
}

func  (me *Mesh) GetSurfaceCount()  {
  panic("TODO: implement")
}

func  (me *Mesh) SurfaceGetArrays(surf_idx int, )  {
  panic("TODO: implement")
}

func  (me *Mesh) SurfaceGetBlendShapeArrays(surf_idx int, )  {
  panic("TODO: implement")
}

func  (me *Mesh) SurfaceSetMaterial(surf_idx int, material Material, )  {
  panic("TODO: implement")
}

func  (me *Mesh) SurfaceGetMaterial(surf_idx int, )  {
  panic("TODO: implement")
}

func  (me *Mesh) CreatePlaceholder()  {
  panic("TODO: implement")
}

func  (me *Mesh) CreateTrimeshShape()  {
  panic("TODO: implement")
}

func  (me *Mesh) CreateConvexShape(clean bool, simplify bool, )  {
  panic("TODO: implement")
}

func  (me *Mesh) CreateOutline(margin float32, )  {
  panic("TODO: implement")
}

func  (me *Mesh) GenerateTriangleMesh()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
