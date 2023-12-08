// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WorldEnvironment struct {
  obj gdc.ObjectPtr
}

func createWorldEnvironment(obj gdc.ObjectPtr) *WorldEnvironment {
  return &WorldEnvironment{
    obj: obj,
  }
}

func (me *WorldEnvironment) BaseClass() string {
  return "WorldEnvironment"
}
