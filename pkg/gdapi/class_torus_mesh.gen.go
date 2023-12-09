// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TorusMesh struct {
  obj gdc.ObjectPtr
}

func (me *TorusMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TorusMesh) BaseClass() string {
  return "TorusMesh"
}



// Enums

func (me *TorusMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TorusMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TorusMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *TorusMesh) SetInnerRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *TorusMesh) GetInnerRadius()  {
  panic("TODO: implement")
}

func  (me *TorusMesh) SetOuterRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *TorusMesh) GetOuterRadius()  {
  panic("TODO: implement")
}

func  (me *TorusMesh) SetRings(rings int, )  {
  panic("TODO: implement")
}

func  (me *TorusMesh) GetRings()  {
  panic("TODO: implement")
}

func  (me *TorusMesh) SetRingSegments(rings int, )  {
  panic("TODO: implement")
}

func  (me *TorusMesh) GetRingSegments()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
