// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDPipelineColorBlendStateAttachment struct {
  obj gdc.ObjectPtr
}

func createRDPipelineColorBlendStateAttachment(obj gdc.ObjectPtr) *RDPipelineColorBlendStateAttachment {
  return &RDPipelineColorBlendStateAttachment{
    obj: obj,
  }
}

func (me *RDPipelineColorBlendStateAttachment) BaseClass() string {
  return "RDPipelineColorBlendStateAttachment"
}
