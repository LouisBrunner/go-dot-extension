// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTransformVecMult struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTransformVecMult) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTransformVecMult) BaseClass() string {
  return "VisualShaderNodeTransformVecMult"
}



// Enums

type VisualShaderNodeTransformVecMultOperator int
const (
  VisualShaderNodeTransformVecMultOperatorOpAxb VisualShaderNodeTransformVecMultOperator = 0
  VisualShaderNodeTransformVecMultOperatorOpBxa VisualShaderNodeTransformVecMultOperator = 1
  VisualShaderNodeTransformVecMultOperatorOp3X3Axb VisualShaderNodeTransformVecMultOperator = 2
  VisualShaderNodeTransformVecMultOperatorOp3X3Bxa VisualShaderNodeTransformVecMultOperator = 3
  VisualShaderNodeTransformVecMultOperatorOpMax VisualShaderNodeTransformVecMultOperator = 4
)

func (me *VisualShaderNodeTransformVecMult) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTransformVecMult) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeTransformVecMult) SetOperator(op VisualShaderNodeTransformVecMultOperator, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTransformVecMult) GetOperator()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
