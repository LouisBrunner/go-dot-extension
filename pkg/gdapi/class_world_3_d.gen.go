// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type World3D struct {
  obj gdc.ObjectPtr
}

func createWorld3D(obj gdc.ObjectPtr) *World3D {
  return &World3D{
    obj: obj,
  }
}

func (me *World3D) BaseClass() string {
  return "World3D"
}
