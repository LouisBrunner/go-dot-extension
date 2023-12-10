// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeColorOp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeColorOp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeColorOp) BaseClass() string {
  return "VisualShaderNodeColorOp"
}



// Enums

type VisualShaderNodeColorOpOperator int
const (
  VisualShaderNodeColorOpOperatorOpScreen VisualShaderNodeColorOpOperator = 0
  VisualShaderNodeColorOpOperatorOpDifference VisualShaderNodeColorOpOperator = 1
  VisualShaderNodeColorOpOperatorOpDarken VisualShaderNodeColorOpOperator = 2
  VisualShaderNodeColorOpOperatorOpLighten VisualShaderNodeColorOpOperator = 3
  VisualShaderNodeColorOpOperatorOpOverlay VisualShaderNodeColorOpOperator = 4
  VisualShaderNodeColorOpOperatorOpDodge VisualShaderNodeColorOpOperator = 5
  VisualShaderNodeColorOpOperatorOpBurn VisualShaderNodeColorOpOperator = 6
  VisualShaderNodeColorOpOperatorOpSoftLight VisualShaderNodeColorOpOperator = 7
  VisualShaderNodeColorOpOperatorOpHardLight VisualShaderNodeColorOpOperator = 8
  VisualShaderNodeColorOpOperatorOpMax VisualShaderNodeColorOpOperator = 9
)

func (me *VisualShaderNodeColorOp) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeColorOp) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeColorOp) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeColorOp) SetOperator(op VisualShaderNodeColorOpOperator, )  {
  classNameV := StringNameFromStr("VisualShaderNodeColorOp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4260370673) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeColorOp) GetOperator() VisualShaderNodeColorOpOperator {
  classNameV := StringNameFromStr("VisualShaderNodeColorOp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1950956529) // FIXME: should cache?
  var ret VisualShaderNodeColorOpOperator
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeColorOp) GetPropOperator() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeColorOp) SetPropOperator(value int) {
  panic("TODO: implement")
}