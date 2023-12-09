// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventKey struct {
  obj gdc.ObjectPtr
}

func (me *InputEventKey) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventKey) BaseClass() string {
  return "InputEventKey"
}



// Enums

func (me *InputEventKey) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventKey) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventKey) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InputEventKey) SetPressed(pressed bool, )  {
  panic("TODO: implement")
}

func  (me *InputEventKey) SetKeycode(keycode Key, )  {
  panic("TODO: implement")
}

func  (me *InputEventKey) GetKeycode()  {
  panic("TODO: implement")
}

func  (me *InputEventKey) SetPhysicalKeycode(physical_keycode Key, )  {
  panic("TODO: implement")
}

func  (me *InputEventKey) GetPhysicalKeycode()  {
  panic("TODO: implement")
}

func  (me *InputEventKey) SetKeyLabel(key_label Key, )  {
  panic("TODO: implement")
}

func  (me *InputEventKey) GetKeyLabel()  {
  panic("TODO: implement")
}

func  (me *InputEventKey) SetUnicode(unicode int, )  {
  panic("TODO: implement")
}

func  (me *InputEventKey) GetUnicode()  {
  panic("TODO: implement")
}

func  (me *InputEventKey) SetEcho(echo bool, )  {
  panic("TODO: implement")
}

func  (me *InputEventKey) GetKeycodeWithModifiers()  {
  panic("TODO: implement")
}

func  (me *InputEventKey) GetPhysicalKeycodeWithModifiers()  {
  panic("TODO: implement")
}

func  (me *InputEventKey) GetKeyLabelWithModifiers()  {
  panic("TODO: implement")
}

func  (me *InputEventKey) AsTextKeycode()  {
  panic("TODO: implement")
}

func  (me *InputEventKey) AsTextPhysicalKeycode()  {
  panic("TODO: implement")
}

func  (me *InputEventKey) AsTextKeyLabel()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
