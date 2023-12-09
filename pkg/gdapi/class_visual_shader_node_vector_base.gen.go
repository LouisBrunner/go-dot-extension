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



// Enums

type VisualShaderNodeVectorBaseOpType int
const (
  VisualShaderNodeVectorBaseOpTypeOpTypeVector2D VisualShaderNodeVectorBaseOpType = 0
  VisualShaderNodeVectorBaseOpTypeOpTypeVector3D VisualShaderNodeVectorBaseOpType = 1
  VisualShaderNodeVectorBaseOpTypeOpTypeVector4D VisualShaderNodeVectorBaseOpType = 2
  VisualShaderNodeVectorBaseOpTypeOpTypeMax VisualShaderNodeVectorBaseOpType = 3
)

func (me *VisualShaderNodeVectorBase) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVectorBase) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVectorBase) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeVectorBase) SetOpType(type_ VisualShaderNodeVectorBaseOpType, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeVectorBase) GetOpType()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
