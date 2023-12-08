// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  VisualShaderNodeFloatOpOperatorOpAdd VisualShaderNodeFloatOpOperator = 0
  VisualShaderNodeFloatOpOperatorOpSub VisualShaderNodeFloatOpOperator = 1
  VisualShaderNodeFloatOpOperatorOpMul VisualShaderNodeFloatOpOperator = 2
  VisualShaderNodeFloatOpOperatorOpDiv VisualShaderNodeFloatOpOperator = 3
  VisualShaderNodeFloatOpOperatorOpMod VisualShaderNodeFloatOpOperator = 4
  VisualShaderNodeFloatOpOperatorOpPow VisualShaderNodeFloatOpOperator = 5
  VisualShaderNodeFloatOpOperatorOpMax VisualShaderNodeFloatOpOperator = 6
  VisualShaderNodeFloatOpOperatorOpMin VisualShaderNodeFloatOpOperator = 7
  VisualShaderNodeFloatOpOperatorOpAtan2 VisualShaderNodeFloatOpOperator = 8
  VisualShaderNodeFloatOpOperatorOpStep VisualShaderNodeFloatOpOperator = 9
  VisualShaderNodeFloatOpOperatorOpEnumSize VisualShaderNodeFloatOpOperator = 10
)

func  (me *VisualShaderNodeFloatOp) SetOperator(op VisualShaderNodeFloatOpOperator, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeFloatOp) GetOperator() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
