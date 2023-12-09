// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PrimitiveMesh struct {
  obj gdc.ObjectPtr
}

func (me *PrimitiveMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PrimitiveMesh) BaseClass() string {
  return "PrimitiveMesh"
}



// Enums

func (me *PrimitiveMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PrimitiveMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PrimitiveMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PrimitiveMesh) XCreateMeshArray()  {
  panic("TODO: implement")
}

func  (me *PrimitiveMesh) SetMaterial(material Material, )  {
  panic("TODO: implement")
}

func  (me *PrimitiveMesh) GetMaterial()  {
  panic("TODO: implement")
}

func  (me *PrimitiveMesh) GetMeshArrays()  {
  panic("TODO: implement")
}

func  (me *PrimitiveMesh) SetCustomAabb(aabb AABB, )  {
  panic("TODO: implement")
}

func  (me *PrimitiveMesh) GetCustomAabb()  {
  panic("TODO: implement")
}

func  (me *PrimitiveMesh) SetFlipFaces(flip_faces bool, )  {
  panic("TODO: implement")
}

func  (me *PrimitiveMesh) GetFlipFaces()  {
  panic("TODO: implement")
}

func  (me *PrimitiveMesh) SetAddUv2(add_uv2 bool, )  {
  panic("TODO: implement")
}

func  (me *PrimitiveMesh) GetAddUv2()  {
  panic("TODO: implement")
}

func  (me *PrimitiveMesh) SetUv2Padding(uv2_padding float32, )  {
  panic("TODO: implement")
}

func  (me *PrimitiveMesh) GetUv2Padding()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
