// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *ImporterMesh) AddBlendShape(name String, )  {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_blend_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImporterMesh) GetBlendShapeCount() int64 {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_shape_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ImporterMesh) GetBlendShapeName(blend_shape_idx int64, ) String {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_shape_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&blend_shape_idx), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ImporterMesh) SetBlendShapeMode(mode MeshBlendShapeMode, )  {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_shape_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 227983991) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImporterMesh) GetBlendShapeMode() MeshBlendShapeMode {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_shape_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 836485024) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret MeshBlendShapeMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ImporterMesh) AddSurface(primitive MeshPrimitiveType, arrays Array, blend_shapes []Array, lods Dictionary, material Material, name String, flags int64, )  {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_surface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740448849) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&primitive), gdc.ConstTypePtr(arrays.AsCTypePtr()), gdc.ConstTypePtr(&blend_shapes), gdc.ConstTypePtr(lods.AsCTypePtr()), gdc.ConstTypePtr(material.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&flags), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImporterMesh) GetSurfaceCount() int64 {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ImporterMesh) GetSurfacePrimitiveType(surface_idx int64, ) MeshPrimitiveType {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_primitive_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3552571330) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), }
  var ret MeshPrimitiveType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ImporterMesh) GetSurfaceName(surface_idx int64, ) String {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ImporterMesh) GetSurfaceArrays(surface_idx int64, ) Array {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_arrays")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 663333327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), }
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ImporterMesh) GetSurfaceBlendShapeArrays(surface_idx int64, blend_shape_idx int64, ) Array {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_blend_shape_arrays")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2345056839) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), gdc.ConstTypePtr(&blend_shape_idx), }
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ImporterMesh) GetSurfaceLodCount(surface_idx int64, ) int64 {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_lod_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ImporterMesh) GetSurfaceLodSize(surface_idx int64, lod_idx int64, ) float64 {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_lod_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3085491603) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), gdc.ConstTypePtr(&lod_idx), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ImporterMesh) GetSurfaceLodIndices(surface_idx int64, lod_idx int64, ) PackedInt32Array {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_lod_indices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265128013) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), gdc.ConstTypePtr(&lod_idx), }
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ImporterMesh) GetSurfaceMaterial(surface_idx int64, ) Material {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2897466400) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), }
  ret := NewMaterial()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ImporterMesh) GetSurfaceFormat(surface_idx int64, ) int64 {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ImporterMesh) SetSurfaceName(surface_idx int64, name String, )  {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_surface_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), gdc.ConstTypePtr(name.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImporterMesh) SetSurfaceMaterial(surface_idx int64, material Material, )  {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_surface_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3671737478) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_idx), gdc.ConstTypePtr(material.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImporterMesh) GenerateLods(normal_merge_angle float64, normal_split_angle float64, bone_transform_array Array, )  {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_lods")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2491878677) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normal_merge_angle), gdc.ConstTypePtr(&normal_split_angle), gdc.ConstTypePtr(bone_transform_array.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImporterMesh) GetMesh(base_mesh ArrayMesh, ) ArrayMesh {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1457573577) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(base_mesh.AsCTypePtr()), }
  ret := NewArrayMesh()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ImporterMesh) Clear()  {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImporterMesh) SetLightmapSizeHint(size Vector2i, )  {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lightmap_size_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImporterMesh) GetLightmapSizeHint() Vector2i {
  classNameV := StringNameFromStr("ImporterMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lightmap_size_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
