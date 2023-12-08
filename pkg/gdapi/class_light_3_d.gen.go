// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Light3D struct {
  obj gdc.ObjectPtr
}

func createLight3D(obj gdc.ObjectPtr) *Light3D {
  return &Light3D{
    obj: obj,
  }
}

func (me *Light3D) BaseClass() string {
  return "Light3D"
}
