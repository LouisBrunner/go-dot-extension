// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CapsuleMesh struct {
  obj gdc.ObjectPtr
}

func createCapsuleMesh(obj gdc.ObjectPtr) *CapsuleMesh {
  return &CapsuleMesh{
    obj: obj,
  }
}

func (me *CapsuleMesh) BaseClass() string {
  return "CapsuleMesh"
}
