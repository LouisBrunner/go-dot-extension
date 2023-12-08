// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVectorOp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeVectorOp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeVectorOp) BaseClass() string {
  return "VisualShaderNodeVectorOp"
}

type VisualShaderNodeVectorOpOperator int
const (
  VisualShaderNodeVectorOpOpAdd VisualShaderNodeVectorOpOperator = 0
  VisualShaderNodeVectorOpOpSub VisualShaderNodeVectorOpOperator = 1
  VisualShaderNodeVectorOpOpMul VisualShaderNodeVectorOpOperator = 2
  VisualShaderNodeVectorOpOpDiv VisualShaderNodeVectorOpOperator = 3
  VisualShaderNodeVectorOpOpMod VisualShaderNodeVectorOpOperator = 4
  VisualShaderNodeVectorOpOpPow VisualShaderNodeVectorOpOperator = 5
  VisualShaderNodeVectorOpOpMax VisualShaderNodeVectorOpOperator = 6
  VisualShaderNodeVectorOpOpMin VisualShaderNodeVectorOpOperator = 7
  VisualShaderNodeVectorOpOpCross VisualShaderNodeVectorOpOperator = 8
  VisualShaderNodeVectorOpOpAtan2 VisualShaderNodeVectorOpOperator = 9
  VisualShaderNodeVectorOpOpReflect VisualShaderNodeVectorOpOperator = 10
  VisualShaderNodeVectorOpOpStep VisualShaderNodeVectorOpOperator = 11
  VisualShaderNodeVectorOpOpEnumSize VisualShaderNodeVectorOpOperator = 12
)
