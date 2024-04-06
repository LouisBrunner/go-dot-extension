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

type ptrsForVisualShaderNodeVec3ParameterList struct {
	fnSetDefaultValueEnabled gdc.MethodBindPtr
	fnIsDefaultValueEnabled  gdc.MethodBindPtr
	fnSetDefaultValue        gdc.MethodBindPtr
	fnGetDefaultValue        gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeVec3Parameter ptrsForVisualShaderNodeVec3ParameterList

func initVisualShaderNodeVec3ParameterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeVec3Parameter")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_default_value_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVec3Parameter.fnSetDefaultValueEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_default_value_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVec3Parameter.fnIsDefaultValueEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_default_value")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVec3Parameter.fnSetDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_default_value")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVec3Parameter.fnGetDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
}

type VisualShaderNodeVec3Parameter struct {
	VisualShaderNodeParameter
}

func (me *VisualShaderNodeVec3Parameter) BaseClass() string {
	return "VisualShaderNodeVec3Parameter"
}

func NewVisualShaderNodeVec3Parameter() *VisualShaderNodeVec3Parameter {
	str := StringNameFromStr("VisualShaderNodeVec3Parameter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeVec3Parameter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeVec3Parameter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVec3Parameter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVec3Parameter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeVec3Parameter) SetDefaultValueEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVec3Parameter.fnSetDefaultValueEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeVec3Parameter) IsDefaultValueEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVec3Parameter.fnIsDefaultValueEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShaderNodeVec3Parameter) SetDefaultValue(value Vector3) {
	cargs := []gdc.ConstTypePtr{value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVec3Parameter.fnSetDefaultValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeVec3Parameter) GetDefaultValue() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVec3Parameter.fnGetDefaultValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
