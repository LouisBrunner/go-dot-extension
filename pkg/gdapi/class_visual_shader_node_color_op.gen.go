// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeColorOp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeColorOp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeColorOp) BaseClass() string {
  return "VisualShaderNodeColorOp"
}

type VisualShaderNodeColorOpOperator int
const (
  VisualShaderNodeColorOpOpScreen VisualShaderNodeColorOpOperator = 0
  VisualShaderNodeColorOpOpDifference VisualShaderNodeColorOpOperator = 1
  VisualShaderNodeColorOpOpDarken VisualShaderNodeColorOpOperator = 2
  VisualShaderNodeColorOpOpLighten VisualShaderNodeColorOpOperator = 3
  VisualShaderNodeColorOpOpOverlay VisualShaderNodeColorOpOperator = 4
  VisualShaderNodeColorOpOpDodge VisualShaderNodeColorOpOperator = 5
  VisualShaderNodeColorOpOpBurn VisualShaderNodeColorOpOperator = 6
  VisualShaderNodeColorOpOpSoftLight VisualShaderNodeColorOpOperator = 7
  VisualShaderNodeColorOpOpHardLight VisualShaderNodeColorOpOperator = 8
  VisualShaderNodeColorOpOpMax VisualShaderNodeColorOpOperator = 9
)
