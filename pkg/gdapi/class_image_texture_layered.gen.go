// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImageTextureLayered struct {
  obj gdc.ObjectPtr
}

func createImageTextureLayered(obj gdc.ObjectPtr) *ImageTextureLayered {
  return &ImageTextureLayered{
    obj: obj,
  }
}

func (me *ImageTextureLayered) BaseClass() string {
  return "ImageTextureLayered"
}
