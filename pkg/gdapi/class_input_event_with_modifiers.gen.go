// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InputEventWithModifiers struct {
  InputEventFromWindow
}

func (me *InputEventWithModifiers) BaseClass() string {
  return "InputEventWithModifiers"
}

func NewInputEventWithModifiers() *InputEventWithModifiers {
  str := StringNameFromStr("InputEventWithModifiers") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &InputEventWithModifiers{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *InputEventWithModifiers) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventWithModifiers) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventWithModifiers) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InputEventWithModifiers) SetCommandOrControlAutoremap(enable bool, )  {
  classNameV := StringNameFromStr("InputEventWithModifiers")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_command_or_control_autoremap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventWithModifiers) IsCommandOrControlAutoremap() bool {
  classNameV := StringNameFromStr("InputEventWithModifiers")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_command_or_control_autoremap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventWithModifiers) IsCommandOrControlPressed() bool {
  classNameV := StringNameFromStr("InputEventWithModifiers")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_command_or_control_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventWithModifiers) SetAltPressed(pressed bool, )  {
  classNameV := StringNameFromStr("InputEventWithModifiers")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alt_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventWithModifiers) IsAltPressed() bool {
  classNameV := StringNameFromStr("InputEventWithModifiers")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_alt_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventWithModifiers) SetShiftPressed(pressed bool, )  {
  classNameV := StringNameFromStr("InputEventWithModifiers")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shift_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventWithModifiers) IsShiftPressed() bool {
  classNameV := StringNameFromStr("InputEventWithModifiers")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shift_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventWithModifiers) SetCtrlPressed(pressed bool, )  {
  classNameV := StringNameFromStr("InputEventWithModifiers")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ctrl_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventWithModifiers) IsCtrlPressed() bool {
  classNameV := StringNameFromStr("InputEventWithModifiers")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ctrl_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventWithModifiers) SetMetaPressed(pressed bool, )  {
  classNameV := StringNameFromStr("InputEventWithModifiers")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_meta_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pressed), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputEventWithModifiers) IsMetaPressed() bool {
  classNameV := StringNameFromStr("InputEventWithModifiers")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_meta_pressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputEventWithModifiers) GetModifiersMask() KeyModifierMask {
  classNameV := StringNameFromStr("InputEventWithModifiers")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modifiers_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1258259499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret KeyModifierMask

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
