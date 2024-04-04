// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type SceneTreeTimer struct {
  RefCounted
}

func (me *SceneTreeTimer) BaseClass() string {
  return "SceneTreeTimer"
}

func NewSceneTreeTimer() *SceneTreeTimer {
  str := StringNameFromStr("SceneTreeTimer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SceneTreeTimer{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *SceneTreeTimer) SetTimeLeft(time float64, )  {
  classNameV := StringNameFromStr("SceneTreeTimer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_time_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SceneTreeTimer) GetTimeLeft() float64 {
  classNameV := StringNameFromStr("SceneTreeTimer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type SceneTreeTimerTimeoutSignalFn func()

func (me *SceneTreeTimer) ConnectTimeout(subs SignalSubscribers, fn SceneTreeTimerTimeoutSignalFn) {
  sig := StringNameFromStr("timeout")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *SceneTreeTimer) DisconnectTimeout(subs SignalSubscribers, fn SceneTreeTimerTimeoutSignalFn) {
  sig := StringNameFromStr("timeout")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
