// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeFloatOp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeFloatOp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeFloatOp) BaseClass() string {
  return "VisualShaderNodeFloatOp"
}

type VisualShaderNodeFloatOpOperator int
const (
  VisualShaderNodeFloatOpOpAdd VisualShaderNodeFloatOpOperator = 0
  VisualShaderNodeFloatOpOpSub VisualShaderNodeFloatOpOperator = 1
  VisualShaderNodeFloatOpOpMul VisualShaderNodeFloatOpOperator = 2
  VisualShaderNodeFloatOpOpDiv VisualShaderNodeFloatOpOperator = 3
  VisualShaderNodeFloatOpOpMod VisualShaderNodeFloatOpOperator = 4
  VisualShaderNodeFloatOpOpPow VisualShaderNodeFloatOpOperator = 5
  VisualShaderNodeFloatOpOpMax VisualShaderNodeFloatOpOperator = 6
  VisualShaderNodeFloatOpOpMin VisualShaderNodeFloatOpOperator = 7
  VisualShaderNodeFloatOpOpAtan2 VisualShaderNodeFloatOpOperator = 8
  VisualShaderNodeFloatOpOpStep VisualShaderNodeFloatOpOperator = 9
  VisualShaderNodeFloatOpOpEnumSize VisualShaderNodeFloatOpOperator = 10
)
