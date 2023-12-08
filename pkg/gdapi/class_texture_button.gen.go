// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  TextureButtonStretchScale TextureButtonStretchMode = 0
  TextureButtonStretchTile TextureButtonStretchMode = 1
  TextureButtonStretchKeep TextureButtonStretchMode = 2
  TextureButtonStretchKeepCentered TextureButtonStretchMode = 3
  TextureButtonStretchKeepAspect TextureButtonStretchMode = 4
  TextureButtonStretchKeepAspectCentered TextureButtonStretchMode = 5
  TextureButtonStretchKeepAspectCovered TextureButtonStretchMode = 6
)
