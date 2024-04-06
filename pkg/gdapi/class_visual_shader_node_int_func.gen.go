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

type ptrsForVisualShaderNodeIntFuncList struct {
	fnSetFunction gdc.MethodBindPtr
	fnGetFunction gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeIntFunc ptrsForVisualShaderNodeIntFuncList

func initVisualShaderNodeIntFuncPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeIntFunc")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_function")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntFunc.fnSetFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 424195284))
	}
	{
		methodName := StringNameFromStr("get_function")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeIntFunc.fnGetFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2753496911))
	}

}

type VisualShaderNodeIntFunc struct {
	VisualShaderNode
}

func (me *VisualShaderNodeIntFunc) BaseClass() string {
	return "VisualShaderNodeIntFunc"
}

func NewVisualShaderNodeIntFunc() *VisualShaderNodeIntFunc {
	str := StringNameFromStr("VisualShaderNodeIntFunc") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeIntFunc{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeIntFuncFunction int

const (
	VisualShaderNodeIntFuncFunctionFuncAbs        VisualShaderNodeIntFuncFunction = 0
	VisualShaderNodeIntFuncFunctionFuncNegate     VisualShaderNodeIntFuncFunction = 1
	VisualShaderNodeIntFuncFunctionFuncSign       VisualShaderNodeIntFuncFunction = 2
	VisualShaderNodeIntFuncFunctionFuncBitwiseNot VisualShaderNodeIntFuncFunction = 3
	VisualShaderNodeIntFuncFunctionFuncMax        VisualShaderNodeIntFuncFunction = 4
)

func (me *VisualShaderNodeIntFunc) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeIntFunc) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeIntFunc) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeIntFunc) SetFunction(func_ VisualShaderNodeIntFuncFunction) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntFunc.fnSetFunction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeIntFunc) GetFunction() VisualShaderNodeIntFuncFunction {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeIntFuncFunction

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntFunc.fnGetFunction), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
