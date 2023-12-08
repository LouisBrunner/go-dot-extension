// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Gradient struct {
  obj gdc.ObjectPtr
}

func createGradient(obj gdc.ObjectPtr) *Gradient {
  return &Gradient{
    obj: obj,
  }
}

func (me *Gradient) BaseClass() string {
  return "Gradient"
}
