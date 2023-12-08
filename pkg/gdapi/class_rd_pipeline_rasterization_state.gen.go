// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDPipelineRasterizationState struct {
  obj gdc.ObjectPtr
}

func createRDPipelineRasterizationState(obj gdc.ObjectPtr) *RDPipelineRasterizationState {
  return &RDPipelineRasterizationState{
    obj: obj,
  }
}

func (me *RDPipelineRasterizationState) BaseClass() string {
  return "RDPipelineRasterizationState"
}
