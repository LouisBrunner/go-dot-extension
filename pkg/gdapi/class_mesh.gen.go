// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type Mesh struct {
  Resource
}

func (me *Mesh) BaseClass() string {
  return "Mesh"
}

func NewMesh() *Mesh {
  str := StringNameFromStr("Mesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Mesh{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Mesh) GetLightmapSizeHint() Vector2i {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lightmap_size_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Mesh) GetAabb() AABB {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Mesh) GetFaces() PackedVector3Array {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Mesh) GetSurfaceCount() int64 {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Mesh) SurfaceGetArrays(surf_idx int64, ) Array {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_arrays")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  pinner.Pin(&surf_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Mesh) SurfaceGetBlendShapeArrays(surf_idx int64, ) []Array {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_blend_shape_arrays")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&surf_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Array](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *Mesh) SurfaceSetMaterial(surf_idx int64, material Material, )  {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_set_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3671737478) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Mesh) SurfaceGetMaterial(surf_idx int64, ) Material {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2897466400) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()
  pinner.Pin(&surf_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Mesh) CreatePlaceholder() Resource {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 121922552) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewResource()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Mesh) CreateTrimeshShape() ConcavePolygonShape3D {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_trimesh_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4160111210) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewConcavePolygonShape3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Mesh) CreateConvexShape(clean bool, simplify bool, ) ConvexPolygonShape3D {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_convex_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2529984628) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clean) , gdc.ConstTypePtr(&simplify) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewConvexPolygonShape3D()
  pinner.Pin(&clean)
  pinner.Pin(&simplify)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Mesh) CreateOutline(margin float64, ) Mesh {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1208642001) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMesh()
  pinner.Pin(&margin)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Mesh) GenerateTriangleMesh() TriangleMesh {
  classNameV := StringNameFromStr("Mesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_triangle_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3476533166) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTriangleMesh()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
