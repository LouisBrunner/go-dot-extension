// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Marker2D struct {
  obj gdc.ObjectPtr
}

func (me *Marker2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Marker2D) BaseClass() string {
  return "Marker2D"
}



// Enums

func (me *Marker2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Marker2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Marker2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Marker2D) SetGizmoExtents(extents float32, )  {
  panic("TODO: implement")
}

func  (me *Marker2D) GetGizmoExtents()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
