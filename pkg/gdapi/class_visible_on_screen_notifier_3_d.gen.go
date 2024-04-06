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

type ptrsForVisibleOnScreenNotifier3DList struct {
	fnSetAabb    gdc.MethodBindPtr
	fnIsOnScreen gdc.MethodBindPtr
}

var ptrsForVisibleOnScreenNotifier3D ptrsForVisibleOnScreenNotifier3DList

func initVisibleOnScreenNotifier3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisibleOnScreenNotifier3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_aabb")
		defer methodName.Destroy()
		ptrsForVisibleOnScreenNotifier3D.fnSetAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 259215842))
	}
	{
		methodName := StringNameFromStr("is_on_screen")
		defer methodName.Destroy()
		ptrsForVisibleOnScreenNotifier3D.fnIsOnScreen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type VisibleOnScreenNotifier3D struct {
	VisualInstance3D
}

func (me *VisibleOnScreenNotifier3D) BaseClass() string {
	return "VisibleOnScreenNotifier3D"
}

func NewVisibleOnScreenNotifier3D() *VisibleOnScreenNotifier3D {
	str := StringNameFromStr("VisibleOnScreenNotifier3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisibleOnScreenNotifier3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisibleOnScreenNotifier3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisibleOnScreenNotifier3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisibleOnScreenNotifier3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisibleOnScreenNotifier3D) SetAabb(rect AABB) {
	cargs := []gdc.ConstTypePtr{rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisibleOnScreenNotifier3D.fnSetAabb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisibleOnScreenNotifier3D) IsOnScreen() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisibleOnScreenNotifier3D.fnIsOnScreen), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type VisibleOnScreenNotifier3DScreenEnteredSignalFn func()

func (me *VisibleOnScreenNotifier3D) ConnectScreenEntered(subs SignalSubscribers, fn VisibleOnScreenNotifier3DScreenEnteredSignalFn) {
	sig := StringNameFromStr("screen_entered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *VisibleOnScreenNotifier3D) DisconnectScreenEntered(subs SignalSubscribers, fn VisibleOnScreenNotifier3DScreenEnteredSignalFn) {
	sig := StringNameFromStr("screen_entered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type VisibleOnScreenNotifier3DScreenExitedSignalFn func()

func (me *VisibleOnScreenNotifier3D) ConnectScreenExited(subs SignalSubscribers, fn VisibleOnScreenNotifier3DScreenExitedSignalFn) {
	sig := StringNameFromStr("screen_exited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *VisibleOnScreenNotifier3D) DisconnectScreenExited(subs SignalSubscribers, fn VisibleOnScreenNotifier3DScreenExitedSignalFn) {
	sig := StringNameFromStr("screen_exited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
