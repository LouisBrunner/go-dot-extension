// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  VisualShaderNodeMixOpTypeScalar VisualShaderNodeMixOpType = 0
  VisualShaderNodeMixOpTypeVector2D VisualShaderNodeMixOpType = 1
  VisualShaderNodeMixOpTypeVector2DScalar VisualShaderNodeMixOpType = 2
  VisualShaderNodeMixOpTypeVector3D VisualShaderNodeMixOpType = 3
  VisualShaderNodeMixOpTypeVector3DScalar VisualShaderNodeMixOpType = 4
  VisualShaderNodeMixOpTypeVector4D VisualShaderNodeMixOpType = 5
  VisualShaderNodeMixOpTypeVector4DScalar VisualShaderNodeMixOpType = 6
  VisualShaderNodeMixOpTypeMax VisualShaderNodeMixOpType = 7
)
