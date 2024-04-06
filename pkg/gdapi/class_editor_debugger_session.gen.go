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

type ptrsForEditorDebuggerSessionList struct {
	fnSendMessage      gdc.MethodBindPtr
	fnToggleProfiler   gdc.MethodBindPtr
	fnIsBreaked        gdc.MethodBindPtr
	fnIsDebuggable     gdc.MethodBindPtr
	fnIsActive         gdc.MethodBindPtr
	fnAddSessionTab    gdc.MethodBindPtr
	fnRemoveSessionTab gdc.MethodBindPtr
}

var ptrsForEditorDebuggerSession ptrsForEditorDebuggerSessionList

func initEditorDebuggerSessionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorDebuggerSession")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("send_message")
		defer methodName.Destroy()
		ptrsForEditorDebuggerSession.fnSendMessage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 85656714))
	}
	{
		methodName := StringNameFromStr("toggle_profiler")
		defer methodName.Destroy()
		ptrsForEditorDebuggerSession.fnToggleProfiler = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1198443697))
	}
	{
		methodName := StringNameFromStr("is_breaked")
		defer methodName.Destroy()
		ptrsForEditorDebuggerSession.fnIsBreaked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("is_debuggable")
		defer methodName.Destroy()
		ptrsForEditorDebuggerSession.fnIsDebuggable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("is_active")
		defer methodName.Destroy()
		ptrsForEditorDebuggerSession.fnIsActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("add_session_tab")
		defer methodName.Destroy()
		ptrsForEditorDebuggerSession.fnAddSessionTab = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1496901182))
	}
	{
		methodName := StringNameFromStr("remove_session_tab")
		defer methodName.Destroy()
		ptrsForEditorDebuggerSession.fnRemoveSessionTab = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1496901182))
	}

}

type EditorDebuggerSession struct {
	RefCounted
}

func (me *EditorDebuggerSession) BaseClass() string {
	return "EditorDebuggerSession"
}

func NewEditorDebuggerSession() *EditorDebuggerSession {
	str := StringNameFromStr("EditorDebuggerSession") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorDebuggerSession{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorDebuggerSession) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorDebuggerSession) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorDebuggerSession) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EditorDebuggerSession) SendMessage(message String, data Array) {
	cargs := []gdc.ConstTypePtr{message.AsCTypePtr(), data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorDebuggerSession.fnSendMessage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorDebuggerSession) ToggleProfiler(profiler String, enable bool, data Array) {
	cargs := []gdc.ConstTypePtr{profiler.AsCTypePtr(), gdc.ConstTypePtr(&enable), data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorDebuggerSession.fnToggleProfiler), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorDebuggerSession) IsBreaked() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorDebuggerSession.fnIsBreaked), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorDebuggerSession) IsDebuggable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorDebuggerSession.fnIsDebuggable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorDebuggerSession) IsActive() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorDebuggerSession.fnIsActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorDebuggerSession) AddSessionTab(control Control) {
	cargs := []gdc.ConstTypePtr{control.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorDebuggerSession.fnAddSessionTab), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorDebuggerSession) RemoveSessionTab(control Control) {
	cargs := []gdc.ConstTypePtr{control.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorDebuggerSession.fnRemoveSessionTab), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals

type EditorDebuggerSessionStartedSignalFn func()

func (me *EditorDebuggerSession) ConnectStarted(subs SignalSubscribers, fn EditorDebuggerSessionStartedSignalFn) {
	sig := StringNameFromStr("started")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorDebuggerSession) DisconnectStarted(subs SignalSubscribers, fn EditorDebuggerSessionStartedSignalFn) {
	sig := StringNameFromStr("started")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type EditorDebuggerSessionStoppedSignalFn func()

func (me *EditorDebuggerSession) ConnectStopped(subs SignalSubscribers, fn EditorDebuggerSessionStoppedSignalFn) {
	sig := StringNameFromStr("stopped")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorDebuggerSession) DisconnectStopped(subs SignalSubscribers, fn EditorDebuggerSessionStoppedSignalFn) {
	sig := StringNameFromStr("stopped")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type EditorDebuggerSessionBreakedSignalFn func(can_debug bool)

func (me *EditorDebuggerSession) ConnectBreaked(subs SignalSubscribers, fn EditorDebuggerSessionBreakedSignalFn) {
	sig := StringNameFromStr("breaked")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorDebuggerSession) DisconnectBreaked(subs SignalSubscribers, fn EditorDebuggerSessionBreakedSignalFn) {
	sig := StringNameFromStr("breaked")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type EditorDebuggerSessionContinuedSignalFn func()

func (me *EditorDebuggerSession) ConnectContinued(subs SignalSubscribers, fn EditorDebuggerSessionContinuedSignalFn) {
	sig := StringNameFromStr("continued")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorDebuggerSession) DisconnectContinued(subs SignalSubscribers, fn EditorDebuggerSessionContinuedSignalFn) {
	sig := StringNameFromStr("continued")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
