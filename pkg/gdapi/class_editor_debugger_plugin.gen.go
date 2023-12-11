// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *EditorDebuggerPlugin) GetSession(id int, ) EditorDebuggerSession {
  classNameV := StringNameFromStr("EditorDebuggerPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_session")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3061968499) // FIXME: should cache?
  var ret EditorDebuggerSession
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorDebuggerPlugin) GetSessions() Array {
  classNameV := StringNameFromStr("EditorDebuggerPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sessions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
