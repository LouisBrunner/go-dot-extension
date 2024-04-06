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

type ptrsForSpinBoxList struct {
	fnSetHorizontalAlignment gdc.MethodBindPtr
	fnGetHorizontalAlignment gdc.MethodBindPtr
	fnSetSuffix              gdc.MethodBindPtr
	fnGetSuffix              gdc.MethodBindPtr
	fnSetPrefix              gdc.MethodBindPtr
	fnGetPrefix              gdc.MethodBindPtr
	fnSetEditable            gdc.MethodBindPtr
	fnSetCustomArrowStep     gdc.MethodBindPtr
	fnGetCustomArrowStep     gdc.MethodBindPtr
	fnIsEditable             gdc.MethodBindPtr
	fnSetUpdateOnTextChanged gdc.MethodBindPtr
	fnGetUpdateOnTextChanged gdc.MethodBindPtr
	fnSetSelectAllOnFocus    gdc.MethodBindPtr
	fnIsSelectAllOnFocus     gdc.MethodBindPtr
	fnApply                  gdc.MethodBindPtr
	fnGetLineEdit            gdc.MethodBindPtr
}

var ptrsForSpinBox ptrsForSpinBoxList

func initSpinBoxPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SpinBox")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_horizontal_alignment")
		defer methodName.Destroy()
		ptrsForSpinBox.fnSetHorizontalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2312603777))
	}
	{
		methodName := StringNameFromStr("get_horizontal_alignment")
		defer methodName.Destroy()
		ptrsForSpinBox.fnGetHorizontalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 341400642))
	}
	{
		methodName := StringNameFromStr("set_suffix")
		defer methodName.Destroy()
		ptrsForSpinBox.fnSetSuffix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_suffix")
		defer methodName.Destroy()
		ptrsForSpinBox.fnGetSuffix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_prefix")
		defer methodName.Destroy()
		ptrsForSpinBox.fnSetPrefix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_prefix")
		defer methodName.Destroy()
		ptrsForSpinBox.fnGetPrefix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_editable")
		defer methodName.Destroy()
		ptrsForSpinBox.fnSetEditable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_custom_arrow_step")
		defer methodName.Destroy()
		ptrsForSpinBox.fnSetCustomArrowStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_custom_arrow_step")
		defer methodName.Destroy()
		ptrsForSpinBox.fnGetCustomArrowStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("is_editable")
		defer methodName.Destroy()
		ptrsForSpinBox.fnIsEditable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_update_on_text_changed")
		defer methodName.Destroy()
		ptrsForSpinBox.fnSetUpdateOnTextChanged = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_update_on_text_changed")
		defer methodName.Destroy()
		ptrsForSpinBox.fnGetUpdateOnTextChanged = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_select_all_on_focus")
		defer methodName.Destroy()
		ptrsForSpinBox.fnSetSelectAllOnFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_select_all_on_focus")
		defer methodName.Destroy()
		ptrsForSpinBox.fnIsSelectAllOnFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("apply")
		defer methodName.Destroy()
		ptrsForSpinBox.fnApply = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_line_edit")
		defer methodName.Destroy()
		ptrsForSpinBox.fnGetLineEdit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4071694264))
	}
}

type SpinBox struct {
	Range
}

func (me *SpinBox) BaseClass() string {
	return "SpinBox"
}

func NewSpinBox() *SpinBox {
	str := StringNameFromStr("SpinBox") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SpinBox{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *SpinBox) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SpinBox) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SpinBox) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SpinBox) SetHorizontalAlignment(alignment HorizontalAlignment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnSetHorizontalAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpinBox) GetHorizontalAlignment() HorizontalAlignment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret HorizontalAlignment

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnGetHorizontalAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *SpinBox) SetSuffix(suffix String) {
	cargs := []gdc.ConstTypePtr{suffix.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnSetSuffix), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpinBox) GetSuffix() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnGetSuffix), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SpinBox) SetPrefix(prefix String) {
	cargs := []gdc.ConstTypePtr{prefix.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnSetPrefix), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpinBox) GetPrefix() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnGetPrefix), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SpinBox) SetEditable(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnSetEditable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpinBox) SetCustomArrowStep(arrow_step float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&arrow_step)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnSetCustomArrowStep), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpinBox) GetCustomArrowStep() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnGetCustomArrowStep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SpinBox) IsEditable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnIsEditable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SpinBox) SetUpdateOnTextChanged(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnSetUpdateOnTextChanged), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpinBox) GetUpdateOnTextChanged() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnGetUpdateOnTextChanged), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SpinBox) SetSelectAllOnFocus(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnSetSelectAllOnFocus), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpinBox) IsSelectAllOnFocus() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnIsSelectAllOnFocus), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SpinBox) Apply() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnApply), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpinBox) GetLineEdit() LineEdit {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewLineEdit()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpinBox.fnGetLineEdit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
