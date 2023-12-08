// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GraphEdit struct {
  obj gdc.ObjectPtr
}

func createGraphEdit(obj gdc.ObjectPtr) *GraphEdit {
  return &GraphEdit{
    obj: obj,
  }
}

func (me *GraphEdit) BaseClass() string {
  return "GraphEdit"
}
