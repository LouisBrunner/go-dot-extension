// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDPipelineMultisampleState struct {
  obj gdc.ObjectPtr
}

func createRDPipelineMultisampleState(obj gdc.ObjectPtr) *RDPipelineMultisampleState {
  return &RDPipelineMultisampleState{
    obj: obj,
  }
}

func (me *RDPipelineMultisampleState) BaseClass() string {
  return "RDPipelineMultisampleState"
}
