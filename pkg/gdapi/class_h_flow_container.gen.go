// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HFlowContainer struct {
  obj gdc.ObjectPtr
}

func createHFlowContainer(obj gdc.ObjectPtr) *HFlowContainer {
  return &HFlowContainer{
    obj: obj,
  }
}

func (me *HFlowContainer) BaseClass() string {
  return "HFlowContainer"
}
