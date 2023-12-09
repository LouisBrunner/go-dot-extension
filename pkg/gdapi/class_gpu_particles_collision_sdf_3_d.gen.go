// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesCollisionSDF3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesCollisionSDF3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesCollisionSDF3D) BaseClass() string {
  return "GPUParticlesCollisionSDF3D"
}



// Enums

type GPUParticlesCollisionSDF3DResolution int
const (
  GPUParticlesCollisionSDF3DResolutionResolution16 GPUParticlesCollisionSDF3DResolution = 0
  GPUParticlesCollisionSDF3DResolutionResolution32 GPUParticlesCollisionSDF3DResolution = 1
  GPUParticlesCollisionSDF3DResolutionResolution64 GPUParticlesCollisionSDF3DResolution = 2
  GPUParticlesCollisionSDF3DResolutionResolution128 GPUParticlesCollisionSDF3DResolution = 3
  GPUParticlesCollisionSDF3DResolutionResolution256 GPUParticlesCollisionSDF3DResolution = 4
  GPUParticlesCollisionSDF3DResolutionResolution512 GPUParticlesCollisionSDF3DResolution = 5
  GPUParticlesCollisionSDF3DResolutionResolutionMax GPUParticlesCollisionSDF3DResolution = 6
)

func (me *GPUParticlesCollisionSDF3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesCollisionSDF3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GPUParticlesCollisionSDF3D) SetSize(size Vector3, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionSDF3D) GetSize()  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionSDF3D) SetResolution(resolution GPUParticlesCollisionSDF3DResolution, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionSDF3D) GetResolution()  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionSDF3D) SetTexture(texture Texture3D, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionSDF3D) GetTexture()  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionSDF3D) SetThickness(thickness float32, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionSDF3D) GetThickness()  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionSDF3D) SetBakeMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionSDF3D) GetBakeMask()  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionSDF3D) SetBakeMaskValue(layer_number int, value bool, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionSDF3D) GetBakeMaskValue(layer_number int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
