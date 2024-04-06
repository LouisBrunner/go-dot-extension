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

type ptrsForGraphElementList struct {
	fnSetResizable      gdc.MethodBindPtr
	fnIsResizable       gdc.MethodBindPtr
	fnSetDraggable      gdc.MethodBindPtr
	fnIsDraggable       gdc.MethodBindPtr
	fnSetSelectable     gdc.MethodBindPtr
	fnIsSelectable      gdc.MethodBindPtr
	fnSetSelected       gdc.MethodBindPtr
	fnIsSelected        gdc.MethodBindPtr
	fnSetPositionOffset gdc.MethodBindPtr
	fnGetPositionOffset gdc.MethodBindPtr
}

var ptrsForGraphElement ptrsForGraphElementList

func initGraphElementPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GraphElement")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_resizable")
		defer methodName.Destroy()
		ptrsForGraphElement.fnSetResizable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_resizable")
		defer methodName.Destroy()
		ptrsForGraphElement.fnIsResizable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_draggable")
		defer methodName.Destroy()
		ptrsForGraphElement.fnSetDraggable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_draggable")
		defer methodName.Destroy()
		ptrsForGraphElement.fnIsDraggable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_selectable")
		defer methodName.Destroy()
		ptrsForGraphElement.fnSetSelectable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_selectable")
		defer methodName.Destroy()
		ptrsForGraphElement.fnIsSelectable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_selected")
		defer methodName.Destroy()
		ptrsForGraphElement.fnSetSelected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_selected")
		defer methodName.Destroy()
		ptrsForGraphElement.fnIsSelected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_position_offset")
		defer methodName.Destroy()
		ptrsForGraphElement.fnSetPositionOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_position_offset")
		defer methodName.Destroy()
		ptrsForGraphElement.fnGetPositionOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
}

type GraphElement struct {
	Container
}

func (me *GraphElement) BaseClass() string {
	return "GraphElement"
}

func NewGraphElement() *GraphElement {
	str := StringNameFromStr("GraphElement") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GraphElement{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GraphElement) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GraphElement) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GraphElement) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GraphElement) SetResizable(resizable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&resizable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphElement.fnSetResizable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphElement) IsResizable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphElement.fnIsResizable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphElement) SetDraggable(draggable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draggable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphElement.fnSetDraggable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphElement) IsDraggable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphElement.fnIsDraggable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphElement) SetSelectable(selectable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&selectable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphElement.fnSetSelectable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphElement) IsSelectable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphElement.fnIsSelectable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphElement) SetSelected(selected bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&selected)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphElement.fnSetSelected), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphElement) IsSelected() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphElement.fnIsSelected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphElement) SetPositionOffset(offset Vector2) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphElement.fnSetPositionOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphElement) GetPositionOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphElement.fnGetPositionOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type GraphElementNodeSelectedSignalFn func()

func (me *GraphElement) ConnectNodeSelected(subs SignalSubscribers, fn GraphElementNodeSelectedSignalFn) {
	sig := StringNameFromStr("node_selected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphElement) DisconnectNodeSelected(subs SignalSubscribers, fn GraphElementNodeSelectedSignalFn) {
	sig := StringNameFromStr("node_selected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphElementNodeDeselectedSignalFn func()

func (me *GraphElement) ConnectNodeDeselected(subs SignalSubscribers, fn GraphElementNodeDeselectedSignalFn) {
	sig := StringNameFromStr("node_deselected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphElement) DisconnectNodeDeselected(subs SignalSubscribers, fn GraphElementNodeDeselectedSignalFn) {
	sig := StringNameFromStr("node_deselected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphElementRaiseRequestSignalFn func()

func (me *GraphElement) ConnectRaiseRequest(subs SignalSubscribers, fn GraphElementRaiseRequestSignalFn) {
	sig := StringNameFromStr("raise_request")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphElement) DisconnectRaiseRequest(subs SignalSubscribers, fn GraphElementRaiseRequestSignalFn) {
	sig := StringNameFromStr("raise_request")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphElementDeleteRequestSignalFn func()

func (me *GraphElement) ConnectDeleteRequest(subs SignalSubscribers, fn GraphElementDeleteRequestSignalFn) {
	sig := StringNameFromStr("delete_request")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphElement) DisconnectDeleteRequest(subs SignalSubscribers, fn GraphElementDeleteRequestSignalFn) {
	sig := StringNameFromStr("delete_request")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphElementResizeRequestSignalFn func(new_minsize Vector2)

func (me *GraphElement) ConnectResizeRequest(subs SignalSubscribers, fn GraphElementResizeRequestSignalFn) {
	sig := StringNameFromStr("resize_request")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphElement) DisconnectResizeRequest(subs SignalSubscribers, fn GraphElementResizeRequestSignalFn) {
	sig := StringNameFromStr("resize_request")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphElementDraggedSignalFn func(from Vector2, to Vector2)

func (me *GraphElement) ConnectDragged(subs SignalSubscribers, fn GraphElementDraggedSignalFn) {
	sig := StringNameFromStr("dragged")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphElement) DisconnectDragged(subs SignalSubscribers, fn GraphElementDraggedSignalFn) {
	sig := StringNameFromStr("dragged")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type GraphElementPositionOffsetChangedSignalFn func()

func (me *GraphElement) ConnectPositionOffsetChanged(subs SignalSubscribers, fn GraphElementPositionOffsetChangedSignalFn) {
	sig := StringNameFromStr("position_offset_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphElement) DisconnectPositionOffsetChanged(subs SignalSubscribers, fn GraphElementPositionOffsetChangedSignalFn) {
	sig := StringNameFromStr("position_offset_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
