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

func  (me *TextureButton) SetTextureNormal(texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) SetTexturePressed(texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) SetTextureHover(texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) SetTextureDisabled(texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) SetTextureFocused(texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) SetClickMask(mask BitMap, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) SetIgnoreTextureSize(ignore bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) SetStretchMode(mode TextureButtonStretchMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) SetFlipH(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) IsFlippedH() { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) SetFlipV(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) IsFlippedV() { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) GetTextureNormal() { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) GetTexturePressed() { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) GetTextureHover() { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) GetTextureDisabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) GetTextureFocused() { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) GetClickMask() { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) GetIgnoreTextureSize() { // TODO: return value
  // TODO: implement
}

func  (me *TextureButton) GetStretchMode() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
