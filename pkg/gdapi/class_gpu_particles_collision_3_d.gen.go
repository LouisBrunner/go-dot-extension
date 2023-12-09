// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesCollision3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesCollision3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesCollision3D) BaseClass() string {
  return "GPUParticlesCollision3D"
}



// Enums

func (me *GPUParticlesCollision3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GPUParticlesCollision3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesCollision3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GPUParticlesCollision3D) SetCullMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollision3D) GetCullMask()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
