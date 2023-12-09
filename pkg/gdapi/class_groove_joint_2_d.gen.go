// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GrooveJoint2D struct {
  obj gdc.ObjectPtr
}

func (me *GrooveJoint2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GrooveJoint2D) BaseClass() string {
  return "GrooveJoint2D"
}



// Enums

func (me *GrooveJoint2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GrooveJoint2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GrooveJoint2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GrooveJoint2D) SetLength(length float32, )  {
  panic("TODO: implement")
}

func  (me *GrooveJoint2D) GetLength()  {
  panic("TODO: implement")
}

func  (me *GrooveJoint2D) SetInitialOffset(offset float32, )  {
  panic("TODO: implement")
}

func  (me *GrooveJoint2D) GetInitialOffset()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
