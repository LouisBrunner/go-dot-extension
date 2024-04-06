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

type ptrsForBoxContainerList struct {
	fnAddSpacer    gdc.MethodBindPtr
	fnSetAlignment gdc.MethodBindPtr
	fnGetAlignment gdc.MethodBindPtr
	fnSetVertical  gdc.MethodBindPtr
	fnIsVertical   gdc.MethodBindPtr
}

var ptrsForBoxContainer ptrsForBoxContainerList

func initBoxContainerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("BoxContainer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_spacer")
		defer methodName.Destroy()
		ptrsForBoxContainer.fnAddSpacer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1326660695))
	}
	{
		methodName := StringNameFromStr("set_alignment")
		defer methodName.Destroy()
		ptrsForBoxContainer.fnSetAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2456745134))
	}
	{
		methodName := StringNameFromStr("get_alignment")
		defer methodName.Destroy()
		ptrsForBoxContainer.fnGetAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1915476527))
	}
	{
		methodName := StringNameFromStr("set_vertical")
		defer methodName.Destroy()
		ptrsForBoxContainer.fnSetVertical = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_vertical")
		defer methodName.Destroy()
		ptrsForBoxContainer.fnIsVertical = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type BoxContainer struct {
	Container
}

func (me *BoxContainer) BaseClass() string {
	return "BoxContainer"
}

func NewBoxContainer() *BoxContainer {
	str := StringNameFromStr("BoxContainer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &BoxContainer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type BoxContainerAlignmentMode int

const (
	BoxContainerAlignmentModeAlignmentBegin  BoxContainerAlignmentMode = 0
	BoxContainerAlignmentModeAlignmentCenter BoxContainerAlignmentMode = 1
	BoxContainerAlignmentModeAlignmentEnd    BoxContainerAlignmentMode = 2
)

func (me *BoxContainer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *BoxContainer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *BoxContainer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *BoxContainer) AddSpacer(begin bool) Control {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&begin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewControl()
	pinner.Pin(&begin)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxContainer.fnAddSpacer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BoxContainer) SetAlignment(alignment BoxContainerAlignmentMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxContainer.fnSetAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BoxContainer) GetAlignment() BoxContainerAlignmentMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret BoxContainerAlignmentMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxContainer.fnGetAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *BoxContainer) SetVertical(vertical bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertical)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxContainer.fnSetVertical), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BoxContainer) IsVertical() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxContainer.fnIsVertical), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
