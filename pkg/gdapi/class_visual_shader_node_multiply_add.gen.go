// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeMultiplyAdd struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeMultiplyAdd) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeMultiplyAdd) BaseClass() string {
  return "VisualShaderNodeMultiplyAdd"
}



// Enums

type VisualShaderNodeMultiplyAddOpType int
const (
  VisualShaderNodeMultiplyAddOpTypeOpTypeScalar VisualShaderNodeMultiplyAddOpType = 0
  VisualShaderNodeMultiplyAddOpTypeOpTypeVector2D VisualShaderNodeMultiplyAddOpType = 1
  VisualShaderNodeMultiplyAddOpTypeOpTypeVector3D VisualShaderNodeMultiplyAddOpType = 2
  VisualShaderNodeMultiplyAddOpTypeOpTypeVector4D VisualShaderNodeMultiplyAddOpType = 3
  VisualShaderNodeMultiplyAddOpTypeOpTypeMax VisualShaderNodeMultiplyAddOpType = 4
)

func (me *VisualShaderNodeMultiplyAdd) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeMultiplyAdd) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeMultiplyAdd) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeMultiplyAdd) SetOpType(type_ VisualShaderNodeMultiplyAddOpType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeMultiplyAdd")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1409862380) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeMultiplyAdd) GetOpType() VisualShaderNodeMultiplyAddOpType {
  classNameV := StringNameFromStr("VisualShaderNodeMultiplyAdd")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2823201991) // FIXME: should cache?
  var ret VisualShaderNodeMultiplyAddOpType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
