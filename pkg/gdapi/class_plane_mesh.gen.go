// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PlaneMesh struct {
  obj gdc.ObjectPtr
}

func (me *PlaneMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PlaneMesh) BaseClass() string {
  return "PlaneMesh"
}



// Enums

type PlaneMeshOrientation int
const (
  PlaneMeshOrientationFaceX PlaneMeshOrientation = 0
  PlaneMeshOrientationFaceY PlaneMeshOrientation = 1
  PlaneMeshOrientationFaceZ PlaneMeshOrientation = 2
)

func (me *PlaneMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaneMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PlaneMesh) SetSize(size Vector2, )  {
  panic("TODO: implement")
}

func  (me *PlaneMesh) GetSize()  {
  panic("TODO: implement")
}

func  (me *PlaneMesh) SetSubdivideWidth(subdivide int, )  {
  panic("TODO: implement")
}

func  (me *PlaneMesh) GetSubdivideWidth()  {
  panic("TODO: implement")
}

func  (me *PlaneMesh) SetSubdivideDepth(subdivide int, )  {
  panic("TODO: implement")
}

func  (me *PlaneMesh) GetSubdivideDepth()  {
  panic("TODO: implement")
}

func  (me *PlaneMesh) SetCenterOffset(offset Vector3, )  {
  panic("TODO: implement")
}

func  (me *PlaneMesh) GetCenterOffset()  {
  panic("TODO: implement")
}

func  (me *PlaneMesh) SetOrientation(orientation PlaneMeshOrientation, )  {
  panic("TODO: implement")
}

func  (me *PlaneMesh) GetOrientation()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
