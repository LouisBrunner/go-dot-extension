// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForInputEventMIDIList struct {
  fnSetChannel gdc.MethodBindPtr
  fnGetChannel gdc.MethodBindPtr
  fnSetMessage gdc.MethodBindPtr
  fnGetMessage gdc.MethodBindPtr
  fnSetPitch gdc.MethodBindPtr
  fnGetPitch gdc.MethodBindPtr
  fnSetVelocity gdc.MethodBindPtr
  fnGetVelocity gdc.MethodBindPtr
  fnSetInstrument gdc.MethodBindPtr
  fnGetInstrument gdc.MethodBindPtr
  fnSetPressure gdc.MethodBindPtr
  fnGetPressure gdc.MethodBindPtr
  fnSetControllerNumber gdc.MethodBindPtr
  fnGetControllerNumber gdc.MethodBindPtr
  fnSetControllerValue gdc.MethodBindPtr
  fnGetControllerValue gdc.MethodBindPtr
}

var ptrsForInputEventMIDI ptrsForInputEventMIDIList

func initInputEventMIDIPtrs(iface gdc.Interface) {

  className := StringNameFromStr("InputEventMIDI")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_channel")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnSetChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_channel")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnGetChannel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_message")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnSetMessage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1064271510))
  }
  {
    methodName := StringNameFromStr("get_message")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnGetMessage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1936512097))
  }
  {
    methodName := StringNameFromStr("set_pitch")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnSetPitch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_pitch")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnGetPitch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_velocity")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnSetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_velocity")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnGetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_instrument")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnSetInstrument = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_instrument")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnGetInstrument = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_pressure")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnSetPressure = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_pressure")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnGetPressure = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_controller_number")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnSetControllerNumber = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_controller_number")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnGetControllerNumber = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_controller_value")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnSetControllerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_controller_value")
    defer methodName.Destroy()
    ptrsForInputEventMIDI.fnGetControllerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

type InputEventMIDI struct {
  InputEvent
}

func (me *InputEventMIDI) BaseClass() string {
  return "InputEventMIDI"
}

func NewInputEventMIDI() *InputEventMIDI {
  str := StringNameFromStr("InputEventMIDI") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &InputEventMIDI{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *InputEventMIDI) SetChannel(channel int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&channel) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnSetChannel), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMIDI) GetChannel() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnGetChannel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventMIDI) SetMessage(message MIDIMessage, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&message) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnSetMessage), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMIDI) GetMessage() MIDIMessage {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret MIDIMessage

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnGetMessage), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *InputEventMIDI) SetPitch(pitch int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pitch) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnSetPitch), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMIDI) GetPitch() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnGetPitch), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventMIDI) SetVelocity(velocity int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&velocity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnSetVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMIDI) GetVelocity() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnGetVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventMIDI) SetInstrument(instrument int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instrument) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnSetInstrument), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMIDI) GetInstrument() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnGetInstrument), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventMIDI) SetPressure(pressure int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressure) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnSetPressure), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMIDI) GetPressure() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnGetPressure), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventMIDI) SetControllerNumber(controller_number int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&controller_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnSetControllerNumber), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMIDI) GetControllerNumber() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnGetControllerNumber), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventMIDI) SetControllerValue(controller_value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&controller_value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnSetControllerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventMIDI) GetControllerValue() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForInputEventMIDI.fnGetControllerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
