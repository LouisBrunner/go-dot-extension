// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImageTexture3D struct {
  obj gdc.ObjectPtr
}

func createImageTexture3D(obj gdc.ObjectPtr) *ImageTexture3D {
  return &ImageTexture3D{
    obj: obj,
  }
}

func (me *ImageTexture3D) BaseClass() string {
  return "ImageTexture3D"
}
