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

type InputEventScreenTouch struct {
  InputEventFromWindow
}

func (me *InputEventScreenTouch) BaseClass() string {
  return "InputEventScreenTouch"
}

func NewInputEventScreenTouch() *InputEventScreenTouch {
  str := StringNameFromStr("InputEventScreenTouch") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &InputEventScreenTouch{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *InputEventScreenTouch) SetIndex(index int64, )  {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventScreenTouch) GetIndex() int64 {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventScreenTouch) SetPosition(position Vector2, )  {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventScreenTouch) GetPosition() Vector2 {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *InputEventScreenTouch) SetPressed(pressed bool, )  {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventScreenTouch) SetCanceled(canceled bool, )  {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_canceled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&canceled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventScreenTouch) SetDoubleTap(double_tap bool, )  {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_double_tap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&double_tap) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventScreenTouch) IsDoubleTap() bool {
  classNameV := StringNameFromStr("InputEventScreenTouch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_double_tap")
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
