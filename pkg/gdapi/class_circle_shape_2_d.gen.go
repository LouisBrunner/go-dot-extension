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

type CircleShape2D struct {
  Shape2D
}

func (me *CircleShape2D) BaseClass() string {
  return "CircleShape2D"
}

func NewCircleShape2D() *CircleShape2D {
  str := StringNameFromStr("CircleShape2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CircleShape2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CircleShape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CircleShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CircleShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CircleShape2D) SetRadius(radius float64, )  {
  classNameV := StringNameFromStr("CircleShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CircleShape2D) GetRadius() float64 {
  classNameV := StringNameFromStr("CircleShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
