// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type VisualShaderNodeSmoothStep struct {
  VisualShaderNode
}

func (me *VisualShaderNodeSmoothStep) BaseClass() string {
  return "VisualShaderNodeSmoothStep"
}

func NewVisualShaderNodeSmoothStep() *VisualShaderNodeSmoothStep {
  str := StringNameFromStr("VisualShaderNodeSmoothStep") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeSmoothStep{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeSmoothStep) GetOpType() VisualShaderNodeSmoothStepOpType {
  classNameV := StringNameFromStr("VisualShaderNodeSmoothStep")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 359640855) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeSmoothStepOpType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
