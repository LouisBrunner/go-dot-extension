// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  VisualShaderNodeTransformFuncFuncInverse VisualShaderNodeTransformFuncFunction = 0
  VisualShaderNodeTransformFuncFuncTranspose VisualShaderNodeTransformFuncFunction = 1
  VisualShaderNodeTransformFuncFuncMax VisualShaderNodeTransformFuncFunction = 2
)
