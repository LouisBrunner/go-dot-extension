// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeTransformVecMult struct {
  VisualShaderNode
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

func (me *VisualShaderNodeTransformVecMult) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTransformVecMult) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTransformVecMult) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeTransformVecMult) SetOperator(op VisualShaderNodeTransformVecMultOperator, )  {
  classNameV := StringNameFromStr("VisualShaderNodeTransformVecMult")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1785665912) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeTransformVecMult) GetOperator() VisualShaderNodeTransformVecMultOperator {
  classNameV := StringNameFromStr("VisualShaderNodeTransformVecMult")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1622088722) // FIXME: should cache?
  var ret VisualShaderNodeTransformVecMultOperator
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
