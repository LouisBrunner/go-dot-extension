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

type ptrsForScrollBarList struct {
	fnSetCustomStep gdc.MethodBindPtr
	fnGetCustomStep gdc.MethodBindPtr
}

var ptrsForScrollBar ptrsForScrollBarList

func initScrollBarPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ScrollBar")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_custom_step")
		defer methodName.Destroy()
		ptrsForScrollBar.fnSetCustomStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_custom_step")
		defer methodName.Destroy()
		ptrsForScrollBar.fnGetCustomStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type ScrollBar struct {
	Range
}

func (me *ScrollBar) BaseClass() string {
	return "ScrollBar"
}

func NewScrollBar() *ScrollBar {
	str := StringNameFromStr("ScrollBar") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ScrollBar{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ScrollBar) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ScrollBar) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ScrollBar) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ScrollBar) SetCustomStep(step float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&step)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollBar.fnSetCustomStep), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ScrollBar) GetCustomStep() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForScrollBar.fnGetCustomStep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ScrollBarScrollingSignalFn func()

func (me *ScrollBar) ConnectScrolling(subs SignalSubscribers, fn ScrollBarScrollingSignalFn) {
	sig := StringNameFromStr("scrolling")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *ScrollBar) DisconnectScrolling(subs SignalSubscribers, fn ScrollBarScrollingSignalFn) {
	sig := StringNameFromStr("scrolling")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
