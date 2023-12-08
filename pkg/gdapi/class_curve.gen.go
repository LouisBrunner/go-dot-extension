// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Curve struct {
  obj gdc.ObjectPtr
}

func createCurve(obj gdc.ObjectPtr) *Curve {
  return &Curve{
    obj: obj,
  }
}

func (me *Curve) BaseClass() string {
  return "Curve"
}
