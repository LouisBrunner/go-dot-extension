// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeColorFunc struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeColorFunc) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeColorFunc) BaseClass() string {
  return "VisualShaderNodeColorFunc"
}

type VisualShaderNodeColorFuncFunction int
const (
  VisualShaderNodeColorFuncFuncGrayscale VisualShaderNodeColorFuncFunction = 0
  VisualShaderNodeColorFuncFuncHsv2Rgb VisualShaderNodeColorFuncFunction = 1
  VisualShaderNodeColorFuncFuncRgb2Hsv VisualShaderNodeColorFuncFunction = 2
  VisualShaderNodeColorFuncFuncSepia VisualShaderNodeColorFuncFunction = 3
  VisualShaderNodeColorFuncFuncMax VisualShaderNodeColorFuncFunction = 4
)
