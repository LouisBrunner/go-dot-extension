// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeClamp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeClamp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeClamp) BaseClass() string {
  return "VisualShaderNodeClamp"
}

type VisualShaderNodeClampOpType int
const (
  VisualShaderNodeClampOpTypeFloat VisualShaderNodeClampOpType = 0
  VisualShaderNodeClampOpTypeInt VisualShaderNodeClampOpType = 1
  VisualShaderNodeClampOpTypeUint VisualShaderNodeClampOpType = 2
  VisualShaderNodeClampOpTypeVector2D VisualShaderNodeClampOpType = 3
  VisualShaderNodeClampOpTypeVector3D VisualShaderNodeClampOpType = 4
  VisualShaderNodeClampOpTypeVector4D VisualShaderNodeClampOpType = 5
  VisualShaderNodeClampOpTypeMax VisualShaderNodeClampOpType = 6
)
