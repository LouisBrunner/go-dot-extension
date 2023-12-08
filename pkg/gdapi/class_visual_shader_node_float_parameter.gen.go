// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeFloatParameter struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeFloatParameter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeFloatParameter) BaseClass() string {
  return "VisualShaderNodeFloatParameter"
}

type VisualShaderNodeFloatParameterHint int
const (
  VisualShaderNodeFloatParameterHintNone VisualShaderNodeFloatParameterHint = 0
  VisualShaderNodeFloatParameterHintRange VisualShaderNodeFloatParameterHint = 1
  VisualShaderNodeFloatParameterHintRangeStep VisualShaderNodeFloatParameterHint = 2
  VisualShaderNodeFloatParameterHintMax VisualShaderNodeFloatParameterHint = 3
)
