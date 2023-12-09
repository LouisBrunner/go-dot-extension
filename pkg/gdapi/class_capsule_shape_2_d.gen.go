// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CapsuleShape2D struct {
  obj gdc.ObjectPtr
}

func (me *CapsuleShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CapsuleShape2D) BaseClass() string {
  return "CapsuleShape2D"
}



// Enums

func (me *CapsuleShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CapsuleShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CapsuleShape2D) SetRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *CapsuleShape2D) GetRadius()  {
  panic("TODO: implement")
}

func  (me *CapsuleShape2D) SetHeight(height float32, )  {
  panic("TODO: implement")
}

func  (me *CapsuleShape2D) GetHeight()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
