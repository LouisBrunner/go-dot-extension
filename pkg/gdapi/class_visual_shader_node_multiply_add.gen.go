// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  VisualShaderNodeMultiplyAddOpTypeOpTypeScalar VisualShaderNodeMultiplyAddOpType = 0
  VisualShaderNodeMultiplyAddOpTypeOpTypeVector2D VisualShaderNodeMultiplyAddOpType = 1
  VisualShaderNodeMultiplyAddOpTypeOpTypeVector3D VisualShaderNodeMultiplyAddOpType = 2
  VisualShaderNodeMultiplyAddOpTypeOpTypeVector4D VisualShaderNodeMultiplyAddOpType = 3
  VisualShaderNodeMultiplyAddOpTypeOpTypeMax VisualShaderNodeMultiplyAddOpType = 4
)

func  (me *VisualShaderNodeMultiplyAdd) SetOpType(type_ VisualShaderNodeMultiplyAddOpType, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeMultiplyAdd) GetOpType() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
