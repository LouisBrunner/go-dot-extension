// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeClamp struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeClamp) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeClamp) BaseClass() string {
  return "VisualShaderNodeClamp"
}



// Enums

type VisualShaderNodeClampOpType int
const (
  VisualShaderNodeClampOpTypeOpTypeFloat VisualShaderNodeClampOpType = 0
  VisualShaderNodeClampOpTypeOpTypeInt VisualShaderNodeClampOpType = 1
  VisualShaderNodeClampOpTypeOpTypeUint VisualShaderNodeClampOpType = 2
  VisualShaderNodeClampOpTypeOpTypeVector2D VisualShaderNodeClampOpType = 3
  VisualShaderNodeClampOpTypeOpTypeVector3D VisualShaderNodeClampOpType = 4
  VisualShaderNodeClampOpTypeOpTypeVector4D VisualShaderNodeClampOpType = 5
  VisualShaderNodeClampOpTypeOpTypeMax VisualShaderNodeClampOpType = 6
)

func (me *VisualShaderNodeClamp) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeClamp) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeClamp) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeClamp) SetOpType(op_type VisualShaderNodeClampOpType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeClamp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 405010749) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeClamp) GetOpType() VisualShaderNodeClampOpType {
  classNameV := StringNameFromStr("VisualShaderNodeClamp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 233276050) // FIXME: should cache?
  var ret VisualShaderNodeClampOpType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeClamp) GetPropOpType() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeClamp) SetPropOpType(value int) {
  panic("TODO: implement")
}