// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  VisualShaderNodeStepOpTypeOpTypeScalar VisualShaderNodeStepOpType = 0
  VisualShaderNodeStepOpTypeOpTypeVector2D VisualShaderNodeStepOpType = 1
  VisualShaderNodeStepOpTypeOpTypeVector2DScalar VisualShaderNodeStepOpType = 2
  VisualShaderNodeStepOpTypeOpTypeVector3D VisualShaderNodeStepOpType = 3
  VisualShaderNodeStepOpTypeOpTypeVector3DScalar VisualShaderNodeStepOpType = 4
  VisualShaderNodeStepOpTypeOpTypeVector4D VisualShaderNodeStepOpType = 5
  VisualShaderNodeStepOpTypeOpTypeVector4DScalar VisualShaderNodeStepOpType = 6
  VisualShaderNodeStepOpTypeOpTypeMax VisualShaderNodeStepOpType = 7
)

func  (me *VisualShaderNodeStep) SetOpType(op_type VisualShaderNodeStepOpType, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeStep) GetOpType() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
