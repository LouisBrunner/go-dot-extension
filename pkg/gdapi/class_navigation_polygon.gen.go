// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationPolygon struct {
  obj gdc.ObjectPtr
}

func (me *NavigationPolygon) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationPolygon) BaseClass() string {
  return "NavigationPolygon"
}



// Enums

func (me *NavigationPolygon) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationPolygon) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationPolygon) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *NavigationPolygon) SetVertices(vertices PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) GetVertices()  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) AddPolygon(polygon PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) GetPolygonCount()  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) GetPolygon(idx int, )  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) ClearPolygons()  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) GetNavigationMesh()  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) AddOutline(outline PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) AddOutlineAtIndex(outline PackedVector2Array, index int, )  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) GetOutlineCount()  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) SetOutline(idx int, outline PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) GetOutline(idx int, )  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) RemoveOutline(idx int, )  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) ClearOutlines()  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) MakePolygonsFromOutlines()  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) SetCellSize(cell_size float32, )  {
  panic("TODO: implement")
}

func  (me *NavigationPolygon) GetCellSize()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
