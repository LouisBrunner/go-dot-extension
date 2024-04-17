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

type ptrsForFlowContainerList struct {
	fnGetLineCount   gdc.MethodBindPtr
	fnSetAlignment   gdc.MethodBindPtr
	fnGetAlignment   gdc.MethodBindPtr
	fnSetVertical    gdc.MethodBindPtr
	fnIsVertical     gdc.MethodBindPtr
	fnSetReverseFill gdc.MethodBindPtr
	fnIsReverseFill  gdc.MethodBindPtr
}

var ptrsForFlowContainer ptrsForFlowContainerList

func initFlowContainerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("FlowContainer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_line_count")
		defer methodName.Destroy()
		ptrsForFlowContainer.fnGetLineCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_alignment")
		defer methodName.Destroy()
		ptrsForFlowContainer.fnSetAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 575250951))
	}
	{
		methodName := StringNameFromStr("get_alignment")
		defer methodName.Destroy()
		ptrsForFlowContainer.fnGetAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3749743559))
	}
	{
		methodName := StringNameFromStr("set_vertical")
		defer methodName.Destroy()
		ptrsForFlowContainer.fnSetVertical = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_vertical")
		defer methodName.Destroy()
		ptrsForFlowContainer.fnIsVertical = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_reverse_fill")
		defer methodName.Destroy()
		ptrsForFlowContainer.fnSetReverseFill = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_reverse_fill")
		defer methodName.Destroy()
		ptrsForFlowContainer.fnIsReverseFill = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type FlowContainer struct {
	Container
}

func (me *FlowContainer) BaseClass() string {
	return "FlowContainer"
}

func NewFlowContainer() *FlowContainer {
	str := StringNameFromStr("FlowContainer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &FlowContainer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type FlowContainerAlignmentMode int

const (
	FlowContainerAlignmentModeAlignmentBegin  FlowContainerAlignmentMode = 0
	FlowContainerAlignmentModeAlignmentCenter FlowContainerAlignmentMode = 1
	FlowContainerAlignmentModeAlignmentEnd    FlowContainerAlignmentMode = 2
)

func (me *FlowContainer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *FlowContainer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *FlowContainer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *FlowContainer) GetLineCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFlowContainer.fnGetLineCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FlowContainer) SetAlignment(alignment FlowContainerAlignmentMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFlowContainer.fnSetAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FlowContainer) GetAlignment() FlowContainerAlignmentMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret FlowContainerAlignmentMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFlowContainer.fnGetAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *FlowContainer) SetVertical(vertical bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertical)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFlowContainer.fnSetVertical), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FlowContainer) IsVertical() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFlowContainer.fnIsVertical), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *FlowContainer) SetReverseFill(reverse_fill bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&reverse_fill)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFlowContainer.fnSetReverseFill), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *FlowContainer) IsReverseFill() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForFlowContainer.fnIsReverseFill), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
