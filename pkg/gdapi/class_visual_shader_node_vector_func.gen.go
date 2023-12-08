// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVectorFunc struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeVectorFunc) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeVectorFunc) BaseClass() string {
  return "VisualShaderNodeVectorFunc"
}

type VisualShaderNodeVectorFuncFunction int
const (
  VisualShaderNodeVectorFuncFuncNormalize VisualShaderNodeVectorFuncFunction = 0
  VisualShaderNodeVectorFuncFuncSaturate VisualShaderNodeVectorFuncFunction = 1
  VisualShaderNodeVectorFuncFuncNegate VisualShaderNodeVectorFuncFunction = 2
  VisualShaderNodeVectorFuncFuncReciprocal VisualShaderNodeVectorFuncFunction = 3
  VisualShaderNodeVectorFuncFuncAbs VisualShaderNodeVectorFuncFunction = 4
  VisualShaderNodeVectorFuncFuncAcos VisualShaderNodeVectorFuncFunction = 5
  VisualShaderNodeVectorFuncFuncAcosh VisualShaderNodeVectorFuncFunction = 6
  VisualShaderNodeVectorFuncFuncAsin VisualShaderNodeVectorFuncFunction = 7
  VisualShaderNodeVectorFuncFuncAsinh VisualShaderNodeVectorFuncFunction = 8
  VisualShaderNodeVectorFuncFuncAtan VisualShaderNodeVectorFuncFunction = 9
  VisualShaderNodeVectorFuncFuncAtanh VisualShaderNodeVectorFuncFunction = 10
  VisualShaderNodeVectorFuncFuncCeil VisualShaderNodeVectorFuncFunction = 11
  VisualShaderNodeVectorFuncFuncCos VisualShaderNodeVectorFuncFunction = 12
  VisualShaderNodeVectorFuncFuncCosh VisualShaderNodeVectorFuncFunction = 13
  VisualShaderNodeVectorFuncFuncDegrees VisualShaderNodeVectorFuncFunction = 14
  VisualShaderNodeVectorFuncFuncExp VisualShaderNodeVectorFuncFunction = 15
  VisualShaderNodeVectorFuncFuncExp2 VisualShaderNodeVectorFuncFunction = 16
  VisualShaderNodeVectorFuncFuncFloor VisualShaderNodeVectorFuncFunction = 17
  VisualShaderNodeVectorFuncFuncFract VisualShaderNodeVectorFuncFunction = 18
  VisualShaderNodeVectorFuncFuncInverseSqrt VisualShaderNodeVectorFuncFunction = 19
  VisualShaderNodeVectorFuncFuncLog VisualShaderNodeVectorFuncFunction = 20
  VisualShaderNodeVectorFuncFuncLog2 VisualShaderNodeVectorFuncFunction = 21
  VisualShaderNodeVectorFuncFuncRadians VisualShaderNodeVectorFuncFunction = 22
  VisualShaderNodeVectorFuncFuncRound VisualShaderNodeVectorFuncFunction = 23
  VisualShaderNodeVectorFuncFuncRoundeven VisualShaderNodeVectorFuncFunction = 24
  VisualShaderNodeVectorFuncFuncSign VisualShaderNodeVectorFuncFunction = 25
  VisualShaderNodeVectorFuncFuncSin VisualShaderNodeVectorFuncFunction = 26
  VisualShaderNodeVectorFuncFuncSinh VisualShaderNodeVectorFuncFunction = 27
  VisualShaderNodeVectorFuncFuncSqrt VisualShaderNodeVectorFuncFunction = 28
  VisualShaderNodeVectorFuncFuncTan VisualShaderNodeVectorFuncFunction = 29
  VisualShaderNodeVectorFuncFuncTanh VisualShaderNodeVectorFuncFunction = 30
  VisualShaderNodeVectorFuncFuncTrunc VisualShaderNodeVectorFuncFunction = 31
  VisualShaderNodeVectorFuncFuncOneminus VisualShaderNodeVectorFuncFunction = 32
  VisualShaderNodeVectorFuncFuncMax VisualShaderNodeVectorFuncFunction = 33
)
