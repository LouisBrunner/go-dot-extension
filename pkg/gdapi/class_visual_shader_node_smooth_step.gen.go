// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeSmoothStep struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeSmoothStep) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeSmoothStep) BaseClass() string {
  return "VisualShaderNodeSmoothStep"
}



// Enums

type VisualShaderNodeSmoothStepOpType int
const (
  VisualShaderNodeSmoothStepOpTypeOpTypeScalar VisualShaderNodeSmoothStepOpType = 0
  VisualShaderNodeSmoothStepOpTypeOpTypeVector2D VisualShaderNodeSmoothStepOpType = 1
  VisualShaderNodeSmoothStepOpTypeOpTypeVector2DScalar VisualShaderNodeSmoothStepOpType = 2
  VisualShaderNodeSmoothStepOpTypeOpTypeVector3D VisualShaderNodeSmoothStepOpType = 3
  VisualShaderNodeSmoothStepOpTypeOpTypeVector3DScalar VisualShaderNodeSmoothStepOpType = 4
  VisualShaderNodeSmoothStepOpTypeOpTypeVector4D VisualShaderNodeSmoothStepOpType = 5
  VisualShaderNodeSmoothStepOpTypeOpTypeVector4DScalar VisualShaderNodeSmoothStepOpType = 6
  VisualShaderNodeSmoothStepOpTypeOpTypeMax VisualShaderNodeSmoothStepOpType = 7
)

func (me *VisualShaderNodeSmoothStep) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeSmoothStep) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeSmoothStep) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeSmoothStep) SetOpType(op_type VisualShaderNodeSmoothStepOpType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeSmoothStep")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2427426148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeSmoothStep) GetOpType() VisualShaderNodeSmoothStepOpType {
  classNameV := StringNameFromStr("VisualShaderNodeSmoothStep")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 359640855) // FIXME: should cache?
  var ret VisualShaderNodeSmoothStepOpType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeSmoothStep) GetPropOpType() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeSmoothStep) SetPropOpType(value int) {
  panic("TODO: implement")
}