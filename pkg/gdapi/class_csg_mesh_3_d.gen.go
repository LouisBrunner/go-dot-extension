// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGMesh3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGMesh3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGMesh3D) BaseClass() string {
  return "CSGMesh3D"
}



// Enums

func (me *CSGMesh3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGMesh3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CSGMesh3D) SetMesh(mesh Mesh, )  {
  panic("TODO: implement")
}

func  (me *CSGMesh3D) GetMesh()  {
  panic("TODO: implement")
}

func  (me *CSGMesh3D) SetMaterial(material Material, )  {
  panic("TODO: implement")
}

func  (me *CSGMesh3D) GetMaterial()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
