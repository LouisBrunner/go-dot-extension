// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Camera2D struct {
  obj gdc.ObjectPtr
}

func createCamera2D(obj gdc.ObjectPtr) *Camera2D {
  return &Camera2D{
    obj: obj,
  }
}

func (me *Camera2D) BaseClass() string {
  return "Camera2D"
}
