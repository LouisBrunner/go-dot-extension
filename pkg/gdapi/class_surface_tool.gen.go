// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SurfaceTool struct {
  obj gdc.ObjectPtr
}

func (me *SurfaceTool) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SurfaceTool) BaseClass() string {
  return "SurfaceTool"
}



// Enums

type SurfaceToolCustomFormat int
const (
  SurfaceToolCustomFormatCustomRgba8Unorm SurfaceToolCustomFormat = 0
  SurfaceToolCustomFormatCustomRgba8Snorm SurfaceToolCustomFormat = 1
  SurfaceToolCustomFormatCustomRgHalf SurfaceToolCustomFormat = 2
  SurfaceToolCustomFormatCustomRgbaHalf SurfaceToolCustomFormat = 3
  SurfaceToolCustomFormatCustomRFloat SurfaceToolCustomFormat = 4
  SurfaceToolCustomFormatCustomRgFloat SurfaceToolCustomFormat = 5
  SurfaceToolCustomFormatCustomRgbFloat SurfaceToolCustomFormat = 6
  SurfaceToolCustomFormatCustomRgbaFloat SurfaceToolCustomFormat = 7
  SurfaceToolCustomFormatCustomMax SurfaceToolCustomFormat = 8
)

type SurfaceToolSkinWeightCount int
const (
  SurfaceToolSkinWeightCountSkin4Weights SurfaceToolSkinWeightCount = 0
  SurfaceToolSkinWeightCountSkin8Weights SurfaceToolSkinWeightCount = 1
)

func (me *SurfaceTool) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SurfaceTool) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SurfaceTool) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SurfaceTool) SetSkinWeightCount(count SurfaceToolSkinWeightCount, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skin_weight_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 618679515) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) GetSkinWeightCount() SurfaceToolSkinWeightCount {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skin_weight_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1072401130) // FIXME: should cache?
  var ret SurfaceToolSkinWeightCount
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SurfaceTool) SetCustomFormat(channel_index int, format SurfaceToolCustomFormat, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4087759856) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel_index), gdc.ConstTypePtr(&format), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) GetCustomFormat(channel_index int, ) SurfaceToolCustomFormat {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 839863283) // FIXME: should cache?
  var ret SurfaceToolCustomFormat
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SurfaceTool) Begin(primitive MeshPrimitiveType, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2230304113) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&primitive), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) AddVertex(vertex Vector3, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_vertex")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(vertex.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) SetColor(color Color, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) SetNormal(normal Vector3, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(normal.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) SetTangent(tangent Plane, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tangent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3505987427) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tangent.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) SetUv(uv Vector2, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(uv.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) SetUv2(uv2 Vector2, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uv2")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(uv2.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) SetBones(bones PackedInt32Array, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bones")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bones.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) SetWeights(weights PackedFloat32Array, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_weights")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2899603908) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(weights.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) SetCustom(channel_index int, custom_color Color, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel_index), gdc.ConstTypePtr(custom_color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) SetSmoothGroup(index int, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_smooth_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) AddTriangleFan(vertices PackedVector3Array, uvs PackedVector2Array, colors PackedColorArray, uv2s PackedVector2Array, normals PackedVector3Array, tangents Plane, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_triangle_fan")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 297960074) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(vertices.AsCTypePtr()), gdc.ConstTypePtr(uvs.AsCTypePtr()), gdc.ConstTypePtr(colors.AsCTypePtr()), gdc.ConstTypePtr(uv2s.AsCTypePtr()), gdc.ConstTypePtr(normals.AsCTypePtr()), gdc.ConstTypePtr(tangents.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) AddIndex(index int, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) Index()  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) Deindex()  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("deindex")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) GenerateNormals(flip bool, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_normals")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 107499316) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) GenerateTangents()  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_tangents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) OptimizeIndicesForCache()  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("optimize_indices_for_cache")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) GetAabb() AABB {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  var ret AABB
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SurfaceTool) GenerateLod(nd_threshold float32, target_index_count int, ) PackedInt32Array {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_lod")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1894448909) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&nd_threshold), gdc.ConstTypePtr(&target_index_count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SurfaceTool) SetMaterial(material Material, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(material.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) GetPrimitiveType() MeshPrimitiveType {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_primitive_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 768822145) // FIXME: should cache?
  var ret MeshPrimitiveType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SurfaceTool) Clear()  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) CreateFrom(existing Mesh, surface int, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_from")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1767024570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(existing.AsCTypePtr()), gdc.ConstTypePtr(&surface), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) CreateFromBlendShape(existing Mesh, surface int, blend_shape String, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_from_blend_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1306185582) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(existing.AsCTypePtr()), gdc.ConstTypePtr(&surface), gdc.ConstTypePtr(blend_shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) AppendFrom(existing Mesh, surface int, transform Transform3D, )  {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("append_from")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2217967155) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(existing.AsCTypePtr()), gdc.ConstTypePtr(&surface), gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SurfaceTool) Commit(existing ArrayMesh, flags int, ) ArrayMesh {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("commit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4107864055) // FIXME: should cache?
  var ret ArrayMesh
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(existing.AsCTypePtr()), gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SurfaceTool) CommitToArrays() Array {
  classNameV := StringNameFromStr("SurfaceTool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("commit_to_arrays")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
