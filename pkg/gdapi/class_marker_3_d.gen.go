// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForMarker3DList struct {
  fnSetGizmoExtents gdc.MethodBindPtr
  fnGetGizmoExtents gdc.MethodBindPtr
}

var ptrsForMarker3D ptrsForMarker3DList

func initMarker3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Marker3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_gizmo_extents")
    defer methodName.Destroy()
    ptrsForMarker3D.fnSetGizmoExtents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_gizmo_extents")
    defer methodName.Destroy()
    ptrsForMarker3D.fnGetGizmoExtents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&extents) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMarker3D.fnSetGizmoExtents), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Marker3D) GetGizmoExtents() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMarker3D.fnGetGizmoExtents), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
