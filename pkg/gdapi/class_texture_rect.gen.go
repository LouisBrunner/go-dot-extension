// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextureRect struct {
  obj gdc.ObjectPtr
}

func (me *TextureRect) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextureRect) BaseClass() string {
  return "TextureRect"
}

type TextureRectExpandMode int
const (
  TextureRectExpandModeExpandKeepSize TextureRectExpandMode = 0
  TextureRectExpandModeExpandIgnoreSize TextureRectExpandMode = 1
  TextureRectExpandModeExpandFitWidth TextureRectExpandMode = 2
  TextureRectExpandModeExpandFitWidthProportional TextureRectExpandMode = 3
  TextureRectExpandModeExpandFitHeight TextureRectExpandMode = 4
  TextureRectExpandModeExpandFitHeightProportional TextureRectExpandMode = 5
)

type TextureRectStretchMode int
const (
  TextureRectStretchModeStretchScale TextureRectStretchMode = 0
  TextureRectStretchModeStretchTile TextureRectStretchMode = 1
  TextureRectStretchModeStretchKeep TextureRectStretchMode = 2
  TextureRectStretchModeStretchKeepCentered TextureRectStretchMode = 3
  TextureRectStretchModeStretchKeepAspect TextureRectStretchMode = 4
  TextureRectStretchModeStretchKeepAspectCentered TextureRectStretchMode = 5
  TextureRectStretchModeStretchKeepAspectCovered TextureRectStretchMode = 6
)

func  (me *TextureRect) SetTexture(texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureRect) GetTexture() { // TODO: return value
  // TODO: implement
}

func  (me *TextureRect) SetExpandMode(expand_mode TextureRectExpandMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureRect) GetExpandMode() { // TODO: return value
  // TODO: implement
}

func  (me *TextureRect) SetFlipH(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureRect) IsFlippedH() { // TODO: return value
  // TODO: implement
}

func  (me *TextureRect) SetFlipV(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureRect) IsFlippedV() { // TODO: return value
  // TODO: implement
}

func  (me *TextureRect) SetStretchMode(stretch_mode TextureRectStretchMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextureRect) GetStretchMode() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
