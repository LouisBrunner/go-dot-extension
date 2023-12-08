// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFTexture struct {
  obj gdc.ObjectPtr
}

func createGLTFTexture(obj gdc.ObjectPtr) *GLTFTexture {
  return &GLTFTexture{
    obj: obj,
  }
}

func (me *GLTFTexture) BaseClass() string {
  return "GLTFTexture"
}
