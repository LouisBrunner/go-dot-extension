// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDTextureView struct {
  obj gdc.ObjectPtr
}

func createRDTextureView(obj gdc.ObjectPtr) *RDTextureView {
  return &RDTextureView{
    obj: obj,
  }
}

func (me *RDTextureView) BaseClass() string {
  return "RDTextureView"
}
