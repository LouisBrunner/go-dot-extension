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

type DecalDecalTexture int
const (
  DecalDecalTextureTextureAlbedo DecalDecalTexture = 0
  DecalDecalTextureTextureNormal DecalDecalTexture = 1
  DecalDecalTextureTextureOrm DecalDecalTexture = 2
  DecalDecalTextureTextureEmission DecalDecalTexture = 3
  DecalDecalTextureTextureMax DecalDecalTexture = 4
)

func  (me *Decal) SetSize(size Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Decal) GetSize() { // TODO: return value
  // TODO: implement
}

func  (me *Decal) SetTexture(type_ DecalDecalTexture, texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Decal) GetTexture(type_ DecalDecalTexture, ) { // TODO: return value
  // TODO: implement
}

func  (me *Decal) SetEmissionEnergy(energy float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Decal) GetEmissionEnergy() { // TODO: return value
  // TODO: implement
}

func  (me *Decal) SetAlbedoMix(energy float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Decal) GetAlbedoMix() { // TODO: return value
  // TODO: implement
}

func  (me *Decal) SetModulate(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Decal) GetModulate() { // TODO: return value
  // TODO: implement
}

func  (me *Decal) SetUpperFade(fade float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Decal) GetUpperFade() { // TODO: return value
  // TODO: implement
}

func  (me *Decal) SetLowerFade(fade float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Decal) GetLowerFade() { // TODO: return value
  // TODO: implement
}

func  (me *Decal) SetNormalFade(fade float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Decal) GetNormalFade() { // TODO: return value
  // TODO: implement
}

func  (me *Decal) SetEnableDistanceFade(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Decal) IsDistanceFadeEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *Decal) SetDistanceFadeBegin(distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Decal) GetDistanceFadeBegin() { // TODO: return value
  // TODO: implement
}

func  (me *Decal) SetDistanceFadeLength(distance float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Decal) GetDistanceFadeLength() { // TODO: return value
  // TODO: implement
}

func  (me *Decal) SetCullMask(mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Decal) GetCullMask() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
