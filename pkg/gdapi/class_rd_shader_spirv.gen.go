// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDShaderSPIRV struct {
  obj gdc.ObjectPtr
}

func (me *RDShaderSPIRV) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDShaderSPIRV) BaseClass() string {
  return "RDShaderSPIRV"
}
