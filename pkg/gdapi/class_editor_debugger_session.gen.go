// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("EditorDebuggerSession")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("send_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3780025912) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(message.AsCTypePtr()), gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorDebuggerSession) ToggleProfiler(profiler String, enable bool, data Array, )  {
  classNameV := StringNameFromStr("EditorDebuggerSession")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("toggle_profiler")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 35674246) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(profiler.AsCTypePtr()), gdc.ConstTypePtr(&enable), gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorDebuggerSession) IsBreaked() bool {
  classNameV := StringNameFromStr("EditorDebuggerSession")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_breaked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorDebuggerSession) IsDebuggable() bool {
  classNameV := StringNameFromStr("EditorDebuggerSession")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_debuggable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorDebuggerSession) IsActive() bool {
  classNameV := StringNameFromStr("EditorDebuggerSession")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorDebuggerSession) AddSessionTab(control Control, )  {
  classNameV := StringNameFromStr("EditorDebuggerSession")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_session_tab")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1496901182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(control.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorDebuggerSession) RemoveSessionTab(control Control, )  {
  classNameV := StringNameFromStr("EditorDebuggerSession")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_session_tab")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1496901182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(control.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API