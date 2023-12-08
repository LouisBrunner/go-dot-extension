// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  VisualShaderNodeSmoothStepOpTypeOpTypeScalar VisualShaderNodeSmoothStepOpType = 0
  VisualShaderNodeSmoothStepOpTypeOpTypeVector2D VisualShaderNodeSmoothStepOpType = 1
  VisualShaderNodeSmoothStepOpTypeOpTypeVector2DScalar VisualShaderNodeSmoothStepOpType = 2
  VisualShaderNodeSmoothStepOpTypeOpTypeVector3D VisualShaderNodeSmoothStepOpType = 3
  VisualShaderNodeSmoothStepOpTypeOpTypeVector3DScalar VisualShaderNodeSmoothStepOpType = 4
  VisualShaderNodeSmoothStepOpTypeOpTypeVector4D VisualShaderNodeSmoothStepOpType = 5
  VisualShaderNodeSmoothStepOpTypeOpTypeVector4DScalar VisualShaderNodeSmoothStepOpType = 6
  VisualShaderNodeSmoothStepOpTypeOpTypeMax VisualShaderNodeSmoothStepOpType = 7
)

func  (me *VisualShaderNodeSmoothStep) SetOpType(op_type VisualShaderNodeSmoothStepOpType, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeSmoothStep) GetOpType() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
