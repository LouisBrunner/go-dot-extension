// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CompressedTextureLayered struct {
  obj gdc.ObjectPtr
}

func createCompressedTextureLayered(obj gdc.ObjectPtr) *CompressedTextureLayered {
  return &CompressedTextureLayered{
    obj: obj,
  }
}

func (me *CompressedTextureLayered) BaseClass() string {
  return "CompressedTextureLayered"
}
