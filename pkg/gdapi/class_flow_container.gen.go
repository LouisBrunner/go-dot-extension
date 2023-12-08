// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FlowContainer struct {
  obj gdc.ObjectPtr
}

func createFlowContainer(obj gdc.ObjectPtr) *FlowContainer {
  return &FlowContainer{
    obj: obj,
  }
}

func (me *FlowContainer) BaseClass() string {
  return "FlowContainer"
}
