// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ReferenceRect struct {
  obj gdc.ObjectPtr
}

func createReferenceRect(obj gdc.ObjectPtr) *ReferenceRect {
  return &ReferenceRect{
    obj: obj,
  }
}

func (me *ReferenceRect) BaseClass() string {
  return "ReferenceRect"
}
