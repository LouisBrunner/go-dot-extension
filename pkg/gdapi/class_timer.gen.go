// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *Timer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Timer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Timer) SetWaitTime(time_sec float32, )  {
  panic("TODO: implement")
}

func  (me *Timer) GetWaitTime()  {
  panic("TODO: implement")
}

func  (me *Timer) SetOneShot(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Timer) IsOneShot()  {
  panic("TODO: implement")
}

func  (me *Timer) SetAutostart(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Timer) HasAutostart()  {
  panic("TODO: implement")
}

func  (me *Timer) Start(time_sec float32, )  {
  panic("TODO: implement")
}

func  (me *Timer) Stop()  {
  panic("TODO: implement")
}

func  (me *Timer) SetPaused(paused bool, )  {
  panic("TODO: implement")
}

func  (me *Timer) IsPaused()  {
  panic("TODO: implement")
}

func  (me *Timer) IsStopped()  {
  panic("TODO: implement")
}

func  (me *Timer) GetTimeLeft()  {
  panic("TODO: implement")
}

func  (me *Timer) SetTimerProcessCallback(callback TimerTimerProcessCallback, )  {
  panic("TODO: implement")
}

func  (me *Timer) GetTimerProcessCallback()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
