// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorDebuggerSession struct {
  obj gdc.ObjectPtr
}

func (me *EditorDebuggerSession) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorDebuggerSession) BaseClass() string {
  return "EditorDebuggerSession"
}



// Enums

func (me *EditorDebuggerSession) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorDebuggerSession) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorDebuggerSession) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorDebuggerSession) SendMessage(message String, data Array, )  {
  panic("TODO: implement")
}

func  (me *EditorDebuggerSession) ToggleProfiler(profiler String, enable bool, data Array, )  {
  panic("TODO: implement")
}

func  (me *EditorDebuggerSession) IsBreaked()  {
  panic("TODO: implement")
}

func  (me *EditorDebuggerSession) IsDebuggable()  {
  panic("TODO: implement")
}

func  (me *EditorDebuggerSession) IsActive()  {
  panic("TODO: implement")
}

func  (me *EditorDebuggerSession) AddSessionTab(control Control, )  {
  panic("TODO: implement")
}

func  (me *EditorDebuggerSession) RemoveSessionTab(control Control, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
