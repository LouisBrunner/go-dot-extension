// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("Marker3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gizmo_extents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&extents), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Marker3D) GetGizmoExtents() float32 {
  classNameV := StringNameFromStr("Marker3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gizmo_extents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
