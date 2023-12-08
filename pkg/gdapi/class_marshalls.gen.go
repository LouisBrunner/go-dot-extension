// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Marshalls struct {
  obj gdc.ObjectPtr
}

func createMarshalls(obj gdc.ObjectPtr) *Marshalls {
  return &Marshalls{
    obj: obj,
  }
}

func (me *Marshalls) BaseClass() string {
  return "Marshalls"
}
