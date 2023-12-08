// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Performance struct {
  obj gdc.ObjectPtr
}

func createPerformance(obj gdc.ObjectPtr) *Performance {
  return &Performance{
    obj: obj,
  }
}

func (me *Performance) BaseClass() string {
  return "Performance"
}
