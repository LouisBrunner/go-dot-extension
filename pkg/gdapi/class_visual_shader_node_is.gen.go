// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeIs struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeIs) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeIs) BaseClass() string {
  return "VisualShaderNodeIs"
}

type VisualShaderNodeIsFunction int
const (
  VisualShaderNodeIsFuncIsInf VisualShaderNodeIsFunction = 0
  VisualShaderNodeIsFuncIsNan VisualShaderNodeIsFunction = 1
  VisualShaderNodeIsFuncMax VisualShaderNodeIsFunction = 2
)
