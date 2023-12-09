// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RibbonTrailMesh struct {
  obj gdc.ObjectPtr
}

func (me *RibbonTrailMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RibbonTrailMesh) BaseClass() string {
  return "RibbonTrailMesh"
}



// Enums

type RibbonTrailMeshShape int
const (
  RibbonTrailMeshShapeShapeFlat RibbonTrailMeshShape = 0
  RibbonTrailMeshShapeShapeCross RibbonTrailMeshShape = 1
)

func (me *RibbonTrailMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RibbonTrailMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *RibbonTrailMesh) SetSize(size float32, )  {
  panic("TODO: implement")
}

func  (me *RibbonTrailMesh) GetSize()  {
  panic("TODO: implement")
}

func  (me *RibbonTrailMesh) SetSections(sections int, )  {
  panic("TODO: implement")
}

func  (me *RibbonTrailMesh) GetSections()  {
  panic("TODO: implement")
}

func  (me *RibbonTrailMesh) SetSectionLength(section_length float32, )  {
  panic("TODO: implement")
}

func  (me *RibbonTrailMesh) GetSectionLength()  {
  panic("TODO: implement")
}

func  (me *RibbonTrailMesh) SetSectionSegments(section_segments int, )  {
  panic("TODO: implement")
}

func  (me *RibbonTrailMesh) GetSectionSegments()  {
  panic("TODO: implement")
}

func  (me *RibbonTrailMesh) SetCurve(curve Curve, )  {
  panic("TODO: implement")
}

func  (me *RibbonTrailMesh) GetCurve()  {
  panic("TODO: implement")
}

func  (me *RibbonTrailMesh) SetShape(shape RibbonTrailMeshShape, )  {
  panic("TODO: implement")
}

func  (me *RibbonTrailMesh) GetShape()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
