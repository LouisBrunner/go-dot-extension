// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeVectorBase struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeVectorBase) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeVectorBase) BaseClass() string {
  return "VisualShaderNodeVectorBase"
}



// Enums

type VisualShaderNodeVectorBaseOpType int
const (
  VisualShaderNodeVectorBaseOpTypeOpTypeVector2D VisualShaderNodeVectorBaseOpType = 0
  VisualShaderNodeVectorBaseOpTypeOpTypeVector3D VisualShaderNodeVectorBaseOpType = 1
  VisualShaderNodeVectorBaseOpTypeOpTypeVector4D VisualShaderNodeVectorBaseOpType = 2
  VisualShaderNodeVectorBaseOpTypeOpTypeMax VisualShaderNodeVectorBaseOpType = 3
)

func (me *VisualShaderNodeVectorBase) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVectorBase) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVectorBase) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeVectorBase) SetOpType(type_ VisualShaderNodeVectorBaseOpType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeVectorBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1692596998) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeVectorBase) GetOpType() VisualShaderNodeVectorBaseOpType {
  classNameV := StringNameFromStr("VisualShaderNodeVectorBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2568738462) // FIXME: should cache?
  var ret VisualShaderNodeVectorBaseOpType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeVectorBase) GetPropOpType() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeVectorBase) SetPropOpType(value int) {
  panic("TODO: implement")
}