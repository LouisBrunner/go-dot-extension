// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  VisualShaderNodeDerivativeFuncOpTypeOpTypeScalar VisualShaderNodeDerivativeFuncOpType = 0
  VisualShaderNodeDerivativeFuncOpTypeOpTypeVector2D VisualShaderNodeDerivativeFuncOpType = 1
  VisualShaderNodeDerivativeFuncOpTypeOpTypeVector3D VisualShaderNodeDerivativeFuncOpType = 2
  VisualShaderNodeDerivativeFuncOpTypeOpTypeVector4D VisualShaderNodeDerivativeFuncOpType = 3
  VisualShaderNodeDerivativeFuncOpTypeOpTypeMax VisualShaderNodeDerivativeFuncOpType = 4
)

type VisualShaderNodeDerivativeFuncFunction int
const (
  VisualShaderNodeDerivativeFuncFunctionFuncSum VisualShaderNodeDerivativeFuncFunction = 0
  VisualShaderNodeDerivativeFuncFunctionFuncX VisualShaderNodeDerivativeFuncFunction = 1
  VisualShaderNodeDerivativeFuncFunctionFuncY VisualShaderNodeDerivativeFuncFunction = 2
  VisualShaderNodeDerivativeFuncFunctionFuncMax VisualShaderNodeDerivativeFuncFunction = 3
)

type VisualShaderNodeDerivativeFuncPrecision int
const (
  VisualShaderNodeDerivativeFuncPrecisionPrecisionNone VisualShaderNodeDerivativeFuncPrecision = 0
  VisualShaderNodeDerivativeFuncPrecisionPrecisionCoarse VisualShaderNodeDerivativeFuncPrecision = 1
  VisualShaderNodeDerivativeFuncPrecisionPrecisionFine VisualShaderNodeDerivativeFuncPrecision = 2
  VisualShaderNodeDerivativeFuncPrecisionPrecisionMax VisualShaderNodeDerivativeFuncPrecision = 3
)

func  (me *VisualShaderNodeDerivativeFunc) SetOpType(type_ VisualShaderNodeDerivativeFuncOpType, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeDerivativeFunc) GetOpType() { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeDerivativeFunc) SetFunction(func_ VisualShaderNodeDerivativeFuncFunction, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeDerivativeFunc) GetFunction() { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeDerivativeFunc) SetPrecision(precision VisualShaderNodeDerivativeFuncPrecision, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeDerivativeFunc) GetPrecision() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
