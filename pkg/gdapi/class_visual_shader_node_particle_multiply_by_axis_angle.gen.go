// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParticleMultiplyByAxisAngle struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) BaseClass() string {
  return "VisualShaderNodeParticleMultiplyByAxisAngle"
}



// Enums

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeParticleMultiplyByAxisAngle) SetDegreesMode(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeParticleMultiplyByAxisAngle) IsDegreesMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
