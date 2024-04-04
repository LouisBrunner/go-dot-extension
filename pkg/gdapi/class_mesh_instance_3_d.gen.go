// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

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
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 194775623) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) GetMesh() Mesh {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1808005922) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMesh()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MeshInstance3D) SetSkeletonPath(skeleton_path NodePath, )  {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skeleton_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{skeleton_path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) GetSkeletonPath() NodePath {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skeleton_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 277076166) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MeshInstance3D) SetSkin(skin Skin, )  {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3971435618) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{skin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) GetSkin() Skin {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2074563878) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSkin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MeshInstance3D) GetSurfaceOverrideMaterialCount() int64 {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_override_material_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshInstance3D) SetSurfaceOverrideMaterial(surface int64, material Material, )  {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_surface_override_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3671737478) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface) , material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) GetSurfaceOverrideMaterial(surface int64, ) Material {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_override_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2897466400) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()
  pinner.Pin(&surface)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MeshInstance3D) GetActiveMaterial(surface int64, ) Material {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_active_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2897466400) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()
  pinner.Pin(&surface)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *MeshInstance3D) CreateTrimeshCollision()  {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_trimesh_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) CreateConvexCollision(clean bool, simplify bool, )  {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_convex_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2751962654) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clean) , gdc.ConstTypePtr(&simplify) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) CreateMultipleConvexCollisions(settings MeshConvexDecompositionSettings, )  {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_multiple_convex_collisions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 628789669) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{settings.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) GetBlendShapeCount() int64 {
  classNameV := StringNameFromStr("MeshInstance3D")
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

func  (me *MeshInstance3D) FindBlendShapeByName(name StringName, ) int64 {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_blend_shape_by_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4150868206) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshInstance3D) GetBlendShapeValue(blend_shape_idx int64, ) float64 {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_shape_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&blend_shape_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&blend_shape_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *MeshInstance3D) SetBlendShapeValue(blend_shape_idx int64, value float64, )  {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_shape_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&blend_shape_idx) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *MeshInstance3D) CreateDebugTangents()  {
  classNameV := StringNameFromStr("MeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_debug_tangents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
