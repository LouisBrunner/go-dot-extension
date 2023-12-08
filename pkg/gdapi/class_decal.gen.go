// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  DecalTextureAlbedo DecalDecalTexture = 0
  DecalTextureNormal DecalDecalTexture = 1
  DecalTextureOrm DecalDecalTexture = 2
  DecalTextureEmission DecalDecalTexture = 3
  DecalTextureMax DecalDecalTexture = 4
)
