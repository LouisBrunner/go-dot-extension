// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *VisualShaderNodeColorOp) GetOperator()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
