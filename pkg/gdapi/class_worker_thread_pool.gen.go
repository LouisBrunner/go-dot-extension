// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WorkerThreadPool struct {
  obj gdc.ObjectPtr
}

func (me *WorkerThreadPool) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WorkerThreadPool) BaseClass() string {
  return "WorkerThreadPool"
}



// Enums

func (me *WorkerThreadPool) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WorkerThreadPool) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *WorkerThreadPool) AddTask(action Callable, high_priority bool, description String, )  {
  panic("TODO: implement")
}

func  (me *WorkerThreadPool) IsTaskCompleted(task_id int, )  {
  panic("TODO: implement")
}

func  (me *WorkerThreadPool) WaitForTaskCompletion(task_id int, )  {
  panic("TODO: implement")
}

func  (me *WorkerThreadPool) AddGroupTask(action Callable, elements int, tasks_needed int, high_priority bool, description String, )  {
  panic("TODO: implement")
}

func  (me *WorkerThreadPool) IsGroupTaskCompleted(group_id int, )  {
  panic("TODO: implement")
}

func  (me *WorkerThreadPool) GetGroupProcessedElementCount(group_id int, )  {
  panic("TODO: implement")
}

func  (me *WorkerThreadPool) WaitForGroupTaskCompletion(group_id int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
