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

type ptrsForImporterMeshList struct {
	fnAddBlendShape              gdc.MethodBindPtr
	fnGetBlendShapeCount         gdc.MethodBindPtr
	fnGetBlendShapeName          gdc.MethodBindPtr
	fnSetBlendShapeMode          gdc.MethodBindPtr
	fnGetBlendShapeMode          gdc.MethodBindPtr
	fnAddSurface                 gdc.MethodBindPtr
	fnGetSurfaceCount            gdc.MethodBindPtr
	fnGetSurfacePrimitiveType    gdc.MethodBindPtr
	fnGetSurfaceName             gdc.MethodBindPtr
	fnGetSurfaceArrays           gdc.MethodBindPtr
	fnGetSurfaceBlendShapeArrays gdc.MethodBindPtr
	fnGetSurfaceLodCount         gdc.MethodBindPtr
	fnGetSurfaceLodSize          gdc.MethodBindPtr
	fnGetSurfaceLodIndices       gdc.MethodBindPtr
	fnGetSurfaceMaterial         gdc.MethodBindPtr
	fnGetSurfaceFormat           gdc.MethodBindPtr
	fnSetSurfaceName             gdc.MethodBindPtr
	fnSetSurfaceMaterial         gdc.MethodBindPtr
	fnGenerateLods               gdc.MethodBindPtr
	fnGetMesh                    gdc.MethodBindPtr
	fnClear                      gdc.MethodBindPtr
	fnSetLightmapSizeHint        gdc.MethodBindPtr
	fnGetLightmapSizeHint        gdc.MethodBindPtr
}

var ptrsForImporterMesh ptrsForImporterMeshList

func initImporterMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ImporterMesh")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_blend_shape")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnAddBlendShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_blend_shape_count")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetBlendShapeCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_blend_shape_name")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetBlendShapeName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_blend_shape_mode")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnSetBlendShapeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 227983991))
	}
	{
		methodName := StringNameFromStr("get_blend_shape_mode")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetBlendShapeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 836485024))
	}
	{
		methodName := StringNameFromStr("add_surface")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnAddSurface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740448849))
	}
	{
		methodName := StringNameFromStr("get_surface_count")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetSurfaceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_surface_primitive_type")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetSurfacePrimitiveType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3552571330))
	}
	{
		methodName := StringNameFromStr("get_surface_name")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetSurfaceName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_surface_arrays")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetSurfaceArrays = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 663333327))
	}
	{
		methodName := StringNameFromStr("get_surface_blend_shape_arrays")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetSurfaceBlendShapeArrays = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2345056839))
	}
	{
		methodName := StringNameFromStr("get_surface_lod_count")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetSurfaceLodCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_surface_lod_size")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetSurfaceLodSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
	}
	{
		methodName := StringNameFromStr("get_surface_lod_indices")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetSurfaceLodIndices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1265128013))
	}
	{
		methodName := StringNameFromStr("get_surface_material")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetSurfaceMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2897466400))
	}
	{
		methodName := StringNameFromStr("get_surface_format")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetSurfaceFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_surface_name")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnSetSurfaceName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("set_surface_material")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnSetSurfaceMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3671737478))
	}
	{
		methodName := StringNameFromStr("generate_lods")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGenerateLods = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2491878677))
	}
	{
		methodName := StringNameFromStr("get_mesh")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1457573577))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_lightmap_size_hint")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnSetLightmapSizeHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
	}
	{
		methodName := StringNameFromStr("get_lightmap_size_hint")
		defer methodName.Destroy()
		ptrsForImporterMesh.fnGetLightmapSizeHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}

}

type ImporterMesh struct {
	Resource
}

func (me *ImporterMesh) BaseClass() string {
	return "ImporterMesh"
}

func NewImporterMesh() *ImporterMesh {
	str := StringNameFromStr("ImporterMesh") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ImporterMesh{}
	obj.SetBaseObject(objPtr)
	return obj
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

func (me *ImporterMesh) AddBlendShape(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnAddBlendShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMesh) GetBlendShapeCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetBlendShapeCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ImporterMesh) GetBlendShapeName(blend_shape_idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&blend_shape_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&blend_shape_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetBlendShapeName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ImporterMesh) SetBlendShapeMode(mode MeshBlendShapeMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnSetBlendShapeMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMesh) GetBlendShapeMode() MeshBlendShapeMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret MeshBlendShapeMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetBlendShapeMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ImporterMesh) AddSurface(primitive MeshPrimitiveType, arrays Array, blend_shapes []Array, lods Dictionary, material Material, name String, flags int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&primitive), arrays.AsCTypePtr(), gdc.ConstTypePtr(&blend_shapes), lods.AsCTypePtr(), material.AsCTypePtr(), name.AsCTypePtr(), gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnAddSurface), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMesh) GetSurfaceCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetSurfaceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ImporterMesh) GetSurfacePrimitiveType(surface_idx int64) MeshPrimitiveType {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret MeshPrimitiveType
	pinner.Pin(&surface_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetSurfacePrimitiveType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ImporterMesh) GetSurfaceName(surface_idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&surface_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetSurfaceName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ImporterMesh) GetSurfaceArrays(surface_idx int64) Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	pinner.Pin(&surface_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetSurfaceArrays), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ImporterMesh) GetSurfaceBlendShapeArrays(surface_idx int64, blend_shape_idx int64) Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), gdc.ConstTypePtr(&blend_shape_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	pinner.Pin(&surface_idx)
	pinner.Pin(&blend_shape_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetSurfaceBlendShapeArrays), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ImporterMesh) GetSurfaceLodCount(surface_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&surface_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetSurfaceLodCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ImporterMesh) GetSurfaceLodSize(surface_idx int64, lod_idx int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), gdc.ConstTypePtr(&lod_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&surface_idx)
	pinner.Pin(&lod_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetSurfaceLodSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ImporterMesh) GetSurfaceLodIndices(surface_idx int64, lod_idx int64) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), gdc.ConstTypePtr(&lod_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()
	pinner.Pin(&surface_idx)
	pinner.Pin(&lod_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetSurfaceLodIndices), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ImporterMesh) GetSurfaceMaterial(surface_idx int64) Material {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMaterial()
	pinner.Pin(&surface_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetSurfaceMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ImporterMesh) GetSurfaceFormat(surface_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&surface_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetSurfaceFormat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ImporterMesh) SetSurfaceName(surface_idx int64, name String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnSetSurfaceName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMesh) SetSurfaceMaterial(surface_idx int64, material Material) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnSetSurfaceMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMesh) GenerateLods(normal_merge_angle float64, normal_split_angle float64, bone_transform_array Array) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normal_merge_angle), gdc.ConstTypePtr(&normal_split_angle), bone_transform_array.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGenerateLods), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMesh) GetMesh(base_mesh ArrayMesh) ArrayMesh {
	cargs := []gdc.ConstTypePtr{base_mesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArrayMesh()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ImporterMesh) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMesh) SetLightmapSizeHint(size Vector2i) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnSetLightmapSizeHint), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ImporterMesh) GetLightmapSizeHint() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImporterMesh.fnGetLightmapSizeHint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
