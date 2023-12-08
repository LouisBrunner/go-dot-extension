// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HSplitContainer struct {
  obj gdc.ObjectPtr
}

func createHSplitContainer(obj gdc.ObjectPtr) *HSplitContainer {
  return &HSplitContainer{
    obj: obj,
  }
}

func (me *HSplitContainer) BaseClass() string {
  return "HSplitContainer"
}
