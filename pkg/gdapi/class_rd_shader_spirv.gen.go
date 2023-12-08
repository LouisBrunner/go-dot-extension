// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDShaderSPIRV struct {
  obj gdc.ObjectPtr
}

func createRDShaderSPIRV(obj gdc.ObjectPtr) *RDShaderSPIRV {
  return &RDShaderSPIRV{
    obj: obj,
  }
}

func (me *RDShaderSPIRV) BaseClass() string {
  return "RDShaderSPIRV"
}
