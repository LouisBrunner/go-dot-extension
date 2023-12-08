// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HSeparator struct {
  obj gdc.ObjectPtr
}

func createHSeparator(obj gdc.ObjectPtr) *HSeparator {
  return &HSeparator{
    obj: obj,
  }
}

func (me *HSeparator) BaseClass() string {
  return "HSeparator"
}
