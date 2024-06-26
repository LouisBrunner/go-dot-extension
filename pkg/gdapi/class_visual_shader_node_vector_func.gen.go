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

type ptrsForVisualShaderNodeVectorFuncList struct {
	fnSetFunction gdc.MethodBindPtr
	fnGetFunction gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeVectorFunc ptrsForVisualShaderNodeVectorFuncList

func initVisualShaderNodeVectorFuncPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualShaderNodeVectorFunc")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_function")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVectorFunc.fnSetFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 629964457))
	}
	{
		methodName := StringNameFromStr("get_function")
		defer methodName.Destroy()
		ptrsForVisualShaderNodeVectorFunc.fnGetFunction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4047776843))
	}

}

type VisualShaderNodeVectorFunc struct {
	VisualShaderNodeVectorBase
}

func (me *VisualShaderNodeVectorFunc) BaseClass() string {
	return "VisualShaderNodeVectorFunc"
}

func NewVisualShaderNodeVectorFunc() *VisualShaderNodeVectorFunc {
	str := StringNameFromStr("VisualShaderNodeVectorFunc") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualShaderNodeVectorFunc{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type VisualShaderNodeVectorFuncFunction int

const (
	VisualShaderNodeVectorFuncFunctionFuncNormalize   VisualShaderNodeVectorFuncFunction = 0
	VisualShaderNodeVectorFuncFunctionFuncSaturate    VisualShaderNodeVectorFuncFunction = 1
	VisualShaderNodeVectorFuncFunctionFuncNegate      VisualShaderNodeVectorFuncFunction = 2
	VisualShaderNodeVectorFuncFunctionFuncReciprocal  VisualShaderNodeVectorFuncFunction = 3
	VisualShaderNodeVectorFuncFunctionFuncAbs         VisualShaderNodeVectorFuncFunction = 4
	VisualShaderNodeVectorFuncFunctionFuncAcos        VisualShaderNodeVectorFuncFunction = 5
	VisualShaderNodeVectorFuncFunctionFuncAcosh       VisualShaderNodeVectorFuncFunction = 6
	VisualShaderNodeVectorFuncFunctionFuncAsin        VisualShaderNodeVectorFuncFunction = 7
	VisualShaderNodeVectorFuncFunctionFuncAsinh       VisualShaderNodeVectorFuncFunction = 8
	VisualShaderNodeVectorFuncFunctionFuncAtan        VisualShaderNodeVectorFuncFunction = 9
	VisualShaderNodeVectorFuncFunctionFuncAtanh       VisualShaderNodeVectorFuncFunction = 10
	VisualShaderNodeVectorFuncFunctionFuncCeil        VisualShaderNodeVectorFuncFunction = 11
	VisualShaderNodeVectorFuncFunctionFuncCos         VisualShaderNodeVectorFuncFunction = 12
	VisualShaderNodeVectorFuncFunctionFuncCosh        VisualShaderNodeVectorFuncFunction = 13
	VisualShaderNodeVectorFuncFunctionFuncDegrees     VisualShaderNodeVectorFuncFunction = 14
	VisualShaderNodeVectorFuncFunctionFuncExp         VisualShaderNodeVectorFuncFunction = 15
	VisualShaderNodeVectorFuncFunctionFuncExp2        VisualShaderNodeVectorFuncFunction = 16
	VisualShaderNodeVectorFuncFunctionFuncFloor       VisualShaderNodeVectorFuncFunction = 17
	VisualShaderNodeVectorFuncFunctionFuncFract       VisualShaderNodeVectorFuncFunction = 18
	VisualShaderNodeVectorFuncFunctionFuncInverseSqrt VisualShaderNodeVectorFuncFunction = 19
	VisualShaderNodeVectorFuncFunctionFuncLog         VisualShaderNodeVectorFuncFunction = 20
	VisualShaderNodeVectorFuncFunctionFuncLog2        VisualShaderNodeVectorFuncFunction = 21
	VisualShaderNodeVectorFuncFunctionFuncRadians     VisualShaderNodeVectorFuncFunction = 22
	VisualShaderNodeVectorFuncFunctionFuncRound       VisualShaderNodeVectorFuncFunction = 23
	VisualShaderNodeVectorFuncFunctionFuncRoundeven   VisualShaderNodeVectorFuncFunction = 24
	VisualShaderNodeVectorFuncFunctionFuncSign        VisualShaderNodeVectorFuncFunction = 25
	VisualShaderNodeVectorFuncFunctionFuncSin         VisualShaderNodeVectorFuncFunction = 26
	VisualShaderNodeVectorFuncFunctionFuncSinh        VisualShaderNodeVectorFuncFunction = 27
	VisualShaderNodeVectorFuncFunctionFuncSqrt        VisualShaderNodeVectorFuncFunction = 28
	VisualShaderNodeVectorFuncFunctionFuncTan         VisualShaderNodeVectorFuncFunction = 29
	VisualShaderNodeVectorFuncFunctionFuncTanh        VisualShaderNodeVectorFuncFunction = 30
	VisualShaderNodeVectorFuncFunctionFuncTrunc       VisualShaderNodeVectorFuncFunction = 31
	VisualShaderNodeVectorFuncFunctionFuncOneminus    VisualShaderNodeVectorFuncFunction = 32
	VisualShaderNodeVectorFuncFunctionFuncMax         VisualShaderNodeVectorFuncFunction = 33
)

func (me *VisualShaderNodeVectorFunc) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVectorFunc) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVectorFunc) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualShaderNodeVectorFunc) SetFunction(func_ VisualShaderNodeVectorFuncFunction) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVectorFunc.fnSetFunction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualShaderNodeVectorFunc) GetFunction() VisualShaderNodeVectorFuncFunction {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VisualShaderNodeVectorFuncFunction

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVectorFunc.fnGetFunction), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
