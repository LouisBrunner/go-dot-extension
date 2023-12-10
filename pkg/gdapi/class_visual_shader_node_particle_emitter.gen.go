// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeParticleEmitter struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeParticleEmitter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeParticleEmitter) BaseClass() string {
  return "VisualShaderNodeParticleEmitter"
}



// Enums

func (me *VisualShaderNodeParticleEmitter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleEmitter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleEmitter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeParticleEmitter) SetMode2D(enabled bool, )  {
  classNameV := StringNameFromStr("VisualShaderNodeParticleEmitter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mode_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeParticleEmitter) IsMode2D() bool {
  classNameV := StringNameFromStr("VisualShaderNodeParticleEmitter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_mode_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeParticleEmitter) GetPropMode2D() bool {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParticleEmitter) SetPropMode2D(value bool) {
  panic("TODO: implement")
}