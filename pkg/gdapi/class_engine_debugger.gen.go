// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EngineDebugger struct {
  obj gdc.ObjectPtr
}

func (me *EngineDebugger) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EngineDebugger) BaseClass() string {
  return "EngineDebugger"
}



// Enums

func (me *EngineDebugger) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EngineDebugger) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EngineDebugger) IsActive()  {
  panic("TODO: implement")
}

func  (me *EngineDebugger) RegisterProfiler(name StringName, profiler EngineProfiler, )  {
  panic("TODO: implement")
}

func  (me *EngineDebugger) UnregisterProfiler(name StringName, )  {
  panic("TODO: implement")
}

func  (me *EngineDebugger) IsProfiling(name StringName, )  {
  panic("TODO: implement")
}

func  (me *EngineDebugger) HasProfiler(name StringName, )  {
  panic("TODO: implement")
}

func  (me *EngineDebugger) ProfilerAddFrameData(name StringName, data Array, )  {
  panic("TODO: implement")
}

func  (me *EngineDebugger) ProfilerEnable(name StringName, enable bool, arguments Array, )  {
  panic("TODO: implement")
}

func  (me *EngineDebugger) RegisterMessageCapture(name StringName, callable Callable, )  {
  panic("TODO: implement")
}

func  (me *EngineDebugger) UnregisterMessageCapture(name StringName, )  {
  panic("TODO: implement")
}

func  (me *EngineDebugger) HasCapture(name StringName, )  {
  panic("TODO: implement")
}

func  (me *EngineDebugger) SendMessage(message String, data Array, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
