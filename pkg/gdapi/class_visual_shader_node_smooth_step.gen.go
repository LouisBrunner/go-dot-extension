// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeSmoothStep struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeSmoothStep) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeSmoothStep) BaseClass() string {
  return "VisualShaderNodeSmoothStep"
}

type VisualShaderNodeSmoothStepOpType int
const (
  VisualShaderNodeSmoothStepOpTypeScalar VisualShaderNodeSmoothStepOpType = 0
  VisualShaderNodeSmoothStepOpTypeVector2D VisualShaderNodeSmoothStepOpType = 1
  VisualShaderNodeSmoothStepOpTypeVector2DScalar VisualShaderNodeSmoothStepOpType = 2
  VisualShaderNodeSmoothStepOpTypeVector3D VisualShaderNodeSmoothStepOpType = 3
  VisualShaderNodeSmoothStepOpTypeVector3DScalar VisualShaderNodeSmoothStepOpType = 4
  VisualShaderNodeSmoothStepOpTypeVector4D VisualShaderNodeSmoothStepOpType = 5
  VisualShaderNodeSmoothStepOpTypeVector4DScalar VisualShaderNodeSmoothStepOpType = 6
  VisualShaderNodeSmoothStepOpTypeMax VisualShaderNodeSmoothStepOpType = 7
)
