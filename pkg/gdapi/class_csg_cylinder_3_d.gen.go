// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGCylinder3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGCylinder3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGCylinder3D) BaseClass() string {
  return "CSGCylinder3D"
}



// Enums

func (me *CSGCylinder3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGCylinder3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGCylinder3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CSGCylinder3D) SetRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *CSGCylinder3D) GetRadius()  {
  panic("TODO: implement")
}

func  (me *CSGCylinder3D) SetHeight(height float32, )  {
  panic("TODO: implement")
}

func  (me *CSGCylinder3D) GetHeight()  {
  panic("TODO: implement")
}

func  (me *CSGCylinder3D) SetSides(sides int, )  {
  panic("TODO: implement")
}

func  (me *CSGCylinder3D) GetSides()  {
  panic("TODO: implement")
}

func  (me *CSGCylinder3D) SetCone(cone bool, )  {
  panic("TODO: implement")
}

func  (me *CSGCylinder3D) IsCone()  {
  panic("TODO: implement")
}

func  (me *CSGCylinder3D) SetMaterial(material Material, )  {
  panic("TODO: implement")
}

func  (me *CSGCylinder3D) GetMaterial()  {
  panic("TODO: implement")
}

func  (me *CSGCylinder3D) SetSmoothFaces(smooth_faces bool, )  {
  panic("TODO: implement")
}

func  (me *CSGCylinder3D) GetSmoothFaces()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
