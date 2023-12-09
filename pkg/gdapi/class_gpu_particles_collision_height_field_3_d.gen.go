// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesCollisionHeightField3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesCollisionHeightField3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesCollisionHeightField3D) BaseClass() string {
  return "GPUParticlesCollisionHeightField3D"
}



// Enums

type GPUParticlesCollisionHeightField3DResolution int
const (
  GPUParticlesCollisionHeightField3DResolutionResolution256 GPUParticlesCollisionHeightField3DResolution = 0
  GPUParticlesCollisionHeightField3DResolutionResolution512 GPUParticlesCollisionHeightField3DResolution = 1
  GPUParticlesCollisionHeightField3DResolutionResolution1024 GPUParticlesCollisionHeightField3DResolution = 2
  GPUParticlesCollisionHeightField3DResolutionResolution2048 GPUParticlesCollisionHeightField3DResolution = 3
  GPUParticlesCollisionHeightField3DResolutionResolution4096 GPUParticlesCollisionHeightField3DResolution = 4
  GPUParticlesCollisionHeightField3DResolutionResolution8192 GPUParticlesCollisionHeightField3DResolution = 5
  GPUParticlesCollisionHeightField3DResolutionResolutionMax GPUParticlesCollisionHeightField3DResolution = 6
)

type GPUParticlesCollisionHeightField3DUpdateMode int
const (
  GPUParticlesCollisionHeightField3DUpdateModeUpdateModeWhenMoved GPUParticlesCollisionHeightField3DUpdateMode = 0
  GPUParticlesCollisionHeightField3DUpdateModeUpdateModeAlways GPUParticlesCollisionHeightField3DUpdateMode = 1
)

func (me *GPUParticlesCollisionHeightField3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GPUParticlesCollisionHeightField3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesCollisionHeightField3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GPUParticlesCollisionHeightField3D) SetSize(size Vector3, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionHeightField3D) GetSize()  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionHeightField3D) SetResolution(resolution GPUParticlesCollisionHeightField3DResolution, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionHeightField3D) GetResolution()  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionHeightField3D) SetUpdateMode(update_mode GPUParticlesCollisionHeightField3DUpdateMode, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionHeightField3D) GetUpdateMode()  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionHeightField3D) SetFollowCameraEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionHeightField3D) IsFollowCameraEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
