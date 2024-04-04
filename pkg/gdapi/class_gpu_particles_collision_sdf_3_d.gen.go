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

type GPUParticlesCollisionSDF3D struct {
  GPUParticlesCollision3D
}

func (me *GPUParticlesCollisionSDF3D) BaseClass() string {
  return "GPUParticlesCollisionSDF3D"
}

func NewGPUParticlesCollisionSDF3D() *GPUParticlesCollisionSDF3D {
  str := StringNameFromStr("GPUParticlesCollisionSDF3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GPUParticlesCollisionSDF3D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticlesCollisionSDF3D) GetSize() Vector3 {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
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

func  (me *GPUParticlesCollisionSDF3D) SetResolution(resolution GPUParticlesCollisionSDF3DResolution, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_resolution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1155629297) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&resolution) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticlesCollisionSDF3D) GetResolution() GPUParticlesCollisionSDF3DResolution {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resolution")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2919555867) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret GPUParticlesCollisionSDF3DResolution

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *GPUParticlesCollisionSDF3D) SetTexture(texture Texture3D, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1188404210) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticlesCollisionSDF3D) GetTexture() Texture3D {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373985333) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GPUParticlesCollisionSDF3D) SetThickness(thickness float64, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&thickness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticlesCollisionSDF3D) GetThickness() float64 {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticlesCollisionSDF3D) SetBakeMask(mask int64, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bake_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticlesCollisionSDF3D) GetBakeMask() int64 {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bake_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GPUParticlesCollisionSDF3D) SetBakeMaskValue(layer_number int64, value bool, )  {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bake_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GPUParticlesCollisionSDF3D) GetBakeMaskValue(layer_number int64, ) bool {
  classNameV := StringNameFromStr("GPUParticlesCollisionSDF3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bake_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
