// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDShaderFile struct {
  obj gdc.ObjectPtr
}

func createRDShaderFile(obj gdc.ObjectPtr) *RDShaderFile {
  return &RDShaderFile{
    obj: obj,
  }
}

func (me *RDShaderFile) BaseClass() string {
  return "RDShaderFile"
}
