// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeStep struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeStep) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeStep) BaseClass() string {
  return "VisualShaderNodeStep"
}



// Enums

type VisualShaderNodeStepOpType int
const (
  VisualShaderNodeStepOpTypeOpTypeScalar VisualShaderNodeStepOpType = 0
  VisualShaderNodeStepOpTypeOpTypeVector2D VisualShaderNodeStepOpType = 1
  VisualShaderNodeStepOpTypeOpTypeVector2DScalar VisualShaderNodeStepOpType = 2
  VisualShaderNodeStepOpTypeOpTypeVector3D VisualShaderNodeStepOpType = 3
  VisualShaderNodeStepOpTypeOpTypeVector3DScalar VisualShaderNodeStepOpType = 4
  VisualShaderNodeStepOpTypeOpTypeVector4D VisualShaderNodeStepOpType = 5
  VisualShaderNodeStepOpTypeOpTypeVector4DScalar VisualShaderNodeStepOpType = 6
  VisualShaderNodeStepOpTypeOpTypeMax VisualShaderNodeStepOpType = 7
)

func (me *VisualShaderNodeStep) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeStep) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeStep) SetOpType(op_type VisualShaderNodeStepOpType, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeStep) GetOpType()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
