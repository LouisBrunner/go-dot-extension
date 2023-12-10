// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeParticleRandomness struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeParticleRandomness) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeParticleRandomness) BaseClass() string {
  return "VisualShaderNodeParticleRandomness"
}



// Enums

type VisualShaderNodeParticleRandomnessOpType int
const (
  VisualShaderNodeParticleRandomnessOpTypeOpTypeScalar VisualShaderNodeParticleRandomnessOpType = 0
  VisualShaderNodeParticleRandomnessOpTypeOpTypeVector2D VisualShaderNodeParticleRandomnessOpType = 1
  VisualShaderNodeParticleRandomnessOpTypeOpTypeVector3D VisualShaderNodeParticleRandomnessOpType = 2
  VisualShaderNodeParticleRandomnessOpTypeOpTypeVector4D VisualShaderNodeParticleRandomnessOpType = 3
  VisualShaderNodeParticleRandomnessOpTypeOpTypeMax VisualShaderNodeParticleRandomnessOpType = 4
)

func (me *VisualShaderNodeParticleRandomness) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleRandomness) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleRandomness) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeParticleRandomness) SetOpType(type_ VisualShaderNodeParticleRandomnessOpType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeParticleRandomness")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2060089061) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeParticleRandomness) GetOpType() VisualShaderNodeParticleRandomnessOpType {
  classNameV := StringNameFromStr("VisualShaderNodeParticleRandomness")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3597061078) // FIXME: should cache?
  var ret VisualShaderNodeParticleRandomnessOpType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeParticleRandomness) GetPropOpType() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParticleRandomness) SetPropOpType(value int) {
  panic("TODO: implement")
}