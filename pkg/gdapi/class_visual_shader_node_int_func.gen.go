// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeIntFunc struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeIntFunc) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeIntFunc) BaseClass() string {
  return "VisualShaderNodeIntFunc"
}

type VisualShaderNodeIntFuncFunction int
const (
  VisualShaderNodeIntFuncFunctionFuncAbs VisualShaderNodeIntFuncFunction = 0
  VisualShaderNodeIntFuncFunctionFuncNegate VisualShaderNodeIntFuncFunction = 1
  VisualShaderNodeIntFuncFunctionFuncSign VisualShaderNodeIntFuncFunction = 2
  VisualShaderNodeIntFuncFunctionFuncBitwiseNot VisualShaderNodeIntFuncFunction = 3
  VisualShaderNodeIntFuncFunctionFuncMax VisualShaderNodeIntFuncFunction = 4
)

func  (me *VisualShaderNodeIntFunc) SetFunction(func_ VisualShaderNodeIntFuncFunction, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeIntFunc) GetFunction() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
