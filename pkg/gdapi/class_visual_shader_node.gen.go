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



// Enums

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

func (me *VisualShaderNode) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNode) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNode) SetOutputPortForPreview(port int, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNode) GetOutputPortForPreview()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNode) SetInputPortDefaultValue(port int, value Variant, prev_value Variant, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNode) GetInputPortDefaultValue(port int, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNode) RemoveInputPortDefaultValue(port int, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNode) ClearDefaultInputValues()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNode) SetDefaultInputValues(values Array, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNode) GetDefaultInputValues()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
