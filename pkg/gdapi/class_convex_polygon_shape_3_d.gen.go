// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConvexPolygonShape3D struct {
  obj gdc.ObjectPtr
}

func (me *ConvexPolygonShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConvexPolygonShape3D) BaseClass() string {
  return "ConvexPolygonShape3D"
}



// Enums

func (me *ConvexPolygonShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ConvexPolygonShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ConvexPolygonShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ConvexPolygonShape3D) SetPoints(points PackedVector3Array, )  {
  panic("TODO: implement")
}

func  (me *ConvexPolygonShape3D) GetPoints()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
