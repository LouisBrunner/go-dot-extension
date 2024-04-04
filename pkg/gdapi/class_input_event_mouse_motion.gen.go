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

type InputEventMouseMotion struct {
  InputEventMouse
}

func (me *InputEventMouseMotion) BaseClass() string {
  return "InputEventMouseMotion"
}

func NewInputEventMouseMotion() *InputEventMouseMotion {
  str := StringNameFromStr("InputEventMouseMotion") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &InputEventMouseMotion{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *InputEventMouseMotion) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventMouseMotion) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventMouseMotion) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InputEventMouseMotion) SetTilt(tilt Vector2, )  {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tilt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{tilt.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMouseMotion) GetTilt() Vector2 {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tilt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *InputEventMouseMotion) SetPressure(pressure float64, )  {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pressure")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressure) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMouseMotion) GetPressure() float64 {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pressure")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventMouseMotion) SetPenInverted(pen_inverted bool, )  {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pen_inverted")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pen_inverted) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMouseMotion) GetPenInverted() bool {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pen_inverted")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventMouseMotion) SetRelative(relative Vector2, )  {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_relative")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{relative.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMouseMotion) GetRelative() Vector2 {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_relative")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *InputEventMouseMotion) SetVelocity(velocity Vector2, )  {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMouseMotion) GetVelocity() Vector2 {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
