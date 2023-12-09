// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GradientTexture2D struct {
  obj gdc.ObjectPtr
}

func (me *GradientTexture2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GradientTexture2D) BaseClass() string {
  return "GradientTexture2D"
}



// Enums

type GradientTexture2DFill int
const (
  GradientTexture2DFillFillLinear GradientTexture2DFill = 0
  GradientTexture2DFillFillRadial GradientTexture2DFill = 1
  GradientTexture2DFillFillSquare GradientTexture2DFill = 2
)

type GradientTexture2DRepeat int
const (
  GradientTexture2DRepeatRepeatNone GradientTexture2DRepeat = 0
  GradientTexture2DRepeatRepeat GradientTexture2DRepeat = 1
  GradientTexture2DRepeatRepeatMirror GradientTexture2DRepeat = 2
)

func (me *GradientTexture2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GradientTexture2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GradientTexture2D) SetGradient(gradient Gradient, )  {
  panic("TODO: implement")
}

func  (me *GradientTexture2D) GetGradient()  {
  panic("TODO: implement")
}

func  (me *GradientTexture2D) SetWidth(width int, )  {
  panic("TODO: implement")
}

func  (me *GradientTexture2D) SetHeight(height int, )  {
  panic("TODO: implement")
}

func  (me *GradientTexture2D) SetUseHdr(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *GradientTexture2D) IsUsingHdr()  {
  panic("TODO: implement")
}

func  (me *GradientTexture2D) SetFill(fill GradientTexture2DFill, )  {
  panic("TODO: implement")
}

func  (me *GradientTexture2D) GetFill()  {
  panic("TODO: implement")
}

func  (me *GradientTexture2D) SetFillFrom(fill_from Vector2, )  {
  panic("TODO: implement")
}

func  (me *GradientTexture2D) GetFillFrom()  {
  panic("TODO: implement")
}

func  (me *GradientTexture2D) SetFillTo(fill_to Vector2, )  {
  panic("TODO: implement")
}

func  (me *GradientTexture2D) GetFillTo()  {
  panic("TODO: implement")
}

func  (me *GradientTexture2D) SetRepeat(repeat GradientTexture2DRepeat, )  {
  panic("TODO: implement")
}

func  (me *GradientTexture2D) GetRepeat()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
