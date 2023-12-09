// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *VisualShaderNodeParticleEmitter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleEmitter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeParticleEmitter) SetMode2D(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeParticleEmitter) IsMode2D()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
