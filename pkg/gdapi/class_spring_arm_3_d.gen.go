// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SpringArm3D struct {
  obj gdc.ObjectPtr
}

func (me *SpringArm3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SpringArm3D) BaseClass() string {
  return "SpringArm3D"
}



// Enums

func (me *SpringArm3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SpringArm3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SpringArm3D) GetHitLength()  {
  panic("TODO: implement")
}

func  (me *SpringArm3D) SetLength(length float32, )  {
  panic("TODO: implement")
}

func  (me *SpringArm3D) GetLength()  {
  panic("TODO: implement")
}

func  (me *SpringArm3D) SetShape(shape Shape3D, )  {
  panic("TODO: implement")
}

func  (me *SpringArm3D) GetShape()  {
  panic("TODO: implement")
}

func  (me *SpringArm3D) AddExcludedObject(RID RID, )  {
  panic("TODO: implement")
}

func  (me *SpringArm3D) RemoveExcludedObject(RID RID, )  {
  panic("TODO: implement")
}

func  (me *SpringArm3D) ClearExcludedObjects()  {
  panic("TODO: implement")
}

func  (me *SpringArm3D) SetCollisionMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *SpringArm3D) GetCollisionMask()  {
  panic("TODO: implement")
}

func  (me *SpringArm3D) SetMargin(margin float32, )  {
  panic("TODO: implement")
}

func  (me *SpringArm3D) GetMargin()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
