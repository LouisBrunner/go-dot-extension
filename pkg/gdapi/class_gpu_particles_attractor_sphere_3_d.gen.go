// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesAttractorSphere3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesAttractorSphere3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesAttractorSphere3D) BaseClass() string {
  return "GPUParticlesAttractorSphere3D"
}



// Enums

func (me *GPUParticlesAttractorSphere3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GPUParticlesAttractorSphere3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesAttractorSphere3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GPUParticlesAttractorSphere3D) SetRadius(radius float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesAttractorSphere3D) GetRadius()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
