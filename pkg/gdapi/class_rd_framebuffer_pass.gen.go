// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDFramebufferPass struct {
  obj gdc.ObjectPtr
}

func createRDFramebufferPass(obj gdc.ObjectPtr) *RDFramebufferPass {
  return &RDFramebufferPass{
    obj: obj,
  }
}

func (me *RDFramebufferPass) BaseClass() string {
  return "RDFramebufferPass"
}
