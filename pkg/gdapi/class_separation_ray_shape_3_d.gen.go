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

type ptrsForSeparationRayShape3DList struct {
  fnSetLength gdc.MethodBindPtr
  fnGetLength gdc.MethodBindPtr
  fnSetSlideOnSlope gdc.MethodBindPtr
  fnGetSlideOnSlope gdc.MethodBindPtr
}

var ptrsForSeparationRayShape3D ptrsForSeparationRayShape3DList

func initSeparationRayShape3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("SeparationRayShape3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_length")
    defer methodName.Destroy()
    ptrsForSeparationRayShape3D.fnSetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_length")
    defer methodName.Destroy()
    ptrsForSeparationRayShape3D.fnGetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_slide_on_slope")
    defer methodName.Destroy()
    ptrsForSeparationRayShape3D.fnSetSlideOnSlope = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_slide_on_slope")
    defer methodName.Destroy()
    ptrsForSeparationRayShape3D.fnGetSlideOnSlope = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type SeparationRayShape3D struct {
  Shape3D
}

func (me *SeparationRayShape3D) BaseClass() string {
  return "SeparationRayShape3D"
}

func NewSeparationRayShape3D() *SeparationRayShape3D {
  str := StringNameFromStr("SeparationRayShape3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SeparationRayShape3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *SeparationRayShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SeparationRayShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SeparationRayShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SeparationRayShape3D) SetLength(length float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSeparationRayShape3D.fnSetLength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SeparationRayShape3D) GetLength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSeparationRayShape3D.fnGetLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SeparationRayShape3D) SetSlideOnSlope(active bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSeparationRayShape3D.fnSetSlideOnSlope), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SeparationRayShape3D) GetSlideOnSlope() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSeparationRayShape3D.fnGetSlideOnSlope), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
