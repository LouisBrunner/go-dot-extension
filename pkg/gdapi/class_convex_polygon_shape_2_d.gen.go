// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConvexPolygonShape2D struct {
  obj gdc.ObjectPtr
}

func (me *ConvexPolygonShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConvexPolygonShape2D) BaseClass() string {
  return "ConvexPolygonShape2D"
}



// Enums

func (me *ConvexPolygonShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ConvexPolygonShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ConvexPolygonShape2D) SetPointCloud(point_cloud PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *ConvexPolygonShape2D) SetPoints(points PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *ConvexPolygonShape2D) GetPoints()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
