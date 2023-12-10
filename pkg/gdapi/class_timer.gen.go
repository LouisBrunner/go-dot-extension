// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Timer struct {
  obj gdc.ObjectPtr
}

func (me *Timer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Timer) BaseClass() string {
  return "Timer"
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

func  (me *Timer) SetWaitTime(time_sec float32, )  {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wait_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_sec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Timer) GetWaitTime() float32 {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wait_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Timer) SetOneShot(enable bool, )  {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_one_shot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Timer) IsOneShot() bool {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_one_shot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Timer) SetAutostart(enable bool, )  {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autostart")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Timer) HasAutostart() bool {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_autostart")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Timer) Start(time_sec float32, )  {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1392008558) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_sec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Timer) Stop()  {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Timer) SetPaused(paused bool, )  {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&paused), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Timer) IsPaused() bool {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Timer) IsStopped() bool {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_stopped")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Timer) GetTimeLeft() float32 {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Timer) SetTimerProcessCallback(callback TimerTimerProcessCallback, )  {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_timer_process_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3469495063) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&callback), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Timer) GetTimerProcessCallback() TimerTimerProcessCallback {
  classNameV := StringNameFromStr("Timer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_timer_process_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2672570227) // FIXME: should cache?
  var ret TimerTimerProcessCallback
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Timer) GetPropProcessCallback() int {
  panic("TODO: implement")
}

func (me *Timer) SetPropProcessCallback(value int) {
  panic("TODO: implement")
}

func (me *Timer) GetPropWaitTime() float32 {
  panic("TODO: implement")
}

func (me *Timer) SetPropWaitTime(value float32) {
  panic("TODO: implement")
}

func (me *Timer) GetPropOneShot() bool {
  panic("TODO: implement")
}

func (me *Timer) SetPropOneShot(value bool) {
  panic("TODO: implement")
}

func (me *Timer) GetPropAutostart() bool {
  panic("TODO: implement")
}

func (me *Timer) SetPropAutostart(value bool) {
  panic("TODO: implement")
}

func (me *Timer) GetPropPaused() bool {
  panic("TODO: implement")
}

func (me *Timer) SetPropPaused(value bool) {
  panic("TODO: implement")
}

func (me *Timer) GetPropTimeLeft() float32 {
  panic("TODO: implement")
}

func (me *Timer) SetPropTimeLeft(value float32) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API