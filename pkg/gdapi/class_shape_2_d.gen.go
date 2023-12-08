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

func  (me *Shape2D) SetCustomSolverBias(bias float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Shape2D) GetCustomSolverBias() { // TODO: return value
  // TODO: implement
}

func  (me *Shape2D) Collide(local_xform Transform2D, with_shape Shape2D, shape_xform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Shape2D) CollideWithMotion(local_xform Transform2D, local_motion Vector2, with_shape Shape2D, shape_xform Transform2D, shape_motion Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Shape2D) CollideAndGetContacts(local_xform Transform2D, with_shape Shape2D, shape_xform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Shape2D) CollideWithMotionAndGetContacts(local_xform Transform2D, local_motion Vector2, with_shape Shape2D, shape_xform Transform2D, shape_motion Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Shape2D) Draw(canvas_item RID, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Shape2D) GetRect() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
