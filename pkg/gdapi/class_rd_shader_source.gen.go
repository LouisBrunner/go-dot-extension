// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDShaderSource struct {
  obj gdc.ObjectPtr
}

func createRDShaderSource(obj gdc.ObjectPtr) *RDShaderSource {
  return &RDShaderSource{
    obj: obj,
  }
}

func (me *RDShaderSource) BaseClass() string {
  return "RDShaderSource"
}
