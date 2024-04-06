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

type ptrsForTimerList struct {
	fnSetWaitTime             gdc.MethodBindPtr
	fnGetWaitTime             gdc.MethodBindPtr
	fnSetOneShot              gdc.MethodBindPtr
	fnIsOneShot               gdc.MethodBindPtr
	fnSetAutostart            gdc.MethodBindPtr
	fnHasAutostart            gdc.MethodBindPtr
	fnStart                   gdc.MethodBindPtr
	fnStop                    gdc.MethodBindPtr
	fnSetPaused               gdc.MethodBindPtr
	fnIsPaused                gdc.MethodBindPtr
	fnIsStopped               gdc.MethodBindPtr
	fnGetTimeLeft             gdc.MethodBindPtr
	fnSetTimerProcessCallback gdc.MethodBindPtr
	fnGetTimerProcessCallback gdc.MethodBindPtr
}

var ptrsForTimer ptrsForTimerList

func initTimerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Timer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_wait_time")
		defer methodName.Destroy()
		ptrsForTimer.fnSetWaitTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_wait_time")
		defer methodName.Destroy()
		ptrsForTimer.fnGetWaitTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_one_shot")
		defer methodName.Destroy()
		ptrsForTimer.fnSetOneShot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_one_shot")
		defer methodName.Destroy()
		ptrsForTimer.fnIsOneShot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_autostart")
		defer methodName.Destroy()
		ptrsForTimer.fnSetAutostart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("has_autostart")
		defer methodName.Destroy()
		ptrsForTimer.fnHasAutostart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("start")
		defer methodName.Destroy()
		ptrsForTimer.fnStart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1392008558))
	}
	{
		methodName := StringNameFromStr("stop")
		defer methodName.Destroy()
		ptrsForTimer.fnStop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_paused")
		defer methodName.Destroy()
		ptrsForTimer.fnSetPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_paused")
		defer methodName.Destroy()
		ptrsForTimer.fnIsPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_stopped")
		defer methodName.Destroy()
		ptrsForTimer.fnIsStopped = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_time_left")
		defer methodName.Destroy()
		ptrsForTimer.fnGetTimeLeft = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_timer_process_callback")
		defer methodName.Destroy()
		ptrsForTimer.fnSetTimerProcessCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3469495063))
	}
	{
		methodName := StringNameFromStr("get_timer_process_callback")
		defer methodName.Destroy()
		ptrsForTimer.fnGetTimerProcessCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2672570227))
	}
}

type Timer struct {
	Node
}

func (me *Timer) BaseClass() string {
	return "Timer"
}

func NewTimer() *Timer {
	str := StringNameFromStr("Timer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Timer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TimerTimerProcessCallback int

const (
	TimerTimerProcessCallbackTimerProcessPhysics TimerTimerProcessCallback = 0
	TimerTimerProcessCallbackTimerProcessIdle    TimerTimerProcessCallback = 1
)

func (me *Timer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Timer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Timer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Timer) SetWaitTime(time_sec float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_sec)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTimer.fnSetWaitTime), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Timer) GetWaitTime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTimer.fnGetWaitTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Timer) SetOneShot(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTimer.fnSetOneShot), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Timer) IsOneShot() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTimer.fnIsOneShot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Timer) SetAutostart(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTimer.fnSetAutostart), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Timer) HasAutostart() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTimer.fnHasAutostart), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Timer) Start(time_sec float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_sec)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTimer.fnStart), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Timer) Stop() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTimer.fnStop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Timer) SetPaused(paused bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&paused)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTimer.fnSetPaused), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Timer) IsPaused() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTimer.fnIsPaused), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Timer) IsStopped() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTimer.fnIsStopped), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Timer) GetTimeLeft() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTimer.fnGetTimeLeft), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Timer) SetTimerProcessCallback(callback TimerTimerProcessCallback) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&callback)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTimer.fnSetTimerProcessCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Timer) GetTimerProcessCallback() TimerTimerProcessCallback {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TimerTimerProcessCallback

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTimer.fnGetTimerProcessCallback), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TimerTimeoutSignalFn func()

func (me *Timer) ConnectTimeout(subs SignalSubscribers, fn TimerTimeoutSignalFn) {
	sig := StringNameFromStr("timeout")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Timer) DisconnectTimeout(subs SignalSubscribers, fn TimerTimeoutSignalFn) {
	sig := StringNameFromStr("timeout")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
