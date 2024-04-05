// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForGPUParticlesCollisionSphere3DList struct {
  fnSetRadius gdc.MethodBindPtr
  fnGetRadius gdc.MethodBindPtr
}

var ptrsForGPUParticlesCollisionSphere3D ptrsForGPUParticlesCollisionSphere3DList

func initGPUParticlesCollisionSphere3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("GPUParticlesCollisionSphere3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_radius")
    defer methodName.Destroy()
    ptrsForGPUParticlesCollisionSphere3D.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_radius")
    defer methodName.Destroy()
    ptrsForGPUParticlesCollisionSphere3D.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

type GPUParticlesCollisionSphere3D struct {
  GPUParticlesCollision3D
}

func (me *GPUParticlesCollisionSphere3D) BaseClass() string {
  return "GPUParticlesCollisionSphere3D"
}

func NewGPUParticlesCollisionSphere3D() *GPUParticlesCollisionSphere3D {
  str := StringNameFromStr("GPUParticlesCollisionSphere3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GPUParticlesCollisionSphere3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *GPUParticlesCollisionSphere3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GPUParticlesCollisionSphere3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesCollisionSphere3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GPUParticlesCollisionSphere3D) SetRadius(radius float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionSphere3D.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticlesCollisionSphere3D) GetRadius() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionSphere3D.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
