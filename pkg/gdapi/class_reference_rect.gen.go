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

type ptrsForReferenceRectList struct {
	fnGetBorderColor gdc.MethodBindPtr
	fnSetBorderColor gdc.MethodBindPtr
	fnGetBorderWidth gdc.MethodBindPtr
	fnSetBorderWidth gdc.MethodBindPtr
	fnGetEditorOnly  gdc.MethodBindPtr
	fnSetEditorOnly  gdc.MethodBindPtr
}

var ptrsForReferenceRect ptrsForReferenceRectList

func initReferenceRectPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ReferenceRect")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_border_color")
		defer methodName.Destroy()
		ptrsForReferenceRect.fnGetBorderColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_border_color")
		defer methodName.Destroy()
		ptrsForReferenceRect.fnSetBorderColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_border_width")
		defer methodName.Destroy()
		ptrsForReferenceRect.fnGetBorderWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_border_width")
		defer methodName.Destroy()
		ptrsForReferenceRect.fnSetBorderWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_editor_only")
		defer methodName.Destroy()
		ptrsForReferenceRect.fnGetEditorOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_editor_only")
		defer methodName.Destroy()
		ptrsForReferenceRect.fnSetEditorOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}

}

type ReferenceRect struct {
	Control
}

func (me *ReferenceRect) BaseClass() string {
	return "ReferenceRect"
}

func NewReferenceRect() *ReferenceRect {
	str := StringNameFromStr("ReferenceRect") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &ReferenceRect{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *ReferenceRect) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *ReferenceRect) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *ReferenceRect) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *ReferenceRect) GetBorderColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReferenceRect.fnGetBorderColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ReferenceRect) SetBorderColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReferenceRect.fnSetBorderColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ReferenceRect) GetBorderWidth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReferenceRect.fnGetBorderWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ReferenceRect) SetBorderWidth(width float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReferenceRect.fnSetBorderWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ReferenceRect) GetEditorOnly() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReferenceRect.fnGetEditorOnly), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ReferenceRect) SetEditorOnly(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForReferenceRect.fnSetEditorOnly), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
