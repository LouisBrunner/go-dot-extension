// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PolygonPathFinder struct {
  obj gdc.ObjectPtr
}

func (me *PolygonPathFinder) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PolygonPathFinder) BaseClass() string {
  return "PolygonPathFinder"
}



// Enums

func (me *PolygonPathFinder) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PolygonPathFinder) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PolygonPathFinder) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PolygonPathFinder) Setup(points PackedVector2Array, connections PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *PolygonPathFinder) FindPath(from Vector2, to Vector2, )  {
  panic("TODO: implement")
}

func  (me *PolygonPathFinder) GetIntersections(from Vector2, to Vector2, )  {
  panic("TODO: implement")
}

func  (me *PolygonPathFinder) GetClosestPoint(point Vector2, )  {
  panic("TODO: implement")
}

func  (me *PolygonPathFinder) IsPointInside(point Vector2, )  {
  panic("TODO: implement")
}

func  (me *PolygonPathFinder) SetPointPenalty(idx int, penalty float32, )  {
  panic("TODO: implement")
}

func  (me *PolygonPathFinder) GetPointPenalty(idx int, )  {
  panic("TODO: implement")
}

func  (me *PolygonPathFinder) GetBounds()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
