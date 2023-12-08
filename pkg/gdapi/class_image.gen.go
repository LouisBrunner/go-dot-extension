// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Image struct {
  obj gdc.ObjectPtr
}

func createImage(obj gdc.ObjectPtr) *Image {
  return &Image{
    obj: obj,
  }
}

func (me *Image) BaseClass() string {
  return "Image"
}
