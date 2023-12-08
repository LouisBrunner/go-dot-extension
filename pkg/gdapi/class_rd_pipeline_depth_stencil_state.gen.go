// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDPipelineDepthStencilState struct {
  obj gdc.ObjectPtr
}

func createRDPipelineDepthStencilState(obj gdc.ObjectPtr) *RDPipelineDepthStencilState {
  return &RDPipelineDepthStencilState{
    obj: obj,
  }
}

func (me *RDPipelineDepthStencilState) BaseClass() string {
  return "RDPipelineDepthStencilState"
}
