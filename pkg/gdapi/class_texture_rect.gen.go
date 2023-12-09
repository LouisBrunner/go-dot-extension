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



// Enums

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

func (me *TextureRect) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextureRect) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *TextureRect) SetTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TextureRect) GetTexture()  {
  panic("TODO: implement")
}

func  (me *TextureRect) SetExpandMode(expand_mode TextureRectExpandMode, )  {
  panic("TODO: implement")
}

func  (me *TextureRect) GetExpandMode()  {
  panic("TODO: implement")
}

func  (me *TextureRect) SetFlipH(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TextureRect) IsFlippedH()  {
  panic("TODO: implement")
}

func  (me *TextureRect) SetFlipV(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TextureRect) IsFlippedV()  {
  panic("TODO: implement")
}

func  (me *TextureRect) SetStretchMode(stretch_mode TextureRectStretchMode, )  {
  panic("TODO: implement")
}

func  (me *TextureRect) GetStretchMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
