// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Cubemap struct {
  obj gdc.ObjectPtr
}

func createCubemap(obj gdc.ObjectPtr) *Cubemap {
  return &Cubemap{
    obj: obj,
  }
}

func (me *Cubemap) BaseClass() string {
  return "Cubemap"
}
