// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CapsuleShape3D struct {
  obj gdc.ObjectPtr
}

func (me *CapsuleShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CapsuleShape3D) BaseClass() string {
  return "CapsuleShape3D"
}



// Enums

func (me *CapsuleShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CapsuleShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CapsuleShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CapsuleShape3D) SetRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *CapsuleShape3D) GetRadius()  {
  panic("TODO: implement")
}

func  (me *CapsuleShape3D) SetHeight(height float32, )  {
  panic("TODO: implement")
}

func  (me *CapsuleShape3D) GetHeight()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
