// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDUniform struct {
  obj gdc.ObjectPtr
}

func createRDUniform(obj gdc.ObjectPtr) *RDUniform {
  return &RDUniform{
    obj: obj,
  }
}

func (me *RDUniform) BaseClass() string {
  return "RDUniform"
}
