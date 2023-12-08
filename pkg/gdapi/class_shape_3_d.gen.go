// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Shape3D struct {
  obj gdc.ObjectPtr
}

func createShape3D(obj gdc.ObjectPtr) *Shape3D {
  return &Shape3D{
    obj: obj,
  }
}

func (me *Shape3D) BaseClass() string {
  return "Shape3D"
}
