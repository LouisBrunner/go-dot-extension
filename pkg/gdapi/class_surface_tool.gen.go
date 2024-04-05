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

type ptrsForSurfaceToolList struct {
  fnSetSkinWeightCount gdc.MethodBindPtr
  fnGetSkinWeightCount gdc.MethodBindPtr
  fnSetCustomFormat gdc.MethodBindPtr
  fnGetCustomFormat gdc.MethodBindPtr
  fnBegin gdc.MethodBindPtr
  fnAddVertex gdc.MethodBindPtr
  fnSetColor gdc.MethodBindPtr
  fnSetNormal gdc.MethodBindPtr
  fnSetTangent gdc.MethodBindPtr
  fnSetUv gdc.MethodBindPtr
  fnSetUv2 gdc.MethodBindPtr
  fnSetBones gdc.MethodBindPtr
  fnSetWeights gdc.MethodBindPtr
  fnSetCustom gdc.MethodBindPtr
  fnSetSmoothGroup gdc.MethodBindPtr
  fnAddTriangleFan gdc.MethodBindPtr
  fnAddIndex gdc.MethodBindPtr
  fnIndex gdc.MethodBindPtr
  fnDeindex gdc.MethodBindPtr
  fnGenerateNormals gdc.MethodBindPtr
  fnGenerateTangents gdc.MethodBindPtr
  fnOptimizeIndicesForCache gdc.MethodBindPtr
  fnGetAabb gdc.MethodBindPtr
  fnGenerateLod gdc.MethodBindPtr
  fnSetMaterial gdc.MethodBindPtr
  fnGetPrimitiveType gdc.MethodBindPtr
  fnClear gdc.MethodBindPtr
  fnCreateFrom gdc.MethodBindPtr
  fnCreateFromBlendShape gdc.MethodBindPtr
  fnAppendFrom gdc.MethodBindPtr
  fnCommit gdc.MethodBindPtr
  fnCommitToArrays gdc.MethodBindPtr
}

var ptrsForSurfaceTool ptrsForSurfaceToolList

func initSurfaceToolPtrs(iface gdc.Interface) {

  className := StringNameFromStr("SurfaceTool")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_skin_weight_count")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnSetSkinWeightCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 618679515))
  }
  {
    methodName := StringNameFromStr("get_skin_weight_count")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnGetSkinWeightCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1072401130))
  }
  {
    methodName := StringNameFromStr("set_custom_format")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnSetCustomFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4087759856))
  }
  {
    methodName := StringNameFromStr("get_custom_format")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnGetCustomFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 839863283))
  }
  {
    methodName := StringNameFromStr("begin")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2230304113))
  }
  {
    methodName := StringNameFromStr("add_vertex")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnAddVertex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("set_color")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("set_normal")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnSetNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("set_tangent")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnSetTangent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3505987427))
  }
  {
    methodName := StringNameFromStr("set_uv")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnSetUv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("set_uv2")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnSetUv2 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("set_bones")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnSetBones = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
  }
  {
    methodName := StringNameFromStr("set_weights")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnSetWeights = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2899603908))
  }
  {
    methodName := StringNameFromStr("set_custom")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnSetCustom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
  }
  {
    methodName := StringNameFromStr("set_smooth_group")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnSetSmoothGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("add_triangle_fan")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnAddTriangleFan = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2235017613))
  }
  {
    methodName := StringNameFromStr("add_index")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnAddIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("index")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("deindex")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnDeindex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("generate_normals")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnGenerateNormals = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 107499316))
  }
  {
    methodName := StringNameFromStr("generate_tangents")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnGenerateTangents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("optimize_indices_for_cache")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnOptimizeIndicesForCache = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_aabb")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnGetAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1068685055))
  }
  {
    methodName := StringNameFromStr("generate_lod")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnGenerateLod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1938056459))
  }
  {
    methodName := StringNameFromStr("set_material")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
  }
  {
    methodName := StringNameFromStr("get_primitive_type")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnGetPrimitiveType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 768822145))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("create_from")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnCreateFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1767024570))
  }
  {
    methodName := StringNameFromStr("create_from_blend_shape")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnCreateFromBlendShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1306185582))
  }
  {
    methodName := StringNameFromStr("append_from")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnAppendFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2217967155))
  }
  {
    methodName := StringNameFromStr("commit")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnCommit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4107864055))
  }
  {
    methodName := StringNameFromStr("commit_to_arrays")
    defer methodName.Destroy()
    ptrsForSurfaceTool.fnCommitToArrays = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
}

