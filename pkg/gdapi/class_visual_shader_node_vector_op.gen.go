// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *VisualShaderNodeVectorOp) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVectorOp) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVectorOp) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeVectorOp) SetOperator(op VisualShaderNodeVectorOpOperator, )  {
  classNameV := StringNameFromStr("VisualShaderNodeVectorOp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3371507302) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeVectorOp) GetOperator() VisualShaderNodeVectorOpOperator {
  classNameV := StringNameFromStr("VisualShaderNodeVectorOp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 11793929) // FIXME: should cache?
  var ret VisualShaderNodeVectorOpOperator
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeVectorOp) GetPropOperator() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeVectorOp) SetPropOperator(value int) {
  panic("TODO: implement")
}