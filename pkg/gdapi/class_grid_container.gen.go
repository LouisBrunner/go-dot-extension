// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GridContainer struct {
  obj gdc.ObjectPtr
}

func createGridContainer(obj gdc.ObjectPtr) *GridContainer {
  return &GridContainer{
    obj: obj,
  }
}

func (me *GridContainer) BaseClass() string {
  return "GridContainer"
}
