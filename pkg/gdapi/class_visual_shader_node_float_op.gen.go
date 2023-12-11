// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeFloatOp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeFloatOp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeFloatOp) BaseClass() string {
  return "VisualShaderNodeFloatOp"
}



// Enums

type VisualShaderNodeFloatOpOperator int
const (
  VisualShaderNodeFloatOpOperatorOpAdd VisualShaderNodeFloatOpOperator = 0
  VisualShaderNodeFloatOpOperatorOpSub VisualShaderNodeFloatOpOperator = 1
  VisualShaderNodeFloatOpOperatorOpMul VisualShaderNodeFloatOpOperator = 2
  VisualShaderNodeFloatOpOperatorOpDiv VisualShaderNodeFloatOpOperator = 3
  VisualShaderNodeFloatOpOperatorOpMod VisualShaderNodeFloatOpOperator = 4
  VisualShaderNodeFloatOpOperatorOpPow VisualShaderNodeFloatOpOperator = 5
  VisualShaderNodeFloatOpOperatorOpMax VisualShaderNodeFloatOpOperator = 6
  VisualShaderNodeFloatOpOperatorOpMin VisualShaderNodeFloatOpOperator = 7
  VisualShaderNodeFloatOpOperatorOpAtan2 VisualShaderNodeFloatOpOperator = 8
  VisualShaderNodeFloatOpOperatorOpStep VisualShaderNodeFloatOpOperator = 9
  VisualShaderNodeFloatOpOperatorOpEnumSize VisualShaderNodeFloatOpOperator = 10
)

func (me *VisualShaderNodeFloatOp) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeFloatOp) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeFloatOp) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeFloatOp) SetOperator(op VisualShaderNodeFloatOpOperator, )  {
  classNameV := StringNameFromStr("VisualShaderNodeFloatOp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2488468047) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeFloatOp) GetOperator() VisualShaderNodeFloatOpOperator {
  classNameV := StringNameFromStr("VisualShaderNodeFloatOp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1867979390) // FIXME: should cache?
  var ret VisualShaderNodeFloatOpOperator
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
