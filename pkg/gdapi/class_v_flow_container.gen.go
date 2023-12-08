// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VFlowContainer struct {
  obj gdc.ObjectPtr
}

func createVFlowContainer(obj gdc.ObjectPtr) *VFlowContainer {
  return &VFlowContainer{
    obj: obj,
  }
}

func (me *VFlowContainer) BaseClass() string {
  return "VFlowContainer"
}
