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

type ptrsForMeshList struct {
	fnXGetSurfaceCount            gdc.MethodBindPtr
	fnXSurfaceGetArrayLen         gdc.MethodBindPtr
	fnXSurfaceGetArrayIndexLen    gdc.MethodBindPtr
	fnXSurfaceGetArrays           gdc.MethodBindPtr
	fnXSurfaceGetBlendShapeArrays gdc.MethodBindPtr
	fnXSurfaceGetLods             gdc.MethodBindPtr
	fnXSurfaceGetFormat           gdc.MethodBindPtr
	fnXSurfaceGetPrimitiveType    gdc.MethodBindPtr
	fnXSurfaceSetMaterial         gdc.MethodBindPtr
	fnXSurfaceGetMaterial         gdc.MethodBindPtr
	fnXGetBlendShapeCount         gdc.MethodBindPtr
	fnXGetBlendShapeName          gdc.MethodBindPtr
	fnXSetBlendShapeName          gdc.MethodBindPtr
	fnXGetAabb                    gdc.MethodBindPtr
	fnSetLightmapSizeHint         gdc.MethodBindPtr
	fnGetLightmapSizeHint         gdc.MethodBindPtr
	fnGetAabb                     gdc.MethodBindPtr
	fnGetFaces                    gdc.MethodBindPtr
	fnGetSurfaceCount             gdc.MethodBindPtr
	fnSurfaceGetArrays            gdc.MethodBindPtr
	fnSurfaceGetBlendShapeArrays  gdc.MethodBindPtr
	fnSurfaceSetMaterial          gdc.MethodBindPtr
	fnSurfaceGetMaterial          gdc.MethodBindPtr
	fnCreatePlaceholder           gdc.MethodBindPtr
	fnCreateTrimeshShape          gdc.MethodBindPtr
	fnCreateConvexShape           gdc.MethodBindPtr
	fnCreateOutline               gdc.MethodBindPtr
	fnGenerateTriangleMesh        gdc.MethodBindPtr
}

var ptrsForMesh ptrsForMeshList

func initMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Mesh")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_lightmap_size_hint")
		defer methodName.Destroy()
		ptrsForMesh.fnSetLightmapSizeHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
	}
	{
		methodName := StringNameFromStr("get_lightmap_size_hint")
		defer methodName.Destroy()
		ptrsForMesh.fnGetLightmapSizeHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("get_aabb")
		defer methodName.Destroy()
		ptrsForMesh.fnGetAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1068685055))
	}
	{
		methodName := StringNameFromStr("get_faces")
		defer methodName.Destroy()
		ptrsForMesh.fnGetFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 497664490))
	}
	{
		methodName := StringNameFromStr("get_surface_count")
		defer methodName.Destroy()
		ptrsForMesh.fnGetSurfaceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("surface_get_arrays")
		defer methodName.Destroy()
		ptrsForMesh.fnSurfaceGetArrays = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 663333327))
	}
	{
		methodName := StringNameFromStr("surface_get_blend_shape_arrays")
		defer methodName.Destroy()
		ptrsForMesh.fnSurfaceGetBlendShapeArrays = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 663333327))
	}
	{
		methodName := StringNameFromStr("surface_set_material")
		defer methodName.Destroy()
		ptrsForMesh.fnSurfaceSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3671737478))
	}
	{
		methodName := StringNameFromStr("surface_get_material")
		defer methodName.Destroy()
		ptrsForMesh.fnSurfaceGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2897466400))
	}
	{
		methodName := StringNameFromStr("create_placeholder")
		defer methodName.Destroy()
		ptrsForMesh.fnCreatePlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 121922552))
	}
	{
		methodName := StringNameFromStr("create_trimesh_shape")
		defer methodName.Destroy()
		ptrsForMesh.fnCreateTrimeshShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4160111210))
	}
	{
		methodName := StringNameFromStr("create_convex_shape")
		defer methodName.Destroy()
		ptrsForMesh.fnCreateConvexShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2529984628))
	}
	{
		methodName := StringNameFromStr("create_outline")
		defer methodName.Destroy()
		ptrsForMesh.fnCreateOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1208642001))
	}
	{
		methodName := StringNameFromStr("generate_triangle_mesh")
		defer methodName.Destroy()
		ptrsForMesh.fnGenerateTriangleMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3476533166))
	}
}

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
	MeshPrimitiveTypePrimitivePoints        MeshPrimitiveType = 0
	MeshPrimitiveTypePrimitiveLines         MeshPrimitiveType = 1
	MeshPrimitiveTypePrimitiveLineStrip     MeshPrimitiveType = 2
	MeshPrimitiveTypePrimitiveTriangles     MeshPrimitiveType = 3
	MeshPrimitiveTypePrimitiveTriangleStrip MeshPrimitiveType = 4
)

type MeshArrayType int

const (
	MeshArrayTypeArrayVertex  MeshArrayType = 0
	MeshArrayTypeArrayNormal  MeshArrayType = 1
	MeshArrayTypeArrayTangent MeshArrayType = 2
	MeshArrayTypeArrayColor   MeshArrayType = 3
	MeshArrayTypeArrayTexUv   MeshArrayType = 4
	MeshArrayTypeArrayTexUv2  MeshArrayType = 5
	MeshArrayTypeArrayCustom0 MeshArrayType = 6
	MeshArrayTypeArrayCustom1 MeshArrayType = 7
	MeshArrayTypeArrayCustom2 MeshArrayType = 8
	MeshArrayTypeArrayCustom3 MeshArrayType = 9
	MeshArrayTypeArrayBones   MeshArrayType = 10
	MeshArrayTypeArrayWeights MeshArrayType = 11
	MeshArrayTypeArrayIndex   MeshArrayType = 12
	MeshArrayTypeArrayMax     MeshArrayType = 13
)

