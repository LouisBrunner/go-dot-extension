// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"log"
	"runtime"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForWorkerThreadPoolList struct {
	fnAddTask                       gdc.MethodBindPtr
	fnIsTaskCompleted               gdc.MethodBindPtr
	fnWaitForTaskCompletion         gdc.MethodBindPtr
	fnAddGroupTask                  gdc.MethodBindPtr
	fnIsGroupTaskCompleted          gdc.MethodBindPtr
	fnGetGroupProcessedElementCount gdc.MethodBindPtr
	fnWaitForGroupTaskCompletion    gdc.MethodBindPtr
}

var ptrsForWorkerThreadPool ptrsForWorkerThreadPoolList

func initWorkerThreadPoolPtrs(iface gdc.Interface) {

	className := StringNameFromStr("WorkerThreadPool")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_task")
		defer methodName.Destroy()
		ptrsForWorkerThreadPool.fnAddTask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3745067146))
	}
	{
		methodName := StringNameFromStr("is_task_completed")
		defer methodName.Destroy()
		ptrsForWorkerThreadPool.fnIsTaskCompleted = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("wait_for_task_completion")
		defer methodName.Destroy()
		ptrsForWorkerThreadPool.fnWaitForTaskCompletion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844576869))
	}
	{
		methodName := StringNameFromStr("add_group_task")
		defer methodName.Destroy()
		ptrsForWorkerThreadPool.fnAddGroupTask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1801953219))
	}
	{
		methodName := StringNameFromStr("is_group_task_completed")
		defer methodName.Destroy()
		ptrsForWorkerThreadPool.fnIsGroupTaskCompleted = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_group_processed_element_count")
		defer methodName.Destroy()
		ptrsForWorkerThreadPool.fnGetGroupProcessedElementCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("wait_for_group_task_completion")
		defer methodName.Destroy()
		ptrsForWorkerThreadPool.fnWaitForGroupTaskCompletion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}

}

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

func (me *WorkerThreadPool) AddTask(action Callable, high_priority bool, description String) int64 {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&high_priority), description.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&high_priority)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorkerThreadPool.fnAddTask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *WorkerThreadPool) IsTaskCompleted(task_id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&task_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&task_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorkerThreadPool.fnIsTaskCompleted), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *WorkerThreadPool) WaitForTaskCompletion(task_id int64) Error {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&task_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&task_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorkerThreadPool.fnWaitForTaskCompletion), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *WorkerThreadPool) AddGroupTask(action Callable, elements int64, tasks_needed int64, high_priority bool, description String) int64 {
	cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&elements), gdc.ConstTypePtr(&tasks_needed), gdc.ConstTypePtr(&high_priority), description.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&elements)
	pinner.Pin(&tasks_needed)
	pinner.Pin(&high_priority)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorkerThreadPool.fnAddGroupTask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *WorkerThreadPool) IsGroupTaskCompleted(group_id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&group_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorkerThreadPool.fnIsGroupTaskCompleted), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *WorkerThreadPool) GetGroupProcessedElementCount(group_id int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&group_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorkerThreadPool.fnGetGroupProcessedElementCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *WorkerThreadPool) WaitForGroupTaskCompletion(group_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWorkerThreadPool.fnWaitForGroupTaskCompletion), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
