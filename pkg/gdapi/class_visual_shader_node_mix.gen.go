// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeMix struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeMix) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeMix) BaseClass() string {
  return "VisualShaderNodeMix"
}

type VisualShaderNodeMixOpType int
const (
  VisualShaderNodeMixOpTypeOpTypeScalar VisualShaderNodeMixOpType = 0
  VisualShaderNodeMixOpTypeOpTypeVector2D VisualShaderNodeMixOpType = 1
  VisualShaderNodeMixOpTypeOpTypeVector2DScalar VisualShaderNodeMixOpType = 2
  VisualShaderNodeMixOpTypeOpTypeVector3D VisualShaderNodeMixOpType = 3
  VisualShaderNodeMixOpTypeOpTypeVector3DScalar VisualShaderNodeMixOpType = 4
  VisualShaderNodeMixOpTypeOpTypeVector4D VisualShaderNodeMixOpType = 5
  VisualShaderNodeMixOpTypeOpTypeVector4DScalar VisualShaderNodeMixOpType = 6
  VisualShaderNodeMixOpTypeOpTypeMax VisualShaderNodeMixOpType = 7
)

func  (me *VisualShaderNodeMix) SetOpType(op_type VisualShaderNodeMixOpType, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeMix) GetOpType() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
