// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  VisualShaderNodeColorFuncFunctionFuncGrayscale VisualShaderNodeColorFuncFunction = 0
  VisualShaderNodeColorFuncFunctionFuncHsv2Rgb VisualShaderNodeColorFuncFunction = 1
  VisualShaderNodeColorFuncFunctionFuncRgb2Hsv VisualShaderNodeColorFuncFunction = 2
  VisualShaderNodeColorFuncFunctionFuncSepia VisualShaderNodeColorFuncFunction = 3
  VisualShaderNodeColorFuncFunctionFuncMax VisualShaderNodeColorFuncFunction = 4
)

func  (me *VisualShaderNodeColorFunc) SetFunction(func_ VisualShaderNodeColorFuncFunction, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeColorFunc) GetFunction() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
