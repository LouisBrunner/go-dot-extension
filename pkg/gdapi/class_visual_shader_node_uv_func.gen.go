// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  VisualShaderNodeUVFuncFuncPanning VisualShaderNodeUVFuncFunction = 0
  VisualShaderNodeUVFuncFuncScaling VisualShaderNodeUVFuncFunction = 1
  VisualShaderNodeUVFuncFuncMax VisualShaderNodeUVFuncFunction = 2
)
