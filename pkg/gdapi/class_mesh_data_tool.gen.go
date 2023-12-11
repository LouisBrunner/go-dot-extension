// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MeshDataTool struct {
  obj gdc.ObjectPtr
}

func (me *MeshDataTool) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MeshDataTool) BaseClass() string {
  return "MeshDataTool"
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

func  (me *MeshDataTool) Clear()  {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshDataTool) CreateFromSurface(mesh ArrayMesh, surface int, ) Error {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_from_surface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2727020678) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mesh.AsCTypePtr()), gdc.ConstTypePtr(&surface), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) CommitToSurface(mesh ArrayMesh, ) Error {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("commit_to_surface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3521099812) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mesh.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) GetFormat() int {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) GetVertexCount() int {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertex_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) GetEdgeCount() int {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edge_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) GetFaceCount() int {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_face_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) SetVertex(idx int, vertex Vector3, )  {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertex")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530502735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(vertex.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshDataTool) GetVertex(idx int, ) Vector3 {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertex")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) SetVertexNormal(idx int, normal Vector3, )  {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertex_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530502735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(normal.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshDataTool) GetVertexNormal(idx int, ) Vector3 {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertex_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) SetVertexTangent(idx int, tangent Plane, )  {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertex_tangent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1104099133) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(tangent.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshDataTool) GetVertexTangent(idx int, ) Plane {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertex_tangent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1372055458) // FIXME: should cache?
  var ret Plane
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) SetVertexUv(idx int, uv Vector2, )  {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertex_uv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 163021252) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(uv.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshDataTool) GetVertexUv(idx int, ) Vector2 {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertex_uv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) SetVertexUv2(idx int, uv2 Vector2, )  {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertex_uv2")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 163021252) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(uv2.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshDataTool) GetVertexUv2(idx int, ) Vector2 {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertex_uv2")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) SetVertexColor(idx int, color Color, )  {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertex_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshDataTool) GetVertexColor(idx int, ) Color {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertex_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) SetVertexBones(idx int, bones PackedInt32Array, )  {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertex_bones")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3500328261) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(bones.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshDataTool) GetVertexBones(idx int, ) PackedInt32Array {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertex_bones")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1706082319) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) SetVertexWeights(idx int, weights PackedFloat32Array, )  {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertex_weights")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1345852415) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(weights.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshDataTool) GetVertexWeights(idx int, ) PackedFloat32Array {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertex_weights")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1542882410) // FIXME: should cache?
  var ret PackedFloat32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) SetVertexMeta(idx int, meta Variant, )  {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertex_meta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2152698145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(meta.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshDataTool) GetVertexMeta(idx int, ) Variant {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertex_meta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4227898402) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) GetVertexEdges(idx int, ) PackedInt32Array {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertex_edges")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1706082319) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) GetVertexFaces(idx int, ) PackedInt32Array {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertex_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1706082319) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) GetEdgeVertex(idx int, vertex int, ) int {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edge_vertex")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175239445) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&vertex), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) GetEdgeFaces(idx int, ) PackedInt32Array {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edge_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1706082319) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) SetEdgeMeta(idx int, meta Variant, )  {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_edge_meta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2152698145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(meta.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshDataTool) GetEdgeMeta(idx int, ) Variant {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_edge_meta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4227898402) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) GetFaceVertex(idx int, vertex int, ) int {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_face_vertex")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175239445) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&vertex), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) GetFaceEdge(idx int, edge int, ) int {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_face_edge")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175239445) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&edge), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) SetFaceMeta(idx int, meta Variant, )  {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_face_meta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2152698145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(meta.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshDataTool) GetFaceMeta(idx int, ) Variant {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_face_meta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4227898402) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) GetFaceNormal(idx int, ) Vector3 {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_face_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *MeshDataTool) SetMaterial(material Material, )  {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(material.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MeshDataTool) GetMaterial() Material {
  classNameV := StringNameFromStr("MeshDataTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 5934680) // FIXME: should cache?
  var ret Material
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
