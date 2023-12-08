// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeStep struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeStep) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeStep) BaseClass() string {
  return "VisualShaderNodeStep"
}

type VisualShaderNodeStepOpType int
const (
  VisualShaderNodeStepOpTypeScalar VisualShaderNodeStepOpType = 0
  VisualShaderNodeStepOpTypeVector2D VisualShaderNodeStepOpType = 1
  VisualShaderNodeStepOpTypeVector2DScalar VisualShaderNodeStepOpType = 2
  VisualShaderNodeStepOpTypeVector3D VisualShaderNodeStepOpType = 3
  VisualShaderNodeStepOpTypeVector3DScalar VisualShaderNodeStepOpType = 4
  VisualShaderNodeStepOpTypeVector4D VisualShaderNodeStepOpType = 5
  VisualShaderNodeStepOpTypeVector4DScalar VisualShaderNodeStepOpType = 6
  VisualShaderNodeStepOpTypeMax VisualShaderNodeStepOpType = 7
)
