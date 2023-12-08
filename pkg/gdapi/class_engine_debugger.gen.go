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

func  (me *EngineDebugger) IsActive() { // TODO: return value
  // TODO: implement
}

func  (me *EngineDebugger) RegisterProfiler(name StringName, profiler EngineProfiler, ) { // TODO: return value
  // TODO: implement
}

func  (me *EngineDebugger) UnregisterProfiler(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *EngineDebugger) IsProfiling(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *EngineDebugger) HasProfiler(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *EngineDebugger) ProfilerAddFrameData(name StringName, data Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *EngineDebugger) ProfilerEnable(name StringName, enable bool, arguments Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *EngineDebugger) RegisterMessageCapture(name StringName, callable Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *EngineDebugger) UnregisterMessageCapture(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *EngineDebugger) HasCapture(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *EngineDebugger) SendMessage(message String, data Array, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
