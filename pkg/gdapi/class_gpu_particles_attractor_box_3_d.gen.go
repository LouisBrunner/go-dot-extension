// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type GPUParticlesAttractorBox3D struct {
  GPUParticlesAttractor3D
}

func (me *GPUParticlesAttractorBox3D) BaseClass() string {
  return "GPUParticlesAttractorBox3D"
}

func NewGPUParticlesAttractorBox3D() *GPUParticlesAttractorBox3D {
  str := StringNameFromStr("GPUParticlesAttractorBox3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GPUParticlesAttractorBox3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *GPUParticlesAttractorBox3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GPUParticlesAttractorBox3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesAttractorBox3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GPUParticlesAttractorBox3D) SetSize(size Vector3, )  {
  classNameV := StringNameFromStr("GPUParticlesAttractorBox3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticlesAttractorBox3D) GetSize() Vector3 {
  classNameV := StringNameFromStr("GPUParticlesAttractorBox3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
