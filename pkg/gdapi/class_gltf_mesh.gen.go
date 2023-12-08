// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFMesh struct {
  obj gdc.ObjectPtr
}

func createGLTFMesh(obj gdc.ObjectPtr) *GLTFMesh {
  return &GLTFMesh{
    obj: obj,
  }
}

func (me *GLTFMesh) BaseClass() string {
  return "GLTFMesh"
}
