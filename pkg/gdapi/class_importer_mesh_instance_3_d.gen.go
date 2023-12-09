// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImporterMeshInstance3D struct {
  obj gdc.ObjectPtr
}

func (me *ImporterMeshInstance3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ImporterMeshInstance3D) BaseClass() string {
  return "ImporterMeshInstance3D"
}



// Enums

func (me *ImporterMeshInstance3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ImporterMeshInstance3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ImporterMeshInstance3D) SetMesh(mesh ImporterMesh, )  {
  panic("TODO: implement")
}

func  (me *ImporterMeshInstance3D) GetMesh()  {
  panic("TODO: implement")
}

func  (me *ImporterMeshInstance3D) SetSkin(skin Skin, )  {
  panic("TODO: implement")
}

func  (me *ImporterMeshInstance3D) GetSkin()  {
  panic("TODO: implement")
}

func  (me *ImporterMeshInstance3D) SetSkeletonPath(skeleton_path NodePath, )  {
  panic("TODO: implement")
}

func  (me *ImporterMeshInstance3D) GetSkeletonPath()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
