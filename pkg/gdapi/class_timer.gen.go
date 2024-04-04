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

type Timer struct {
  Node
}

func (me *Timer) BaseClass() string {
  return "Timer"
}

func NewTimer() *Timer {
  str := StringNameFromStr("Timer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Timer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type TimerTimerProcessCallback int
const (
  TimerTimerProcessCallbackTimerProcessPhysics TimerTimerProcessCallback = 0
  TimerTimerProcessCallbackTimerProcessIdle TimerTimerProcessCallback = 1
)

func (me *Timer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Timer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Timer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Timer) SetWaitTime(time_sec float64, )  {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wait_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Timer) GetWaitTime() float64 {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wait_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Timer) SetOneShot(enable bool, )  {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_one_shot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Timer) IsOneShot() bool {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_one_shot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Timer) SetAutostart(enable bool, )  {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autostart")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Timer) HasAutostart() bool {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_autostart")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Timer) Start(time_sec float64, )  {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1392008558) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Timer) Stop()  {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Timer) SetPaused(paused bool, )  {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&paused) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Timer) IsPaused() bool {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Timer) IsStopped() bool {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_stopped")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Timer) GetTimeLeft() float64 {
  classNameV := StringNameFromStr("Timer")
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

func  (me *Timer) SetTimerProcessCallback(callback TimerTimerProcessCallback, )  {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_timer_process_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3469495063) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&callback) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Timer) GetTimerProcessCallback() TimerTimerProcessCallback {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_timer_process_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2672570227) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TimerTimerProcessCallback

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TimerTimeoutSignalFn func()

func (me *Timer) ConnectTimeout(subs SignalSubscribers, fn TimerTimeoutSignalFn) {
  sig := StringNameFromStr("timeout")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Timer) DisconnectTimeout(subs SignalSubscribers, fn TimerTimeoutSignalFn) {
  sig := StringNameFromStr("timeout")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
