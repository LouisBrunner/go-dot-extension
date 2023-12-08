// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CompressedTexture2D struct {
  obj gdc.ObjectPtr
}

func createCompressedTexture2D(obj gdc.ObjectPtr) *CompressedTexture2D {
  return &CompressedTexture2D{
    obj: obj,
  }
}

func (me *CompressedTexture2D) BaseClass() string {
  return "CompressedTexture2D"
}
