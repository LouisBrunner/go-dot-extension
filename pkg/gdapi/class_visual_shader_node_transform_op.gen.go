// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTransformOp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTransformOp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTransformOp) BaseClass() string {
  return "VisualShaderNodeTransformOp"
}

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

func  (me *VisualShaderNodeTransformOp) SetOperator(op VisualShaderNodeTransformOpOperator, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeTransformOp) GetOperator() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
