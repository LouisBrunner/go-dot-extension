// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeParticleMultiplyByAxisAngle struct {
  VisualShaderNode
}

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) BaseClass() string {
  return "VisualShaderNodeParticleMultiplyByAxisAngle"
}



// Enums

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeParticleMultiplyByAxisAngle) SetDegreesMode(enabled bool, )  {
  classNameV := StringNameFromStr("VisualShaderNodeParticleMultiplyByAxisAngle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_degrees_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeParticleMultiplyByAxisAngle) IsDegreesMode() bool {
  classNameV := StringNameFromStr("VisualShaderNodeParticleMultiplyByAxisAngle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_degrees_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
