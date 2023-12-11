// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeStep struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeStep) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeStep) BaseClass() string {
  return "VisualShaderNodeStep"
}



// Enums

type VisualShaderNodeStepOpType int
const (
  VisualShaderNodeStepOpTypeOpTypeScalar VisualShaderNodeStepOpType = 0
  VisualShaderNodeStepOpTypeOpTypeVector2D VisualShaderNodeStepOpType = 1
  VisualShaderNodeStepOpTypeOpTypeVector2DScalar VisualShaderNodeStepOpType = 2
  VisualShaderNodeStepOpTypeOpTypeVector3D VisualShaderNodeStepOpType = 3
  VisualShaderNodeStepOpTypeOpTypeVector3DScalar VisualShaderNodeStepOpType = 4
  VisualShaderNodeStepOpTypeOpTypeVector4D VisualShaderNodeStepOpType = 5
  VisualShaderNodeStepOpTypeOpTypeVector4DScalar VisualShaderNodeStepOpType = 6
  VisualShaderNodeStepOpTypeOpTypeMax VisualShaderNodeStepOpType = 7
)

func (me *VisualShaderNodeStep) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeStep) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeStep) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeStep) SetOpType(op_type VisualShaderNodeStepOpType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeStep")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 715172489) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeStep) GetOpType() VisualShaderNodeStepOpType {
  classNameV := StringNameFromStr("VisualShaderNodeStep")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3274022781) // FIXME: should cache?
  var ret VisualShaderNodeStepOpType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
