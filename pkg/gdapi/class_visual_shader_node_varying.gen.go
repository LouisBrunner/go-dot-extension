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

type ptrsForVisualShaderNodeVaryingList struct {
	fnSetVaryingName gdc.MethodBindPtr
	fnGetVaryingName gdc.MethodBindPtr
	fnSetVaryingType gdc.MethodBindPtr
	fnGetVaryingType gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeVarying ptrsForVisualShaderNodeVaryingList

func initVisualShaderNodeVaryingPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeVarying")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_varying_name")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVarying.fnSetVaryingName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_varying_name")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVarying.fnGetVaryingName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_varying_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVarying.fnSetVaryingType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3565867981))
	}
	{
		methodName := StringNameFromStr("get_varying_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVarying.fnGetVaryingType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 523183580))
	}
}

type VisualShaderNodeVarying struct {
	VisualShaderNode
}

func (me *VisualShaderNodeVarying) BaseClass() string {
	return "VisualShaderNodeVarying"
}

func NewVisualShaderNodeVarying() *VisualShaderNodeVarying {
	str := StringNameFromStr("VisualShaderNodeVarying") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeVarying{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeVarying) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVarying) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVarying) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeVarying) SetVaryingName(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVarying.fnSetVaryingName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeVarying) GetVaryingName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVarying.fnGetVaryingName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisualShaderNodeVarying) SetVaryingType(type_ VisualShaderVaryingType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVarying.fnSetVaryingType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeVarying) GetVaryingType() VisualShaderVaryingType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderVaryingType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVarying.fnGetVaryingType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
