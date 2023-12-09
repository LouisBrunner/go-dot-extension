// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGPrimitive3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGPrimitive3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGPrimitive3D) BaseClass() string {
  return "CSGPrimitive3D"
}



// Enums

func (me *CSGPrimitive3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGPrimitive3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGPrimitive3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CSGPrimitive3D) SetFlipFaces(flip_faces bool, )  {
  panic("TODO: implement")
}

func  (me *CSGPrimitive3D) GetFlipFaces()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
