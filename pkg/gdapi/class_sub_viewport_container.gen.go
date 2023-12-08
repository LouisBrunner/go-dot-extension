// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SubViewportContainer struct {
  obj gdc.ObjectPtr
}

func createSubViewportContainer(obj gdc.ObjectPtr) *SubViewportContainer {
  return &SubViewportContainer{
    obj: obj,
  }
}

func (me *SubViewportContainer) BaseClass() string {
  return "SubViewportContainer"
}
