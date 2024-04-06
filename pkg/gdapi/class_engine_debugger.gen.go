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

type ptrsForEngineDebuggerList struct {
	fnIsActive                 gdc.MethodBindPtr
	fnRegisterProfiler         gdc.MethodBindPtr
	fnUnregisterProfiler       gdc.MethodBindPtr
	fnIsProfiling              gdc.MethodBindPtr
	fnHasProfiler              gdc.MethodBindPtr
	fnProfilerAddFrameData     gdc.MethodBindPtr
	fnProfilerEnable           gdc.MethodBindPtr
	fnRegisterMessageCapture   gdc.MethodBindPtr
	fnUnregisterMessageCapture gdc.MethodBindPtr
	fnHasCapture               gdc.MethodBindPtr
	fnSendMessage              gdc.MethodBindPtr
}

var ptrsForEngineDebugger ptrsForEngineDebuggerList

func initEngineDebuggerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EngineDebugger")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("is_active")
		defer methodName.Destroy()
		ptrsForEngineDebugger.fnIsActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("register_profiler")
		defer methodName.Destroy()
		ptrsForEngineDebugger.fnRegisterProfiler = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3651669560))
	}
	{
		methodName := StringNameFromStr("unregister_profiler")
		defer methodName.Destroy()
		ptrsForEngineDebugger.fnUnregisterProfiler = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("is_profiling")
		defer methodName.Destroy()
		ptrsForEngineDebugger.fnIsProfiling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2041966384))
	}
	{
		methodName := StringNameFromStr("has_profiler")
		defer methodName.Destroy()
		ptrsForEngineDebugger.fnHasProfiler = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2041966384))
	}
	{
		methodName := StringNameFromStr("profiler_add_frame_data")
		defer methodName.Destroy()
		ptrsForEngineDebugger.fnProfilerAddFrameData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1895267858))
	}
	{
		methodName := StringNameFromStr("profiler_enable")
		defer methodName.Destroy()
		ptrsForEngineDebugger.fnProfilerEnable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3192561009))
	}
	{
		methodName := StringNameFromStr("register_message_capture")
		defer methodName.Destroy()
		ptrsForEngineDebugger.fnRegisterMessageCapture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1874754934))
	}
	{
		methodName := StringNameFromStr("unregister_message_capture")
		defer methodName.Destroy()
		ptrsForEngineDebugger.fnUnregisterMessageCapture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("has_capture")
		defer methodName.Destroy()
		ptrsForEngineDebugger.fnHasCapture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2041966384))
	}
	{
		methodName := StringNameFromStr("send_message")
		defer methodName.Destroy()
		ptrsForEngineDebugger.fnSendMessage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1209351045))
	}
}

type EngineDebugger struct {
	Object
}

func (me *EngineDebugger) BaseClass() string {
	return "EngineDebugger"
}

func NewEngineDebugger() *EngineDebugger {
	str := StringNameFromStr("EngineDebugger") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EngineDebugger{}
	obj.SetBaseObject(objPtr)
	return obj
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

func (me *EngineDebugger) IsActive() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngineDebugger.fnIsActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EngineDebugger) RegisterProfiler(name StringName, profiler EngineProfiler) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), profiler.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngineDebugger.fnRegisterProfiler), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EngineDebugger) UnregisterProfiler(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngineDebugger.fnUnregisterProfiler), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EngineDebugger) IsProfiling(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngineDebugger.fnIsProfiling), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EngineDebugger) HasProfiler(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngineDebugger.fnHasProfiler), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EngineDebugger) ProfilerAddFrameData(name StringName, data Array) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngineDebugger.fnProfilerAddFrameData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EngineDebugger) ProfilerEnable(name StringName, enable bool, arguments Array) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&enable), arguments.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngineDebugger.fnProfilerEnable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EngineDebugger) RegisterMessageCapture(name StringName, callable Callable) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), callable.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngineDebugger.fnRegisterMessageCapture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EngineDebugger) UnregisterMessageCapture(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngineDebugger.fnUnregisterMessageCapture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EngineDebugger) HasCapture(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngineDebugger.fnHasCapture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EngineDebugger) SendMessage(message String, data Array) {
	cargs := []gdc.ConstTypePtr{message.AsCTypePtr(), data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEngineDebugger.fnSendMessage), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
