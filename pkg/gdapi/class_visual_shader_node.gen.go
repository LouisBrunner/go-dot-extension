// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNode struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNode) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNode) BaseClass() string {
  return "VisualShaderNode"
}

type VisualShaderNodePortType int
const (
  VisualShaderNodePortTypeScalar VisualShaderNodePortType = 0
  VisualShaderNodePortTypeScalarInt VisualShaderNodePortType = 1
  VisualShaderNodePortTypeScalarUint VisualShaderNodePortType = 2
  VisualShaderNodePortTypeVector2D VisualShaderNodePortType = 3
  VisualShaderNodePortTypeVector3D VisualShaderNodePortType = 4
  VisualShaderNodePortTypeVector4D VisualShaderNodePortType = 5
  VisualShaderNodePortTypeBoolean VisualShaderNodePortType = 6
  VisualShaderNodePortTypeTransform VisualShaderNodePortType = 7
  VisualShaderNodePortTypeSampler VisualShaderNodePortType = 8
  VisualShaderNodePortTypeMax VisualShaderNodePortType = 9
)
