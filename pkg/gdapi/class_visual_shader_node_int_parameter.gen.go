// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeIntParameter struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeIntParameter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeIntParameter) BaseClass() string {
  return "VisualShaderNodeIntParameter"
}

type VisualShaderNodeIntParameterHint int
const (
  VisualShaderNodeIntParameterHintNone VisualShaderNodeIntParameterHint = 0
  VisualShaderNodeIntParameterHintRange VisualShaderNodeIntParameterHint = 1
  VisualShaderNodeIntParameterHintRangeStep VisualShaderNodeIntParameterHint = 2
  VisualShaderNodeIntParameterHintMax VisualShaderNodeIntParameterHint = 3
)
