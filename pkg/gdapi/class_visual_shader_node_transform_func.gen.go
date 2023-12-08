// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTransformFunc struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTransformFunc) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTransformFunc) BaseClass() string {
  return "VisualShaderNodeTransformFunc"
}

type VisualShaderNodeTransformFuncFunction int
const (
  VisualShaderNodeTransformFuncFunctionFuncInverse VisualShaderNodeTransformFuncFunction = 0
  VisualShaderNodeTransformFuncFunctionFuncTranspose VisualShaderNodeTransformFuncFunction = 1
  VisualShaderNodeTransformFuncFunctionFuncMax VisualShaderNodeTransformFuncFunction = 2
)

func  (me *VisualShaderNodeTransformFunc) SetFunction(func_ VisualShaderNodeTransformFuncFunction, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeTransformFunc) GetFunction() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
