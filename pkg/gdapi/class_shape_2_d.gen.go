// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Shape2D struct {
  obj gdc.ObjectPtr
}

func (me *Shape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Shape2D) BaseClass() string {
  return "Shape2D"
}



// Enums

func (me *Shape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Shape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Shape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Shape2D) SetCustomSolverBias(bias float32, )  {
  panic("TODO: implement")
}

func  (me *Shape2D) GetCustomSolverBias()  {
  panic("TODO: implement")
}

func  (me *Shape2D) Collide(local_xform Transform2D, with_shape Shape2D, shape_xform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *Shape2D) CollideWithMotion(local_xform Transform2D, local_motion Vector2, with_shape Shape2D, shape_xform Transform2D, shape_motion Vector2, )  {
  panic("TODO: implement")
}

func  (me *Shape2D) CollideAndGetContacts(local_xform Transform2D, with_shape Shape2D, shape_xform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *Shape2D) CollideWithMotionAndGetContacts(local_xform Transform2D, local_motion Vector2, with_shape Shape2D, shape_xform Transform2D, shape_motion Vector2, )  {
  panic("TODO: implement")
}

func  (me *Shape2D) Draw(canvas_item RID, color Color, )  {
  panic("TODO: implement")
}

func  (me *Shape2D) GetRect()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
