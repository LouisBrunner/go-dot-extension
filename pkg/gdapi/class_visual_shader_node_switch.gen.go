// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeSwitch struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeSwitch) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeSwitch) BaseClass() string {
  return "VisualShaderNodeSwitch"
}

type VisualShaderNodeSwitchOpType int
const (
  VisualShaderNodeSwitchOpTypeOpTypeFloat VisualShaderNodeSwitchOpType = 0
  VisualShaderNodeSwitchOpTypeOpTypeInt VisualShaderNodeSwitchOpType = 1
  VisualShaderNodeSwitchOpTypeOpTypeUint VisualShaderNodeSwitchOpType = 2
  VisualShaderNodeSwitchOpTypeOpTypeVector2D VisualShaderNodeSwitchOpType = 3
  VisualShaderNodeSwitchOpTypeOpTypeVector3D VisualShaderNodeSwitchOpType = 4
  VisualShaderNodeSwitchOpTypeOpTypeVector4D VisualShaderNodeSwitchOpType = 5
  VisualShaderNodeSwitchOpTypeOpTypeBoolean VisualShaderNodeSwitchOpType = 6
  VisualShaderNodeSwitchOpTypeOpTypeTransform VisualShaderNodeSwitchOpType = 7
  VisualShaderNodeSwitchOpTypeOpTypeMax VisualShaderNodeSwitchOpType = 8
)

func  (me *VisualShaderNodeSwitch) SetOpType(type_ VisualShaderNodeSwitchOpType, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeSwitch) GetOpType() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
