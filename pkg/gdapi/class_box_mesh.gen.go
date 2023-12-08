// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BoxMesh struct {
  obj gdc.ObjectPtr
}

func createBoxMesh(obj gdc.ObjectPtr) *BoxMesh {
  return &BoxMesh{
    obj: obj,
  }
}

func (me *BoxMesh) BaseClass() string {
  return "BoxMesh"
}
