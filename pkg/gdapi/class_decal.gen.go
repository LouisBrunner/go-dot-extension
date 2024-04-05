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

type ptrsForDecalList struct {
  fnSetSize gdc.MethodBindPtr
  fnGetSize gdc.MethodBindPtr
  fnSetTexture gdc.MethodBindPtr
  fnGetTexture gdc.MethodBindPtr
  fnSetEmissionEnergy gdc.MethodBindPtr
  fnGetEmissionEnergy gdc.MethodBindPtr
  fnSetAlbedoMix gdc.MethodBindPtr
  fnGetAlbedoMix gdc.MethodBindPtr
  fnSetModulate gdc.MethodBindPtr
  fnGetModulate gdc.MethodBindPtr
  fnSetUpperFade gdc.MethodBindPtr
  fnGetUpperFade gdc.MethodBindPtr
  fnSetLowerFade gdc.MethodBindPtr
  fnGetLowerFade gdc.MethodBindPtr
  fnSetNormalFade gdc.MethodBindPtr
  fnGetNormalFade gdc.MethodBindPtr
  fnSetEnableDistanceFade gdc.MethodBindPtr
  fnIsDistanceFadeEnabled gdc.MethodBindPtr
  fnSetDistanceFadeBegin gdc.MethodBindPtr
  fnGetDistanceFadeBegin gdc.MethodBindPtr
  fnSetDistanceFadeLength gdc.MethodBindPtr
  fnGetDistanceFadeLength gdc.MethodBindPtr
  fnSetCullMask gdc.MethodBindPtr
  fnGetCullMask gdc.MethodBindPtr
}

var ptrsForDecal ptrsForDecalList

func initDecalPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Decal")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_size")
    defer methodName.Destroy()
    ptrsForDecal.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_size")
    defer methodName.Destroy()
    ptrsForDecal.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_texture")
    defer methodName.Destroy()
    ptrsForDecal.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2086764391))
  }
  {
    methodName := StringNameFromStr("get_texture")
    defer methodName.Destroy()
    ptrsForDecal.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3244119503))
  }
  {
    methodName := StringNameFromStr("set_emission_energy")
    defer methodName.Destroy()
    ptrsForDecal.fnSetEmissionEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_emission_energy")
    defer methodName.Destroy()
    ptrsForDecal.fnGetEmissionEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_albedo_mix")
    defer methodName.Destroy()
    ptrsForDecal.fnSetAlbedoMix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_albedo_mix")
    defer methodName.Destroy()
    ptrsForDecal.fnGetAlbedoMix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_modulate")
    defer methodName.Destroy()
    ptrsForDecal.fnSetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_modulate")
    defer methodName.Destroy()
    ptrsForDecal.fnGetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_upper_fade")
    defer methodName.Destroy()
    ptrsForDecal.fnSetUpperFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_upper_fade")
    defer methodName.Destroy()
    ptrsForDecal.fnGetUpperFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_lower_fade")
    defer methodName.Destroy()
    ptrsForDecal.fnSetLowerFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_lower_fade")
    defer methodName.Destroy()
    ptrsForDecal.fnGetLowerFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_normal_fade")
    defer methodName.Destroy()
    ptrsForDecal.fnSetNormalFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_normal_fade")
    defer methodName.Destroy()
    ptrsForDecal.fnGetNormalFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_enable_distance_fade")
    defer methodName.Destroy()
    ptrsForDecal.fnSetEnableDistanceFade = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_distance_fade_enabled")
    defer methodName.Destroy()
    ptrsForDecal.fnIsDistanceFadeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_distance_fade_begin")
    defer methodName.Destroy()
    ptrsForDecal.fnSetDistanceFadeBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_distance_fade_begin")
    defer methodName.Destroy()
    ptrsForDecal.fnGetDistanceFadeBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_distance_fade_length")
    defer methodName.Destroy()
    ptrsForDecal.fnSetDistanceFadeLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_distance_fade_length")
    defer methodName.Destroy()
    ptrsForDecal.fnGetDistanceFadeLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_cull_mask")
    defer methodName.Destroy()
    ptrsForDecal.fnSetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_cull_mask")
    defer methodName.Destroy()
    ptrsForDecal.fnGetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

type Decal struct {
  VisualInstance3D
}

func (me *Decal) BaseClass() string {
  return "Decal"
}

func NewDecal() *Decal {
  str := StringNameFromStr("Decal") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Decal{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type DecalDecalTexture int
const (
  DecalDecalTextureTextureAlbedo DecalDecalTexture = 0
  DecalDecalTextureTextureNormal DecalDecalTexture = 1
  DecalDecalTextureTextureOrm DecalDecalTexture = 2
  DecalDecalTextureTextureEmission DecalDecalTexture = 3
  DecalDecalTextureTextureMax DecalDecalTexture = 4
)

func (me *Decal) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Decal) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Decal) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Decal) SetSize(size Vector3, )  {
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Decal) GetSize() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Decal) SetTexture(type_ DecalDecalTexture, texture Texture2D, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Decal) GetTexture(type_ DecalDecalTexture, ) Texture2D {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()
  pinner.Pin(&type_)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Decal) SetEmissionEnergy(energy float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnSetEmissionEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Decal) GetEmissionEnergy() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnGetEmissionEnergy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Decal) SetAlbedoMix(energy float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnSetAlbedoMix), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Decal) GetAlbedoMix() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnGetAlbedoMix), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Decal) SetModulate(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnSetModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Decal) GetModulate() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnGetModulate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Decal) SetUpperFade(fade float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fade) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnSetUpperFade), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Decal) GetUpperFade() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnGetUpperFade), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Decal) SetLowerFade(fade float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fade) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnSetLowerFade), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Decal) GetLowerFade() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnGetLowerFade), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Decal) SetNormalFade(fade float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fade) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnSetNormalFade), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Decal) GetNormalFade() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnGetNormalFade), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Decal) SetEnableDistanceFade(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnSetEnableDistanceFade), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Decal) IsDistanceFadeEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnIsDistanceFadeEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Decal) SetDistanceFadeBegin(distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnSetDistanceFadeBegin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Decal) GetDistanceFadeBegin() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnGetDistanceFadeBegin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Decal) SetDistanceFadeLength(distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnSetDistanceFadeLength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Decal) GetDistanceFadeLength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnGetDistanceFadeLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Decal) SetCullMask(mask int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnSetCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Decal) GetCullMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDecal.fnGetCullMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
