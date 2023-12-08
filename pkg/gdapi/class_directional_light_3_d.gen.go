// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DirectionalLight3D struct {
  obj gdc.ObjectPtr
}

func createDirectionalLight3D(obj gdc.ObjectPtr) *DirectionalLight3D {
  return &DirectionalLight3D{
    obj: obj,
  }
}

func (me *DirectionalLight3D) BaseClass() string {
  return "DirectionalLight3D"
}
