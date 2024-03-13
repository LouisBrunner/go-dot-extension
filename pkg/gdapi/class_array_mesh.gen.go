// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ArrayMesh struct {
  obj gdc.ObjectPtr
}

func (me *ArrayMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ArrayMesh) BaseClass() string {
  return "ArrayMesh"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ArrayMesh) GetBlendShapeCount() int {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_shape_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ArrayMesh) GetBlendShapeName(index int, ) StringName {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_shape_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ArrayMesh) SetBlendShapeName(index int, name StringName, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_shape_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3780747571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ArrayMesh) ClearBlendShapes()  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_blend_shapes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ArrayMesh) SetBlendShapeMode(mode MeshBlendShapeMode, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_shape_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 227983991) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ArrayMesh) GetBlendShapeMode() MeshBlendShapeMode {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_shape_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 836485024) // FIXME: should cache?
  var ret MeshBlendShapeMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ArrayMesh) AddSurfaceFromArrays(primitive MeshPrimitiveType, arrays Array, blend_shapes Array, lods Dictionary, flags MeshArrayFormat, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_surface_from_arrays")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1796411378) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&primitive), gdc.ConstTypePtr(arrays.AsCTypePtr()), gdc.ConstTypePtr(blend_shapes.AsCTypePtr()), gdc.ConstTypePtr(lods.AsCTypePtr()), gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ArrayMesh) ClearSurfaces()  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_surfaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ArrayMesh) SurfaceUpdateVertexRegion(surf_idx int, offset int, data PackedByteArray, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_update_vertex_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3837166854) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx), gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ArrayMesh) SurfaceUpdateAttributeRegion(surf_idx int, offset int, data PackedByteArray, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_update_attribute_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3837166854) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx), gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ArrayMesh) SurfaceUpdateSkinRegion(surf_idx int, offset int, data PackedByteArray, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_update_skin_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3837166854) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx), gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ArrayMesh) SurfaceGetArrayLen(surf_idx int, ) int {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_array_len")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ArrayMesh) SurfaceGetArrayIndexLen(surf_idx int, ) int {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_array_index_len")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ArrayMesh) SurfaceGetFormat(surf_idx int, ) MeshArrayFormat {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3718287884) // FIXME: should cache?
  var ret MeshArrayFormat
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ArrayMesh) SurfaceGetPrimitiveType(surf_idx int, ) MeshPrimitiveType {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_primitive_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4141943888) // FIXME: should cache?
  var ret MeshPrimitiveType
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ArrayMesh) SurfaceFindByName(name String, ) int {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_find_by_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321353865) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ArrayMesh) SurfaceSetName(surf_idx int, name String, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_set_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ArrayMesh) SurfaceGetName(surf_idx int, ) String {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("surface_get_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ArrayMesh) RegenNormalMaps()  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("regen_normal_maps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ArrayMesh) LightmapUnwrap(transform Transform3D, texel_size float32, ) Error {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("lightmap_unwrap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1476641071) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(transform.AsCTypePtr()), gdc.ConstTypePtr(&texel_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ArrayMesh) SetCustomAabb(aabb AABB, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 259215842) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(aabb.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ArrayMesh) GetCustomAabb() AABB {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  var ret AABB
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ArrayMesh) SetShadowMesh(mesh ArrayMesh, )  {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shadow_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3377897901) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mesh.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ArrayMesh) GetShadowMesh() ArrayMesh {
  classNameV := StringNameFromStr("ArrayMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shadow_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3206942465) // FIXME: should cache?
  var ret ArrayMesh
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
