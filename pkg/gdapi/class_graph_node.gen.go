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

type ptrsForGraphNodeList struct {
	fnXDrawPort             gdc.MethodBindPtr
	fnSetTitle              gdc.MethodBindPtr
	fnGetTitle              gdc.MethodBindPtr
	fnGetTitlebarHbox       gdc.MethodBindPtr
	fnSetSlot               gdc.MethodBindPtr
	fnClearSlot             gdc.MethodBindPtr
	fnClearAllSlots         gdc.MethodBindPtr
	fnIsSlotEnabledLeft     gdc.MethodBindPtr
	fnSetSlotEnabledLeft    gdc.MethodBindPtr
	fnSetSlotTypeLeft       gdc.MethodBindPtr
	fnGetSlotTypeLeft       gdc.MethodBindPtr
	fnSetSlotColorLeft      gdc.MethodBindPtr
	fnGetSlotColorLeft      gdc.MethodBindPtr
	fnIsSlotEnabledRight    gdc.MethodBindPtr
	fnSetSlotEnabledRight   gdc.MethodBindPtr
	fnSetSlotTypeRight      gdc.MethodBindPtr
	fnGetSlotTypeRight      gdc.MethodBindPtr
	fnSetSlotColorRight     gdc.MethodBindPtr
	fnGetSlotColorRight     gdc.MethodBindPtr
	fnIsSlotDrawStylebox    gdc.MethodBindPtr
	fnSetSlotDrawStylebox   gdc.MethodBindPtr
	fnGetInputPortCount     gdc.MethodBindPtr
	fnGetInputPortPosition  gdc.MethodBindPtr
	fnGetInputPortType      gdc.MethodBindPtr
	fnGetInputPortColor     gdc.MethodBindPtr
	fnGetInputPortSlot      gdc.MethodBindPtr
	fnGetOutputPortCount    gdc.MethodBindPtr
	fnGetOutputPortPosition gdc.MethodBindPtr
	fnGetOutputPortType     gdc.MethodBindPtr
	fnGetOutputPortColor    gdc.MethodBindPtr
	fnGetOutputPortSlot     gdc.MethodBindPtr
}

var ptrsForGraphNode ptrsForGraphNodeList

func initGraphNodePtrs(iface gdc.Interface) {

	className := StringNameFromStr("GraphNode")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_title")
		defer methodName.Destroy()
		ptrsForGraphNode.fnSetTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_title")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_titlebar_hbox")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetTitlebarHbox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3590609951))
	}
	{
		methodName := StringNameFromStr("set_slot")
		defer methodName.Destroy()
		ptrsForGraphNode.fnSetSlot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2873310869))
	}
	{
		methodName := StringNameFromStr("clear_slot")
		defer methodName.Destroy()
		ptrsForGraphNode.fnClearSlot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("clear_all_slots")
		defer methodName.Destroy()
		ptrsForGraphNode.fnClearAllSlots = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("is_slot_enabled_left")
		defer methodName.Destroy()
		ptrsForGraphNode.fnIsSlotEnabledLeft = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_slot_enabled_left")
		defer methodName.Destroy()
		ptrsForGraphNode.fnSetSlotEnabledLeft = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("set_slot_type_left")
		defer methodName.Destroy()
		ptrsForGraphNode.fnSetSlotTypeLeft = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_slot_type_left")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetSlotTypeLeft = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_slot_color_left")
		defer methodName.Destroy()
		ptrsForGraphNode.fnSetSlotColorLeft = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
	}
	{
		methodName := StringNameFromStr("get_slot_color_left")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetSlotColorLeft = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457211756))
	}
	{
		methodName := StringNameFromStr("is_slot_enabled_right")
		defer methodName.Destroy()
		ptrsForGraphNode.fnIsSlotEnabledRight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_slot_enabled_right")
		defer methodName.Destroy()
		ptrsForGraphNode.fnSetSlotEnabledRight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("set_slot_type_right")
		defer methodName.Destroy()
		ptrsForGraphNode.fnSetSlotTypeRight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_slot_type_right")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetSlotTypeRight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_slot_color_right")
		defer methodName.Destroy()
		ptrsForGraphNode.fnSetSlotColorRight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
	}
	{
		methodName := StringNameFromStr("get_slot_color_right")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetSlotColorRight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457211756))
	}
	{
		methodName := StringNameFromStr("is_slot_draw_stylebox")
		defer methodName.Destroy()
		ptrsForGraphNode.fnIsSlotDrawStylebox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_slot_draw_stylebox")
		defer methodName.Destroy()
		ptrsForGraphNode.fnSetSlotDrawStylebox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_input_port_count")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetInputPortCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("get_input_port_position")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetInputPortPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3114997196))
	}
	{
		methodName := StringNameFromStr("get_input_port_type")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetInputPortType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3744713108))
	}
	{
		methodName := StringNameFromStr("get_input_port_color")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetInputPortColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2624840992))
	}
	{
		methodName := StringNameFromStr("get_input_port_slot")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetInputPortSlot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3744713108))
	}
	{
		methodName := StringNameFromStr("get_output_port_count")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetOutputPortCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("get_output_port_position")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetOutputPortPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3114997196))
	}
	{
		methodName := StringNameFromStr("get_output_port_type")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetOutputPortType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3744713108))
	}
	{
		methodName := StringNameFromStr("get_output_port_color")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetOutputPortColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2624840992))
	}
	{
		methodName := StringNameFromStr("get_output_port_slot")
		defer methodName.Destroy()
		ptrsForGraphNode.fnGetOutputPortSlot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3744713108))
	}

}

