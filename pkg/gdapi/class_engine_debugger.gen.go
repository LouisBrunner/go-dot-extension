// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *EngineDebugger) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EngineDebugger) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EngineDebugger) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EngineDebugger) IsActive() bool {
  classNameV := StringNameFromStr("EngineDebugger")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EngineDebugger) RegisterProfiler(name StringName, profiler EngineProfiler, )  {
  classNameV := StringNameFromStr("EngineDebugger")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_profiler")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3651669560) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(profiler.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EngineDebugger) UnregisterProfiler(name StringName, )  {
  classNameV := StringNameFromStr("EngineDebugger")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unregister_profiler")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EngineDebugger) IsProfiling(name StringName, ) bool {
  classNameV := StringNameFromStr("EngineDebugger")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_profiling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2041966384) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EngineDebugger) HasProfiler(name StringName, ) bool {
  classNameV := StringNameFromStr("EngineDebugger")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_profiler")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2041966384) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EngineDebugger) ProfilerAddFrameData(name StringName, data Array, )  {
  classNameV := StringNameFromStr("EngineDebugger")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("profiler_add_frame_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1895267858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EngineDebugger) ProfilerEnable(name StringName, enable bool, arguments Array, )  {
  classNameV := StringNameFromStr("EngineDebugger")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("profiler_enable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 438160728) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&enable), gdc.ConstTypePtr(arguments.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EngineDebugger) RegisterMessageCapture(name StringName, callable Callable, )  {
  classNameV := StringNameFromStr("EngineDebugger")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_message_capture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1874754934) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(callable.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EngineDebugger) UnregisterMessageCapture(name StringName, )  {
  classNameV := StringNameFromStr("EngineDebugger")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unregister_message_capture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EngineDebugger) HasCapture(name StringName, ) bool {
  classNameV := StringNameFromStr("EngineDebugger")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_capture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2041966384) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EngineDebugger) SendMessage(message String, data Array, )  {
  classNameV := StringNameFromStr("EngineDebugger")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("send_message")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1209351045) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(message.AsCTypePtr()), gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties