// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type WorkerThreadPool struct {
  Object
}

func (me *WorkerThreadPool) BaseClass() string {
  return "WorkerThreadPool"
}

func NewWorkerThreadPool() *WorkerThreadPool {
  str := StringNameFromStr("WorkerThreadPool") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &WorkerThreadPool{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *WorkerThreadPool) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WorkerThreadPool) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WorkerThreadPool) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *WorkerThreadPool) AddTask(action Callable, high_priority bool, description String, ) int64 {
  classNameV := StringNameFromStr("WorkerThreadPool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_task")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3745067146) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&high_priority) , description.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&high_priority)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WorkerThreadPool) IsTaskCompleted(task_id int64, ) bool {
  classNameV := StringNameFromStr("WorkerThreadPool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_task_completed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&task_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&task_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WorkerThreadPool) WaitForTaskCompletion(task_id int64, ) Error {
  classNameV := StringNameFromStr("WorkerThreadPool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("wait_for_task_completion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844576869) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&task_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&task_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *WorkerThreadPool) AddGroupTask(action Callable, elements int64, tasks_needed int64, high_priority bool, description String, ) int64 {
  classNameV := StringNameFromStr("WorkerThreadPool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_group_task")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1801953219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&elements) , gdc.ConstTypePtr(&tasks_needed) , gdc.ConstTypePtr(&high_priority) , description.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&elements)
  pinner.Pin(&tasks_needed)
  pinner.Pin(&high_priority)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WorkerThreadPool) IsGroupTaskCompleted(group_id int64, ) bool {
  classNameV := StringNameFromStr("WorkerThreadPool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_group_task_completed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&group_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WorkerThreadPool) GetGroupProcessedElementCount(group_id int64, ) int64 {
  classNameV := StringNameFromStr("WorkerThreadPool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_group_processed_element_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&group_id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *WorkerThreadPool) WaitForGroupTaskCompletion(group_id int64, )  {
  classNameV := StringNameFromStr("WorkerThreadPool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("wait_for_group_task_completion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
