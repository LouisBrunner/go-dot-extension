// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VSplitContainer struct {
  obj gdc.ObjectPtr
}

func createVSplitContainer(obj gdc.ObjectPtr) *VSplitContainer {
  return &VSplitContainer{
    obj: obj,
  }
}

func (me *VSplitContainer) BaseClass() string {
  return "VSplitContainer"
}
