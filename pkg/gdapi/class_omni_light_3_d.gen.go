// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OmniLight3D struct {
  obj gdc.ObjectPtr
}

func createOmniLight3D(obj gdc.ObjectPtr) *OmniLight3D {
  return &OmniLight3D{
    obj: obj,
  }
}

func (me *OmniLight3D) BaseClass() string {
  return "OmniLight3D"
}
