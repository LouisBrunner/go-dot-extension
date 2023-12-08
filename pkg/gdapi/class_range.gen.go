// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Range struct {
  obj gdc.ObjectPtr
}

func createRange(obj gdc.ObjectPtr) *Range {
  return &Range{
    obj: obj,
  }
}

func (me *Range) BaseClass() string {
  return "Range"
}
