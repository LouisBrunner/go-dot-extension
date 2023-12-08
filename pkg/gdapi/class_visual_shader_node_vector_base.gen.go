// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVectorBase struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeVectorBase) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeVectorBase) BaseClass() string {
  return "VisualShaderNodeVectorBase"
}

type VisualShaderNodeVectorBaseOpType int
const (
  VisualShaderNodeVectorBaseOpTypeVector2D VisualShaderNodeVectorBaseOpType = 0
  VisualShaderNodeVectorBaseOpTypeVector3D VisualShaderNodeVectorBaseOpType = 1
  VisualShaderNodeVectorBaseOpTypeVector4D VisualShaderNodeVectorBaseOpType = 2
  VisualShaderNodeVectorBaseOpTypeMax VisualShaderNodeVectorBaseOpType = 3
)
