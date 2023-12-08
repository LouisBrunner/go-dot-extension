// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  TimerTimerProcessPhysics TimerTimerProcessCallback = 0
  TimerTimerProcessIdle TimerTimerProcessCallback = 1
)
