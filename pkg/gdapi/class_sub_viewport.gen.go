// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SubViewport struct {
  obj gdc.ObjectPtr
}

func createSubViewport(obj gdc.ObjectPtr) *SubViewport {
  return &SubViewport{
    obj: obj,
  }
}

func (me *SubViewport) BaseClass() string {
  return "SubViewport"
}
