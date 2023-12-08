// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDSamplerState struct {
  obj gdc.ObjectPtr
}

func createRDSamplerState(obj gdc.ObjectPtr) *RDSamplerState {
  return &RDSamplerState{
    obj: obj,
  }
}

func (me *RDSamplerState) BaseClass() string {
  return "RDSamplerState"
}
