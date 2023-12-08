// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeMultiplyAdd struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeMultiplyAdd) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeMultiplyAdd) BaseClass() string {
  return "VisualShaderNodeMultiplyAdd"
}

type VisualShaderNodeMultiplyAddOpType int
const (
  VisualShaderNodeMultiplyAddOpTypeScalar VisualShaderNodeMultiplyAddOpType = 0
  VisualShaderNodeMultiplyAddOpTypeVector2D VisualShaderNodeMultiplyAddOpType = 1
  VisualShaderNodeMultiplyAddOpTypeVector3D VisualShaderNodeMultiplyAddOpType = 2
  VisualShaderNodeMultiplyAddOpTypeVector4D VisualShaderNodeMultiplyAddOpType = 3
  VisualShaderNodeMultiplyAddOpTypeMax VisualShaderNodeMultiplyAddOpType = 4
)
