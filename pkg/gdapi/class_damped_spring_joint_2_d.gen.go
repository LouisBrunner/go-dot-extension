// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DampedSpringJoint2D struct {
  obj gdc.ObjectPtr
}

func (me *DampedSpringJoint2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *DampedSpringJoint2D) BaseClass() string {
  return "DampedSpringJoint2D"
}



// Enums

func (me *DampedSpringJoint2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *DampedSpringJoint2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *DampedSpringJoint2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *DampedSpringJoint2D) SetLength(length float32, )  {
  panic("TODO: implement")
}

func  (me *DampedSpringJoint2D) GetLength()  {
  panic("TODO: implement")
}

func  (me *DampedSpringJoint2D) SetRestLength(rest_length float32, )  {
  panic("TODO: implement")
}

func  (me *DampedSpringJoint2D) GetRestLength()  {
  panic("TODO: implement")
}

func  (me *DampedSpringJoint2D) SetStiffness(stiffness float32, )  {
  panic("TODO: implement")
}

func  (me *DampedSpringJoint2D) GetStiffness()  {
  panic("TODO: implement")
}

func  (me *DampedSpringJoint2D) SetDamping(damping float32, )  {
  panic("TODO: implement")
}

func  (me *DampedSpringJoint2D) GetDamping()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
