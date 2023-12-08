// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Camera3D struct {
  obj gdc.ObjectPtr
}

func createCamera3D(obj gdc.ObjectPtr) *Camera3D {
  return &Camera3D{
    obj: obj,
  }
}

func (me *Camera3D) BaseClass() string {
  return "Camera3D"
}
