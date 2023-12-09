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



// Enums

type VisualShaderNodeMultiplyAddOpType int
const (
  VisualShaderNodeMultiplyAddOpTypeOpTypeScalar VisualShaderNodeMultiplyAddOpType = 0
  VisualShaderNodeMultiplyAddOpTypeOpTypeVector2D VisualShaderNodeMultiplyAddOpType = 1
  VisualShaderNodeMultiplyAddOpTypeOpTypeVector3D VisualShaderNodeMultiplyAddOpType = 2
  VisualShaderNodeMultiplyAddOpTypeOpTypeVector4D VisualShaderNodeMultiplyAddOpType = 3
  VisualShaderNodeMultiplyAddOpTypeOpTypeMax VisualShaderNodeMultiplyAddOpType = 4
)

func (me *VisualShaderNodeMultiplyAdd) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeMultiplyAdd) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeMultiplyAdd) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeMultiplyAdd) SetOpType(type_ VisualShaderNodeMultiplyAddOpType, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeMultiplyAdd) GetOpType()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
