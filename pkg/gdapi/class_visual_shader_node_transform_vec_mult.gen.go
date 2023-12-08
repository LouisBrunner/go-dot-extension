// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTransformVecMult struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTransformVecMult) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTransformVecMult) BaseClass() string {
  return "VisualShaderNodeTransformVecMult"
}

type VisualShaderNodeTransformVecMultOperator int
const (
  VisualShaderNodeTransformVecMultOpAxb VisualShaderNodeTransformVecMultOperator = 0
  VisualShaderNodeTransformVecMultOpBxa VisualShaderNodeTransformVecMultOperator = 1
  VisualShaderNodeTransformVecMultOp3X3Axb VisualShaderNodeTransformVecMultOperator = 2
  VisualShaderNodeTransformVecMultOp3X3Bxa VisualShaderNodeTransformVecMultOperator = 3
  VisualShaderNodeTransformVecMultOpMax VisualShaderNodeTransformVecMultOperator = 4
)
