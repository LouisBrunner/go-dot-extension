// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Path2D struct {
  obj gdc.ObjectPtr
}

func (me *Path2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Path2D) BaseClass() string {
  return "Path2D"
}



// Enums

func (me *Path2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Path2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Path2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Path2D) SetCurve(curve Curve2D, )  {
  panic("TODO: implement")
}

func  (me *Path2D) GetCurve()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
