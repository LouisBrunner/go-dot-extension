// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Decal struct {
  obj gdc.ObjectPtr
}

func (me *Decal) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Decal) BaseClass() string {
  return "Decal"
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
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Decal) GetSize() Vector3 {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Decal) SetTexture(type_ DecalDecalTexture, texture Texture2D, )  {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2086764391) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Decal) GetTexture(type_ DecalDecalTexture, ) Texture2D {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3244119503) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Decal) SetEmissionEnergy(energy float32, )  {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Decal) GetEmissionEnergy() float32 {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Decal) SetAlbedoMix(energy float32, )  {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_albedo_mix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Decal) GetAlbedoMix() float32 {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_albedo_mix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Decal) SetModulate(color Color, )  {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Decal) GetModulate() Color {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Decal) SetUpperFade(fade float32, )  {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_upper_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fade), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Decal) GetUpperFade() float32 {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_upper_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Decal) SetLowerFade(fade float32, )  {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lower_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fade), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Decal) GetLowerFade() float32 {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lower_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Decal) SetNormalFade(fade float32, )  {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normal_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fade), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Decal) GetNormalFade() float32 {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_normal_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Decal) SetEnableDistanceFade(enable bool, )  {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_distance_fade")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Decal) IsDistanceFadeEnabled() bool {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_distance_fade_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Decal) SetDistanceFadeBegin(distance float32, )  {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distance_fade_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Decal) GetDistanceFadeBegin() float32 {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_distance_fade_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Decal) SetDistanceFadeLength(distance float32, )  {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distance_fade_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Decal) GetDistanceFadeLength() float32 {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_distance_fade_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Decal) SetCullMask(mask int, )  {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Decal) GetCullMask() int {
  classNameV := StringNameFromStr("Decal")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cull_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Decal) GetPropSize() Vector3 {
  panic("TODO: implement")
}

func (me *Decal) SetPropSize(value Vector3) {
  panic("TODO: implement")
}

func (me *Decal) GetPropTextureAlbedo() Texture {
  panic("TODO: implement")
}

func (me *Decal) SetPropTextureAlbedo(value Texture) {
  panic("TODO: implement")
}

func (me *Decal) GetPropTextureNormal() Texture {
  panic("TODO: implement")
}

func (me *Decal) SetPropTextureNormal(value Texture) {
  panic("TODO: implement")
}

func (me *Decal) GetPropTextureOrm() Texture {
  panic("TODO: implement")
}

func (me *Decal) SetPropTextureOrm(value Texture) {
  panic("TODO: implement")
}

func (me *Decal) GetPropTextureEmission() Texture {
  panic("TODO: implement")
}

func (me *Decal) SetPropTextureEmission(value Texture) {
  panic("TODO: implement")
}

func (me *Decal) GetPropEmissionEnergy() float32 {
  panic("TODO: implement")
}

func (me *Decal) SetPropEmissionEnergy(value float32) {
  panic("TODO: implement")
}

func (me *Decal) GetPropModulate() Color {
  panic("TODO: implement")
}

func (me *Decal) SetPropModulate(value Color) {
  panic("TODO: implement")
}

func (me *Decal) GetPropAlbedoMix() float32 {
  panic("TODO: implement")
}

func (me *Decal) SetPropAlbedoMix(value float32) {
  panic("TODO: implement")
}

func (me *Decal) GetPropNormalFade() float32 {
  panic("TODO: implement")
}

func (me *Decal) SetPropNormalFade(value float32) {
  panic("TODO: implement")
}

func (me *Decal) GetPropUpperFade() float32 {
  panic("TODO: implement")
}

func (me *Decal) SetPropUpperFade(value float32) {
  panic("TODO: implement")
}

func (me *Decal) GetPropLowerFade() float32 {
  panic("TODO: implement")
}

func (me *Decal) SetPropLowerFade(value float32) {
  panic("TODO: implement")
}

func (me *Decal) GetPropDistanceFadeEnabled() bool {
  panic("TODO: implement")
}

func (me *Decal) SetPropDistanceFadeEnabled(value bool) {
  panic("TODO: implement")
}

func (me *Decal) GetPropDistanceFadeBegin() float32 {
  panic("TODO: implement")
}

func (me *Decal) SetPropDistanceFadeBegin(value float32) {
  panic("TODO: implement")
}

func (me *Decal) GetPropDistanceFadeLength() float32 {
  panic("TODO: implement")
}

func (me *Decal) SetPropDistanceFadeLength(value float32) {
  panic("TODO: implement")
}

func (me *Decal) GetPropCullMask() int {
  panic("TODO: implement")
}

func (me *Decal) SetPropCullMask(value int) {
  panic("TODO: implement")
}