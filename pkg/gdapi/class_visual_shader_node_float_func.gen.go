// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeFloatFunc struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeFloatFunc) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeFloatFunc) BaseClass() string {
  return "VisualShaderNodeFloatFunc"
}

type VisualShaderNodeFloatFuncFunction int
const (
  VisualShaderNodeFloatFuncFuncSin VisualShaderNodeFloatFuncFunction = 0
  VisualShaderNodeFloatFuncFuncCos VisualShaderNodeFloatFuncFunction = 1
  VisualShaderNodeFloatFuncFuncTan VisualShaderNodeFloatFuncFunction = 2
  VisualShaderNodeFloatFuncFuncAsin VisualShaderNodeFloatFuncFunction = 3
  VisualShaderNodeFloatFuncFuncAcos VisualShaderNodeFloatFuncFunction = 4
  VisualShaderNodeFloatFuncFuncAtan VisualShaderNodeFloatFuncFunction = 5
  VisualShaderNodeFloatFuncFuncSinh VisualShaderNodeFloatFuncFunction = 6
  VisualShaderNodeFloatFuncFuncCosh VisualShaderNodeFloatFuncFunction = 7
  VisualShaderNodeFloatFuncFuncTanh VisualShaderNodeFloatFuncFunction = 8
  VisualShaderNodeFloatFuncFuncLog VisualShaderNodeFloatFuncFunction = 9
  VisualShaderNodeFloatFuncFuncExp VisualShaderNodeFloatFuncFunction = 10
  VisualShaderNodeFloatFuncFuncSqrt VisualShaderNodeFloatFuncFunction = 11
  VisualShaderNodeFloatFuncFuncAbs VisualShaderNodeFloatFuncFunction = 12
  VisualShaderNodeFloatFuncFuncSign VisualShaderNodeFloatFuncFunction = 13
  VisualShaderNodeFloatFuncFuncFloor VisualShaderNodeFloatFuncFunction = 14
  VisualShaderNodeFloatFuncFuncRound VisualShaderNodeFloatFuncFunction = 15
  VisualShaderNodeFloatFuncFuncCeil VisualShaderNodeFloatFuncFunction = 16
  VisualShaderNodeFloatFuncFuncFract VisualShaderNodeFloatFuncFunction = 17
  VisualShaderNodeFloatFuncFuncSaturate VisualShaderNodeFloatFuncFunction = 18
  VisualShaderNodeFloatFuncFuncNegate VisualShaderNodeFloatFuncFunction = 19
  VisualShaderNodeFloatFuncFuncAcosh VisualShaderNodeFloatFuncFunction = 20
  VisualShaderNodeFloatFuncFuncAsinh VisualShaderNodeFloatFuncFunction = 21
  VisualShaderNodeFloatFuncFuncAtanh VisualShaderNodeFloatFuncFunction = 22
  VisualShaderNodeFloatFuncFuncDegrees VisualShaderNodeFloatFuncFunction = 23
  VisualShaderNodeFloatFuncFuncExp2 VisualShaderNodeFloatFuncFunction = 24
  VisualShaderNodeFloatFuncFuncInverseSqrt VisualShaderNodeFloatFuncFunction = 25
  VisualShaderNodeFloatFuncFuncLog2 VisualShaderNodeFloatFuncFunction = 26
  VisualShaderNodeFloatFuncFuncRadians VisualShaderNodeFloatFuncFunction = 27
  VisualShaderNodeFloatFuncFuncReciprocal VisualShaderNodeFloatFuncFunction = 28
  VisualShaderNodeFloatFuncFuncRoundeven VisualShaderNodeFloatFuncFunction = 29
  VisualShaderNodeFloatFuncFuncTrunc VisualShaderNodeFloatFuncFunction = 30
  VisualShaderNodeFloatFuncFuncOneminus VisualShaderNodeFloatFuncFunction = 31
  VisualShaderNodeFloatFuncFuncMax VisualShaderNodeFloatFuncFunction = 32
)
