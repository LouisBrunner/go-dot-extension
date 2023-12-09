// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *VisualShaderNodeParticleAccelerator) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleAccelerator) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeParticleAccelerator) SetMode(mode VisualShaderNodeParticleAcceleratorMode, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeParticleAccelerator) GetMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
