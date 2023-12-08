// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeDerivativeFunc struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeDerivativeFunc) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeDerivativeFunc) BaseClass() string {
  return "VisualShaderNodeDerivativeFunc"
}

type VisualShaderNodeDerivativeFuncOpType int
const (
  VisualShaderNodeDerivativeFuncOpTypeScalar VisualShaderNodeDerivativeFuncOpType = 0
  VisualShaderNodeDerivativeFuncOpTypeVector2D VisualShaderNodeDerivativeFuncOpType = 1
  VisualShaderNodeDerivativeFuncOpTypeVector3D VisualShaderNodeDerivativeFuncOpType = 2
  VisualShaderNodeDerivativeFuncOpTypeVector4D VisualShaderNodeDerivativeFuncOpType = 3
  VisualShaderNodeDerivativeFuncOpTypeMax VisualShaderNodeDerivativeFuncOpType = 4
)

type VisualShaderNodeDerivativeFuncFunction int
const (
  VisualShaderNodeDerivativeFuncFuncSum VisualShaderNodeDerivativeFuncFunction = 0
  VisualShaderNodeDerivativeFuncFuncX VisualShaderNodeDerivativeFuncFunction = 1
  VisualShaderNodeDerivativeFuncFuncY VisualShaderNodeDerivativeFuncFunction = 2
  VisualShaderNodeDerivativeFuncFuncMax VisualShaderNodeDerivativeFuncFunction = 3
)

type VisualShaderNodeDerivativeFuncPrecision int
const (
  VisualShaderNodeDerivativeFuncPrecisionNone VisualShaderNodeDerivativeFuncPrecision = 0
  VisualShaderNodeDerivativeFuncPrecisionCoarse VisualShaderNodeDerivativeFuncPrecision = 1
  VisualShaderNodeDerivativeFuncPrecisionFine VisualShaderNodeDerivativeFuncPrecision = 2
  VisualShaderNodeDerivativeFuncPrecisionMax VisualShaderNodeDerivativeFuncPrecision = 3
)
