// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextureRect struct {
  obj gdc.ObjectPtr
}

func createTextureRect(obj gdc.ObjectPtr) *TextureRect {
  return &TextureRect{
    obj: obj,
  }
}

func (me *TextureRect) BaseClass() string {
  return "TextureRect"
}
