// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  TextureRectExpandKeepSize TextureRectExpandMode = 0
  TextureRectExpandIgnoreSize TextureRectExpandMode = 1
  TextureRectExpandFitWidth TextureRectExpandMode = 2
  TextureRectExpandFitWidthProportional TextureRectExpandMode = 3
  TextureRectExpandFitHeight TextureRectExpandMode = 4
  TextureRectExpandFitHeightProportional TextureRectExpandMode = 5
)

type TextureRectStretchMode int
const (
  TextureRectStretchScale TextureRectStretchMode = 0
  TextureRectStretchTile TextureRectStretchMode = 1
  TextureRectStretchKeep TextureRectStretchMode = 2
  TextureRectStretchKeepCentered TextureRectStretchMode = 3
  TextureRectStretchKeepAspect TextureRectStretchMode = 4
  TextureRectStretchKeepAspectCentered TextureRectStretchMode = 5
  TextureRectStretchKeepAspectCovered TextureRectStretchMode = 6
)
