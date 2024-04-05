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

type ptrsForGPUParticlesAttractorVectorField3DList struct {
  fnSetSize gdc.MethodBindPtr
  fnGetSize gdc.MethodBindPtr
  fnSetTexture gdc.MethodBindPtr
  fnGetTexture gdc.MethodBindPtr
}

var ptrsForGPUParticlesAttractorVectorField3D ptrsForGPUParticlesAttractorVectorField3DList

func initGPUParticlesAttractorVectorField3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("GPUParticlesAttractorVectorField3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_size")
    defer methodName.Destroy()
    ptrsForGPUParticlesAttractorVectorField3D.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_size")
    defer methodName.Destroy()
    ptrsForGPUParticlesAttractorVectorField3D.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_texture")
    defer methodName.Destroy()
    ptrsForGPUParticlesAttractorVectorField3D.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1188404210))
  }
  {
    methodName := StringNameFromStr("get_texture")
    defer methodName.Destroy()
    ptrsForGPUParticlesAttractorVectorField3D.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373985333))
  }
}

type GPUParticlesAttractorVectorField3D struct {
  GPUParticlesAttractor3D
}

func (me *GPUParticlesAttractorVectorField3D) BaseClass() string {
  return "GPUParticlesAttractorVectorField3D"
}

func NewGPUParticlesAttractorVectorField3D() *GPUParticlesAttractorVectorField3D {
  str := StringNameFromStr("GPUParticlesAttractorVectorField3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GPUParticlesAttractorVectorField3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *GPUParticlesAttractorVectorField3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GPUParticlesAttractorVectorField3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesAttractorVectorField3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GPUParticlesAttractorVectorField3D) SetSize(size Vector3, )  {
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractorVectorField3D.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticlesAttractorVectorField3D) GetSize() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractorVectorField3D.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GPUParticlesAttractorVectorField3D) SetTexture(texture Texture3D, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractorVectorField3D.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticlesAttractorVectorField3D) GetTexture() Texture3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractorVectorField3D.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