type SurfaceTool struct {
  RefCounted
}

func (me *SurfaceTool) BaseClass() string {
  return "SurfaceTool"
}

func NewSurfaceTool() *SurfaceTool {
  str := StringNameFromStr("SurfaceTool") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SurfaceTool{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnSetSkinWeightCount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) GetSkinWeightCount() SurfaceToolSkinWeightCount {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret SurfaceToolSkinWeightCount

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnGetSkinWeightCount), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SurfaceTool) SetCustomFormat(channel_index int64, format SurfaceToolCustomFormat, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel_index) , gdc.ConstTypePtr(&format) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnSetCustomFormat), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) GetCustomFormat(channel_index int64, ) SurfaceToolCustomFormat {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret SurfaceToolCustomFormat
  pinner.Pin(&channel_index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnGetCustomFormat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SurfaceTool) Begin(primitive MeshPrimitiveType, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&primitive) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnBegin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) AddVertex(vertex Vector3, )  {
  cargs := []gdc.ConstTypePtr{vertex.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnAddVertex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) SetColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) SetNormal(normal Vector3, )  {
  cargs := []gdc.ConstTypePtr{normal.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnSetNormal), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) SetTangent(tangent Plane, )  {
  cargs := []gdc.ConstTypePtr{tangent.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnSetTangent), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) SetUv(uv Vector2, )  {
  cargs := []gdc.ConstTypePtr{uv.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnSetUv), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) SetUv2(uv2 Vector2, )  {
  cargs := []gdc.ConstTypePtr{uv2.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnSetUv2), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) SetBones(bones PackedInt32Array, )  {
  cargs := []gdc.ConstTypePtr{bones.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnSetBones), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) SetWeights(weights PackedFloat32Array, )  {
  cargs := []gdc.ConstTypePtr{weights.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnSetWeights), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) SetCustom(channel_index int64, custom_color Color, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel_index) , custom_color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnSetCustom), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) SetSmoothGroup(index int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnSetSmoothGroup), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) AddTriangleFan(vertices PackedVector3Array, uvs PackedVector2Array, colors PackedColorArray, uv2s PackedVector2Array, normals PackedVector3Array, tangents []Plane, )  {
  cargs := []gdc.ConstTypePtr{vertices.AsCTypePtr(), uvs.AsCTypePtr(), colors.AsCTypePtr(), uv2s.AsCTypePtr(), normals.AsCTypePtr(), gdc.ConstTypePtr(&tangents) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnAddTriangleFan), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) AddIndex(index int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnAddIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) Index()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) Deindex()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnDeindex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) GenerateNormals(flip bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnGenerateNormals), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) GenerateTangents()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnGenerateTangents), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) OptimizeIndicesForCache()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnOptimizeIndicesForCache), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) GetAabb() AABB {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnGetAabb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SurfaceTool) GenerateLod(nd_threshold float64, target_index_count int64, ) PackedInt32Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&nd_threshold) , gdc.ConstTypePtr(&target_index_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()
  pinner.Pin(&nd_threshold)
  pinner.Pin(&target_index_count)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnGenerateLod), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SurfaceTool) SetMaterial(material Material, )  {
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) GetPrimitiveType() MeshPrimitiveType {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MeshPrimitiveType

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnGetPrimitiveType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *SurfaceTool) Clear()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) CreateFrom(existing Mesh, surface int64, )  {
  cargs := []gdc.ConstTypePtr{existing.AsCTypePtr(), gdc.ConstTypePtr(&surface) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnCreateFrom), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) CreateFromBlendShape(existing Mesh, surface int64, blend_shape String, )  {
  cargs := []gdc.ConstTypePtr{existing.AsCTypePtr(), gdc.ConstTypePtr(&surface) , blend_shape.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnCreateFromBlendShape), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) AppendFrom(existing Mesh, surface int64, transform Transform3D, )  {
  cargs := []gdc.ConstTypePtr{existing.AsCTypePtr(), gdc.ConstTypePtr(&surface) , transform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnAppendFrom), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SurfaceTool) Commit(existing ArrayMesh, flags int64, ) ArrayMesh {
  cargs := []gdc.ConstTypePtr{existing.AsCTypePtr(), gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArrayMesh()
  pinner.Pin(&flags)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnCommit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SurfaceTool) CommitToArrays() Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSurfaceTool.fnCommitToArrays), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
