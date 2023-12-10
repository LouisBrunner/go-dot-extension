// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InputEventGesture struct {
  obj gdc.ObjectPtr
}

func (me *InputEventGesture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventGesture) BaseClass() string {
  return "InputEventGesture"
}



// Enums

func (me *InputEventGesture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventGesture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventGesture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InputEventGesture) SetPosition(position Vector2, )  {
  classNameV := StringNameFromStr("InputEventGesture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventGesture) GetPosition() Vector2 {
  classNameV := StringNameFromStr("InputEventGesture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *InputEventGesture) GetPropPosition() Vector2 {
  panic("TODO: implement")
}

func (me *InputEventGesture) SetPropPosition(value Vector2) {
  panic("TODO: implement")
}