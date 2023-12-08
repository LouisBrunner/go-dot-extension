// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  VisualShaderNodeIntFuncFuncAbs VisualShaderNodeIntFuncFunction = 0
  VisualShaderNodeIntFuncFuncNegate VisualShaderNodeIntFuncFunction = 1
  VisualShaderNodeIntFuncFuncSign VisualShaderNodeIntFuncFunction = 2
  VisualShaderNodeIntFuncFuncBitwiseNot VisualShaderNodeIntFuncFunction = 3
  VisualShaderNodeIntFuncFuncMax VisualShaderNodeIntFuncFunction = 4
)
