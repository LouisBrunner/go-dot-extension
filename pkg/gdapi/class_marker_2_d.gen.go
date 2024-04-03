// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Marker2D struct {
  Node2D
}

func (me *Marker2D) BaseClass() string {
  return "Marker2D"
}

func NewMarker2D() *Marker2D {
  str := StringNameFromStr("Marker2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Marker2D{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *Marker2D) SetGizmoExtents(extents float64, )  {
  classNameV := StringNameFromStr("Marker2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gizmo_extents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&extents), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Marker2D) GetGizmoExtents() float64 {
  classNameV := StringNameFromStr("Marker2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gizmo_extents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
