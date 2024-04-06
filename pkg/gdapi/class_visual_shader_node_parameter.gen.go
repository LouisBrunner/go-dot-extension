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

type ptrsForVisualShaderNodeParameterList struct {
	fnSetParameterName gdc.MethodBindPtr
	fnGetParameterName gdc.MethodBindPtr
	fnSetQualifier     gdc.MethodBindPtr
	fnGetQualifier     gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeParameter ptrsForVisualShaderNodeParameterList

func initVisualShaderNodeParameterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeParameter")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_parameter_name")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParameter.fnSetParameterName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_parameter_name")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParameter.fnGetParameterName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_qualifier")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParameter.fnSetQualifier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1276489447))
	}
	{
		methodName := StringNameFromStr("get_qualifier")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParameter.fnGetQualifier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3558406205))
	}

}

type VisualShaderNodeParameter struct {
	VisualShaderNode
}

func (me *VisualShaderNodeParameter) BaseClass() string {
	return "VisualShaderNodeParameter"
}

func NewVisualShaderNodeParameter() *VisualShaderNodeParameter {
	str := StringNameFromStr("VisualShaderNodeParameter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeParameter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeParameterQualifier int

const (
	VisualShaderNodeParameterQualifierQualNone     VisualShaderNodeParameterQualifier = 0
	VisualShaderNodeParameterQualifierQualGlobal   VisualShaderNodeParameterQualifier = 1
	VisualShaderNodeParameterQualifierQualInstance VisualShaderNodeParameterQualifier = 2
	VisualShaderNodeParameterQualifierQualMax      VisualShaderNodeParameterQualifier = 3
)

func (me *VisualShaderNodeParameter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParameter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParameter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeParameter) SetParameterName(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParameter.fnSetParameterName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeParameter) GetParameterName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParameter.fnGetParameterName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisualShaderNodeParameter) SetQualifier(qualifier VisualShaderNodeParameterQualifier) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&qualifier)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParameter.fnSetQualifier), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeParameter) GetQualifier() VisualShaderNodeParameterQualifier {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeParameterQualifier

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParameter.fnGetQualifier), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
