// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextureProgressBar struct {
  obj gdc.ObjectPtr
}

func (me *TextureProgressBar) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextureProgressBar) BaseClass() string {
  return "TextureProgressBar"
}



// Enums

type TextureProgressBarFillMode int
const (
  TextureProgressBarFillModeFillLeftToRight TextureProgressBarFillMode = 0
  TextureProgressBarFillModeFillRightToLeft TextureProgressBarFillMode = 1
  TextureProgressBarFillModeFillTopToBottom TextureProgressBarFillMode = 2
  TextureProgressBarFillModeFillBottomToTop TextureProgressBarFillMode = 3
  TextureProgressBarFillModeFillClockwise TextureProgressBarFillMode = 4
  TextureProgressBarFillModeFillCounterClockwise TextureProgressBarFillMode = 5
  TextureProgressBarFillModeFillBilinearLeftAndRight TextureProgressBarFillMode = 6
  TextureProgressBarFillModeFillBilinearTopAndBottom TextureProgressBarFillMode = 7
  TextureProgressBarFillModeFillClockwiseAndCounterClockwise TextureProgressBarFillMode = 8
)

func (me *TextureProgressBar) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextureProgressBar) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextureProgressBar) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *TextureProgressBar) SetUnderTexture(tex Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) GetUnderTexture()  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) SetProgressTexture(tex Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) GetProgressTexture()  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) SetOverTexture(tex Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) GetOverTexture()  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) SetFillMode(mode int, )  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) GetFillMode()  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) SetTintUnder(tint Color, )  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) GetTintUnder()  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) SetTintProgress(tint Color, )  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) GetTintProgress()  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) SetTintOver(tint Color, )  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) GetTintOver()  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) SetTextureProgressOffset(offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) GetTextureProgressOffset()  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) SetRadialInitialAngle(mode float32, )  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) GetRadialInitialAngle()  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) SetRadialCenterOffset(mode Vector2, )  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) GetRadialCenterOffset()  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) SetFillDegrees(mode float32, )  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) GetFillDegrees()  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) SetStretchMargin(margin Side, value int, )  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) GetStretchMargin(margin Side, )  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) SetNinePatchStretch(stretch bool, )  {
  panic("TODO: implement")
}

func  (me *TextureProgressBar) GetNinePatchStretch()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
