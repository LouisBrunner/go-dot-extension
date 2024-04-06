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

type ptrsForMeshDataToolList struct {
	fnClear             gdc.MethodBindPtr
	fnCreateFromSurface gdc.MethodBindPtr
	fnCommitToSurface   gdc.MethodBindPtr
	fnGetFormat         gdc.MethodBindPtr
	fnGetVertexCount    gdc.MethodBindPtr
	fnGetEdgeCount      gdc.MethodBindPtr
	fnGetFaceCount      gdc.MethodBindPtr
	fnSetVertex         gdc.MethodBindPtr
	fnGetVertex         gdc.MethodBindPtr
	fnSetVertexNormal   gdc.MethodBindPtr
	fnGetVertexNormal   gdc.MethodBindPtr
	fnSetVertexTangent  gdc.MethodBindPtr
	fnGetVertexTangent  gdc.MethodBindPtr
	fnSetVertexUv       gdc.MethodBindPtr
	fnGetVertexUv       gdc.MethodBindPtr
	fnSetVertexUv2      gdc.MethodBindPtr
	fnGetVertexUv2      gdc.MethodBindPtr
	fnSetVertexColor    gdc.MethodBindPtr
	fnGetVertexColor    gdc.MethodBindPtr
	fnSetVertexBones    gdc.MethodBindPtr
	fnGetVertexBones    gdc.MethodBindPtr
	fnSetVertexWeights  gdc.MethodBindPtr
	fnGetVertexWeights  gdc.MethodBindPtr
	fnSetVertexMeta     gdc.MethodBindPtr
	fnGetVertexMeta     gdc.MethodBindPtr
	fnGetVertexEdges    gdc.MethodBindPtr
	fnGetVertexFaces    gdc.MethodBindPtr
	fnGetEdgeVertex     gdc.MethodBindPtr
	fnGetEdgeFaces      gdc.MethodBindPtr
	fnSetEdgeMeta       gdc.MethodBindPtr
	fnGetEdgeMeta       gdc.MethodBindPtr
	fnGetFaceVertex     gdc.MethodBindPtr
	fnGetFaceEdge       gdc.MethodBindPtr
	fnSetFaceMeta       gdc.MethodBindPtr
	fnGetFaceMeta       gdc.MethodBindPtr
	fnGetFaceNormal     gdc.MethodBindPtr
	fnSetMaterial       gdc.MethodBindPtr
	fnGetMaterial       gdc.MethodBindPtr
}

var ptrsForMeshDataTool ptrsForMeshDataToolList

func initMeshDataToolPtrs(iface gdc.Interface) {

	className := StringNameFromStr("MeshDataTool")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("create_from_surface")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnCreateFromSurface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2727020678))
	}
	{
		methodName := StringNameFromStr("commit_to_surface")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnCommitToSurface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2021686445))
	}
	{
		methodName := StringNameFromStr("get_format")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_vertex_count")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetVertexCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_edge_count")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetEdgeCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_face_count")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetFaceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_vertex")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnSetVertex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1530502735))
	}
	{
		methodName := StringNameFromStr("get_vertex")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetVertex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
	}
	{
		methodName := StringNameFromStr("set_vertex_normal")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnSetVertexNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1530502735))
	}
	{
		methodName := StringNameFromStr("get_vertex_normal")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetVertexNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
	}
	{
		methodName := StringNameFromStr("set_vertex_tangent")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnSetVertexTangent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1104099133))
	}
	{
		methodName := StringNameFromStr("get_vertex_tangent")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetVertexTangent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1372055458))
	}
	{
		methodName := StringNameFromStr("set_vertex_uv")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnSetVertexUv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 163021252))
	}
	{
		methodName := StringNameFromStr("get_vertex_uv")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetVertexUv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
	{
		methodName := StringNameFromStr("set_vertex_uv2")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnSetVertexUv2 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 163021252))
	}
	{
		methodName := StringNameFromStr("get_vertex_uv2")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetVertexUv2 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
	{
		methodName := StringNameFromStr("set_vertex_color")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnSetVertexColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
	}
	{
		methodName := StringNameFromStr("get_vertex_color")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetVertexColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457211756))
	}
	{
		methodName := StringNameFromStr("set_vertex_bones")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnSetVertexBones = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3500328261))
	}
	{
		methodName := StringNameFromStr("get_vertex_bones")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetVertexBones = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1706082319))
	}
	{
		methodName := StringNameFromStr("set_vertex_weights")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnSetVertexWeights = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1345852415))
	}
	{
		methodName := StringNameFromStr("get_vertex_weights")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetVertexWeights = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1542882410))
	}
	{
		methodName := StringNameFromStr("set_vertex_meta")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnSetVertexMeta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2152698145))
	}
	{
		methodName := StringNameFromStr("get_vertex_meta")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetVertexMeta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4227898402))
	}
	{
		methodName := StringNameFromStr("get_vertex_edges")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetVertexEdges = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1706082319))
	}
	{
		methodName := StringNameFromStr("get_vertex_faces")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetVertexFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1706082319))
	}
	{
		methodName := StringNameFromStr("get_edge_vertex")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetEdgeVertex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3175239445))
	}
	{
		methodName := StringNameFromStr("get_edge_faces")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetEdgeFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1706082319))
	}
	{
		methodName := StringNameFromStr("set_edge_meta")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnSetEdgeMeta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2152698145))
	}
	{
		methodName := StringNameFromStr("get_edge_meta")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetEdgeMeta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4227898402))
	}
	{
		methodName := StringNameFromStr("get_face_vertex")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetFaceVertex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3175239445))
	}
	{
		methodName := StringNameFromStr("get_face_edge")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetFaceEdge = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3175239445))
	}
	{
		methodName := StringNameFromStr("set_face_meta")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnSetFaceMeta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2152698145))
	}
	{
		methodName := StringNameFromStr("get_face_meta")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetFaceMeta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4227898402))
	}
	{
		methodName := StringNameFromStr("get_face_normal")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetFaceNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
	}
	{
		methodName := StringNameFromStr("set_material")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
	}
	{
		methodName := StringNameFromStr("get_material")
		defer methodName.Destroy()
		ptrsForMeshDataTool.fnGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
	}

}

