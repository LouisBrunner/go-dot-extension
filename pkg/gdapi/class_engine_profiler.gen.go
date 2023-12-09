// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EngineProfiler struct {
  obj gdc.ObjectPtr
}

func (me *EngineProfiler) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EngineProfiler) BaseClass() string {
  return "EngineProfiler"
}



// Enums

func (me *EngineProfiler) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EngineProfiler) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EngineProfiler) XToggle(enable bool, options Array, )  {
  panic("TODO: implement")
}

func  (me *EngineProfiler) XAddFrame(data Array, )  {
  panic("TODO: implement")
}

func  (me *EngineProfiler) XTick(frame_time float32, process_time float32, physics_time float32, physics_frame_time float32, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
