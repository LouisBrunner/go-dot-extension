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

type ptrsForVisualShaderNodeTransformParameterList struct {
	fnSetDefaultValueEnabled gdc.MethodBindPtr
	fnIsDefaultValueEnabled  gdc.MethodBindPtr
	fnSetDefaultValue        gdc.MethodBindPtr
	fnGetDefaultValue        gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeTransformParameter ptrsForVisualShaderNodeTransformParameterList

func initVisualShaderNodeTransformParameterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeTransformParameter")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_default_value_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTransformParameter.fnSetDefaultValueEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_default_value_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTransformParameter.fnIsDefaultValueEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_default_value")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTransformParameter.fnSetDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2952846383))
	}
	{
		methodName := StringNameFromStr("get_default_value")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeTransformParameter.fnGetDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
	}
}

type VisualShaderNodeTransformParameter struct {
	VisualShaderNodeParameter
}

func (me *VisualShaderNodeTransformParameter) BaseClass() string {
	return "VisualShaderNodeTransformParameter"
}

func NewVisualShaderNodeTransformParameter() *VisualShaderNodeTransformParameter {
	str := StringNameFromStr("VisualShaderNodeTransformParameter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeTransformParameter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeTransformParameter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTransformParameter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTransformParameter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeTransformParameter) SetDefaultValueEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTransformParameter.fnSetDefaultValueEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeTransformParameter) IsDefaultValueEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTransformParameter.fnIsDefaultValueEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShaderNodeTransformParameter) SetDefaultValue(value Transform3D) {
	cargs := []gdc.ConstTypePtr{value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTransformParameter.fnSetDefaultValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeTransformParameter) GetDefaultValue() Transform3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTransformParameter.fnGetDefaultValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
