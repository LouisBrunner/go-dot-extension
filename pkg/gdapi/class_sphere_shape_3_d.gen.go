// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SphereShape3D struct {
  obj gdc.ObjectPtr
}

func (me *SphereShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SphereShape3D) BaseClass() string {
  return "SphereShape3D"
}



// Enums

func (me *SphereShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SphereShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SphereShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SphereShape3D) SetRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *SphereShape3D) GetRadius()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
