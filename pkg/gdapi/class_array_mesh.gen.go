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

type ArrayMesh struct {
  Mesh
}

func (me *ArrayMesh) BaseClass() string {
  return "ArrayMesh"
}

func NewArrayMesh() *ArrayMesh {
  str := StringNameFromStr("ArrayMesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ArrayMesh{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ArrayMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ArrayMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ArrayMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ArrayMesh) AddBlendShape(name StringName, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_blend_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) GetBlendShapeCount() int64 {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_shape_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ArrayMesh) GetBlendShapeName(index int64, ) StringName {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_shape_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ArrayMesh) SetBlendShapeName(index int64, name StringName, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_shape_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3780747571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) ClearBlendShapes()  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_blend_shapes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) SetBlendShapeMode(mode MeshBlendShapeMode, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_shape_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 227983991) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) GetBlendShapeMode() MeshBlendShapeMode {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_shape_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 836485024) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MeshBlendShapeMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ArrayMesh) AddSurfaceFromArrays(primitive MeshPrimitiveType, arrays Array, blend_shapes []Array, lods Dictionary, flags MeshArrayFormat, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_surface_from_arrays")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1796411378) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&primitive) , arrays.AsCTypePtr(), gdc.ConstTypePtr(&blend_shapes) , lods.AsCTypePtr(), gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) ClearSurfaces()  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_surfaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) SurfaceUpdateVertexRegion(surf_idx int64, offset int64, data PackedByteArray, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_update_vertex_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3837166854) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , gdc.ConstTypePtr(&offset) , data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) SurfaceUpdateAttributeRegion(surf_idx int64, offset int64, data PackedByteArray, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_update_attribute_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3837166854) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , gdc.ConstTypePtr(&offset) , data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) SurfaceUpdateSkinRegion(surf_idx int64, offset int64, data PackedByteArray, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_update_skin_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3837166854) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , gdc.ConstTypePtr(&offset) , data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) SurfaceGetArrayLen(surf_idx int64, ) int64 {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_array_len")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&surf_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ArrayMesh) SurfaceGetArrayIndexLen(surf_idx int64, ) int64 {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_array_index_len")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&surf_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ArrayMesh) SurfaceGetFormat(surf_idx int64, ) MeshArrayFormat {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3718287884) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MeshArrayFormat
  pinner.Pin(&surf_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ArrayMesh) SurfaceGetPrimitiveType(surf_idx int64, ) MeshPrimitiveType {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_primitive_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4141943888) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MeshPrimitiveType
  pinner.Pin(&surf_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ArrayMesh) SurfaceFindByName(name String, ) int64 {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_find_by_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321353865) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ArrayMesh) SurfaceSetName(surf_idx int64, name String, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_set_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) SurfaceGetName(surf_idx int64, ) String {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&surf_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ArrayMesh) RegenNormalMaps()  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("regen_normal_maps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) LightmapUnwrap(transform Transform3D, texel_size float64, ) Error {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("lightmap_unwrap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1476641071) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{transform.AsCTypePtr(), gdc.ConstTypePtr(&texel_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&texel_size)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ArrayMesh) SetCustomAabb(aabb AABB, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 259215842) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{aabb.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) GetCustomAabb() AABB {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ArrayMesh) SetShadowMesh(mesh ArrayMesh, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3377897901) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) GetShadowMesh() ArrayMesh {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3206942465) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArrayMesh()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
