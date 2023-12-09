// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVectorOp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeVectorOp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeVectorOp) BaseClass() string {
  return "VisualShaderNodeVectorOp"
}



// Enums

type VisualShaderNodeVectorOpOperator int
const (
  VisualShaderNodeVectorOpOperatorOpAdd VisualShaderNodeVectorOpOperator = 0
  VisualShaderNodeVectorOpOperatorOpSub VisualShaderNodeVectorOpOperator = 1
  VisualShaderNodeVectorOpOperatorOpMul VisualShaderNodeVectorOpOperator = 2
  VisualShaderNodeVectorOpOperatorOpDiv VisualShaderNodeVectorOpOperator = 3
  VisualShaderNodeVectorOpOperatorOpMod VisualShaderNodeVectorOpOperator = 4
  VisualShaderNodeVectorOpOperatorOpPow VisualShaderNodeVectorOpOperator = 5
  VisualShaderNodeVectorOpOperatorOpMax VisualShaderNodeVectorOpOperator = 6
  VisualShaderNodeVectorOpOperatorOpMin VisualShaderNodeVectorOpOperator = 7
  VisualShaderNodeVectorOpOperatorOpCross VisualShaderNodeVectorOpOperator = 8
  VisualShaderNodeVectorOpOperatorOpAtan2 VisualShaderNodeVectorOpOperator = 9
  VisualShaderNodeVectorOpOperatorOpReflect VisualShaderNodeVectorOpOperator = 10
  VisualShaderNodeVectorOpOperatorOpStep VisualShaderNodeVectorOpOperator = 11
  VisualShaderNodeVectorOpOperatorOpEnumSize VisualShaderNodeVectorOpOperator = 12
)

func (me *VisualShaderNodeVectorOp) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVectorOp) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeVectorOp) SetOperator(op VisualShaderNodeVectorOpOperator, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeVectorOp) GetOperator()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
