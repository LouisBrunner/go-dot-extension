// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDPipelineColorBlendState struct {
  obj gdc.ObjectPtr
}

func createRDPipelineColorBlendState(obj gdc.ObjectPtr) *RDPipelineColorBlendState {
  return &RDPipelineColorBlendState{
    obj: obj,
  }
}

func (me *RDPipelineColorBlendState) BaseClass() string {
  return "RDPipelineColorBlendState"
}
