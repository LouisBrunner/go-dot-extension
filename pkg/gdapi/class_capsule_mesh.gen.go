// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CapsuleMesh struct {
  obj gdc.ObjectPtr
}

func (me *CapsuleMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CapsuleMesh) BaseClass() string {
  return "CapsuleMesh"
}



// Enums

func (me *CapsuleMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CapsuleMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CapsuleMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CapsuleMesh) SetRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *CapsuleMesh) GetRadius()  {
  panic("TODO: implement")
}

func  (me *CapsuleMesh) SetHeight(height float32, )  {
  panic("TODO: implement")
}

func  (me *CapsuleMesh) GetHeight()  {
  panic("TODO: implement")
}

func  (me *CapsuleMesh) SetRadialSegments(segments int, )  {
  panic("TODO: implement")
}

func  (me *CapsuleMesh) GetRadialSegments()  {
  panic("TODO: implement")
}

func  (me *CapsuleMesh) SetRings(rings int, )  {
  panic("TODO: implement")
}

func  (me *CapsuleMesh) GetRings()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
