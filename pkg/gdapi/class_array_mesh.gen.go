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

type ptrsForArrayMeshList struct {
  fnAddBlendShape gdc.MethodBindPtr
  fnGetBlendShapeCount gdc.MethodBindPtr
  fnGetBlendShapeName gdc.MethodBindPtr
  fnSetBlendShapeName gdc.MethodBindPtr
  fnClearBlendShapes gdc.MethodBindPtr
  fnSetBlendShapeMode gdc.MethodBindPtr
  fnGetBlendShapeMode gdc.MethodBindPtr
  fnAddSurfaceFromArrays gdc.MethodBindPtr
  fnClearSurfaces gdc.MethodBindPtr
  fnSurfaceUpdateVertexRegion gdc.MethodBindPtr
  fnSurfaceUpdateAttributeRegion gdc.MethodBindPtr
  fnSurfaceUpdateSkinRegion gdc.MethodBindPtr
  fnSurfaceGetArrayLen gdc.MethodBindPtr
  fnSurfaceGetArrayIndexLen gdc.MethodBindPtr
  fnSurfaceGetFormat gdc.MethodBindPtr
  fnSurfaceGetPrimitiveType gdc.MethodBindPtr
  fnSurfaceFindByName gdc.MethodBindPtr
  fnSurfaceSetName gdc.MethodBindPtr
  fnSurfaceGetName gdc.MethodBindPtr
  fnRegenNormalMaps gdc.MethodBindPtr
  fnLightmapUnwrap gdc.MethodBindPtr
  fnSetCustomAabb gdc.MethodBindPtr
  fnGetCustomAabb gdc.MethodBindPtr
  fnSetShadowMesh gdc.MethodBindPtr
  fnGetShadowMesh gdc.MethodBindPtr
}

var ptrsForArrayMesh ptrsForArrayMeshList

func initArrayMeshPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ArrayMesh")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_blend_shape")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnAddBlendShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("get_blend_shape_count")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnGetBlendShapeCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_blend_shape_name")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnGetBlendShapeName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
  }
  {
    methodName := StringNameFromStr("set_blend_shape_name")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnSetBlendShapeName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3780747571))
  }
  {
    methodName := StringNameFromStr("clear_blend_shapes")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnClearBlendShapes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_blend_shape_mode")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnSetBlendShapeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 227983991))
  }
  {
    methodName := StringNameFromStr("get_blend_shape_mode")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnGetBlendShapeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 836485024))
  }
  {
    methodName := StringNameFromStr("add_surface_from_arrays")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnAddSurfaceFromArrays = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1796411378))
  }
  {
    methodName := StringNameFromStr("clear_surfaces")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnClearSurfaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("surface_update_vertex_region")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnSurfaceUpdateVertexRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3837166854))
  }
  {
    methodName := StringNameFromStr("surface_update_attribute_region")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnSurfaceUpdateAttributeRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3837166854))
  }
  {
    methodName := StringNameFromStr("surface_update_skin_region")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnSurfaceUpdateSkinRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3837166854))
  }
  {
    methodName := StringNameFromStr("surface_get_array_len")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnSurfaceGetArrayLen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
  }
  {
    methodName := StringNameFromStr("surface_get_array_index_len")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnSurfaceGetArrayIndexLen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
  }
  {
    methodName := StringNameFromStr("surface_get_format")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnSurfaceGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3718287884))
  }
  {
    methodName := StringNameFromStr("surface_get_primitive_type")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnSurfaceGetPrimitiveType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4141943888))
  }
  {
    methodName := StringNameFromStr("surface_find_by_name")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnSurfaceFindByName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1321353865))
  }
  {
    methodName := StringNameFromStr("surface_set_name")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnSurfaceSetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
  }
  {
    methodName := StringNameFromStr("surface_get_name")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnSurfaceGetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
  }
  {
    methodName := StringNameFromStr("regen_normal_maps")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnRegenNormalMaps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("lightmap_unwrap")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnLightmapUnwrap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1476641071))
  }
  {
    methodName := StringNameFromStr("set_custom_aabb")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnSetCustomAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 259215842))
  }
  {
    methodName := StringNameFromStr("get_custom_aabb")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnGetCustomAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1068685055))
  }
  {
    methodName := StringNameFromStr("set_shadow_mesh")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnSetShadowMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3377897901))
  }
  {
    methodName := StringNameFromStr("get_shadow_mesh")
    defer methodName.Destroy()
    ptrsForArrayMesh.fnGetShadowMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3206942465))
  }
}

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
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnAddBlendShape), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) GetBlendShapeCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnGetBlendShapeCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ArrayMesh) GetBlendShapeName(index int64, ) StringName {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnGetBlendShapeName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ArrayMesh) SetBlendShapeName(index int64, name StringName, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnSetBlendShapeName), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) ClearBlendShapes()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnClearBlendShapes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) SetBlendShapeMode(mode MeshBlendShapeMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnSetBlendShapeMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) GetBlendShapeMode() MeshBlendShapeMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MeshBlendShapeMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnGetBlendShapeMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ArrayMesh) AddSurfaceFromArrays(primitive MeshPrimitiveType, arrays Array, blend_shapes []Array, lods Dictionary, flags MeshArrayFormat, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&primitive) , arrays.AsCTypePtr(), gdc.ConstTypePtr(&blend_shapes) , lods.AsCTypePtr(), gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnAddSurfaceFromArrays), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) ClearSurfaces()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnClearSurfaces), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) SurfaceUpdateVertexRegion(surf_idx int64, offset int64, data PackedByteArray, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , gdc.ConstTypePtr(&offset) , data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnSurfaceUpdateVertexRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) SurfaceUpdateAttributeRegion(surf_idx int64, offset int64, data PackedByteArray, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , gdc.ConstTypePtr(&offset) , data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnSurfaceUpdateAttributeRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) SurfaceUpdateSkinRegion(surf_idx int64, offset int64, data PackedByteArray, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , gdc.ConstTypePtr(&offset) , data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnSurfaceUpdateSkinRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) SurfaceGetArrayLen(surf_idx int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&surf_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnSurfaceGetArrayLen), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ArrayMesh) SurfaceGetArrayIndexLen(surf_idx int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&surf_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnSurfaceGetArrayIndexLen), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ArrayMesh) SurfaceGetFormat(surf_idx int64, ) MeshArrayFormat {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MeshArrayFormat
  pinner.Pin(&surf_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnSurfaceGetFormat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ArrayMesh) SurfaceGetPrimitiveType(surf_idx int64, ) MeshPrimitiveType {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MeshPrimitiveType
  pinner.Pin(&surf_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnSurfaceGetPrimitiveType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ArrayMesh) SurfaceFindByName(name String, ) int64 {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnSurfaceFindByName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ArrayMesh) SurfaceSetName(surf_idx int64, name String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnSurfaceSetName), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) SurfaceGetName(surf_idx int64, ) String {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surf_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&surf_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnSurfaceGetName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ArrayMesh) RegenNormalMaps()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnRegenNormalMaps), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) LightmapUnwrap(transform Transform3D, texel_size float64, ) Error {
  cargs := []gdc.ConstTypePtr{transform.AsCTypePtr(), gdc.ConstTypePtr(&texel_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&texel_size)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnLightmapUnwrap), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ArrayMesh) SetCustomAabb(aabb AABB, )  {
  cargs := []gdc.ConstTypePtr{aabb.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnSetCustomAabb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) GetCustomAabb() AABB {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnGetCustomAabb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ArrayMesh) SetShadowMesh(mesh ArrayMesh, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnSetShadowMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayMesh) GetShadowMesh() ArrayMesh {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArrayMesh()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayMesh.fnGetShadowMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
