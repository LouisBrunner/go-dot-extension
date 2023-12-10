// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GPUParticlesAttractorVectorField3D struct {
  obj gdc.ObjectPtr
}

func (me *GPUParticlesAttractorVectorField3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GPUParticlesAttractorVectorField3D) BaseClass() string {
  return "GPUParticlesAttractorVectorField3D"
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
  classNameV := StringNameFromStr("GPUParticlesAttractorVectorField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticlesAttractorVectorField3D) GetSize() Vector3 {
  classNameV := StringNameFromStr("GPUParticlesAttractorVectorField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticlesAttractorVectorField3D) SetTexture(texture Texture3D, )  {
  classNameV := StringNameFromStr("GPUParticlesAttractorVectorField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1188404210) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticlesAttractorVectorField3D) GetTexture() Texture3D {
  classNameV := StringNameFromStr("GPUParticlesAttractorVectorField3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373985333) // FIXME: should cache?
  var ret Texture3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *GPUParticlesAttractorVectorField3D) GetPropSize() Vector3 {
  panic("TODO: implement")
}

func (me *GPUParticlesAttractorVectorField3D) SetPropSize(value Vector3) {
  panic("TODO: implement")
}

func (me *GPUParticlesAttractorVectorField3D) GetPropTexture() Texture3D {
  panic("TODO: implement")
}

func (me *GPUParticlesAttractorVectorField3D) SetPropTexture(value Texture3D) {
  panic("TODO: implement")
}