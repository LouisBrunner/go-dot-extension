// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InputEvent struct {
  obj gdc.ObjectPtr
}

func (me *InputEvent) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEvent) BaseClass() string {
  return "InputEvent"
}



// Enums

func (me *InputEvent) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEvent) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEvent) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InputEvent) SetDevice(device int, )  {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_device")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&device), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEvent) GetDevice() int {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_device")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEvent) IsAction(action StringName, exact_match bool, ) bool {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1558498928) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), gdc.ConstTypePtr(&exact_match), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEvent) IsActionPressed(action StringName, allow_echo bool, exact_match bool, ) bool {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_action_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1631499404) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), gdc.ConstTypePtr(&allow_echo), gdc.ConstTypePtr(&exact_match), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEvent) IsActionReleased(action StringName, exact_match bool, ) bool {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_action_released")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1558498928) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), gdc.ConstTypePtr(&exact_match), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEvent) GetActionStrength(action StringName, exact_match bool, ) float32 {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_action_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 801543509) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), gdc.ConstTypePtr(&exact_match), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEvent) IsCanceled() bool {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_canceled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEvent) IsPressed() bool {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEvent) IsReleased() bool {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_released")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEvent) IsEcho() bool {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_echo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEvent) AsText() String {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("as_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEvent) IsMatch(event InputEvent, exact_match bool, ) bool {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_match")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3392494811) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(event.AsCTypePtr()), gdc.ConstTypePtr(&exact_match), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEvent) IsActionType() bool {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_action_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEvent) Accumulate(with_event InputEvent, ) bool {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("accumulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1062211774) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(with_event.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEvent) XformedBy(xform Transform2D, local_ofs Vector2, ) InputEvent {
  classNameV := StringNameFromStr("InputEvent")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("xformed_by")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2747409789) // FIXME: should cache?
  var ret InputEvent
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(xform.AsCTypePtr()), gdc.ConstTypePtr(local_ofs.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *InputEvent) GetPropDevice() int {
  panic("TODO: implement")
}

func (me *InputEvent) SetPropDevice(value int) {
  panic("TODO: implement")
}