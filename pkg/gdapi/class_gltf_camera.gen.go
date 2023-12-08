// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFCamera struct {
  obj gdc.ObjectPtr
}

func createGLTFCamera(obj gdc.ObjectPtr) *GLTFCamera {
  return &GLTFCamera{
    obj: obj,
  }
}

func (me *GLTFCamera) BaseClass() string {
  return "GLTFCamera"
}
