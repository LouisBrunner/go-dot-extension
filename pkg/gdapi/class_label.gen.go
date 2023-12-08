// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Label struct {
  obj gdc.ObjectPtr
}

func createLabel(obj gdc.ObjectPtr) *Label {
  return &Label{
    obj: obj,
  }
}

func (me *Label) BaseClass() string {
  return "Label"
}
