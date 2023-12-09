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



// Enums

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

func (me *SpriteBase3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SpriteBase3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SpriteBase3D) SetCentered(centered bool, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) IsCentered()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetOffset(offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GetOffset()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetFlipH(flip_h bool, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) IsFlippedH()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetFlipV(flip_v bool, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) IsFlippedV()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetModulate(modulate Color, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GetModulate()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetRenderPriority(priority int, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GetRenderPriority()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetPixelSize(pixel_size float32, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GetPixelSize()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetAxis(axis Vector3Axis, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GetAxis()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetDrawFlag(flag SpriteBase3DDrawFlags, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GetDrawFlag(flag SpriteBase3DDrawFlags, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetAlphaCutMode(mode SpriteBase3DAlphaCutMode, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GetAlphaCutMode()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetAlphaScissorThreshold(threshold float32, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GetAlphaScissorThreshold()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetAlphaHashScale(threshold float32, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GetAlphaHashScale()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetAlphaAntialiasing(alpha_aa BaseMaterial3DAlphaAntiAliasing, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GetAlphaAntialiasing()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetAlphaAntialiasingEdge(edge float32, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GetAlphaAntialiasingEdge()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetBillboardMode(mode BaseMaterial3DBillboardMode, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GetBillboardMode()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) SetTextureFilter(mode BaseMaterial3DTextureFilter, )  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GetTextureFilter()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GetItemRect()  {
  panic("TODO: implement")
}

func  (me *SpriteBase3D) GenerateTriangleMesh()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
