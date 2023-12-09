// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CylinderShape3D struct {
  obj gdc.ObjectPtr
}

func (me *CylinderShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CylinderShape3D) BaseClass() string {
  return "CylinderShape3D"
}



// Enums

func (me *CylinderShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CylinderShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CylinderShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CylinderShape3D) SetRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *CylinderShape3D) GetRadius()  {
  panic("TODO: implement")
}

func  (me *CylinderShape3D) SetHeight(height float32, )  {
  panic("TODO: implement")
}

func  (me *CylinderShape3D) GetHeight()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
