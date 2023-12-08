// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CompressedCubemapArray struct {
  obj gdc.ObjectPtr
}

func createCompressedCubemapArray(obj gdc.ObjectPtr) *CompressedCubemapArray {
  return &CompressedCubemapArray{
    obj: obj,
  }
}

func (me *CompressedCubemapArray) BaseClass() string {
  return "CompressedCubemapArray"
}
