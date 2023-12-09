// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorDebuggerPlugin struct {
  obj gdc.ObjectPtr
}

func (me *EditorDebuggerPlugin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorDebuggerPlugin) BaseClass() string {
  return "EditorDebuggerPlugin"
}



// Enums

func (me *EditorDebuggerPlugin) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorDebuggerPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorDebuggerPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorDebuggerPlugin) XSetupSession(session_id int, )  {
  panic("TODO: implement")
}

func  (me *EditorDebuggerPlugin) XHasCapture(capture String, )  {
  panic("TODO: implement")
}

func  (me *EditorDebuggerPlugin) XCapture(message String, data Array, session_id int, )  {
  panic("TODO: implement")
}

func  (me *EditorDebuggerPlugin) GetSession(id int, )  {
  panic("TODO: implement")
}

func  (me *EditorDebuggerPlugin) GetSessions()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
