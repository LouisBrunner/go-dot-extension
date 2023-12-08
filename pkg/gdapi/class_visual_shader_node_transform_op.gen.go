// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTransformOp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTransformOp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTransformOp) BaseClass() string {
  return "VisualShaderNodeTransformOp"
}

type VisualShaderNodeTransformOpOperator int
const (
  VisualShaderNodeTransformOpOpAxb VisualShaderNodeTransformOpOperator = 0
  VisualShaderNodeTransformOpOpBxa VisualShaderNodeTransformOpOperator = 1
  VisualShaderNodeTransformOpOpAxbComp VisualShaderNodeTransformOpOperator = 2
  VisualShaderNodeTransformOpOpBxaComp VisualShaderNodeTransformOpOperator = 3
  VisualShaderNodeTransformOpOpAdd VisualShaderNodeTransformOpOperator = 4
  VisualShaderNodeTransformOpOpAMinusB VisualShaderNodeTransformOpOperator = 5
  VisualShaderNodeTransformOpOpBMinusA VisualShaderNodeTransformOpOperator = 6
  VisualShaderNodeTransformOpOpADivB VisualShaderNodeTransformOpOperator = 7
  VisualShaderNodeTransformOpOpBDivA VisualShaderNodeTransformOpOperator = 8
  VisualShaderNodeTransformOpOpMax VisualShaderNodeTransformOpOperator = 9
)
