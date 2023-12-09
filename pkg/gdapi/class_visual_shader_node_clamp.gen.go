// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeClamp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeClamp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeClamp) BaseClass() string {
  return "VisualShaderNodeClamp"
}



// Enums

type VisualShaderNodeClampOpType int
const (
  VisualShaderNodeClampOpTypeOpTypeFloat VisualShaderNodeClampOpType = 0
  VisualShaderNodeClampOpTypeOpTypeInt VisualShaderNodeClampOpType = 1
  VisualShaderNodeClampOpTypeOpTypeUint VisualShaderNodeClampOpType = 2
  VisualShaderNodeClampOpTypeOpTypeVector2D VisualShaderNodeClampOpType = 3
  VisualShaderNodeClampOpTypeOpTypeVector3D VisualShaderNodeClampOpType = 4
  VisualShaderNodeClampOpTypeOpTypeVector4D VisualShaderNodeClampOpType = 5
  VisualShaderNodeClampOpTypeOpTypeMax VisualShaderNodeClampOpType = 6
)

func (me *VisualShaderNodeClamp) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeClamp) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeClamp) SetOpType(op_type VisualShaderNodeClampOpType, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeClamp) GetOpType()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
