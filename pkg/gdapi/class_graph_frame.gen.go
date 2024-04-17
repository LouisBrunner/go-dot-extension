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

type ptrsForGraphFrameList struct {
	fnSetTitle             gdc.MethodBindPtr
	fnGetTitle             gdc.MethodBindPtr
	fnGetTitlebarHbox      gdc.MethodBindPtr
	fnSetAutoshrinkEnabled gdc.MethodBindPtr
	fnIsAutoshrinkEnabled  gdc.MethodBindPtr
	fnSetAutoshrinkMargin  gdc.MethodBindPtr
	fnGetAutoshrinkMargin  gdc.MethodBindPtr
	fnSetDragMargin        gdc.MethodBindPtr
	fnGetDragMargin        gdc.MethodBindPtr
	fnSetTintColorEnabled  gdc.MethodBindPtr
	fnIsTintColorEnabled   gdc.MethodBindPtr
	fnSetTintColor         gdc.MethodBindPtr
	fnGetTintColor         gdc.MethodBindPtr
}

var ptrsForGraphFrame ptrsForGraphFrameList

func initGraphFramePtrs(iface gdc.Interface) {

	className := StringNameFromStr("GraphFrame")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_title")
		defer methodName.Destroy()
		ptrsForGraphFrame.fnSetTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_title")
		defer methodName.Destroy()
		ptrsForGraphFrame.fnGetTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_titlebar_hbox")
		defer methodName.Destroy()
		ptrsForGraphFrame.fnGetTitlebarHbox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3590609951))
	}
	{
		methodName := StringNameFromStr("set_autoshrink_enabled")
		defer methodName.Destroy()
		ptrsForGraphFrame.fnSetAutoshrinkEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_autoshrink_enabled")
		defer methodName.Destroy()
		ptrsForGraphFrame.fnIsAutoshrinkEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_autoshrink_margin")
		defer methodName.Destroy()
		ptrsForGraphFrame.fnSetAutoshrinkMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_autoshrink_margin")
		defer methodName.Destroy()
		ptrsForGraphFrame.fnGetAutoshrinkMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_drag_margin")
		defer methodName.Destroy()
		ptrsForGraphFrame.fnSetDragMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_drag_margin")
		defer methodName.Destroy()
		ptrsForGraphFrame.fnGetDragMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_tint_color_enabled")
		defer methodName.Destroy()
		ptrsForGraphFrame.fnSetTintColorEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_tint_color_enabled")
		defer methodName.Destroy()
		ptrsForGraphFrame.fnIsTintColorEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_tint_color")
		defer methodName.Destroy()
		ptrsForGraphFrame.fnSetTintColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_tint_color")
		defer methodName.Destroy()
		ptrsForGraphFrame.fnGetTintColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}

}

type GraphFrame struct {
	GraphElement
}

func (me *GraphFrame) BaseClass() string {
	return "GraphFrame"
}

func NewGraphFrame() *GraphFrame {
	str := StringNameFromStr("GraphFrame") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GraphFrame{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GraphFrame) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GraphFrame) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GraphFrame) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GraphFrame) SetTitle(title String) {
	cargs := []gdc.ConstTypePtr{title.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphFrame.fnSetTitle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphFrame) GetTitle() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphFrame.fnGetTitle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GraphFrame) GetTitlebarHbox() HBoxContainer {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewHBoxContainer()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphFrame.fnGetTitlebarHbox), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GraphFrame) SetAutoshrinkEnabled(shrink bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shrink)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphFrame.fnSetAutoshrinkEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphFrame) IsAutoshrinkEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphFrame.fnIsAutoshrinkEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphFrame) SetAutoshrinkMargin(autoshrink_margin int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&autoshrink_margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphFrame.fnSetAutoshrinkMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphFrame) GetAutoshrinkMargin() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphFrame.fnGetAutoshrinkMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphFrame) SetDragMargin(drag_margin int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&drag_margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphFrame.fnSetDragMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphFrame) GetDragMargin() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphFrame.fnGetDragMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphFrame) SetTintColorEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphFrame.fnSetTintColorEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphFrame) IsTintColorEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphFrame.fnIsTintColorEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphFrame) SetTintColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphFrame.fnSetTintColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphFrame) GetTintColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphFrame.fnGetTintColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type GraphFrameAutoshrinkChangedSignalFn func()

func (me *GraphFrame) ConnectAutoshrinkChanged(subs SignalSubscribers, fn GraphFrameAutoshrinkChangedSignalFn) {
	sig := StringNameFromStr("autoshrink_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphFrame) DisconnectAutoshrinkChanged(subs SignalSubscribers, fn GraphFrameAutoshrinkChangedSignalFn) {
	sig := StringNameFromStr("autoshrink_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
