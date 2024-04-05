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

type ptrsForMeshInstance3DList struct {
  fnSetMesh gdc.MethodBindPtr
  fnGetMesh gdc.MethodBindPtr
  fnSetSkeletonPath gdc.MethodBindPtr
  fnGetSkeletonPath gdc.MethodBindPtr
  fnSetSkin gdc.MethodBindPtr
  fnGetSkin gdc.MethodBindPtr
  fnGetSurfaceOverrideMaterialCount gdc.MethodBindPtr
  fnSetSurfaceOverrideMaterial gdc.MethodBindPtr
  fnGetSurfaceOverrideMaterial gdc.MethodBindPtr
  fnGetActiveMaterial gdc.MethodBindPtr
  fnCreateTrimeshCollision gdc.MethodBindPtr
  fnCreateConvexCollision gdc.MethodBindPtr
  fnCreateMultipleConvexCollisions gdc.MethodBindPtr
  fnGetBlendShapeCount gdc.MethodBindPtr
  fnFindBlendShapeByName gdc.MethodBindPtr
  fnGetBlendShapeValue gdc.MethodBindPtr
  fnSetBlendShapeValue gdc.MethodBindPtr
  fnCreateDebugTangents gdc.MethodBindPtr
}

var ptrsForMeshInstance3D ptrsForMeshInstance3DList

func initMeshInstance3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("MeshInstance3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_mesh")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnSetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 194775623))
  }
  {
    methodName := StringNameFromStr("get_mesh")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnGetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1808005922))
  }
  {
    methodName := StringNameFromStr("set_skeleton_path")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnSetSkeletonPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
  }
  {
    methodName := StringNameFromStr("get_skeleton_path")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnGetSkeletonPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 277076166))
  }
  {
    methodName := StringNameFromStr("set_skin")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnSetSkin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3971435618))
  }
  {
    methodName := StringNameFromStr("get_skin")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnGetSkin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2074563878))
  }
  {
    methodName := StringNameFromStr("get_surface_override_material_count")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnGetSurfaceOverrideMaterialCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_surface_override_material")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnSetSurfaceOverrideMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3671737478))
  }
  {
    methodName := StringNameFromStr("get_surface_override_material")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnGetSurfaceOverrideMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2897466400))
  }
  {
    methodName := StringNameFromStr("get_active_material")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnGetActiveMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2897466400))
  }
  {
    methodName := StringNameFromStr("create_trimesh_collision")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnCreateTrimeshCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("create_convex_collision")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnCreateConvexCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2751962654))
  }
  {
    methodName := StringNameFromStr("create_multiple_convex_collisions")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnCreateMultipleConvexCollisions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 628789669))
  }
  {
    methodName := StringNameFromStr("get_blend_shape_count")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnGetBlendShapeCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("find_blend_shape_by_name")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnFindBlendShapeByName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4150868206))
  }
  {
    methodName := StringNameFromStr("get_blend_shape_value")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnGetBlendShapeValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
  }
  {
    methodName := StringNameFromStr("set_blend_shape_value")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnSetBlendShapeValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("create_debug_tangents")
    defer methodName.Destroy()
    ptrsForMeshInstance3D.fnCreateDebugTangents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
}

type MeshInstance3D struct {
  GeometryInstance3D
}

func (me *MeshInstance3D) BaseClass() string {
  return "MeshInstance3D"
}

func NewMeshInstance3D() *MeshInstance3D {
  str := StringNameFromStr("MeshInstance3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &MeshInstance3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *MeshInstance3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MeshInstance3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MeshInstance3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MeshInstance3D) SetMesh(mesh Mesh, )  {
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnSetMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) GetMesh() Mesh {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMesh()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnGetMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MeshInstance3D) SetSkeletonPath(skeleton_path NodePath, )  {
  cargs := []gdc.ConstTypePtr{skeleton_path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnSetSkeletonPath), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) GetSkeletonPath() NodePath {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnGetSkeletonPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MeshInstance3D) SetSkin(skin Skin, )  {
  cargs := []gdc.ConstTypePtr{skin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnSetSkin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) GetSkin() Skin {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSkin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnGetSkin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MeshInstance3D) GetSurfaceOverrideMaterialCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnGetSurfaceOverrideMaterialCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshInstance3D) SetSurfaceOverrideMaterial(surface int64, material Material, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface) , material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnSetSurfaceOverrideMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) GetSurfaceOverrideMaterial(surface int64, ) Material {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()
  pinner.Pin(&surface)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnGetSurfaceOverrideMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MeshInstance3D) GetActiveMaterial(surface int64, ) Material {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()
  pinner.Pin(&surface)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnGetActiveMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MeshInstance3D) CreateTrimeshCollision()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnCreateTrimeshCollision), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) CreateConvexCollision(clean bool, simplify bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clean) , gdc.ConstTypePtr(&simplify) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnCreateConvexCollision), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) CreateMultipleConvexCollisions(settings MeshConvexDecompositionSettings, )  {
  cargs := []gdc.ConstTypePtr{settings.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnCreateMultipleConvexCollisions), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) GetBlendShapeCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnGetBlendShapeCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshInstance3D) FindBlendShapeByName(name StringName, ) int64 {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnFindBlendShapeByName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshInstance3D) GetBlendShapeValue(blend_shape_idx int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&blend_shape_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&blend_shape_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnGetBlendShapeValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshInstance3D) SetBlendShapeValue(blend_shape_idx int64, value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&blend_shape_idx) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnSetBlendShapeValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) CreateDebugTangents()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMeshInstance3D.fnCreateDebugTangents), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
