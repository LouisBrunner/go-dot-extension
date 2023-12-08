// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  VisualShaderNodeVectorBaseOpTypeOpTypeVector2D VisualShaderNodeVectorBaseOpType = 0
  VisualShaderNodeVectorBaseOpTypeOpTypeVector3D VisualShaderNodeVectorBaseOpType = 1
  VisualShaderNodeVectorBaseOpTypeOpTypeVector4D VisualShaderNodeVectorBaseOpType = 2
  VisualShaderNodeVectorBaseOpTypeOpTypeMax VisualShaderNodeVectorBaseOpType = 3
)

func  (me *VisualShaderNodeVectorBase) SetOpType(type_ VisualShaderNodeVectorBaseOpType, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeVectorBase) GetOpType() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
