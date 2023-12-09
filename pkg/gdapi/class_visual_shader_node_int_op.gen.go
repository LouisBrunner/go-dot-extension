// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeIntOp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeIntOp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeIntOp) BaseClass() string {
  return "VisualShaderNodeIntOp"
}



// Enums

type VisualShaderNodeIntOpOperator int
const (
  VisualShaderNodeIntOpOperatorOpAdd VisualShaderNodeIntOpOperator = 0
  VisualShaderNodeIntOpOperatorOpSub VisualShaderNodeIntOpOperator = 1
  VisualShaderNodeIntOpOperatorOpMul VisualShaderNodeIntOpOperator = 2
  VisualShaderNodeIntOpOperatorOpDiv VisualShaderNodeIntOpOperator = 3
  VisualShaderNodeIntOpOperatorOpMod VisualShaderNodeIntOpOperator = 4
  VisualShaderNodeIntOpOperatorOpMax VisualShaderNodeIntOpOperator = 5
  VisualShaderNodeIntOpOperatorOpMin VisualShaderNodeIntOpOperator = 6
  VisualShaderNodeIntOpOperatorOpBitwiseAnd VisualShaderNodeIntOpOperator = 7
  VisualShaderNodeIntOpOperatorOpBitwiseOr VisualShaderNodeIntOpOperator = 8
  VisualShaderNodeIntOpOperatorOpBitwiseXor VisualShaderNodeIntOpOperator = 9
  VisualShaderNodeIntOpOperatorOpBitwiseLeftShift VisualShaderNodeIntOpOperator = 10
  VisualShaderNodeIntOpOperatorOpBitwiseRightShift VisualShaderNodeIntOpOperator = 11
  VisualShaderNodeIntOpOperatorOpEnumSize VisualShaderNodeIntOpOperator = 12
)

func (me *VisualShaderNodeIntOp) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeIntOp) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeIntOp) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeIntOp) SetOperator(op VisualShaderNodeIntOpOperator, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeIntOp) GetOperator()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
