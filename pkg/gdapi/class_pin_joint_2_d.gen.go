// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PinJoint2D struct {
  obj gdc.ObjectPtr
}

func (me *PinJoint2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PinJoint2D) BaseClass() string {
  return "PinJoint2D"
}



// Enums

func (me *PinJoint2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PinJoint2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PinJoint2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PinJoint2D) SetSoftness(softness float32, )  {
  panic("TODO: implement")
}

func  (me *PinJoint2D) GetSoftness()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
