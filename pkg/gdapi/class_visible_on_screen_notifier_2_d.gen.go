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

type ptrsForVisibleOnScreenNotifier2DList struct {
	fnSetRect    gdc.MethodBindPtr
	fnGetRect    gdc.MethodBindPtr
	fnIsOnScreen gdc.MethodBindPtr
}

var ptrsForVisibleOnScreenNotifier2D ptrsForVisibleOnScreenNotifier2DList

func initVisibleOnScreenNotifier2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisibleOnScreenNotifier2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_rect")
		defer methodName.Destroy()
		ptrsForVisibleOnScreenNotifier2D.fnSetRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2046264180))
	}
	{
		methodName := StringNameFromStr("get_rect")
		defer methodName.Destroy()
		ptrsForVisibleOnScreenNotifier2D.fnGetRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
	}
	{
		methodName := StringNameFromStr("is_on_screen")
		defer methodName.Destroy()
		ptrsForVisibleOnScreenNotifier2D.fnIsOnScreen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type VisibleOnScreenNotifier2D struct {
	Node2D
}

func (me *VisibleOnScreenNotifier2D) BaseClass() string {
	return "VisibleOnScreenNotifier2D"
}

func NewVisibleOnScreenNotifier2D() *VisibleOnScreenNotifier2D {
	str := StringNameFromStr("VisibleOnScreenNotifier2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisibleOnScreenNotifier2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisibleOnScreenNotifier2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisibleOnScreenNotifier2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisibleOnScreenNotifier2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisibleOnScreenNotifier2D) SetRect(rect Rect2) {
	cargs := []gdc.ConstTypePtr{rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisibleOnScreenNotifier2D.fnSetRect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisibleOnScreenNotifier2D) GetRect() Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisibleOnScreenNotifier2D.fnGetRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisibleOnScreenNotifier2D) IsOnScreen() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisibleOnScreenNotifier2D.fnIsOnScreen), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type VisibleOnScreenNotifier2DScreenEnteredSignalFn func()

func (me *VisibleOnScreenNotifier2D) ConnectScreenEntered(subs SignalSubscribers, fn VisibleOnScreenNotifier2DScreenEnteredSignalFn) {
	sig := StringNameFromStr("screen_entered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *VisibleOnScreenNotifier2D) DisconnectScreenEntered(subs SignalSubscribers, fn VisibleOnScreenNotifier2DScreenEnteredSignalFn) {
	sig := StringNameFromStr("screen_entered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type VisibleOnScreenNotifier2DScreenExitedSignalFn func()

func (me *VisibleOnScreenNotifier2D) ConnectScreenExited(subs SignalSubscribers, fn VisibleOnScreenNotifier2DScreenExitedSignalFn) {
	sig := StringNameFromStr("screen_exited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *VisibleOnScreenNotifier2D) DisconnectScreenExited(subs SignalSubscribers, fn VisibleOnScreenNotifier2DScreenExitedSignalFn) {
	sig := StringNameFromStr("screen_exited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
