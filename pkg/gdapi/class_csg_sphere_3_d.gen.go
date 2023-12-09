// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGSphere3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGSphere3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGSphere3D) BaseClass() string {
  return "CSGSphere3D"
}



// Enums

func (me *CSGSphere3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGSphere3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGSphere3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CSGSphere3D) SetRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *CSGSphere3D) GetRadius()  {
  panic("TODO: implement")
}

func  (me *CSGSphere3D) SetRadialSegments(radial_segments int, )  {
  panic("TODO: implement")
}

func  (me *CSGSphere3D) GetRadialSegments()  {
  panic("TODO: implement")
}

func  (me *CSGSphere3D) SetRings(rings int, )  {
  panic("TODO: implement")
}

func  (me *CSGSphere3D) GetRings()  {
  panic("TODO: implement")
}

func  (me *CSGSphere3D) SetSmoothFaces(smooth_faces bool, )  {
  panic("TODO: implement")
}

func  (me *CSGSphere3D) GetSmoothFaces()  {
  panic("TODO: implement")
}

func  (me *CSGSphere3D) SetMaterial(material Material, )  {
  panic("TODO: implement")
}

func  (me *CSGSphere3D) GetMaterial()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
