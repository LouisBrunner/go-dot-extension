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

type ptrsForVisualShaderNodeParameterRefList struct {
	fnSetParameterName gdc.MethodBindPtr
	fnGetParameterName gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeParameterRef ptrsForVisualShaderNodeParameterRefList

func initVisualShaderNodeParameterRefPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeParameterRef")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_parameter_name")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParameterRef.fnSetParameterName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_parameter_name")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeParameterRef.fnGetParameterName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}

}

type VisualShaderNodeParameterRef struct {
	VisualShaderNode
}

func (me *VisualShaderNodeParameterRef) BaseClass() string {
	return "VisualShaderNodeParameterRef"
}

func NewVisualShaderNodeParameterRef() *VisualShaderNodeParameterRef {
	str := StringNameFromStr("VisualShaderNodeParameterRef") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeParameterRef{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualShaderNodeParameterRef) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParameterRef) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParameterRef) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeParameterRef) SetParameterName(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParameterRef.fnSetParameterName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeParameterRef) GetParameterName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParameterRef.fnGetParameterName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
