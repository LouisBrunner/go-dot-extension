// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeExpression struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeExpression(obj gdc.ObjectPtr) *VisualShaderNodeExpression {
  return &VisualShaderNodeExpression{
    obj: obj,
  }
}

func (me *VisualShaderNodeExpression) BaseClass() string {
  return "VisualShaderNodeExpression"
}
