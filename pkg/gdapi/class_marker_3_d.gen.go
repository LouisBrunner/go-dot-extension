// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Marker3D struct {
  Node3D
}

func (me *Marker3D) BaseClass() string {
  return "Marker3D"
}

func NewMarker3D() *Marker3D {
  str := StringNameFromStr("Marker3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Marker3D{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *Marker3D) SetGizmoExtents(extents float64, )  {
  classNameV := StringNameFromStr("Marker3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gizmo_extents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&extents), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Marker3D) GetGizmoExtents() float64 {
  classNameV := StringNameFromStr("Marker3D")
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
