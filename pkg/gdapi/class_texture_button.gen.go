// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextureButton struct {
  obj gdc.ObjectPtr
}

func createTextureButton(obj gdc.ObjectPtr) *TextureButton {
  return &TextureButton{
    obj: obj,
  }
}

func (me *TextureButton) BaseClass() string {
  return "TextureButton"
}
