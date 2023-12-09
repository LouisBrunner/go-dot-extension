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



// Enums

type ThreadPriority int
const (
  ThreadPriorityPriorityLow ThreadPriority = 0
  ThreadPriorityPriorityNormal ThreadPriority = 1
  ThreadPriorityPriorityHigh ThreadPriority = 2
)

func (me *Thread) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Thread) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Thread) Start(callable Callable, priority ThreadPriority, )  {
  panic("TODO: implement")
}

func  (me *Thread) GetId()  {
  panic("TODO: implement")
}

func  (me *Thread) IsStarted()  {
  panic("TODO: implement")
}

func  (me *Thread) IsAlive()  {
  panic("TODO: implement")
}

func  (me *Thread) WaitToFinish()  {
  panic("TODO: implement")
}

func  ThreadSetThreadSafetyChecksEnabled(enabled bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
