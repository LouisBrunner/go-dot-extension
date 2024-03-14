// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeUIntOp struct {
  VisualShaderNode
}

func (me *VisualShaderNodeUIntOp) BaseClass() string {
  return "VisualShaderNodeUIntOp"
}



// Enums

type VisualShaderNodeUIntOpOperator int
const (
  VisualShaderNodeUIntOpOperatorOpAdd VisualShaderNodeUIntOpOperator = 0
  VisualShaderNodeUIntOpOperatorOpSub VisualShaderNodeUIntOpOperator = 1
  VisualShaderNodeUIntOpOperatorOpMul VisualShaderNodeUIntOpOperator = 2
  VisualShaderNodeUIntOpOperatorOpDiv VisualShaderNodeUIntOpOperator = 3
  VisualShaderNodeUIntOpOperatorOpMod VisualShaderNodeUIntOpOperator = 4
  VisualShaderNodeUIntOpOperatorOpMax VisualShaderNodeUIntOpOperator = 5
  VisualShaderNodeUIntOpOperatorOpMin VisualShaderNodeUIntOpOperator = 6
  VisualShaderNodeUIntOpOperatorOpBitwiseAnd VisualShaderNodeUIntOpOperator = 7
  VisualShaderNodeUIntOpOperatorOpBitwiseOr VisualShaderNodeUIntOpOperator = 8
  VisualShaderNodeUIntOpOperatorOpBitwiseXor VisualShaderNodeUIntOpOperator = 9
  VisualShaderNodeUIntOpOperatorOpBitwiseLeftShift VisualShaderNodeUIntOpOperator = 10
  VisualShaderNodeUIntOpOperatorOpBitwiseRightShift VisualShaderNodeUIntOpOperator = 11
  VisualShaderNodeUIntOpOperatorOpEnumSize VisualShaderNodeUIntOpOperator = 12
)

func (me *VisualShaderNodeUIntOp) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeUIntOp) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeUIntOp) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeUIntOp) SetOperator(op VisualShaderNodeUIntOpOperator, )  {
  classNameV := StringNameFromStr("VisualShaderNodeUIntOp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3463048345) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeUIntOp) GetOperator() VisualShaderNodeUIntOpOperator {
  classNameV := StringNameFromStr("VisualShaderNodeUIntOp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 256631461) // FIXME: should cache?
  var ret VisualShaderNodeUIntOpOperator
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
