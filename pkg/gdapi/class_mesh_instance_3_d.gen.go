// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MeshInstance3D struct {
  obj gdc.ObjectPtr
}

func (me *MeshInstance3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MeshInstance3D) BaseClass() string {
  return "MeshInstance3D"
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
  panic("TODO: implement")
}

func  (me *MeshInstance3D) GetMesh()  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) SetSkeletonPath(skeleton_path NodePath, )  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) GetSkeletonPath()  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) SetSkin(skin Skin, )  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) GetSkin()  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) GetSurfaceOverrideMaterialCount()  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) SetSurfaceOverrideMaterial(surface int, material Material, )  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) GetSurfaceOverrideMaterial(surface int, )  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) GetActiveMaterial(surface int, )  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) CreateTrimeshCollision()  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) CreateConvexCollision(clean bool, simplify bool, )  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) CreateMultipleConvexCollisions(settings MeshConvexDecompositionSettings, )  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) GetBlendShapeCount()  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) FindBlendShapeByName(name StringName, )  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) GetBlendShapeValue(blend_shape_idx int, )  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) SetBlendShapeValue(blend_shape_idx int, value float32, )  {
  panic("TODO: implement")
}

func  (me *MeshInstance3D) CreateDebugTangents()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
