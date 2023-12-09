// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MeshInstance2D struct {
  obj gdc.ObjectPtr
}

func (me *MeshInstance2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MeshInstance2D) BaseClass() string {
  return "MeshInstance2D"
}



// Enums

func (me *MeshInstance2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MeshInstance2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MeshInstance2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MeshInstance2D) SetMesh(mesh Mesh, )  {
  panic("TODO: implement")
}

func  (me *MeshInstance2D) GetMesh()  {
  panic("TODO: implement")
}

func  (me *MeshInstance2D) SetTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *MeshInstance2D) GetTexture()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