type MeshArrayCustomFormat int

const (
	MeshArrayCustomFormatArrayCustomRgba8Unorm MeshArrayCustomFormat = 0
	MeshArrayCustomFormatArrayCustomRgba8Snorm MeshArrayCustomFormat = 1
	MeshArrayCustomFormatArrayCustomRgHalf     MeshArrayCustomFormat = 2
	MeshArrayCustomFormatArrayCustomRgbaHalf   MeshArrayCustomFormat = 3
	MeshArrayCustomFormatArrayCustomRFloat     MeshArrayCustomFormat = 4
	MeshArrayCustomFormatArrayCustomRgFloat    MeshArrayCustomFormat = 5
	MeshArrayCustomFormatArrayCustomRgbFloat   MeshArrayCustomFormat = 6
	MeshArrayCustomFormatArrayCustomRgbaFloat  MeshArrayCustomFormat = 7
	MeshArrayCustomFormatArrayCustomMax        MeshArrayCustomFormat = 8
)

type MeshArrayFormat int

const (
	MeshArrayFormatArrayFormatVertex             MeshArrayFormat = 1
	MeshArrayFormatArrayFormatNormal             MeshArrayFormat = 2
	MeshArrayFormatArrayFormatTangent            MeshArrayFormat = 4
	MeshArrayFormatArrayFormatColor              MeshArrayFormat = 8
	MeshArrayFormatArrayFormatTexUv              MeshArrayFormat = 16
	MeshArrayFormatArrayFormatTexUv2             MeshArrayFormat = 32
	MeshArrayFormatArrayFormatCustom0            MeshArrayFormat = 64
	MeshArrayFormatArrayFormatCustom1            MeshArrayFormat = 128
	MeshArrayFormatArrayFormatCustom2            MeshArrayFormat = 256
	MeshArrayFormatArrayFormatCustom3            MeshArrayFormat = 512
	MeshArrayFormatArrayFormatBones              MeshArrayFormat = 1024
	MeshArrayFormatArrayFormatWeights            MeshArrayFormat = 2048
	MeshArrayFormatArrayFormatIndex              MeshArrayFormat = 4096
	MeshArrayFormatArrayFormatBlendShapeMask     MeshArrayFormat = 7
	MeshArrayFormatArrayFormatCustomBase         MeshArrayFormat = 13
	MeshArrayFormatArrayFormatCustomBits         MeshArrayFormat = 3
	MeshArrayFormatArrayFormatCustom0Shift       MeshArrayFormat = 13
	MeshArrayFormatArrayFormatCustom1Shift       MeshArrayFormat = 16
	MeshArrayFormatArrayFormatCustom2Shift       MeshArrayFormat = 19
	MeshArrayFormatArrayFormatCustom3Shift       MeshArrayFormat = 22
	MeshArrayFormatArrayFormatCustomMask         MeshArrayFormat = 7
	MeshArrayFormatArrayCompressFlagsBase        MeshArrayFormat = 25
	MeshArrayFormatArrayFlagUse2DVertices        MeshArrayFormat = 33554432
	MeshArrayFormatArrayFlagUseDynamicUpdate     MeshArrayFormat = 67108864
	MeshArrayFormatArrayFlagUse8BoneWeights      MeshArrayFormat = 134217728
	MeshArrayFormatArrayFlagUsesEmptyVertexArray MeshArrayFormat = 268435456
	MeshArrayFormatArrayFlagCompressAttributes   MeshArrayFormat = 536870912
)

type MeshBlendShapeMode int

const (
	MeshBlendShapeModeBlendShapeModeNormalized MeshBlendShapeMode = 0
	MeshBlendShapeModeBlendShapeModeRelative   MeshBlendShapeMode = 1
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

func (me *Mesh) SetLightmapSizeHint(size Vector2i) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMesh.fnSetLightmapSizeHint), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Mesh) GetLightmapSizeHint() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMesh.fnGetLightmapSizeHint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Mesh) GetAabb() AABB {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAABB()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMesh.fnGetAabb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Mesh) GetFaces() PackedVector3Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector3Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMesh.fnGetFaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Mesh) GetSurfaceCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMesh.fnGetSurfaceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Mesh) SurfaceGetArrays(surf_idx int64) Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	pinner.Pin(&surf_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMesh.fnSurfaceGetArrays), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Mesh) SurfaceGetBlendShapeArrays(surf_idx int64) []Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&surf_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMesh.fnSurfaceGetBlendShapeArrays), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Array](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Mesh) SurfaceSetMaterial(surf_idx int64, material Material) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx), material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMesh.fnSurfaceSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Mesh) SurfaceGetMaterial(surf_idx int64) Material {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMaterial()
	pinner.Pin(&surf_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMesh.fnSurfaceGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Mesh) CreatePlaceholder() Resource {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewResource()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMesh.fnCreatePlaceholder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Mesh) CreateTrimeshShape() ConcavePolygonShape3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewConcavePolygonShape3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMesh.fnCreateTrimeshShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Mesh) CreateConvexShape(clean bool, simplify bool) ConvexPolygonShape3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clean), gdc.ConstTypePtr(&simplify)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewConvexPolygonShape3D()
	pinner.Pin(&clean)
	pinner.Pin(&simplify)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMesh.fnCreateConvexShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Mesh) CreateOutline(margin float64) Mesh {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMesh()
	pinner.Pin(&margin)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMesh.fnCreateOutline), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Mesh) GenerateTriangleMesh() TriangleMesh {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTriangleMesh()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMesh.fnGenerateTriangleMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
