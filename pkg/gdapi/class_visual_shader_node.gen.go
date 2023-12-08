// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  VisualShaderNodePortTypePortTypeScalar VisualShaderNodePortType = 0
  VisualShaderNodePortTypePortTypeScalarInt VisualShaderNodePortType = 1
  VisualShaderNodePortTypePortTypeScalarUint VisualShaderNodePortType = 2
  VisualShaderNodePortTypePortTypeVector2D VisualShaderNodePortType = 3
  VisualShaderNodePortTypePortTypeVector3D VisualShaderNodePortType = 4
  VisualShaderNodePortTypePortTypeVector4D VisualShaderNodePortType = 5
  VisualShaderNodePortTypePortTypeBoolean VisualShaderNodePortType = 6
  VisualShaderNodePortTypePortTypeTransform VisualShaderNodePortType = 7
  VisualShaderNodePortTypePortTypeSampler VisualShaderNodePortType = 8
  VisualShaderNodePortTypePortTypeMax VisualShaderNodePortType = 9
)

func  (me *VisualShaderNode) SetOutputPortForPreview(port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNode) GetOutputPortForPreview() { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNode) SetInputPortDefaultValue(port int, value Variant, prev_value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNode) GetInputPortDefaultValue(port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNode) RemoveInputPortDefaultValue(port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNode) ClearDefaultInputValues() { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNode) SetDefaultInputValues(values Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNode) GetDefaultInputValues() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
