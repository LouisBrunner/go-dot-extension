// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFTextureSampler struct {
  obj gdc.ObjectPtr
}

func createGLTFTextureSampler(obj gdc.ObjectPtr) *GLTFTextureSampler {
  return &GLTFTextureSampler{
    obj: obj,
  }
}

func (me *GLTFTextureSampler) BaseClass() string {
  return "GLTFTextureSampler"
}
