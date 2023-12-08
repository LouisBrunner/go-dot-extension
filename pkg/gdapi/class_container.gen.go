// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Container struct {
  obj gdc.ObjectPtr
}

func createContainer(obj gdc.ObjectPtr) *Container {
  return &Container{
    obj: obj,
  }
}

func (me *Container) BaseClass() string {
  return "Container"
}
