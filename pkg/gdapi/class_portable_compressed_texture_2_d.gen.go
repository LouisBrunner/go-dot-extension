// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PortableCompressedTexture2D struct {
  obj gdc.ObjectPtr
}

func createPortableCompressedTexture2D(obj gdc.ObjectPtr) *PortableCompressedTexture2D {
  return &PortableCompressedTexture2D{
    obj: obj,
  }
}

func (me *PortableCompressedTexture2D) BaseClass() string {
  return "PortableCompressedTexture2D"
}
