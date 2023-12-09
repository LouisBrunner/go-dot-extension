// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConcavePolygonShape3D struct {
  obj gdc.ObjectPtr
}

func (me *ConcavePolygonShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConcavePolygonShape3D) BaseClass() string {
  return "ConcavePolygonShape3D"
}



// Enums

func (me *ConcavePolygonShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ConcavePolygonShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ConcavePolygonShape3D) SetFaces(faces PackedVector3Array, )  {
  panic("TODO: implement")
}

func  (me *ConcavePolygonShape3D) GetFaces()  {
  panic("TODO: implement")
}

func  (me *ConcavePolygonShape3D) SetBackfaceCollisionEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *ConcavePolygonShape3D) IsBackfaceCollisionEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
