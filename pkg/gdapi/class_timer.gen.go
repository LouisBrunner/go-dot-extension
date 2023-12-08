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

type TimerTimerProcessCallback int
const (
  TimerTimerProcessCallbackTimerProcessPhysics TimerTimerProcessCallback = 0
  TimerTimerProcessCallbackTimerProcessIdle TimerTimerProcessCallback = 1
)

func  (me *Timer) SetWaitTime(time_sec float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Timer) GetWaitTime() { // TODO: return value
  // TODO: implement
}

func  (me *Timer) SetOneShot(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Timer) IsOneShot() { // TODO: return value
  // TODO: implement
}

func  (me *Timer) SetAutostart(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Timer) HasAutostart() { // TODO: return value
  // TODO: implement
}

func  (me *Timer) Start(time_sec float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Timer) Stop() { // TODO: return value
  // TODO: implement
}

func  (me *Timer) SetPaused(paused bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Timer) IsPaused() { // TODO: return value
  // TODO: implement
}

func  (me *Timer) IsStopped() { // TODO: return value
  // TODO: implement
}

func  (me *Timer) GetTimeLeft() { // TODO: return value
  // TODO: implement
}

func  (me *Timer) SetTimerProcessCallback(callback TimerTimerProcessCallback, ) { // TODO: return value
  // TODO: implement
}

func  (me *Timer) GetTimerProcessCallback() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
