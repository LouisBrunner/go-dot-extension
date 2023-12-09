// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CylinderMesh struct {
  obj gdc.ObjectPtr
}

func (me *CylinderMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CylinderMesh) BaseClass() string {
  return "CylinderMesh"
}



// Enums

func (me *CylinderMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CylinderMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CylinderMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CylinderMesh) SetTopRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *CylinderMesh) GetTopRadius()  {
  panic("TODO: implement")
}

func  (me *CylinderMesh) SetBottomRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *CylinderMesh) GetBottomRadius()  {
  panic("TODO: implement")
}

func  (me *CylinderMesh) SetHeight(height float32, )  {
  panic("TODO: implement")
}

func  (me *CylinderMesh) GetHeight()  {
  panic("TODO: implement")
}

func  (me *CylinderMesh) SetRadialSegments(segments int, )  {
  panic("TODO: implement")
}

func  (me *CylinderMesh) GetRadialSegments()  {
  panic("TODO: implement")
}

func  (me *CylinderMesh) SetRings(rings int, )  {
  panic("TODO: implement")
}

func  (me *CylinderMesh) GetRings()  {
  panic("TODO: implement")
}

func  (me *CylinderMesh) SetCapTop(cap_top bool, )  {
  panic("TODO: implement")
}

func  (me *CylinderMesh) IsCapTop()  {
  panic("TODO: implement")
}

func  (me *CylinderMesh) SetCapBottom(cap_bottom bool, )  {
  panic("TODO: implement")
}

func  (me *CylinderMesh) IsCapBottom()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
