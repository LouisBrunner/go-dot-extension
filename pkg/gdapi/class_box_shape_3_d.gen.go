// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BoxShape3D struct {
  obj gdc.ObjectPtr
}

func createBoxShape3D(obj gdc.ObjectPtr) *BoxShape3D {
  return &BoxShape3D{
    obj: obj,
  }
}

func (me *BoxShape3D) BaseClass() string {
  return "BoxShape3D"
}
