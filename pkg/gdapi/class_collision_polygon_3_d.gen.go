// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CollisionPolygon3D struct {
  obj gdc.ObjectPtr
}

func (me *CollisionPolygon3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CollisionPolygon3D) BaseClass() string {
  return "CollisionPolygon3D"
}



// Enums

func (me *CollisionPolygon3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CollisionPolygon3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CollisionPolygon3D) SetDepth(depth float32, )  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon3D) GetDepth()  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon3D) SetPolygon(polygon PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon3D) GetPolygon()  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon3D) SetDisabled(disabled bool, )  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon3D) IsDisabled()  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon3D) SetMargin(margin float32, )  {
  panic("TODO: implement")
}

func  (me *CollisionPolygon3D) GetMargin()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
