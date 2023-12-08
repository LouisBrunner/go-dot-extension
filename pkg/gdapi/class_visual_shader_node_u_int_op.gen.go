// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  VisualShaderNodeUIntOpOpAdd VisualShaderNodeUIntOpOperator = 0
  VisualShaderNodeUIntOpOpSub VisualShaderNodeUIntOpOperator = 1
  VisualShaderNodeUIntOpOpMul VisualShaderNodeUIntOpOperator = 2
  VisualShaderNodeUIntOpOpDiv VisualShaderNodeUIntOpOperator = 3
  VisualShaderNodeUIntOpOpMod VisualShaderNodeUIntOpOperator = 4
  VisualShaderNodeUIntOpOpMax VisualShaderNodeUIntOpOperator = 5
  VisualShaderNodeUIntOpOpMin VisualShaderNodeUIntOpOperator = 6
  VisualShaderNodeUIntOpOpBitwiseAnd VisualShaderNodeUIntOpOperator = 7
  VisualShaderNodeUIntOpOpBitwiseOr VisualShaderNodeUIntOpOperator = 8
  VisualShaderNodeUIntOpOpBitwiseXor VisualShaderNodeUIntOpOperator = 9
  VisualShaderNodeUIntOpOpBitwiseLeftShift VisualShaderNodeUIntOpOperator = 10
  VisualShaderNodeUIntOpOpBitwiseRightShift VisualShaderNodeUIntOpOperator = 11
  VisualShaderNodeUIntOpOpEnumSize VisualShaderNodeUIntOpOperator = 12
)
