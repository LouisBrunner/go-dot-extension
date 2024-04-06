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

type ptrsForSliderList struct {
	fnSetTicks          gdc.MethodBindPtr
	fnGetTicks          gdc.MethodBindPtr
	fnGetTicksOnBorders gdc.MethodBindPtr
	fnSetTicksOnBorders gdc.MethodBindPtr
	fnSetEditable       gdc.MethodBindPtr
	fnIsEditable        gdc.MethodBindPtr
	fnSetScrollable     gdc.MethodBindPtr
	fnIsScrollable      gdc.MethodBindPtr
}

var ptrsForSlider ptrsForSliderList

func initSliderPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Slider")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_ticks")
		defer methodName.Destroy()
		ptrsForSlider.fnSetTicks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_ticks")
		defer methodName.Destroy()
		ptrsForSlider.fnGetTicks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_ticks_on_borders")
		defer methodName.Destroy()
		ptrsForSlider.fnGetTicksOnBorders = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_ticks_on_borders")
		defer methodName.Destroy()
		ptrsForSlider.fnSetTicksOnBorders = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_editable")
		defer methodName.Destroy()
		ptrsForSlider.fnSetEditable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_editable")
		defer methodName.Destroy()
		ptrsForSlider.fnIsEditable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_scrollable")
		defer methodName.Destroy()
		ptrsForSlider.fnSetScrollable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_scrollable")
		defer methodName.Destroy()
		ptrsForSlider.fnIsScrollable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type Slider struct {
	Range
}

func (me *Slider) BaseClass() string {
	return "Slider"
}

func NewSlider() *Slider {
	str := StringNameFromStr("Slider") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Slider{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Slider) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Slider) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Slider) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Slider) SetTicks(count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSlider.fnSetTicks), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Slider) GetTicks() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSlider.fnGetTicks), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Slider) GetTicksOnBorders() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSlider.fnGetTicksOnBorders), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Slider) SetTicksOnBorders(ticks_on_border bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ticks_on_border)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSlider.fnSetTicksOnBorders), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Slider) SetEditable(editable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&editable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSlider.fnSetEditable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Slider) IsEditable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSlider.fnIsEditable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Slider) SetScrollable(scrollable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scrollable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSlider.fnSetScrollable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Slider) IsScrollable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSlider.fnIsScrollable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type SliderDragStartedSignalFn func()

func (me *Slider) ConnectDragStarted(subs SignalSubscribers, fn SliderDragStartedSignalFn) {
	sig := StringNameFromStr("drag_started")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Slider) DisconnectDragStarted(subs SignalSubscribers, fn SliderDragStartedSignalFn) {
	sig := StringNameFromStr("drag_started")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type SliderDragEndedSignalFn func(value_changed bool)

func (me *Slider) ConnectDragEnded(subs SignalSubscribers, fn SliderDragEndedSignalFn) {
	sig := StringNameFromStr("drag_ended")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Slider) DisconnectDragEnded(subs SignalSubscribers, fn SliderDragEndedSignalFn) {
	sig := StringNameFromStr("drag_ended")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
