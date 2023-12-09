// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *Decal) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Decal) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Decal) SetSize(size Vector3, )  {
  panic("TODO: implement")
}

func  (me *Decal) GetSize()  {
  panic("TODO: implement")
}

func  (me *Decal) SetTexture(type_ DecalDecalTexture, texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *Decal) GetTexture(type_ DecalDecalTexture, )  {
  panic("TODO: implement")
}

func  (me *Decal) SetEmissionEnergy(energy float32, )  {
  panic("TODO: implement")
}

func  (me *Decal) GetEmissionEnergy()  {
  panic("TODO: implement")
}

func  (me *Decal) SetAlbedoMix(energy float32, )  {
  panic("TODO: implement")
}

func  (me *Decal) GetAlbedoMix()  {
  panic("TODO: implement")
}

func  (me *Decal) SetModulate(color Color, )  {
  panic("TODO: implement")
}

func  (me *Decal) GetModulate()  {
  panic("TODO: implement")
}

func  (me *Decal) SetUpperFade(fade float32, )  {
  panic("TODO: implement")
}

func  (me *Decal) GetUpperFade()  {
  panic("TODO: implement")
}

func  (me *Decal) SetLowerFade(fade float32, )  {
  panic("TODO: implement")
}

func  (me *Decal) GetLowerFade()  {
  panic("TODO: implement")
}

func  (me *Decal) SetNormalFade(fade float32, )  {
  panic("TODO: implement")
}

func  (me *Decal) GetNormalFade()  {
  panic("TODO: implement")
}

func  (me *Decal) SetEnableDistanceFade(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Decal) IsDistanceFadeEnabled()  {
  panic("TODO: implement")
}

func  (me *Decal) SetDistanceFadeBegin(distance float32, )  {
  panic("TODO: implement")
}

func  (me *Decal) GetDistanceFadeBegin()  {
  panic("TODO: implement")
}

func  (me *Decal) SetDistanceFadeLength(distance float32, )  {
  panic("TODO: implement")
}

func  (me *Decal) GetDistanceFadeLength()  {
  panic("TODO: implement")
}

func  (me *Decal) SetCullMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *Decal) GetCullMask()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
