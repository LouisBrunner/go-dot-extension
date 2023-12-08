// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TorusMesh struct {
  obj gdc.ObjectPtr
}

func createTorusMesh(obj gdc.ObjectPtr) *TorusMesh {
  return &TorusMesh{
    obj: obj,
  }
}

func (me *TorusMesh) BaseClass() string {
  return "TorusMesh"
}
