// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

type TextureProgressBarFillMode int
const (
  TextureProgressBarFillLeftToRight TextureProgressBarFillMode = 0
  TextureProgressBarFillRightToLeft TextureProgressBarFillMode = 1
  TextureProgressBarFillTopToBottom TextureProgressBarFillMode = 2
  TextureProgressBarFillBottomToTop TextureProgressBarFillMode = 3
  TextureProgressBarFillClockwise TextureProgressBarFillMode = 4
  TextureProgressBarFillCounterClockwise TextureProgressBarFillMode = 5
  TextureProgressBarFillBilinearLeftAndRight TextureProgressBarFillMode = 6
  TextureProgressBarFillBilinearTopAndBottom TextureProgressBarFillMode = 7
  TextureProgressBarFillClockwiseAndCounterClockwise TextureProgressBarFillMode = 8
)
