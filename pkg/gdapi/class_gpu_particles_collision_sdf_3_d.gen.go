// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *GPUParticlesCollisionSDF3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GPUParticlesCollisionSDF3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesCollisionSDF3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GPUParticlesCollisionSDF3D) SetSize(size Vector3, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticlesCollisionSDF3D) GetSize() Vector3 {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticlesCollisionSDF3D) SetResolution(resolution GPUParticlesCollisionSDF3DResolution, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_resolution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1155629297) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&resolution), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticlesCollisionSDF3D) GetResolution() GPUParticlesCollisionSDF3DResolution {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resolution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2919555867) // FIXME: should cache?
  var ret GPUParticlesCollisionSDF3DResolution
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticlesCollisionSDF3D) SetTexture(texture Texture3D, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1188404210) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticlesCollisionSDF3D) GetTexture() Texture3D {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373985333) // FIXME: should cache?
  var ret Texture3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticlesCollisionSDF3D) SetThickness(thickness float32, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&thickness), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticlesCollisionSDF3D) GetThickness() float32 {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticlesCollisionSDF3D) SetBakeMask(mask int, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bake_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticlesCollisionSDF3D) GetBakeMask() int {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bake_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GPUParticlesCollisionSDF3D) SetBakeMaskValue(layer_number int, value bool, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bake_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GPUParticlesCollisionSDF3D) GetBakeMaskValue(layer_number int, ) bool {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bake_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *GPUParticlesCollisionSDF3D) GetPropSize() Vector3 {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionSDF3D) SetPropSize(value Vector3) {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionSDF3D) GetPropResolution() int {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionSDF3D) SetPropResolution(value int) {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionSDF3D) GetPropThickness() float32 {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionSDF3D) SetPropThickness(value float32) {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionSDF3D) GetPropBakeMask() int {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionSDF3D) SetPropBakeMask(value int) {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionSDF3D) GetPropTexture() Texture3D {
  panic("TODO: implement")
}

func (me *GPUParticlesCollisionSDF3D) SetPropTexture(value Texture3D) {
  panic("TODO: implement")
}