// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Thread struct {
  obj gdc.ObjectPtr
}

func (me *Thread) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Thread) BaseClass() string {
  return "Thread"
}

type ThreadPriority int
const (
  ThreadPriorityPriorityLow ThreadPriority = 0
  ThreadPriorityPriorityNormal ThreadPriority = 1
  ThreadPriorityPriorityHigh ThreadPriority = 2
)

func  (me *Thread) Start(callable Callable, priority ThreadPriority, ) { // TODO: return value
  // TODO: implement
}

func  (me *Thread) GetId() { // TODO: return value
  // TODO: implement
}

func  (me *Thread) IsStarted() { // TODO: return value
  // TODO: implement
}

func  (me *Thread) IsAlive() { // TODO: return value
  // TODO: implement
}

func  (me *Thread) WaitToFinish() { // TODO: return value
  // TODO: implement
}

func  ThreadSetThreadSafetyChecksEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
