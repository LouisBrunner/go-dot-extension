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

type ptrsForVisualShaderNodeDerivativeFuncList struct {
	fnSetOpType    gdc.MethodBindPtr
	fnGetOpType    gdc.MethodBindPtr
	fnSetFunction  gdc.MethodBindPtr
	fnGetFunction  gdc.MethodBindPtr
	fnSetPrecision gdc.MethodBindPtr
	fnGetPrecision gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeDerivativeFunc ptrsForVisualShaderNodeDerivativeFuncList

func initVisualShaderNodeDerivativeFuncPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeDerivativeFunc")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_op_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeDerivativeFunc.fnSetOpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 377800221))
	}
	{
		methodName := StringNameFromStr("get_op_type")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeDerivativeFunc.fnGetOpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3997800514))
	}
	{
		methodName := StringNameFromStr("set_function")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeDerivativeFunc.fnSetFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1944704156))
	}
	{
		methodName := StringNameFromStr("get_function")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeDerivativeFunc.fnGetFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2389093396))
	}
	{
		methodName := StringNameFromStr("set_precision")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeDerivativeFunc.fnSetPrecision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 797270566))
	}
	{
		methodName := StringNameFromStr("get_precision")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeDerivativeFunc.fnGetPrecision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3822547323))
	}
}

type VisualShaderNodeDerivativeFunc struct {
	VisualShaderNode
}

func (me *VisualShaderNodeDerivativeFunc) BaseClass() string {
	return "VisualShaderNodeDerivativeFunc"
}

func NewVisualShaderNodeDerivativeFunc() *VisualShaderNodeDerivativeFunc {
	str := StringNameFromStr("VisualShaderNodeDerivativeFunc") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeDerivativeFunc{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeDerivativeFuncOpType int

const (
	VisualShaderNodeDerivativeFuncOpTypeOpTypeScalar   VisualShaderNodeDerivativeFuncOpType = 0
	VisualShaderNodeDerivativeFuncOpTypeOpTypeVector2D VisualShaderNodeDerivativeFuncOpType = 1
	VisualShaderNodeDerivativeFuncOpTypeOpTypeVector3D VisualShaderNodeDerivativeFuncOpType = 2
	VisualShaderNodeDerivativeFuncOpTypeOpTypeVector4D VisualShaderNodeDerivativeFuncOpType = 3
	VisualShaderNodeDerivativeFuncOpTypeOpTypeMax      VisualShaderNodeDerivativeFuncOpType = 4
)

type VisualShaderNodeDerivativeFuncFunction int

const (
	VisualShaderNodeDerivativeFuncFunctionFuncSum VisualShaderNodeDerivativeFuncFunction = 0
	VisualShaderNodeDerivativeFuncFunctionFuncX   VisualShaderNodeDerivativeFuncFunction = 1
	VisualShaderNodeDerivativeFuncFunctionFuncY   VisualShaderNodeDerivativeFuncFunction = 2
	VisualShaderNodeDerivativeFuncFunctionFuncMax VisualShaderNodeDerivativeFuncFunction = 3
)

type VisualShaderNodeDerivativeFuncPrecision int

const (
	VisualShaderNodeDerivativeFuncPrecisionPrecisionNone   VisualShaderNodeDerivativeFuncPrecision = 0
	VisualShaderNodeDerivativeFuncPrecisionPrecisionCoarse VisualShaderNodeDerivativeFuncPrecision = 1
	VisualShaderNodeDerivativeFuncPrecisionPrecisionFine   VisualShaderNodeDerivativeFuncPrecision = 2
	VisualShaderNodeDerivativeFuncPrecisionPrecisionMax    VisualShaderNodeDerivativeFuncPrecision = 3
)

func (me *VisualShaderNodeDerivativeFunc) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeDerivativeFunc) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeDerivativeFunc) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeDerivativeFunc) SetOpType(type_ VisualShaderNodeDerivativeFuncOpType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeDerivativeFunc.fnSetOpType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeDerivativeFunc) GetOpType() VisualShaderNodeDerivativeFuncOpType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeDerivativeFuncOpType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeDerivativeFunc.fnGetOpType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *VisualShaderNodeDerivativeFunc) SetFunction(func_ VisualShaderNodeDerivativeFuncFunction) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeDerivativeFunc.fnSetFunction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeDerivativeFunc) GetFunction() VisualShaderNodeDerivativeFuncFunction {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeDerivativeFuncFunction

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeDerivativeFunc.fnGetFunction), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *VisualShaderNodeDerivativeFunc) SetPrecision(precision VisualShaderNodeDerivativeFuncPrecision) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&precision)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeDerivativeFunc.fnSetPrecision), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeDerivativeFunc) GetPrecision() VisualShaderNodeDerivativeFuncPrecision {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeDerivativeFuncPrecision

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeDerivativeFunc.fnGetPrecision), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
