// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeParticleEmit struct {
  VisualShaderNode
}

func (me *VisualShaderNodeParticleEmit) BaseClass() string {
  return "VisualShaderNodeParticleEmit"
}

func NewVisualShaderNodeParticleEmit() *VisualShaderNodeParticleEmit {
  str := StringNameFromStr("VisualShaderNodeParticleEmit") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeParticleEmit{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  var ret VisualShaderNodeParticleEmitEmitFlags

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
