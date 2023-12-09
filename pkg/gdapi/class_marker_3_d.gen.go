// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Marker3D struct {
  obj gdc.ObjectPtr
}

func (me *Marker3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Marker3D) BaseClass() string {
  return "Marker3D"
}



// Enums

func (me *Marker3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Marker3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Marker3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Marker3D) SetGizmoExtents(extents float32, )  {
  panic("TODO: implement")
}

func  (me *Marker3D) GetGizmoExtents()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
