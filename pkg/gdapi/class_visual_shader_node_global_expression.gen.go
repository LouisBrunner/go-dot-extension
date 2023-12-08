// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeGlobalExpression struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeGlobalExpression(obj gdc.ObjectPtr) *VisualShaderNodeGlobalExpression {
  return &VisualShaderNodeGlobalExpression{
    obj: obj,
  }
}

func (me *VisualShaderNodeGlobalExpression) BaseClass() string {
  return "VisualShaderNodeGlobalExpression"
}
