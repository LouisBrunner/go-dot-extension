// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeUIntOp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeUIntOp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeUIntOp) BaseClass() string {
  return "VisualShaderNodeUIntOp"
}

type VisualShaderNodeUIntOpOperator int
const (
  VisualShaderNodeUIntOpOperatorOpAdd VisualShaderNodeUIntOpOperator = 0
  VisualShaderNodeUIntOpOperatorOpSub VisualShaderNodeUIntOpOperator = 1
  VisualShaderNodeUIntOpOperatorOpMul VisualShaderNodeUIntOpOperator = 2
  VisualShaderNodeUIntOpOperatorOpDiv VisualShaderNodeUIntOpOperator = 3
  VisualShaderNodeUIntOpOperatorOpMod VisualShaderNodeUIntOpOperator = 4
  VisualShaderNodeUIntOpOperatorOpMax VisualShaderNodeUIntOpOperator = 5
  VisualShaderNodeUIntOpOperatorOpMin VisualShaderNodeUIntOpOperator = 6
  VisualShaderNodeUIntOpOperatorOpBitwiseAnd VisualShaderNodeUIntOpOperator = 7
  VisualShaderNodeUIntOpOperatorOpBitwiseOr VisualShaderNodeUIntOpOperator = 8
  VisualShaderNodeUIntOpOperatorOpBitwiseXor VisualShaderNodeUIntOpOperator = 9
  VisualShaderNodeUIntOpOperatorOpBitwiseLeftShift VisualShaderNodeUIntOpOperator = 10
  VisualShaderNodeUIntOpOperatorOpBitwiseRightShift VisualShaderNodeUIntOpOperator = 11
  VisualShaderNodeUIntOpOperatorOpEnumSize VisualShaderNodeUIntOpOperator = 12
)

func  (me *VisualShaderNodeUIntOp) SetOperator(op VisualShaderNodeUIntOpOperator, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeUIntOp) GetOperator() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
