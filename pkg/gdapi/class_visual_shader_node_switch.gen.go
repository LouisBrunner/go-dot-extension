// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  VisualShaderNodeSwitchOpTypeFloat VisualShaderNodeSwitchOpType = 0
  VisualShaderNodeSwitchOpTypeInt VisualShaderNodeSwitchOpType = 1
  VisualShaderNodeSwitchOpTypeUint VisualShaderNodeSwitchOpType = 2
  VisualShaderNodeSwitchOpTypeVector2D VisualShaderNodeSwitchOpType = 3
  VisualShaderNodeSwitchOpTypeVector3D VisualShaderNodeSwitchOpType = 4
  VisualShaderNodeSwitchOpTypeVector4D VisualShaderNodeSwitchOpType = 5
  VisualShaderNodeSwitchOpTypeBoolean VisualShaderNodeSwitchOpType = 6
  VisualShaderNodeSwitchOpTypeTransform VisualShaderNodeSwitchOpType = 7
  VisualShaderNodeSwitchOpTypeMax VisualShaderNodeSwitchOpType = 8
)
