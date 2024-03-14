// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type WorkerThreadPool struct {
  Object
}

func (me *WorkerThreadPool) BaseClass() string {
  return "WorkerThreadPool"
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

func  (me *WorkerThreadPool) AddTask(action Callable, high_priority bool, description String, ) int {
  classNameV := StringNameFromStr("WorkerThreadPool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_task")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3745067146) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), gdc.ConstTypePtr(&high_priority), gdc.ConstTypePtr(description.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WorkerThreadPool) IsTaskCompleted(task_id int, ) bool {
  classNameV := StringNameFromStr("WorkerThreadPool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_task_completed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&task_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WorkerThreadPool) WaitForTaskCompletion(task_id int, ) Error {
  classNameV := StringNameFromStr("WorkerThreadPool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("wait_for_task_completion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844576869) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&task_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WorkerThreadPool) AddGroupTask(action Callable, elements int, tasks_needed int, high_priority bool, description String, ) int {
  classNameV := StringNameFromStr("WorkerThreadPool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_group_task")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1801953219) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), gdc.ConstTypePtr(&elements), gdc.ConstTypePtr(&tasks_needed), gdc.ConstTypePtr(&high_priority), gdc.ConstTypePtr(description.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WorkerThreadPool) IsGroupTaskCompleted(group_id int, ) bool {
  classNameV := StringNameFromStr("WorkerThreadPool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_group_task_completed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WorkerThreadPool) GetGroupProcessedElementCount(group_id int, ) int {
  classNameV := StringNameFromStr("WorkerThreadPool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_group_processed_element_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *WorkerThreadPool) WaitForGroupTaskCompletion(group_id int, )  {
  classNameV := StringNameFromStr("WorkerThreadPool")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("wait_for_group_task_completion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
