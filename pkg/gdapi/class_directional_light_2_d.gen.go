// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DirectionalLight2D struct {
  obj gdc.ObjectPtr
}

func createDirectionalLight2D(obj gdc.ObjectPtr) *DirectionalLight2D {
  return &DirectionalLight2D{
    obj: obj,
  }
}

func (me *DirectionalLight2D) BaseClass() string {
  return "DirectionalLight2D"
}
