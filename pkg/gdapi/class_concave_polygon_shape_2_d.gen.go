// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConcavePolygonShape2D struct {
  obj gdc.ObjectPtr
}

func (me *ConcavePolygonShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConcavePolygonShape2D) BaseClass() string {
  return "ConcavePolygonShape2D"
}



// Enums

func (me *ConcavePolygonShape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ConcavePolygonShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ConcavePolygonShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ConcavePolygonShape2D) SetSegments(segments PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *ConcavePolygonShape2D) GetSegments()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
