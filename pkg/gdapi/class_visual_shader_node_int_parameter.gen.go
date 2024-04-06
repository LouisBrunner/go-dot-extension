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

type ptrsForVisualShaderNodeIntParameterList struct {
	fnSetHint                gdc.MethodBindPtr
	fnGetHint                gdc.MethodBindPtr
	fnSetMin                 gdc.MethodBindPtr
	fnGetMin                 gdc.MethodBindPtr
	fnSetMax                 gdc.MethodBindPtr
	fnGetMax                 gdc.MethodBindPtr
	fnSetStep                gdc.MethodBindPtr
	fnGetStep                gdc.MethodBindPtr
	fnSetDefaultValueEnabled gdc.MethodBindPtr
	fnIsDefaultValueEnabled  gdc.MethodBindPtr
	fnSetDefaultValue        gdc.MethodBindPtr
	fnGetDefaultValue        gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeIntParameter ptrsForVisualShaderNodeIntParameterList

func initVisualShaderNodeIntParameterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeIntParameter")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_hint")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntParameter.fnSetHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2540512075))
	}
	{
		methodName := StringNameFromStr("get_hint")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntParameter.fnGetHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4250814924))
	}
	{
		methodName := StringNameFromStr("set_min")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntParameter.fnSetMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_min")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntParameter.fnGetMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_max")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntParameter.fnSetMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_max")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntParameter.fnGetMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_step")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntParameter.fnSetStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_step")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntParameter.fnGetStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_default_value_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntParameter.fnSetDefaultValueEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_default_value_enabled")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntParameter.fnIsDefaultValueEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_default_value")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntParameter.fnSetDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_default_value")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntParameter.fnGetDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type VisualShaderNodeIntParameter struct {
	VisualShaderNodeParameter
}

func (me *VisualShaderNodeIntParameter) BaseClass() string {
	return "VisualShaderNodeIntParameter"
}

func NewVisualShaderNodeIntParameter() *VisualShaderNodeIntParameter {
	str := StringNameFromStr("VisualShaderNodeIntParameter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeIntParameter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeIntParameterHint int

const (
	VisualShaderNodeIntParameterHintHintNone      VisualShaderNodeIntParameterHint = 0
	VisualShaderNodeIntParameterHintHintRange     VisualShaderNodeIntParameterHint = 1
	VisualShaderNodeIntParameterHintHintRangeStep VisualShaderNodeIntParameterHint = 2
	VisualShaderNodeIntParameterHintHintMax       VisualShaderNodeIntParameterHint = 3
)

func (me *VisualShaderNodeIntParameter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeIntParameter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeIntParameter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeIntParameter) SetHint(hint VisualShaderNodeIntParameterHint) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntParameter.fnSetHint), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeIntParameter) GetHint() VisualShaderNodeIntParameterHint {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeIntParameterHint

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntParameter.fnGetHint), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *VisualShaderNodeIntParameter) SetMin(value int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntParameter.fnSetMin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeIntParameter) GetMin() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntParameter.fnGetMin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShaderNodeIntParameter) SetMax(value int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntParameter.fnSetMax), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeIntParameter) GetMax() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntParameter.fnGetMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShaderNodeIntParameter) SetStep(value int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntParameter.fnSetStep), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeIntParameter) GetStep() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntParameter.fnGetStep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShaderNodeIntParameter) SetDefaultValueEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntParameter.fnSetDefaultValueEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeIntParameter) IsDefaultValueEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntParameter.fnIsDefaultValueEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualShaderNodeIntParameter) SetDefaultValue(value int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntParameter.fnSetDefaultValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeIntParameter) GetDefaultValue() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntParameter.fnGetDefaultValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
