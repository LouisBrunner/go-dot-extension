// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDTextureFormat struct {
  obj gdc.ObjectPtr
}

func createRDTextureFormat(obj gdc.ObjectPtr) *RDTextureFormat {
  return &RDTextureFormat{
    obj: obj,
  }
}

func (me *RDTextureFormat) BaseClass() string {
  return "RDTextureFormat"
}
