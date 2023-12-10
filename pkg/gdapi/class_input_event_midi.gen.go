// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InputEventMIDI struct {
  obj gdc.ObjectPtr
}

func (me *InputEventMIDI) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventMIDI) BaseClass() string {
  return "InputEventMIDI"
}



// Enums

func (me *InputEventMIDI) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventMIDI) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventMIDI) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InputEventMIDI) SetChannel(channel int, )  {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_channel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventMIDI) GetChannel() int {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_channel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventMIDI) SetMessage(message MIDIMessage, )  {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1064271510) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&message), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventMIDI) GetMessage() MIDIMessage {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1936512097) // FIXME: should cache?
  var ret MIDIMessage
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventMIDI) SetPitch(pitch int, )  {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pitch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pitch), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventMIDI) GetPitch() int {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pitch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventMIDI) SetVelocity(velocity int, )  {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&velocity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventMIDI) GetVelocity() int {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventMIDI) SetInstrument(instrument int, )  {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_instrument")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instrument), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventMIDI) GetInstrument() int {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instrument")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventMIDI) SetPressure(pressure int, )  {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pressure")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressure), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventMIDI) GetPressure() int {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pressure")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventMIDI) SetControllerNumber(controller_number int, )  {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_controller_number")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&controller_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventMIDI) GetControllerNumber() int {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_controller_number")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputEventMIDI) SetControllerValue(controller_value int, )  {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_controller_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&controller_value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputEventMIDI) GetControllerValue() int {
  classNameV := StringNameFromStr("InputEventMIDI")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_controller_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *InputEventMIDI) GetPropChannel() int {
  panic("TODO: implement")
}

func (me *InputEventMIDI) SetPropChannel(value int) {
  panic("TODO: implement")
}

func (me *InputEventMIDI) GetPropMessage() int {
  panic("TODO: implement")
}

func (me *InputEventMIDI) SetPropMessage(value int) {
  panic("TODO: implement")
}

func (me *InputEventMIDI) GetPropPitch() int {
  panic("TODO: implement")
}

func (me *InputEventMIDI) SetPropPitch(value int) {
  panic("TODO: implement")
}

func (me *InputEventMIDI) GetPropVelocity() int {
  panic("TODO: implement")
}

func (me *InputEventMIDI) SetPropVelocity(value int) {
  panic("TODO: implement")
}

func (me *InputEventMIDI) GetPropInstrument() int {
  panic("TODO: implement")
}

func (me *InputEventMIDI) SetPropInstrument(value int) {
  panic("TODO: implement")
}

func (me *InputEventMIDI) GetPropPressure() int {
  panic("TODO: implement")
}

func (me *InputEventMIDI) SetPropPressure(value int) {
  panic("TODO: implement")
}

func (me *InputEventMIDI) GetPropControllerNumber() int {
  panic("TODO: implement")
}

func (me *InputEventMIDI) SetPropControllerNumber(value int) {
  panic("TODO: implement")
}

func (me *InputEventMIDI) GetPropControllerValue() int {
  panic("TODO: implement")
}

func (me *InputEventMIDI) SetPropControllerValue(value int) {
  panic("TODO: implement")
}