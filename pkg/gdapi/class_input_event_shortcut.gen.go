// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventShortcut struct {
  obj gdc.ObjectPtr
}

func (me *InputEventShortcut) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventShortcut) BaseClass() string {
  return "InputEventShortcut"
}



// Enums

func (me *InputEventShortcut) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventShortcut) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InputEventShortcut) SetShortcut(shortcut Shortcut, )  {
  panic("TODO: implement")
}

func  (me *InputEventShortcut) GetShortcut()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
