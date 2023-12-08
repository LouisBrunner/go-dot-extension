// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  SpriteBase3DDrawFlagsFlagTransparent SpriteBase3DDrawFlags = 0
  SpriteBase3DDrawFlagsFlagShaded SpriteBase3DDrawFlags = 1
  SpriteBase3DDrawFlagsFlagDoubleSided SpriteBase3DDrawFlags = 2
  SpriteBase3DDrawFlagsFlagDisableDepthTest SpriteBase3DDrawFlags = 3
  SpriteBase3DDrawFlagsFlagFixedSize SpriteBase3DDrawFlags = 4
  SpriteBase3DDrawFlagsFlagMax SpriteBase3DDrawFlags = 5
)

type SpriteBase3DAlphaCutMode int
const (
  SpriteBase3DAlphaCutModeAlphaCutDisabled SpriteBase3DAlphaCutMode = 0
  SpriteBase3DAlphaCutModeAlphaCutDiscard SpriteBase3DAlphaCutMode = 1
  SpriteBase3DAlphaCutModeAlphaCutOpaquePrepass SpriteBase3DAlphaCutMode = 2
  SpriteBase3DAlphaCutModeAlphaCutHash SpriteBase3DAlphaCutMode = 3
)

func  (me *SpriteBase3D) SetCentered(centered bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) IsCentered() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetOffset(offset Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GetOffset() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetFlipH(flip_h bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) IsFlippedH() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetFlipV(flip_v bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) IsFlippedV() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetModulate(modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GetModulate() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetRenderPriority(priority int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GetRenderPriority() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetPixelSize(pixel_size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GetPixelSize() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetAxis(axis Vector3Axis, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GetAxis() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetDrawFlag(flag SpriteBase3DDrawFlags, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GetDrawFlag(flag SpriteBase3DDrawFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetAlphaCutMode(mode SpriteBase3DAlphaCutMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GetAlphaCutMode() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetAlphaScissorThreshold(threshold float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GetAlphaScissorThreshold() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetAlphaHashScale(threshold float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GetAlphaHashScale() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetAlphaAntialiasing(alpha_aa BaseMaterial3DAlphaAntiAliasing, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GetAlphaAntialiasing() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetAlphaAntialiasingEdge(edge float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GetAlphaAntialiasingEdge() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetBillboardMode(mode BaseMaterial3DBillboardMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GetBillboardMode() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) SetTextureFilter(mode BaseMaterial3DTextureFilter, ) { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GetTextureFilter() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GetItemRect() { // TODO: return value
  // TODO: implement
}

func  (me *SpriteBase3D) GenerateTriangleMesh() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
