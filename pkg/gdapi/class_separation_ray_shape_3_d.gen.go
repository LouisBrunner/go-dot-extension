// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SeparationRayShape3D struct {
  obj gdc.ObjectPtr
}

func (me *SeparationRayShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SeparationRayShape3D) BaseClass() string {
  return "SeparationRayShape3D"
}



// Enums

func (me *SeparationRayShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SeparationRayShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SeparationRayShape3D) SetLength(length float32, )  {
  panic("TODO: implement")
}

func  (me *SeparationRayShape3D) GetLength()  {
  panic("TODO: implement")
}

func  (me *SeparationRayShape3D) SetSlideOnSlope(active bool, )  {
  panic("TODO: implement")
}

func  (me *SeparationRayShape3D) GetSlideOnSlope()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
