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

type ptrsForVisualShaderNodeUVFuncList struct {
	fnSetFunction gdc.MethodBindPtr
	fnGetFunction gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeUVFunc ptrsForVisualShaderNodeUVFuncList

func initVisualShaderNodeUVFuncPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeUVFunc")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_function")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeUVFunc.fnSetFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 765791915))
	}
	{
		methodName := StringNameFromStr("get_function")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeUVFunc.fnGetFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3772902164))
	}

}

type VisualShaderNodeUVFunc struct {
	VisualShaderNode
}

func (me *VisualShaderNodeUVFunc) BaseClass() string {
	return "VisualShaderNodeUVFunc"
}

func NewVisualShaderNodeUVFunc() *VisualShaderNodeUVFunc {
	str := StringNameFromStr("VisualShaderNodeUVFunc") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeUVFunc{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeUVFuncFunction int

const (
	VisualShaderNodeUVFuncFunctionFuncPanning VisualShaderNodeUVFuncFunction = 0
	VisualShaderNodeUVFuncFunctionFuncScaling VisualShaderNodeUVFuncFunction = 1
	VisualShaderNodeUVFuncFunctionFuncMax     VisualShaderNodeUVFuncFunction = 2
)

func (me *VisualShaderNodeUVFunc) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeUVFunc) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeUVFunc) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeUVFunc) SetFunction(func_ VisualShaderNodeUVFuncFunction) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeUVFunc.fnSetFunction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeUVFunc) GetFunction() VisualShaderNodeUVFuncFunction {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeUVFuncFunction

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeUVFunc.fnGetFunction), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
