// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CompressedTexture2DArray struct {
  obj gdc.ObjectPtr
}

func createCompressedTexture2DArray(obj gdc.ObjectPtr) *CompressedTexture2DArray {
  return &CompressedTexture2DArray{
    obj: obj,
  }
}

func (me *CompressedTexture2DArray) BaseClass() string {
  return "CompressedTexture2DArray"
}
