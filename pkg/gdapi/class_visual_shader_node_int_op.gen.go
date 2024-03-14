// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeIntOp struct {
  VisualShaderNode
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
  classNameV := StringNameFromStr("VisualShaderNodeIntOp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1677909323) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeIntOp) GetOperator() VisualShaderNodeIntOpOperator {
  classNameV := StringNameFromStr("VisualShaderNodeIntOp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1236987913) // FIXME: should cache?
  var ret VisualShaderNodeIntOpOperator
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
