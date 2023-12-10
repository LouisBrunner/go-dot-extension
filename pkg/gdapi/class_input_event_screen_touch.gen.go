// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InputEventScreenTouch struct {
  obj gdc.ObjectPtr
}

func (me *InputEventScreenTouch) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventScreenTouch) BaseClass() string {
  return "InputEventScreenTouch"
}



// Enums

func (me *InputEventScreenTouch) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventScreenTouch) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventScreenTouch) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InputEventScreenTouch) SetIndex(index int, )  {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventScreenTouch) GetIndex() int {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventScreenTouch) SetPosition(position Vector2, )  {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventScreenTouch) GetPosition() Vector2 {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventScreenTouch) SetPressed(pressed bool, )  {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventScreenTouch) SetCanceled(canceled bool, )  {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_canceled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&canceled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventScreenTouch) SetDoubleTap(double_tap bool, )  {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_double_tap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&double_tap), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventScreenTouch) IsDoubleTap() bool {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_double_tap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *InputEventScreenTouch) GetPropIndex() int {
  panic("TODO: implement")
}

func (me *InputEventScreenTouch) SetPropIndex(value int) {
  panic("TODO: implement")
}

func (me *InputEventScreenTouch) GetPropPosition() Vector2 {
  panic("TODO: implement")
}

func (me *InputEventScreenTouch) SetPropPosition(value Vector2) {
  panic("TODO: implement")
}

func (me *InputEventScreenTouch) GetPropCanceled() bool {
  panic("TODO: implement")
}

func (me *InputEventScreenTouch) SetPropCanceled(value bool) {
  panic("TODO: implement")
}

func (me *InputEventScreenTouch) GetPropPressed() bool {
  panic("TODO: implement")
}

func (me *InputEventScreenTouch) SetPropPressed(value bool) {
  panic("TODO: implement")
}

func (me *InputEventScreenTouch) GetPropDoubleTap() bool {
  panic("TODO: implement")
}

func (me *InputEventScreenTouch) SetPropDoubleTap(value bool) {
  panic("TODO: implement")
}