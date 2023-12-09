// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GPUParticlesCollisionBox3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesCollisionBox3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesCollisionBox3D) BaseClass() string {
  return "GPUParticlesCollisionBox3D"
}



// Enums

func (me *GPUParticlesCollisionBox3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GPUParticlesCollisionBox3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesCollisionBox3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GPUParticlesCollisionBox3D) SetSize(size Vector3, )  {
  panic("TODO: implement")
}

func  (me *GPUParticlesCollisionBox3D) GetSize()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
