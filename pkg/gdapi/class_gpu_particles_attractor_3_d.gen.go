// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesAttractor3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesAttractor3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesAttractor3D) BaseClass() string {
  return "GPUParticlesAttractor3D"
}



// Enums

func (me *GPUParticlesAttractor3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GPUParticlesAttractor3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesAttractor3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GPUParticlesAttractor3D) SetCullMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesAttractor3D) GetCullMask()  {
  panic("TODO: implement")
}

func  (me *GPUParticlesAttractor3D) SetStrength(strength float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesAttractor3D) GetStrength()  {
  panic("TODO: implement")
}

func  (me *GPUParticlesAttractor3D) SetAttenuation(attenuation float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesAttractor3D) GetAttenuation()  {
  panic("TODO: implement")
}

func  (me *GPUParticlesAttractor3D) SetDirectionality(amount float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesAttractor3D) GetDirectionality()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
