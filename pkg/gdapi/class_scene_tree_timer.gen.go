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

type ptrsForSceneTreeTimerList struct {
	fnSetTimeLeft gdc.MethodBindPtr
	fnGetTimeLeft gdc.MethodBindPtr
}

var ptrsForSceneTreeTimer ptrsForSceneTreeTimerList

func initSceneTreeTimerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SceneTreeTimer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_time_left")
		defer methodName.Destroy()
		ptrsForSceneTreeTimer.fnSetTimeLeft = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_time_left")
		defer methodName.Destroy()
		ptrsForSceneTreeTimer.fnGetTimeLeft = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type SceneTreeTimer struct {
	RefCounted
}

func (me *SceneTreeTimer) BaseClass() string {
	return "SceneTreeTimer"
}

func NewSceneTreeTimer() *SceneTreeTimer {
	str := StringNameFromStr("SceneTreeTimer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SceneTreeTimer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *SceneTreeTimer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SceneTreeTimer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SceneTreeTimer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SceneTreeTimer) SetTimeLeft(time float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTreeTimer.fnSetTimeLeft), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SceneTreeTimer) GetTimeLeft() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSceneTreeTimer.fnGetTimeLeft), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type SceneTreeTimerTimeoutSignalFn func()

func (me *SceneTreeTimer) ConnectTimeout(subs SignalSubscribers, fn SceneTreeTimerTimeoutSignalFn) {
	sig := StringNameFromStr("timeout")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *SceneTreeTimer) DisconnectTimeout(subs SignalSubscribers, fn SceneTreeTimerTimeoutSignalFn) {
	sig := StringNameFromStr("timeout")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
