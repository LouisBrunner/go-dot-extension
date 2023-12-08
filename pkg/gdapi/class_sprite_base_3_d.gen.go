// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SpriteBase3D struct {
  obj gdc.ObjectPtr
}

func (me *SpriteBase3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SpriteBase3D) BaseClass() string {
  return "SpriteBase3D"
}

type SpriteBase3DDrawFlags int
const (
  SpriteBase3DFlagTransparent SpriteBase3DDrawFlags = 0
  SpriteBase3DFlagShaded SpriteBase3DDrawFlags = 1
  SpriteBase3DFlagDoubleSided SpriteBase3DDrawFlags = 2
  SpriteBase3DFlagDisableDepthTest SpriteBase3DDrawFlags = 3
  SpriteBase3DFlagFixedSize SpriteBase3DDrawFlags = 4
  SpriteBase3DFlagMax SpriteBase3DDrawFlags = 5
)

type SpriteBase3DAlphaCutMode int
const (
  SpriteBase3DAlphaCutDisabled SpriteBase3DAlphaCutMode = 0
  SpriteBase3DAlphaCutDiscard SpriteBase3DAlphaCutMode = 1
  SpriteBase3DAlphaCutOpaquePrepass SpriteBase3DAlphaCutMode = 2
  SpriteBase3DAlphaCutHash SpriteBase3DAlphaCutMode = 3
)
