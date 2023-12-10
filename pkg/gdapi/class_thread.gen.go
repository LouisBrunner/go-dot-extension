// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *Thread) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Thread) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Thread) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Thread) Start(callable Callable, priority ThreadPriority, ) Error {
  classNameV := StringNameFromStr("Thread")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2779832528) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(callable.AsCTypePtr()), gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Thread) GetId() String {
  classNameV := StringNameFromStr("Thread")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Thread) IsStarted() bool {
  classNameV := StringNameFromStr("Thread")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_started")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Thread) IsAlive() bool {
  classNameV := StringNameFromStr("Thread")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_alive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Thread) WaitToFinish() Variant {
  classNameV := StringNameFromStr("Thread")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("wait_to_finish")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1460262497) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  ThreadSetThreadSafetyChecksEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("Thread")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_thread_safety_checks_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), nil)
}

// Properties