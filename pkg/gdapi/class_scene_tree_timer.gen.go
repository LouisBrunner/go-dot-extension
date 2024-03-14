// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SceneTreeTimer struct {
  RefCounted
}

func (me *SceneTreeTimer) BaseClass() string {
  return "SceneTreeTimer"
}



// Enums

func (me *SceneTreeTimer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SceneTreeTimer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SceneTreeTimer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SceneTreeTimer) SetTimeLeft(time float32, )  {
  classNameV := StringNameFromStr("SceneTreeTimer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_time_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SceneTreeTimer) GetTimeLeft() float32 {
  classNameV := StringNameFromStr("SceneTreeTimer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type SceneTreeTimerTimeoutSignalFn func()

func (me *SceneTreeTimer) ConnectTimeout(subs SignalSubscribers, fn SceneTreeTimerTimeoutSignalFn) {
  sig := StringNameFromStr("timeout")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *SceneTreeTimer) DisconnectTimeout(subs SignalSubscribers, fn SceneTreeTimerTimeoutSignalFn) {
  sig := StringNameFromStr("timeout")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
