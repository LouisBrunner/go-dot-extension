// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Mesh struct {
  Resource
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
  MeshArrayFormatArrayFlagCompressAttributes MeshArrayFormat = 536870912
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

func  (me *Mesh) SetLightmapSizeHint(size Vector2i, )  {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lightmap_size_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Mesh) GetLightmapSizeHint() Vector2i {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lightmap_size_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Mesh) GetAabb() AABB {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  var ret AABB
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Mesh) GetFaces() PackedVector3Array {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  var ret PackedVector3Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Mesh) GetSurfaceCount() int {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Mesh) SurfaceGetArrays(surf_idx int, ) Array {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_arrays")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Mesh) SurfaceGetBlendShapeArrays(surf_idx int, ) Array {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_blend_shape_arrays")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Mesh) SurfaceSetMaterial(surf_idx int, material Material, )  {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_set_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3671737478) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx), gdc.ConstTypePtr(material.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Mesh) SurfaceGetMaterial(surf_idx int, ) Material {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2897466400) // FIXME: should cache?
  var ret Material
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Mesh) CreatePlaceholder() Resource {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 121922552) // FIXME: should cache?
  var ret Resource
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Mesh) CreateTrimeshShape() ConcavePolygonShape3D {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_trimesh_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4160111210) // FIXME: should cache?
  var ret ConcavePolygonShape3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Mesh) CreateConvexShape(clean bool, simplify bool, ) ConvexPolygonShape3D {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_convex_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2529984628) // FIXME: should cache?
  var ret ConvexPolygonShape3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clean), gdc.ConstTypePtr(&simplify), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Mesh) CreateOutline(margin float32, ) Mesh {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1208642001) // FIXME: should cache?
  var ret Mesh
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Mesh) GenerateTriangleMesh() TriangleMesh {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_triangle_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3476533166) // FIXME: should cache?
  var ret TriangleMesh
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
