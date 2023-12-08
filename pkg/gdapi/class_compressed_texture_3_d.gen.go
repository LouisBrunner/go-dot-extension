// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CompressedTexture3D struct {
  obj gdc.ObjectPtr
}

func createCompressedTexture3D(obj gdc.ObjectPtr) *CompressedTexture3D {
  return &CompressedTexture3D{
    obj: obj,
  }
}

func (me *CompressedTexture3D) BaseClass() string {
  return "CompressedTexture3D"
}
