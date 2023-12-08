// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Separator struct {
  obj gdc.ObjectPtr
}

func createSeparator(obj gdc.ObjectPtr) *Separator {
  return &Separator{
    obj: obj,
  }
}

func (me *Separator) BaseClass() string {
  return "Separator"
}
