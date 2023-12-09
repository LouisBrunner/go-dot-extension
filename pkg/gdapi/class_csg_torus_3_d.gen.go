// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGTorus3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGTorus3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGTorus3D) BaseClass() string {
  return "CSGTorus3D"
}



// Enums

func (me *CSGTorus3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGTorus3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGTorus3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CSGTorus3D) SetInnerRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *CSGTorus3D) GetInnerRadius()  {
  panic("TODO: implement")
}

func  (me *CSGTorus3D) SetOuterRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *CSGTorus3D) GetOuterRadius()  {
  panic("TODO: implement")
}

func  (me *CSGTorus3D) SetSides(sides int, )  {
  panic("TODO: implement")
}

func  (me *CSGTorus3D) GetSides()  {
  panic("TODO: implement")
}

func  (me *CSGTorus3D) SetRingSides(sides int, )  {
  panic("TODO: implement")
}

func  (me *CSGTorus3D) GetRingSides()  {
  panic("TODO: implement")
}

func  (me *CSGTorus3D) SetMaterial(material Material, )  {
  panic("TODO: implement")
}

func  (me *CSGTorus3D) GetMaterial()  {
  panic("TODO: implement")
}

func  (me *CSGTorus3D) SetSmoothFaces(smooth_faces bool, )  {
  panic("TODO: implement")
}

func  (me *CSGTorus3D) GetSmoothFaces()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
