// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImageTexture struct {
  obj gdc.ObjectPtr
}

func createImageTexture(obj gdc.ObjectPtr) *ImageTexture {
  return &ImageTexture{
    obj: obj,
  }
}

func (me *ImageTexture) BaseClass() string {
  return "ImageTexture"
}
