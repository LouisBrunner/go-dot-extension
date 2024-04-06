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

type ptrsForVisualShaderNodeBooleanParameterList struct {
	fnSetDefaultValueEnabled gdc.MethodBindPtr
	fnIsDefaultValueEnabled  gdc.MethodBindPtr
	fnSetDefaultValue        gdc.MethodBindPtr
	fnGetDefaultValue        gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeBooleanParameter ptrsForVisualShaderNodeBooleanParameterList

func initVisualShaderNodeBooleanParameterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeBooleanParameter")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_default_value_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeBooleanParameter.fnSetDefaultValueEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_default_value_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeBooleanParameter.fnIsDefaultValueEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_default_value")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeBooleanParameter.fnSetDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_default_value")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeBooleanParameter.fnGetDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type VisualShaderNodeBooleanParameter struct {
	VisualShaderNodeParameter
}

func (me *VisualShaderNodeBooleanParameter) BaseClass() string {
	return "VisualShaderNodeBooleanParameter"
}

func NewVisualShaderNodeBooleanParameter() *VisualShaderNodeBooleanParameter {
	str := StringNameFromStr("VisualShaderNodeBooleanParameter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeBooleanParameter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeBooleanParameter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeBooleanParameter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeBooleanParameter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeBooleanParameter) SetDefaultValueEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeBooleanParameter.fnSetDefaultValueEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeBooleanParameter) IsDefaultValueEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeBooleanParameter.fnIsDefaultValueEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShaderNodeBooleanParameter) SetDefaultValue(value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeBooleanParameter.fnSetDefaultValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeBooleanParameter) GetDefaultValue() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeBooleanParameter.fnGetDefaultValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
