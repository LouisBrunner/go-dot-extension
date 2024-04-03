// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorDebuggerSession struct {
  RefCounted
}

func (me *EditorDebuggerSession) BaseClass() string {
  return "EditorDebuggerSession"
}

func NewEditorDebuggerSession() *EditorDebuggerSession {
  str := StringNameFromStr("EditorDebuggerSession") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorDebuggerSession{}
  obj.SetBaseObject(objPtr)
  return obj
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
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 85656714) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(message.AsCTypePtr()), gdc.ConstTypePtr(data.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorDebuggerSession) ToggleProfiler(profiler String, enable bool, data Array, )  {
  classNameV := StringNameFromStr("EditorDebuggerSession")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("toggle_profiler")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1198443697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(profiler.AsCTypePtr()), gdc.ConstTypePtr(&enable), gdc.ConstTypePtr(data.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorDebuggerSession) IsBreaked() bool {
  classNameV := StringNameFromStr("EditorDebuggerSession")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_breaked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorDebuggerSession) IsDebuggable() bool {
  classNameV := StringNameFromStr("EditorDebuggerSession")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_debuggable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorDebuggerSession) IsActive() bool {
  classNameV := StringNameFromStr("EditorDebuggerSession")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
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

// Signals

type EditorDebuggerSessionStartedSignalFn func()

func (me *EditorDebuggerSession) ConnectStarted(subs SignalSubscribers, fn EditorDebuggerSessionStartedSignalFn) {
  sig := StringNameFromStr("started")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorDebuggerSession) DisconnectStarted(subs SignalSubscribers, fn EditorDebuggerSessionStartedSignalFn) {
  sig := StringNameFromStr("started")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorDebuggerSessionStoppedSignalFn func()

func (me *EditorDebuggerSession) ConnectStopped(subs SignalSubscribers, fn EditorDebuggerSessionStoppedSignalFn) {
  sig := StringNameFromStr("stopped")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorDebuggerSession) DisconnectStopped(subs SignalSubscribers, fn EditorDebuggerSessionStoppedSignalFn) {
  sig := StringNameFromStr("stopped")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorDebuggerSessionBreakedSignalFn func(can_debug bool, )

func (me *EditorDebuggerSession) ConnectBreaked(subs SignalSubscribers, fn EditorDebuggerSessionBreakedSignalFn) {
  sig := StringNameFromStr("breaked")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorDebuggerSession) DisconnectBreaked(subs SignalSubscribers, fn EditorDebuggerSessionBreakedSignalFn) {
  sig := StringNameFromStr("breaked")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorDebuggerSessionContinuedSignalFn func()

func (me *EditorDebuggerSession) ConnectContinued(subs SignalSubscribers, fn EditorDebuggerSessionContinuedSignalFn) {
  sig := StringNameFromStr("continued")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorDebuggerSession) DisconnectContinued(subs SignalSubscribers, fn EditorDebuggerSessionContinuedSignalFn) {
  sig := StringNameFromStr("continued")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
