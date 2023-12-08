// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeUVFunc struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeUVFunc) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeUVFunc) BaseClass() string {
  return "VisualShaderNodeUVFunc"
}

type VisualShaderNodeUVFuncFunction int
const (
  VisualShaderNodeUVFuncFunctionFuncPanning VisualShaderNodeUVFuncFunction = 0
  VisualShaderNodeUVFuncFunctionFuncScaling VisualShaderNodeUVFuncFunction = 1
  VisualShaderNodeUVFuncFunctionFuncMax VisualShaderNodeUVFuncFunction = 2
)

func  (me *VisualShaderNodeUVFunc) SetFunction(func_ VisualShaderNodeUVFuncFunction, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeUVFunc) GetFunction() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
