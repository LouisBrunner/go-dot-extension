// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type SeparationRayShape2D struct {
  Shape2D
}

func (me *SeparationRayShape2D) BaseClass() string {
  return "SeparationRayShape2D"
}

func NewSeparationRayShape2D() *SeparationRayShape2D {
  str := StringNameFromStr("SeparationRayShape2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SeparationRayShape2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *SeparationRayShape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SeparationRayShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SeparationRayShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SeparationRayShape2D) SetLength(length float64, )  {
  classNameV := StringNameFromStr("SeparationRayShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SeparationRayShape2D) GetLength() float64 {
  classNameV := StringNameFromStr("SeparationRayShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SeparationRayShape2D) SetSlideOnSlope(active bool, )  {
  classNameV := StringNameFromStr("SeparationRayShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slide_on_slope")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SeparationRayShape2D) GetSlideOnSlope() bool {
  classNameV := StringNameFromStr("SeparationRayShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slide_on_slope")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
