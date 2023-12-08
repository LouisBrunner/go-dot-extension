// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

type GradientTexture2DFill int
const (
  GradientTexture2DFillLinear GradientTexture2DFill = 0
  GradientTexture2DFillRadial GradientTexture2DFill = 1
  GradientTexture2DFillSquare GradientTexture2DFill = 2
)

type GradientTexture2DRepeat int
const (
  GradientTexture2DRepeatNone GradientTexture2DRepeat = 0
  GradientTexture2DRepeat_ GradientTexture2DRepeat = 1
  GradientTexture2DRepeatMirror GradientTexture2DRepeat = 2
)
