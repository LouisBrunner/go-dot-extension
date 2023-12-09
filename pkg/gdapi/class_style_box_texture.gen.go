// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StyleBoxTexture struct {
  obj gdc.ObjectPtr
}

func (me *StyleBoxTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StyleBoxTexture) BaseClass() string {
  return "StyleBoxTexture"
}



// Enums

type StyleBoxTextureAxisStretchMode int
const (
  StyleBoxTextureAxisStretchModeAxisStretchModeStretch StyleBoxTextureAxisStretchMode = 0
  StyleBoxTextureAxisStretchModeAxisStretchModeTile StyleBoxTextureAxisStretchMode = 1
  StyleBoxTextureAxisStretchModeAxisStretchModeTileFit StyleBoxTextureAxisStretchMode = 2
)

func (me *StyleBoxTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StyleBoxTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StyleBoxTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *StyleBoxTexture) SetTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) GetTexture()  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) SetTextureMargin(margin Side, size float32, )  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) SetTextureMarginAll(size float32, )  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) GetTextureMargin(margin Side, )  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) SetExpandMargin(margin Side, size float32, )  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) SetExpandMarginAll(size float32, )  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) GetExpandMargin(margin Side, )  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) SetRegionRect(region Rect2, )  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) GetRegionRect()  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) SetDrawCenter(enable bool, )  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) IsDrawCenterEnabled()  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) SetModulate(color Color, )  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) GetModulate()  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) SetHAxisStretchMode(mode StyleBoxTextureAxisStretchMode, )  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) GetHAxisStretchMode()  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) SetVAxisStretchMode(mode StyleBoxTextureAxisStretchMode, )  {
  panic("TODO: implement")
}

func  (me *StyleBoxTexture) GetVAxisStretchMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
