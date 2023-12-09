// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsMaterial struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsMaterial) BaseClass() string {
  return "PhysicsMaterial"
}



// Enums

func (me *PhysicsMaterial) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsMaterial) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PhysicsMaterial) SetFriction(friction float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsMaterial) GetFriction()  {
  panic("TODO: implement")
}

func  (me *PhysicsMaterial) SetRough(rough bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsMaterial) IsRough()  {
  panic("TODO: implement")
}

func  (me *PhysicsMaterial) SetBounce(bounce float32, )  {
  panic("TODO: implement")
}

func  (me *PhysicsMaterial) GetBounce()  {
  panic("TODO: implement")
}

func  (me *PhysicsMaterial) SetAbsorbent(absorbent bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicsMaterial) IsAbsorbent()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
