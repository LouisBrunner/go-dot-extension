// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *InputEventMIDI) GetChannel()  {
  panic("TODO: implement")
}

func  (me *InputEventMIDI) SetMessage(message MIDIMessage, )  {
  panic("TODO: implement")
}

func  (me *InputEventMIDI) GetMessage()  {
  panic("TODO: implement")
}

func  (me *InputEventMIDI) SetPitch(pitch int, )  {
  panic("TODO: implement")
}

func  (me *InputEventMIDI) GetPitch()  {
  panic("TODO: implement")
}

func  (me *InputEventMIDI) SetVelocity(velocity int, )  {
  panic("TODO: implement")
}

func  (me *InputEventMIDI) GetVelocity()  {
  panic("TODO: implement")
}

func  (me *InputEventMIDI) SetInstrument(instrument int, )  {
  panic("TODO: implement")
}

func  (me *InputEventMIDI) GetInstrument()  {
  panic("TODO: implement")
}

func  (me *InputEventMIDI) SetPressure(pressure int, )  {
  panic("TODO: implement")
}

func  (me *InputEventMIDI) GetPressure()  {
  panic("TODO: implement")
}

func  (me *InputEventMIDI) SetControllerNumber(controller_number int, )  {
  panic("TODO: implement")
}

func  (me *InputEventMIDI) GetControllerNumber()  {
  panic("TODO: implement")
}

func  (me *InputEventMIDI) SetControllerValue(controller_value int, )  {
  panic("TODO: implement")
}

func  (me *InputEventMIDI) GetControllerValue()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
