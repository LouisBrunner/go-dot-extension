// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VSeparator struct {
  obj gdc.ObjectPtr
}

func createVSeparator(obj gdc.ObjectPtr) *VSeparator {
  return &VSeparator{
    obj: obj,
  }
}

func (me *VSeparator) BaseClass() string {
  return "VSeparator"
}
