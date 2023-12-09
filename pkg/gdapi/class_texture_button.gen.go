// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextureButton struct {
  obj gdc.ObjectPtr
}

func (me *TextureButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextureButton) BaseClass() string {
  return "TextureButton"
}



// Enums

type TextureButtonStretchMode int
const (
  TextureButtonStretchModeStretchScale TextureButtonStretchMode = 0
  TextureButtonStretchModeStretchTile TextureButtonStretchMode = 1
  TextureButtonStretchModeStretchKeep TextureButtonStretchMode = 2
  TextureButtonStretchModeStretchKeepCentered TextureButtonStretchMode = 3
  TextureButtonStretchModeStretchKeepAspect TextureButtonStretchMode = 4
  TextureButtonStretchModeStretchKeepAspectCentered TextureButtonStretchMode = 5
  TextureButtonStretchModeStretchKeepAspectCovered TextureButtonStretchMode = 6
)

func (me *TextureButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextureButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *TextureButton) SetTextureNormal(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TextureButton) SetTexturePressed(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TextureButton) SetTextureHover(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TextureButton) SetTextureDisabled(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TextureButton) SetTextureFocused(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TextureButton) SetClickMask(mask BitMap, )  {
  panic("TODO: implement")
}

func  (me *TextureButton) SetIgnoreTextureSize(ignore bool, )  {
  panic("TODO: implement")
}

func  (me *TextureButton) SetStretchMode(mode TextureButtonStretchMode, )  {
  panic("TODO: implement")
}

func  (me *TextureButton) SetFlipH(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TextureButton) IsFlippedH()  {
  panic("TODO: implement")
}

func  (me *TextureButton) SetFlipV(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TextureButton) IsFlippedV()  {
  panic("TODO: implement")
}

func  (me *TextureButton) GetTextureNormal()  {
  panic("TODO: implement")
}

func  (me *TextureButton) GetTexturePressed()  {
  panic("TODO: implement")
}

func  (me *TextureButton) GetTextureHover()  {
  panic("TODO: implement")
}

func  (me *TextureButton) GetTextureDisabled()  {
  panic("TODO: implement")
}

func  (me *TextureButton) GetTextureFocused()  {
  panic("TODO: implement")
}

func  (me *TextureButton) GetClickMask()  {
  panic("TODO: implement")
}

func  (me *TextureButton) GetIgnoreTextureSize()  {
  panic("TODO: implement")
}

func  (me *TextureButton) GetStretchMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
