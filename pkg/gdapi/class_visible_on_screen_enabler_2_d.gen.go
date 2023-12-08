// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisibleOnScreenEnabler2D struct {
  obj gdc.ObjectPtr
}

func createVisibleOnScreenEnabler2D(obj gdc.ObjectPtr) *VisibleOnScreenEnabler2D {
  return &VisibleOnScreenEnabler2D{
    obj: obj,
  }
}

func (me *VisibleOnScreenEnabler2D) BaseClass() string {
  return "VisibleOnScreenEnabler2D"
}
