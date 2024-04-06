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

type ptrsForVisualShaderNodeColorParameterList struct {
	fnSetDefaultValueEnabled gdc.MethodBindPtr
	fnIsDefaultValueEnabled  gdc.MethodBindPtr
	fnSetDefaultValue        gdc.MethodBindPtr
	fnGetDefaultValue        gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeColorParameter ptrsForVisualShaderNodeColorParameterList

func initVisualShaderNodeColorParameterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeColorParameter")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_default_value_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeColorParameter.fnSetDefaultValueEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_default_value_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeColorParameter.fnIsDefaultValueEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_default_value")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeColorParameter.fnSetDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_default_value")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeColorParameter.fnGetDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}

}

type VisualShaderNodeColorParameter struct {
	VisualShaderNodeParameter
}

func (me *VisualShaderNodeColorParameter) BaseClass() string {
	return "VisualShaderNodeColorParameter"
}

func NewVisualShaderNodeColorParameter() *VisualShaderNodeColorParameter {
	str := StringNameFromStr("VisualShaderNodeColorParameter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeColorParameter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeColorParameter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeColorParameter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeColorParameter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeColorParameter) SetDefaultValueEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeColorParameter.fnSetDefaultValueEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeColorParameter) IsDefaultValueEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeColorParameter.fnIsDefaultValueEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShaderNodeColorParameter) SetDefaultValue(value Color) {
	cargs := []gdc.ConstTypePtr{value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeColorParameter.fnSetDefaultValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeColorParameter) GetDefaultValue() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeColorParameter.fnGetDefaultValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
