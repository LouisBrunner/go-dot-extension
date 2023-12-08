// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CompressedCubemap struct {
  obj gdc.ObjectPtr
}

func createCompressedCubemap(obj gdc.ObjectPtr) *CompressedCubemap {
  return &CompressedCubemap{
    obj: obj,
  }
}

func (me *CompressedCubemap) BaseClass() string {
  return "CompressedCubemap"
}
