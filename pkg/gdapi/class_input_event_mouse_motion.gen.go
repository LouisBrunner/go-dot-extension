// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InputEventMouseMotion struct {
  obj gdc.ObjectPtr
}

func (me *InputEventMouseMotion) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventMouseMotion) BaseClass() string {
  return "InputEventMouseMotion"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tilt.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventMouseMotion) GetTilt() Vector2 {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tilt")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventMouseMotion) SetPressure(pressure float32, )  {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pressure")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressure), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventMouseMotion) GetPressure() float32 {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pressure")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventMouseMotion) SetPenInverted(pen_inverted bool, )  {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pen_inverted")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pen_inverted), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventMouseMotion) GetPenInverted() bool {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pen_inverted")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventMouseMotion) SetRelative(relative Vector2, )  {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_relative")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(relative.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventMouseMotion) GetRelative() Vector2 {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_relative")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventMouseMotion) SetVelocity(velocity Vector2, )  {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventMouseMotion) GetVelocity() Vector2 {
  classNameV := StringNameFromStr("InputEventMouseMotion")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *InputEventMouseMotion) GetPropTilt() Vector2 {
  panic("TODO: implement")
}

func (me *InputEventMouseMotion) SetPropTilt(value Vector2) {
  panic("TODO: implement")
}

func (me *InputEventMouseMotion) GetPropPressure() float32 {
  panic("TODO: implement")
}

func (me *InputEventMouseMotion) SetPropPressure(value float32) {
  panic("TODO: implement")
}

func (me *InputEventMouseMotion) GetPropPenInverted() bool {
  panic("TODO: implement")
}

func (me *InputEventMouseMotion) SetPropPenInverted(value bool) {
  panic("TODO: implement")
}

func (me *InputEventMouseMotion) GetPropRelative() Vector2 {
  panic("TODO: implement")
}

func (me *InputEventMouseMotion) SetPropRelative(value Vector2) {
  panic("TODO: implement")
}

func (me *InputEventMouseMotion) GetPropVelocity() Vector2 {
  panic("TODO: implement")
}

func (me *InputEventMouseMotion) SetPropVelocity(value Vector2) {
  panic("TODO: implement")
}