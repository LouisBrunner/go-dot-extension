// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ImmediateMesh struct {
  obj gdc.ObjectPtr
}

func createImmediateMesh(obj gdc.ObjectPtr) *ImmediateMesh {
  return &ImmediateMesh{
    obj: obj,
  }
}

func (me *ImmediateMesh) BaseClass() string {
  return "ImmediateMesh"
}
