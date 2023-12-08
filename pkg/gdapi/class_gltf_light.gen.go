// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFLight struct {
  obj gdc.ObjectPtr
}

func createGLTFLight(obj gdc.ObjectPtr) *GLTFLight {
  return &GLTFLight{
    obj: obj,
  }
}

func (me *GLTFLight) BaseClass() string {
  return "GLTFLight"
}
