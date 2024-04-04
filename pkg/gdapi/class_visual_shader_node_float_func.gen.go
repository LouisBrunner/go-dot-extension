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

type VisualShaderNodeFloatFunc struct {
  VisualShaderNode
}

func (me *VisualShaderNodeFloatFunc) BaseClass() string {
  return "VisualShaderNodeFloatFunc"
}

func NewVisualShaderNodeFloatFunc() *VisualShaderNodeFloatFunc {
  str := StringNameFromStr("VisualShaderNodeFloatFunc") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeFloatFunc{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodeFloatFuncFunction int
const (
  VisualShaderNodeFloatFuncFunctionFuncSin VisualShaderNodeFloatFuncFunction = 0
  VisualShaderNodeFloatFuncFunctionFuncCos VisualShaderNodeFloatFuncFunction = 1
  VisualShaderNodeFloatFuncFunctionFuncTan VisualShaderNodeFloatFuncFunction = 2
  VisualShaderNodeFloatFuncFunctionFuncAsin VisualShaderNodeFloatFuncFunction = 3
  VisualShaderNodeFloatFuncFunctionFuncAcos VisualShaderNodeFloatFuncFunction = 4
  VisualShaderNodeFloatFuncFunctionFuncAtan VisualShaderNodeFloatFuncFunction = 5
  VisualShaderNodeFloatFuncFunctionFuncSinh VisualShaderNodeFloatFuncFunction = 6
  VisualShaderNodeFloatFuncFunctionFuncCosh VisualShaderNodeFloatFuncFunction = 7
  VisualShaderNodeFloatFuncFunctionFuncTanh VisualShaderNodeFloatFuncFunction = 8
  VisualShaderNodeFloatFuncFunctionFuncLog VisualShaderNodeFloatFuncFunction = 9
  VisualShaderNodeFloatFuncFunctionFuncExp VisualShaderNodeFloatFuncFunction = 10
  VisualShaderNodeFloatFuncFunctionFuncSqrt VisualShaderNodeFloatFuncFunction = 11
  VisualShaderNodeFloatFuncFunctionFuncAbs VisualShaderNodeFloatFuncFunction = 12
  VisualShaderNodeFloatFuncFunctionFuncSign VisualShaderNodeFloatFuncFunction = 13
  VisualShaderNodeFloatFuncFunctionFuncFloor VisualShaderNodeFloatFuncFunction = 14
  VisualShaderNodeFloatFuncFunctionFuncRound VisualShaderNodeFloatFuncFunction = 15
  VisualShaderNodeFloatFuncFunctionFuncCeil VisualShaderNodeFloatFuncFunction = 16
  VisualShaderNodeFloatFuncFunctionFuncFract VisualShaderNodeFloatFuncFunction = 17
  VisualShaderNodeFloatFuncFunctionFuncSaturate VisualShaderNodeFloatFuncFunction = 18
  VisualShaderNodeFloatFuncFunctionFuncNegate VisualShaderNodeFloatFuncFunction = 19
  VisualShaderNodeFloatFuncFunctionFuncAcosh VisualShaderNodeFloatFuncFunction = 20
  VisualShaderNodeFloatFuncFunctionFuncAsinh VisualShaderNodeFloatFuncFunction = 21
  VisualShaderNodeFloatFuncFunctionFuncAtanh VisualShaderNodeFloatFuncFunction = 22
  VisualShaderNodeFloatFuncFunctionFuncDegrees VisualShaderNodeFloatFuncFunction = 23
  VisualShaderNodeFloatFuncFunctionFuncExp2 VisualShaderNodeFloatFuncFunction = 24
  VisualShaderNodeFloatFuncFunctionFuncInverseSqrt VisualShaderNodeFloatFuncFunction = 25
  VisualShaderNodeFloatFuncFunctionFuncLog2 VisualShaderNodeFloatFuncFunction = 26
  VisualShaderNodeFloatFuncFunctionFuncRadians VisualShaderNodeFloatFuncFunction = 27
  VisualShaderNodeFloatFuncFunctionFuncReciprocal VisualShaderNodeFloatFuncFunction = 28
  VisualShaderNodeFloatFuncFunctionFuncRoundeven VisualShaderNodeFloatFuncFunction = 29
  VisualShaderNodeFloatFuncFunctionFuncTrunc VisualShaderNodeFloatFuncFunction = 30
  VisualShaderNodeFloatFuncFunctionFuncOneminus VisualShaderNodeFloatFuncFunction = 31
  VisualShaderNodeFloatFuncFunctionFuncMax VisualShaderNodeFloatFuncFunction = 32
)

func (me *VisualShaderNodeFloatFunc) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeFloatFunc) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeFloatFunc) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeFloatFunc) SetFunction(func_ VisualShaderNodeFloatFuncFunction, )  {
  classNameV := StringNameFromStr("VisualShaderNodeFloatFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 536026177) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&func_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeFloatFunc) GetFunction() VisualShaderNodeFloatFuncFunction {
  classNameV := StringNameFromStr("VisualShaderNodeFloatFunc")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_function")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2033948868) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeFloatFuncFunction

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
