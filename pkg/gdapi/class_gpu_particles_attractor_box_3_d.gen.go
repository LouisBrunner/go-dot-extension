// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesAttractorBox3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesAttractorBox3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesAttractorBox3D) BaseClass() string {
  return "GPUParticlesAttractorBox3D"
}



// Enums

func (me *GPUParticlesAttractorBox3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesAttractorBox3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GPUParticlesAttractorBox3D) SetSize(size Vector3, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesAttractorBox3D) GetSize()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
