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

type ptrsForVisualShaderNodeVec4ParameterList struct {
	fnSetDefaultValueEnabled gdc.MethodBindPtr
	fnIsDefaultValueEnabled  gdc.MethodBindPtr
	fnSetDefaultValue        gdc.MethodBindPtr
	fnGetDefaultValue        gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeVec4Parameter ptrsForVisualShaderNodeVec4ParameterList

func initVisualShaderNodeVec4ParameterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeVec4Parameter")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_default_value_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVec4Parameter.fnSetDefaultValueEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_default_value_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVec4Parameter.fnIsDefaultValueEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_default_value")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVec4Parameter.fnSetDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 643568085))
	}
	{
		methodName := StringNameFromStr("get_default_value")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVec4Parameter.fnGetDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2435802345))
	}
}

type VisualShaderNodeVec4Parameter struct {
	VisualShaderNodeParameter
}

func (me *VisualShaderNodeVec4Parameter) BaseClass() string {
	return "VisualShaderNodeVec4Parameter"
}

func NewVisualShaderNodeVec4Parameter() *VisualShaderNodeVec4Parameter {
	str := StringNameFromStr("VisualShaderNodeVec4Parameter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeVec4Parameter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeVec4Parameter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVec4Parameter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVec4Parameter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeVec4Parameter) SetDefaultValueEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVec4Parameter.fnSetDefaultValueEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeVec4Parameter) IsDefaultValueEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVec4Parameter.fnIsDefaultValueEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShaderNodeVec4Parameter) SetDefaultValue(value Vector4) {
	cargs := []gdc.ConstTypePtr{value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVec4Parameter.fnSetDefaultValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeVec4Parameter) GetDefaultValue() Vector4 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector4()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVec4Parameter.fnGetDefaultValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
