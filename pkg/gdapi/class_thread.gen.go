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

type ptrsForThreadList struct {
	fnStart                        gdc.MethodBindPtr
	fnGetId                        gdc.MethodBindPtr
	fnIsStarted                    gdc.MethodBindPtr
	fnIsAlive                      gdc.MethodBindPtr
	fnWaitToFinish                 gdc.MethodBindPtr
	fnSetThreadSafetyChecksEnabled gdc.MethodBindPtr
}

var ptrsForThread ptrsForThreadList

func initThreadPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Thread")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("start")
		defer methodName.Destroy()
		ptrsForThread.fnStart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1327203254))
	}
	{
		methodName := StringNameFromStr("get_id")
		defer methodName.Destroy()
		ptrsForThread.fnGetId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("is_started")
		defer methodName.Destroy()
		ptrsForThread.fnIsStarted = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_alive")
		defer methodName.Destroy()
		ptrsForThread.fnIsAlive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("wait_to_finish")
		defer methodName.Destroy()
		ptrsForThread.fnWaitToFinish = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1460262497))
	}
	{
		methodName := StringNameFromStr("set_thread_safety_checks_enabled")
		defer methodName.Destroy()
		ptrsForThread.fnSetThreadSafetyChecksEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
}

type Thread struct {
	RefCounted
}

func (me *Thread) BaseClass() string {
	return "Thread"
}

func NewThread() *Thread {
	str := StringNameFromStr("Thread") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Thread{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type ThreadPriority int

const (
	ThreadPriorityPriorityLow    ThreadPriority = 0
	ThreadPriorityPriorityNormal ThreadPriority = 1
	ThreadPriorityPriorityHigh   ThreadPriority = 2
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

func (me *Thread) Start(callable Callable, priority ThreadPriority) Error {
	cargs := []gdc.ConstTypePtr{callable.AsCTypePtr(), gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&priority)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThread.fnStart), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Thread) GetId() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThread.fnGetId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Thread) IsStarted() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThread.fnIsStarted), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Thread) IsAlive() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThread.fnIsAlive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Thread) WaitToFinish() Variant {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThread.fnWaitToFinish), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func ThreadSetThreadSafetyChecksEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForThread.fnSetThreadSafetyChecksEnabled), nil, unsafe.SliceData(cargs), nil)

}

// Signals
