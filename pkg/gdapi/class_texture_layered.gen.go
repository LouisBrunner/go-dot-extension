// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextureLayered struct {
  obj gdc.ObjectPtr
}

func createTextureLayered(obj gdc.ObjectPtr) *TextureLayered {
  return &TextureLayered{
    obj: obj,
  }
}

func (me *TextureLayered) BaseClass() string {
  return "TextureLayered"
}
