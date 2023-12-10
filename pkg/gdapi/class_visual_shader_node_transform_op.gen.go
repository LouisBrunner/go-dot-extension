// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeTransformOp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTransformOp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTransformOp) BaseClass() string {
  return "VisualShaderNodeTransformOp"
}



// Enums

type VisualShaderNodeTransformOpOperator int
const (
  VisualShaderNodeTransformOpOperatorOpAxb VisualShaderNodeTransformOpOperator = 0
  VisualShaderNodeTransformOpOperatorOpBxa VisualShaderNodeTransformOpOperator = 1
  VisualShaderNodeTransformOpOperatorOpAxbComp VisualShaderNodeTransformOpOperator = 2
  VisualShaderNodeTransformOpOperatorOpBxaComp VisualShaderNodeTransformOpOperator = 3
  VisualShaderNodeTransformOpOperatorOpAdd VisualShaderNodeTransformOpOperator = 4
  VisualShaderNodeTransformOpOperatorOpAMinusB VisualShaderNodeTransformOpOperator = 5
  VisualShaderNodeTransformOpOperatorOpBMinusA VisualShaderNodeTransformOpOperator = 6
  VisualShaderNodeTransformOpOperatorOpADivB VisualShaderNodeTransformOpOperator = 7
  VisualShaderNodeTransformOpOperatorOpBDivA VisualShaderNodeTransformOpOperator = 8
  VisualShaderNodeTransformOpOperatorOpMax VisualShaderNodeTransformOpOperator = 9
)

func (me *VisualShaderNodeTransformOp) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTransformOp) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTransformOp) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeTransformOp) SetOperator(op VisualShaderNodeTransformOpOperator, )  {
  classNameV := StringNameFromStr("VisualShaderNodeTransformOp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2287310733) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeTransformOp) GetOperator() VisualShaderNodeTransformOpOperator {
  classNameV := StringNameFromStr("VisualShaderNodeTransformOp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1238663601) // FIXME: should cache?
  var ret VisualShaderNodeTransformOpOperator
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeTransformOp) GetPropOperator() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeTransformOp) SetPropOperator(value int) {
  panic("TODO: implement")
}