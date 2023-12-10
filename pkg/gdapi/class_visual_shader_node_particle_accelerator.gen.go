// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeParticleAccelerator struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeParticleAccelerator) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeParticleAccelerator) BaseClass() string {
  return "VisualShaderNodeParticleAccelerator"
}



// Enums

type VisualShaderNodeParticleAcceleratorMode int
const (
  VisualShaderNodeParticleAcceleratorModeModeLinear VisualShaderNodeParticleAcceleratorMode = 0
  VisualShaderNodeParticleAcceleratorModeModeRadial VisualShaderNodeParticleAcceleratorMode = 1
  VisualShaderNodeParticleAcceleratorModeModeTangential VisualShaderNodeParticleAcceleratorMode = 2
  VisualShaderNodeParticleAcceleratorModeModeMax VisualShaderNodeParticleAcceleratorMode = 3
)

func (me *VisualShaderNodeParticleAccelerator) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleAccelerator) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleAccelerator) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeParticleAccelerator) SetMode(mode VisualShaderNodeParticleAcceleratorMode, )  {
  classNameV := StringNameFromStr("VisualShaderNodeParticleAccelerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457585749) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeParticleAccelerator) GetMode() VisualShaderNodeParticleAcceleratorMode {
  classNameV := StringNameFromStr("VisualShaderNodeParticleAccelerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2660365633) // FIXME: should cache?
  var ret VisualShaderNodeParticleAcceleratorMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeParticleAccelerator) GetPropMode() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParticleAccelerator) SetPropMode(value int) {
  panic("TODO: implement")
}