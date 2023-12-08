// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeIntOp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeIntOp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeIntOp) BaseClass() string {
  return "VisualShaderNodeIntOp"
}

type VisualShaderNodeIntOpOperator int
const (
  VisualShaderNodeIntOpOpAdd VisualShaderNodeIntOpOperator = 0
  VisualShaderNodeIntOpOpSub VisualShaderNodeIntOpOperator = 1
  VisualShaderNodeIntOpOpMul VisualShaderNodeIntOpOperator = 2
  VisualShaderNodeIntOpOpDiv VisualShaderNodeIntOpOperator = 3
  VisualShaderNodeIntOpOpMod VisualShaderNodeIntOpOperator = 4
  VisualShaderNodeIntOpOpMax VisualShaderNodeIntOpOperator = 5
  VisualShaderNodeIntOpOpMin VisualShaderNodeIntOpOperator = 6
  VisualShaderNodeIntOpOpBitwiseAnd VisualShaderNodeIntOpOperator = 7
  VisualShaderNodeIntOpOpBitwiseOr VisualShaderNodeIntOpOperator = 8
  VisualShaderNodeIntOpOpBitwiseXor VisualShaderNodeIntOpOperator = 9
  VisualShaderNodeIntOpOpBitwiseLeftShift VisualShaderNodeIntOpOperator = 10
  VisualShaderNodeIntOpOpBitwiseRightShift VisualShaderNodeIntOpOperator = 11
  VisualShaderNodeIntOpOpEnumSize VisualShaderNodeIntOpOperator = 12
)
