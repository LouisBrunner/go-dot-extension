// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeParticleEmit struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeParticleEmit) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeParticleEmit) BaseClass() string {
  return "VisualShaderNodeParticleEmit"
}



// Enums

type VisualShaderNodeParticleEmitEmitFlags int
const (
  VisualShaderNodeParticleEmitEmitFlagsEmitFlagPosition VisualShaderNodeParticleEmitEmitFlags = 1
  VisualShaderNodeParticleEmitEmitFlagsEmitFlagRotScale VisualShaderNodeParticleEmitEmitFlags = 2
  VisualShaderNodeParticleEmitEmitFlagsEmitFlagVelocity VisualShaderNodeParticleEmitEmitFlags = 4
  VisualShaderNodeParticleEmitEmitFlagsEmitFlagColor VisualShaderNodeParticleEmitEmitFlags = 8
  VisualShaderNodeParticleEmitEmitFlagsEmitFlagCustom VisualShaderNodeParticleEmitEmitFlags = 16
)

func (me *VisualShaderNodeParticleEmit) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleEmit) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleEmit) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeParticleEmit) SetFlags(flags VisualShaderNodeParticleEmitEmitFlags, )  {
  classNameV := StringNameFromStr("VisualShaderNodeParticleEmit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3960756792) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeParticleEmit) GetFlags() VisualShaderNodeParticleEmitEmitFlags {
  classNameV := StringNameFromStr("VisualShaderNodeParticleEmit")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 171277835) // FIXME: should cache?
  var ret VisualShaderNodeParticleEmitEmitFlags
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeParticleEmit) GetPropFlags() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParticleEmit) SetPropFlags(value int) {
  panic("TODO: implement")
}