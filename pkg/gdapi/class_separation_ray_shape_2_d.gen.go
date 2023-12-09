// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SeparationRayShape2D struct {
  obj gdc.ObjectPtr
}

func (me *SeparationRayShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SeparationRayShape2D) BaseClass() string {
  return "SeparationRayShape2D"
}



// Enums

func (me *SeparationRayShape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SeparationRayShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SeparationRayShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SeparationRayShape2D) SetLength(length float32, )  {
  panic("TODO: implement")
}

func  (me *SeparationRayShape2D) GetLength()  {
  panic("TODO: implement")
}

func  (me *SeparationRayShape2D) SetSlideOnSlope(active bool, )  {
  panic("TODO: implement")
}

func  (me *SeparationRayShape2D) GetSlideOnSlope()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
