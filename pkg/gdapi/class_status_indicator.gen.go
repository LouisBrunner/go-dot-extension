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

type ptrsForStatusIndicatorList struct {
	fnSetTooltip gdc.MethodBindPtr
	fnGetTooltip gdc.MethodBindPtr
	fnSetIcon    gdc.MethodBindPtr
	fnGetIcon    gdc.MethodBindPtr
	fnSetVisible gdc.MethodBindPtr
	fnIsVisible  gdc.MethodBindPtr
}

var ptrsForStatusIndicator ptrsForStatusIndicatorList

func initStatusIndicatorPtrs(iface gdc.Interface) {

	className := StringNameFromStr("StatusIndicator")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_tooltip")
		defer methodName.Destroy()
		ptrsForStatusIndicator.fnSetTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_tooltip")
		defer methodName.Destroy()
		ptrsForStatusIndicator.fnGetTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_icon")
		defer methodName.Destroy()
		ptrsForStatusIndicator.fnSetIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 532598488))
	}
	{
		methodName := StringNameFromStr("get_icon")
		defer methodName.Destroy()
		ptrsForStatusIndicator.fnGetIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4190603485))
	}
	{
		methodName := StringNameFromStr("set_visible")
		defer methodName.Destroy()
		ptrsForStatusIndicator.fnSetVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_visible")
		defer methodName.Destroy()
		ptrsForStatusIndicator.fnIsVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type StatusIndicator struct {
	Node
}

func (me *StatusIndicator) BaseClass() string {
	return "StatusIndicator"
}

func NewStatusIndicator() *StatusIndicator {
	str := StringNameFromStr("StatusIndicator") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &StatusIndicator{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *StatusIndicator) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *StatusIndicator) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *StatusIndicator) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *StatusIndicator) SetTooltip(tooltip String) {
	cargs := []gdc.ConstTypePtr{tooltip.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStatusIndicator.fnSetTooltip), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StatusIndicator) GetTooltip() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStatusIndicator.fnGetTooltip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *StatusIndicator) SetIcon(texture Image) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStatusIndicator.fnSetIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StatusIndicator) GetIcon() Image {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewImage()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStatusIndicator.fnGetIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *StatusIndicator) SetVisible(visible bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStatusIndicator.fnSetVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StatusIndicator) IsVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStatusIndicator.fnIsVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type StatusIndicatorPressedSignalFn func(mouse_button int, mouse_position Vector2i)

func (me *StatusIndicator) ConnectPressed(subs SignalSubscribers, fn StatusIndicatorPressedSignalFn) {
	sig := StringNameFromStr("pressed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *StatusIndicator) DisconnectPressed(subs SignalSubscribers, fn StatusIndicatorPressedSignalFn) {
	sig := StringNameFromStr("pressed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
