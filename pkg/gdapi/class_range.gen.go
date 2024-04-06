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

type ptrsForRangeList struct {
	fnXValueChanged        gdc.MethodBindPtr
	fnGetValue             gdc.MethodBindPtr
	fnGetMin               gdc.MethodBindPtr
	fnGetMax               gdc.MethodBindPtr
	fnGetStep              gdc.MethodBindPtr
	fnGetPage              gdc.MethodBindPtr
	fnGetAsRatio           gdc.MethodBindPtr
	fnSetValue             gdc.MethodBindPtr
	fnSetValueNoSignal     gdc.MethodBindPtr
	fnSetMin               gdc.MethodBindPtr
	fnSetMax               gdc.MethodBindPtr
	fnSetStep              gdc.MethodBindPtr
	fnSetPage              gdc.MethodBindPtr
	fnSetAsRatio           gdc.MethodBindPtr
	fnSetUseRoundedValues  gdc.MethodBindPtr
	fnIsUsingRoundedValues gdc.MethodBindPtr
	fnSetExpRatio          gdc.MethodBindPtr
	fnIsRatioExp           gdc.MethodBindPtr
	fnSetAllowGreater      gdc.MethodBindPtr
	fnIsGreaterAllowed     gdc.MethodBindPtr
	fnSetAllowLesser       gdc.MethodBindPtr
	fnIsLesserAllowed      gdc.MethodBindPtr
	fnShare                gdc.MethodBindPtr
	fnUnshare              gdc.MethodBindPtr
}

var ptrsForRange ptrsForRangeList

func initRangePtrs(iface gdc.Interface) {

	className := StringNameFromStr("Range")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_value")
		defer methodName.Destroy()
		ptrsForRange.fnGetValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_min")
		defer methodName.Destroy()
		ptrsForRange.fnGetMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_max")
		defer methodName.Destroy()
		ptrsForRange.fnGetMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_step")
		defer methodName.Destroy()
		ptrsForRange.fnGetStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_page")
		defer methodName.Destroy()
		ptrsForRange.fnGetPage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_as_ratio")
		defer methodName.Destroy()
		ptrsForRange.fnGetAsRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_value")
		defer methodName.Destroy()
		ptrsForRange.fnSetValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_value_no_signal")
		defer methodName.Destroy()
		ptrsForRange.fnSetValueNoSignal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_min")
		defer methodName.Destroy()
		ptrsForRange.fnSetMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_max")
		defer methodName.Destroy()
		ptrsForRange.fnSetMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_step")
		defer methodName.Destroy()
		ptrsForRange.fnSetStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_page")
		defer methodName.Destroy()
		ptrsForRange.fnSetPage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_as_ratio")
		defer methodName.Destroy()
		ptrsForRange.fnSetAsRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_use_rounded_values")
		defer methodName.Destroy()
		ptrsForRange.fnSetUseRoundedValues = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_rounded_values")
		defer methodName.Destroy()
		ptrsForRange.fnIsUsingRoundedValues = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_exp_ratio")
		defer methodName.Destroy()
		ptrsForRange.fnSetExpRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_ratio_exp")
		defer methodName.Destroy()
		ptrsForRange.fnIsRatioExp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_allow_greater")
		defer methodName.Destroy()
		ptrsForRange.fnSetAllowGreater = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_greater_allowed")
		defer methodName.Destroy()
		ptrsForRange.fnIsGreaterAllowed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_allow_lesser")
		defer methodName.Destroy()
		ptrsForRange.fnSetAllowLesser = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_lesser_allowed")
		defer methodName.Destroy()
		ptrsForRange.fnIsLesserAllowed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("share")
		defer methodName.Destroy()
		ptrsForRange.fnShare = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
	{
		methodName := StringNameFromStr("unshare")
		defer methodName.Destroy()
		ptrsForRange.fnUnshare = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
}

type Range struct {
	Control
}

func (me *Range) BaseClass() string {
	return "Range"
}

func NewRange() *Range {
	str := StringNameFromStr("Range") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Range{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Range) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Range) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Range) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Range) GetValue() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnGetValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Range) GetMin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnGetMin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Range) GetMax() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnGetMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Range) GetStep() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnGetStep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Range) GetPage() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnGetPage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Range) GetAsRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnGetAsRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Range) SetValue(value float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnSetValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Range) SetValueNoSignal(value float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnSetValueNoSignal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Range) SetMin(minimum float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&minimum)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnSetMin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Range) SetMax(maximum float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&maximum)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnSetMax), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Range) SetStep(step float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&step)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnSetStep), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Range) SetPage(pagesize float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pagesize)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnSetPage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Range) SetAsRatio(value float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnSetAsRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Range) SetUseRoundedValues(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnSetUseRoundedValues), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Range) IsUsingRoundedValues() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnIsUsingRoundedValues), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Range) SetExpRatio(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnSetExpRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Range) IsRatioExp() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnIsRatioExp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Range) SetAllowGreater(allow bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnSetAllowGreater), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Range) IsGreaterAllowed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnIsGreaterAllowed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Range) SetAllowLesser(allow bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnSetAllowLesser), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Range) IsLesserAllowed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnIsLesserAllowed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Range) Share(with Node) {
	cargs := []gdc.ConstTypePtr{with.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnShare), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Range) Unshare() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRange.fnUnshare), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type RangeValueChangedSignalFn func(value float32)

func (me *Range) ConnectValueChanged(subs SignalSubscribers, fn RangeValueChangedSignalFn) {
	sig := StringNameFromStr("value_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Range) DisconnectValueChanged(subs SignalSubscribers, fn RangeValueChangedSignalFn) {
	sig := StringNameFromStr("value_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type RangeChangedSignalFn func()

func (me *Range) ConnectChanged(subs SignalSubscribers, fn RangeChangedSignalFn) {
	sig := StringNameFromStr("changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Range) DisconnectChanged(subs SignalSubscribers, fn RangeChangedSignalFn) {
	sig := StringNameFromStr("changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