type MeshDataTool struct {
	RefCounted
}

func (me *MeshDataTool) BaseClass() string {
	return "MeshDataTool"
}

func NewMeshDataTool() *MeshDataTool {
	str := StringNameFromStr("MeshDataTool") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &MeshDataTool{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *MeshDataTool) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *MeshDataTool) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *MeshDataTool) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *MeshDataTool) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshDataTool) CreateFromSurface(mesh ArrayMesh, surface int64) Error {
	cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), gdc.ConstTypePtr(&surface)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&surface)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnCreateFromSurface), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *MeshDataTool) CommitToSurface(mesh ArrayMesh, compression_flags int64) Error {
	cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), gdc.ConstTypePtr(&compression_flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&compression_flags)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnCommitToSurface), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *MeshDataTool) GetFormat() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetFormat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MeshDataTool) GetVertexCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetVertexCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MeshDataTool) GetEdgeCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetEdgeCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MeshDataTool) GetFaceCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetFaceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MeshDataTool) SetVertex(idx int64, vertex Vector3) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), vertex.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnSetVertex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshDataTool) GetVertex(idx int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetVertex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) SetVertexNormal(idx int64, normal Vector3) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), normal.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnSetVertexNormal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshDataTool) GetVertexNormal(idx int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetVertexNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) SetVertexTangent(idx int64, tangent Plane) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), tangent.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnSetVertexTangent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshDataTool) GetVertexTangent(idx int64) Plane {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPlane()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetVertexTangent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) SetVertexUv(idx int64, uv Vector2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), uv.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnSetVertexUv), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshDataTool) GetVertexUv(idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetVertexUv), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) SetVertexUv2(idx int64, uv2 Vector2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), uv2.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnSetVertexUv2), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshDataTool) GetVertexUv2(idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetVertexUv2), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) SetVertexColor(idx int64, color Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnSetVertexColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshDataTool) GetVertexColor(idx int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetVertexColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) SetVertexBones(idx int64, bones PackedInt32Array) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), bones.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnSetVertexBones), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshDataTool) GetVertexBones(idx int64) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetVertexBones), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) SetVertexWeights(idx int64, weights PackedFloat32Array) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), weights.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnSetVertexWeights), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshDataTool) GetVertexWeights(idx int64) PackedFloat32Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedFloat32Array()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetVertexWeights), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) SetVertexMeta(idx int64, meta Variant) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), meta.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnSetVertexMeta), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshDataTool) GetVertexMeta(idx int64) Variant {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetVertexMeta), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) GetVertexEdges(idx int64) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetVertexEdges), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) GetVertexFaces(idx int64) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetVertexFaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) GetEdgeVertex(idx int64, vertex int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&vertex)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&idx)
	pinner.Pin(&vertex)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetEdgeVertex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MeshDataTool) GetEdgeFaces(idx int64) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetEdgeFaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) SetEdgeMeta(idx int64, meta Variant) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), meta.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnSetEdgeMeta), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshDataTool) GetEdgeMeta(idx int64) Variant {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetEdgeMeta), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) GetFaceVertex(idx int64, vertex int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&vertex)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&idx)
	pinner.Pin(&vertex)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetFaceVertex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MeshDataTool) GetFaceEdge(idx int64, edge int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&edge)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&idx)
	pinner.Pin(&edge)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetFaceEdge), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MeshDataTool) SetFaceMeta(idx int64, meta Variant) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), meta.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnSetFaceMeta), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshDataTool) GetFaceMeta(idx int64) Variant {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetFaceMeta), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) GetFaceNormal(idx int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetFaceNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *MeshDataTool) SetMaterial(material Material) {
	cargs := []gdc.ConstTypePtr{material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MeshDataTool) GetMaterial() Material {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshDataTool.fnGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
