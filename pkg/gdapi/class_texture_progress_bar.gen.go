// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextureProgressBar struct {
  obj gdc.ObjectPtr
}

func createTextureProgressBar(obj gdc.ObjectPtr) *TextureProgressBar {
  return &TextureProgressBar{
    obj: obj,
  }
}

func (me *TextureProgressBar) BaseClass() string {
  return "TextureProgressBar"
}
