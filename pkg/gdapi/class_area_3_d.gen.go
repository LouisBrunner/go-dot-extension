// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Area3D struct {
  obj gdc.ObjectPtr
}

func createArea3D(obj gdc.ObjectPtr) *Area3D {
  return &Area3D{
    obj: obj,
  }
}

func (me *Area3D) BaseClass() string {
  return "Area3D"
}
