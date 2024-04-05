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

type ptrsForSphereShape3DList struct {
  fnSetRadius gdc.MethodBindPtr
  fnGetRadius gdc.MethodBindPtr
}

var ptrsForSphereShape3D ptrsForSphereShape3DList

func initSphereShape3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("SphereShape3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_radius")
    defer methodName.Destroy()
    ptrsForSphereShape3D.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_radius")
    defer methodName.Destroy()
    ptrsForSphereShape3D.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

type SphereShape3D struct {
  Shape3D
}

func (me *SphereShape3D) BaseClass() string {
  return "SphereShape3D"
}

func NewSphereShape3D() *SphereShape3D {
  str := StringNameFromStr("SphereShape3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SphereShape3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *SphereShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SphereShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SphereShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SphereShape3D) SetRadius(radius float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSphereShape3D.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SphereShape3D) GetRadius() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSphereShape3D.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
