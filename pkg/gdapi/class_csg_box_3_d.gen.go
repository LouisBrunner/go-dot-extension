// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGBox3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGBox3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGBox3D) BaseClass() string {
  return "CSGBox3D"
}



// Enums

func (me *CSGBox3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGBox3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CSGBox3D) SetSize(size Vector3, )  {
  panic("TODO: implement")
}

func  (me *CSGBox3D) GetSize()  {
  panic("TODO: implement")
}

func  (me *CSGBox3D) SetMaterial(material Material, )  {
  panic("TODO: implement")
}

func  (me *CSGBox3D) GetMaterial()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