type GraphNode struct {
	GraphElement
}

func (me *GraphNode) BaseClass() string {
	return "GraphNode"
}

func NewGraphNode() *GraphNode {
	str := StringNameFromStr("GraphNode") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GraphNode{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GraphNode) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GraphNode) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GraphNode) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GraphNode) SetTitle(title String) {
	cargs := []gdc.ConstTypePtr{title.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnSetTitle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphNode) GetTitle() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetTitle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GraphNode) GetTitlebarHbox() HBoxContainer {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewHBoxContainer()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetTitlebarHbox), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GraphNode) SetSlot(slot_index int64, enable_left_port bool, type_left int64, color_left Color, enable_right_port bool, type_right int64, color_right Color, custom_icon_left Texture2D, custom_icon_right Texture2D, draw_stylebox bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&enable_left_port), gdc.ConstTypePtr(&type_left), color_left.AsCTypePtr(), gdc.ConstTypePtr(&enable_right_port), gdc.ConstTypePtr(&type_right), color_right.AsCTypePtr(), custom_icon_left.AsCTypePtr(), custom_icon_right.AsCTypePtr(), gdc.ConstTypePtr(&draw_stylebox)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnSetSlot), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphNode) ClearSlot(slot_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnClearSlot), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphNode) ClearAllSlots() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnClearAllSlots), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphNode) IsSlotEnabledLeft(slot_index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&slot_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnIsSlotEnabledLeft), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphNode) SetSlotEnabledLeft(slot_index int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnSetSlotEnabledLeft), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphNode) SetSlotTypeLeft(slot_index int64, type_ int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnSetSlotTypeLeft), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphNode) GetSlotTypeLeft(slot_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&slot_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetSlotTypeLeft), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphNode) SetSlotColorLeft(slot_index int64, color Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnSetSlotColorLeft), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphNode) GetSlotColorLeft(slot_index int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&slot_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetSlotColorLeft), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GraphNode) IsSlotEnabledRight(slot_index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&slot_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnIsSlotEnabledRight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphNode) SetSlotEnabledRight(slot_index int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnSetSlotEnabledRight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphNode) SetSlotTypeRight(slot_index int64, type_ int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnSetSlotTypeRight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphNode) GetSlotTypeRight(slot_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&slot_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetSlotTypeRight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphNode) SetSlotColorRight(slot_index int64, color Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnSetSlotColorRight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphNode) GetSlotColorRight(slot_index int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&slot_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetSlotColorRight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GraphNode) IsSlotDrawStylebox(slot_index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&slot_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnIsSlotDrawStylebox), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphNode) SetSlotDrawStylebox(slot_index int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnSetSlotDrawStylebox), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GraphNode) GetInputPortCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetInputPortCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphNode) GetInputPortPosition(port_idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&port_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetInputPortPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GraphNode) GetInputPortType(port_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&port_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetInputPortType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphNode) GetInputPortColor(port_idx int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&port_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetInputPortColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GraphNode) GetInputPortSlot(port_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&port_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetInputPortSlot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphNode) GetOutputPortCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetOutputPortCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphNode) GetOutputPortPosition(port_idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&port_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetOutputPortPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GraphNode) GetOutputPortType(port_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&port_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetOutputPortType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GraphNode) GetOutputPortColor(port_idx int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&port_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetOutputPortColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GraphNode) GetOutputPortSlot(port_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&port_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGraphNode.fnGetOutputPortSlot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type GraphNodeSlotUpdatedSignalFn func(slot_index int)

func (me *GraphNode) ConnectSlotUpdated(subs SignalSubscribers, fn GraphNodeSlotUpdatedSignalFn) {
	sig := StringNameFromStr("slot_updated")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphNode) DisconnectSlotUpdated(subs SignalSubscribers, fn GraphNodeSlotUpdatedSignalFn) {
	sig := StringNameFromStr("slot_updated")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
