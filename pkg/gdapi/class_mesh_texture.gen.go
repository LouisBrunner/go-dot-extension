// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MeshTexture struct {
  obj gdc.ObjectPtr
}

func (me *MeshTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MeshTexture) BaseClass() string {
  return "MeshTexture"
}



// Enums

func (me *MeshTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MeshTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MeshTexture) SetMesh(mesh Mesh, )  {
  panic("TODO: implement")
}

func  (me *MeshTexture) GetMesh()  {
  panic("TODO: implement")
}

func  (me *MeshTexture) SetImageSize(size Vector2, )  {
  panic("TODO: implement")
}

func  (me *MeshTexture) GetImageSize()  {
  panic("TODO: implement")
}

func  (me *MeshTexture) SetBaseTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *MeshTexture) GetBaseTexture()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
